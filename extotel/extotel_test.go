// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

package extotel

import (
	"context"
	"testing"
	"time"

	"github.com/steadybit/extension-kit/extsignals"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// restoreGlobals snapshots the process-wide OTel globals before a test mutates
// them and puts them back afterwards, so an SDK provider installed by one test
// does not leak into the next one.
func restoreGlobals(t *testing.T) {
	t.Helper()
	prevProvider := otel.GetTracerProvider()
	prevPropagator := otel.GetTextMapPropagator()
	t.Cleanup(func() {
		otel.SetTracerProvider(prevProvider)
		otel.SetTextMapPropagator(prevPropagator)
		extsignals.RemoveSignalHandlersByName(signalHandlerName)
	})
}

func TestInitOpenTelemetry_SdkDisabled_InstallsNoopProvider(t *testing.T) {
	t.Setenv("OTEL_SDK_DISABLED", "true")
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4317")
	restoreGlobals(t)

	shutdown := InitOpenTelemetry()
	require.NotNil(t, shutdown)

	_, isSdk := otel.GetTracerProvider().(*sdktrace.TracerProvider)
	assert.False(t, isSdk, "TracerProvider should be noop when OTEL_SDK_DISABLED=true")
}

func TestInitOpenTelemetry_MissingEndpoint_InstallsNoopProvider(t *testing.T) {
	t.Setenv("OTEL_SDK_DISABLED", "")
	t.Setenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT", "")
	t.Setenv("OTEL_EXPORTER_OTLP_TRACES_PROTOCOL", "")
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "")
	t.Setenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT", "")
	restoreGlobals(t)

	shutdown := InitOpenTelemetry()
	require.NotNil(t, shutdown)

	_, isSdk := otel.GetTracerProvider().(*sdktrace.TracerProvider)
	assert.False(t, isSdk, "TracerProvider should be noop when OTEL_EXPORTER_OTLP_ENDPOINT is unset")
}

func TestInitOpenTelemetry_ValidEndpoint_InstallsSdkProvider(t *testing.T) {
	t.Setenv("OTEL_SDK_DISABLED", "")
	t.Setenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT", "")
	t.Setenv("OTEL_EXPORTER_OTLP_TRACES_PROTOCOL", "")
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4317")
	t.Setenv("OTEL_EXPORTER_OTLP_PROTOCOL", "grpc")
	restoreGlobals(t)

	shutdown := InitOpenTelemetry()
	require.NotNil(t, shutdown)
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_ = shutdown(ctx)
	}()

	_, isSdk := otel.GetTracerProvider().(*sdktrace.TracerProvider)
	assert.True(t, isSdk, "TracerProvider should be sdktrace provider when OTLP endpoint is configured")
}

func TestInitOpenTelemetry_ShutdownIdempotent(t *testing.T) {
	t.Setenv("OTEL_SDK_DISABLED", "")
	t.Setenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT", "")
	t.Setenv("OTEL_EXPORTER_OTLP_TRACES_PROTOCOL", "")
	// 4318 is the OTLP/HTTP port, matching the default protocol; pairing 4317
	// with http/protobuf is the misconfiguration this package warns about.
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4318")
	restoreGlobals(t)

	shutdown := InitOpenTelemetry()
	require.NotNil(t, shutdown)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	assert.NoError(t, shutdown(ctx))
	assert.NoError(t, shutdown(ctx), "second shutdown should not error")
}

func TestInitOpenTelemetry_TracesEndpointAloneInstallsSdkProvider(t *testing.T) {
	// The signal-specific variable is enough on its own; an operator that only
	// sets OTEL_EXPORTER_OTLP_TRACES_ENDPOINT must still get a real provider.
	t.Setenv("OTEL_SDK_DISABLED", "")
	t.Setenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT", "")
	t.Setenv("OTEL_EXPORTER_OTLP_TRACES_PROTOCOL", "")
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "")
	t.Setenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT", "http://localhost:4318")
	restoreGlobals(t)

	shutdown := InitOpenTelemetry()
	require.NotNil(t, shutdown)
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_ = shutdown(ctx)
	}()

	_, isSdk := otel.GetTracerProvider().(*sdktrace.TracerProvider)
	assert.True(t, isSdk, "TracerProvider should be the sdktrace provider when only the traces endpoint is set")
}

func TestResolveProtocol(t *testing.T) {
	tests := []struct {
		name           string
		generic        string
		tracesSpecific string
		want           string
	}{
		{name: "unset defaults to the specification default", want: protocolHTTP},
		{name: "generic grpc", generic: "grpc", want: protocolGRPC},
		{name: "generic http/protobuf", generic: "http/protobuf", want: protocolHTTP},
		{name: "traces-specific wins over generic", generic: "http/protobuf", tracesSpecific: "grpc", want: protocolGRPC},
		{name: "case and padding are tolerated", generic: "  GRPC ", want: protocolGRPC},
		{name: "unsupported value falls back", generic: "http/json", want: protocolHTTP},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("OTEL_EXPORTER_OTLP_PROTOCOL", tt.generic)
			t.Setenv("OTEL_EXPORTER_OTLP_TRACES_PROTOCOL", tt.tracesSpecific)
			assert.Equal(t, tt.want, resolveProtocol())
		})
	}
}
