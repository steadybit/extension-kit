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

func resetGlobals(t *testing.T) {
	t.Helper()
	otel.SetTracerProvider(otel.GetTracerProvider())
	extsignals.RemoveSignalHandlersByName(signalHandlerName)
}

func TestInitOpenTelemetry_SdkDisabled_InstallsNoopProvider(t *testing.T) {
	t.Setenv("OTEL_SDK_DISABLED", "true")
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4317")
	defer resetGlobals(t)

	shutdown := InitOpenTelemetry()
	require.NotNil(t, shutdown)

	_, isSdk := otel.GetTracerProvider().(*sdktrace.TracerProvider)
	assert.False(t, isSdk, "TracerProvider should be noop when OTEL_SDK_DISABLED=true")
}

func TestInitOpenTelemetry_MissingEndpoint_InstallsNoopProvider(t *testing.T) {
	t.Setenv("OTEL_SDK_DISABLED", "")
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "")
	defer resetGlobals(t)

	shutdown := InitOpenTelemetry()
	require.NotNil(t, shutdown)

	_, isSdk := otel.GetTracerProvider().(*sdktrace.TracerProvider)
	assert.False(t, isSdk, "TracerProvider should be noop when OTEL_EXPORTER_OTLP_ENDPOINT is unset")
}

func TestInitOpenTelemetry_ValidEndpoint_InstallsSdkProvider(t *testing.T) {
	t.Setenv("OTEL_SDK_DISABLED", "")
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4317")
	t.Setenv("OTEL_EXPORTER_OTLP_PROTOCOL", "grpc")
	defer resetGlobals(t)

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
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://localhost:4317")
	defer resetGlobals(t)

	shutdown := InitOpenTelemetry()
	require.NotNil(t, shutdown)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	assert.NoError(t, shutdown(ctx))
	assert.NoError(t, shutdown(ctx), "second shutdown should not error")
}
