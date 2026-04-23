/*
 * Copyright 2026 steadybit GmbH. All rights reserved.
 */

package exthttp

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

func TestRegisterHttpHandler_EmitsSpanPerRequest(t *testing.T) {
	recorder := tracetest.NewSpanRecorder()
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	prevProvider := otel.GetTracerProvider()
	otel.SetTracerProvider(tp)
	defer otel.SetTracerProvider(prevProvider)

	http.DefaultServeMux = http.NewServeMux()
	RegisterHttpHandler("/traced", func(w http.ResponseWriter, r *http.Request, body []byte) {
		w.WriteHeader(http.StatusNoContent)
	})

	req := httptest.NewRequest(http.MethodPost, "/traced", nil)
	rr := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rr, req)

	require.Equal(t, http.StatusNoContent, rr.Code)

	spans := recorder.Ended()
	require.Len(t, spans, 1, "expected exactly one span")
	assert.Equal(t, "POST /traced", spans[0].Name())
}

func TestRegisterHttpHandler_NoSpanWhenOtelNotConfigured(t *testing.T) {
	prevProvider := otel.GetTracerProvider()
	otel.SetTracerProvider(noop.NewTracerProvider())
	defer otel.SetTracerProvider(prevProvider)

	var capturedSpanCtx trace.SpanContext
	http.DefaultServeMux = http.NewServeMux()
	RegisterHttpHandler("/untraced", func(w http.ResponseWriter, r *http.Request, body []byte) {
		capturedSpanCtx = trace.SpanFromContext(r.Context()).SpanContext()
		w.WriteHeader(http.StatusNoContent)
	})

	req := httptest.NewRequest(http.MethodPost, "/untraced", nil)
	rr := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rr, req)

	require.Equal(t, http.StatusNoContent, rr.Code)
	assert.False(t, capturedSpanCtx.IsValid(), "handler should see a noop span with invalid span context when OTEL is not configured")
}
