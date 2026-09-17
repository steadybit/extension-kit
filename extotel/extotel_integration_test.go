// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

package extotel

import (
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/steadybit/extension-kit/exthttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	tracepb "go.opentelemetry.io/proto/otlp/trace/v1"
	"google.golang.org/protobuf/proto"
)

// otlpCollector is a minimal OTLP/HTTP trace receiver: enough of a collector to
// assert that spans really leave the process over the wire, rather than only
// reaching an in-memory SpanRecorder.
type otlpCollector struct {
	server *httptest.Server

	mu    sync.Mutex
	spans []*tracepb.Span
}

func startOtlpCollector(t *testing.T) *otlpCollector {
	t.Helper()
	c := &otlpCollector{}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/traces", func(w http.ResponseWriter, r *http.Request) {
		body, err := readPossiblyGzipped(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var req coltracepb.ExportTraceServiceRequest
		if err := proto.Unmarshal(body, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		c.mu.Lock()
		for _, rs := range req.GetResourceSpans() {
			for _, ss := range rs.GetScopeSpans() {
				c.spans = append(c.spans, ss.GetSpans()...)
			}
		}
		c.mu.Unlock()

		resp, err := proto.Marshal(&coltracepb.ExportTraceServiceResponse{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/x-protobuf")
		_, _ = w.Write(resp)
	})

	c.server = httptest.NewServer(mux)
	t.Cleanup(c.server.Close)
	return c
}

func readPossiblyGzipped(r *http.Request) ([]byte, error) {
	if r.Header.Get("Content-Encoding") != "gzip" {
		return io.ReadAll(r.Body)
	}
	gz, err := gzip.NewReader(r.Body)
	if err != nil {
		return nil, err
	}
	defer func() { _ = gz.Close() }()
	return io.ReadAll(gz)
}

func (c *otlpCollector) received() []*tracepb.Span {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([]*tracepb.Span(nil), c.spans...)
}

// An HTTP request served through RegisterHttpHandler must end up as a span at
// the configured OTLP endpoint. This covers the whole path the extensions rely
// on: env-var configuration, exporter construction, the otelhttp middleware and
// the flush on shutdown.
func TestInitOpenTelemetry_ExportsHandlerSpansOverOTLP(t *testing.T) {
	collector := startOtlpCollector(t)

	t.Setenv("OTEL_SDK_DISABLED", "")
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", collector.server.URL)
	t.Setenv("OTEL_EXPORTER_OTLP_PROTOCOL", protocolHTTP)
	t.Setenv("OTEL_SERVICE_NAME", "extension-kit-integration-test")
	restoreGlobals(t)

	shutdown := InitOpenTelemetry()
	require.NotNil(t, shutdown)

	http.DefaultServeMux = http.NewServeMux()
	exthttp.RegisterHttpHandler("/exported", func(w http.ResponseWriter, r *http.Request, body []byte) {
		w.WriteHeader(http.StatusNoContent)
	})

	req := httptest.NewRequest(http.MethodPost, "/exported", nil)
	rr := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rr, req)
	require.Equal(t, http.StatusNoContent, rr.Code)

	// Shutdown flushes the batch processor, so no polling for the exporter's
	// schedule delay is needed.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	require.NoError(t, shutdown(ctx))

	spans := collector.received()
	require.Len(t, spans, 1, "expected exactly one exported span")
	assert.Equal(t, "POST /exported", spans[0].GetName())
	assert.Equal(t, tracepb.Span_SPAN_KIND_SERVER, spans[0].GetKind())
}

// With no endpoint configured nothing must be exported, so an extension that is
// not opted in pays no network cost.
func TestInitOpenTelemetry_ExportsNothingWhenUnconfigured(t *testing.T) {
	collector := startOtlpCollector(t)

	t.Setenv("OTEL_SDK_DISABLED", "")
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "")
	t.Setenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT", "")
	restoreGlobals(t)

	shutdown := InitOpenTelemetry()
	require.NotNil(t, shutdown)

	http.DefaultServeMux = http.NewServeMux()
	exthttp.RegisterHttpHandler("/unexported", func(w http.ResponseWriter, r *http.Request, body []byte) {
		w.WriteHeader(http.StatusNoContent)
	})

	req := httptest.NewRequest(http.MethodPost, "/unexported", nil)
	rr := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rr, req)
	require.Equal(t, http.StatusNoContent, rr.Code)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	require.NoError(t, shutdown(ctx))

	assert.Empty(t, collector.received(), "no spans may be exported when no endpoint is configured")
}
