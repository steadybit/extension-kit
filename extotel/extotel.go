// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

// Package extotel provides opinionated initialization of the OpenTelemetry
// SDK for Steadybit extensions. It is configured entirely via standard
// OTEL_* environment variables so operators can swap collectors and
// backends without code changes.
package extotel

import (
	"context"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/steadybit/extension-kit/extsignals"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	signalHandlerName = "ShutdownOpenTelemetry"
	shutdownTimeout   = 5 * time.Second
)

// InitOpenTelemetry configures the global TracerProvider and propagators
// based on standard OTEL_* environment variables. It is safe to call once
// from main(). The returned function shuts the provider down and flushes
// buffered spans; it is idempotent. In production the registered
// extsignals handler performs shutdown on SIGTERM/SIGINT.
func InitOpenTelemetry() func(context.Context) error {
	if strings.EqualFold(os.Getenv("OTEL_SDK_DISABLED"), "true") {
		log.Info().Msg("OpenTelemetry SDK disabled via OTEL_SDK_DISABLED; tracing is a noop")
		return noopShutdown
	}

	endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if endpoint == "" {
		log.Info().Msg("OTEL_EXPORTER_OTLP_ENDPOINT not set; tracing is a noop")
		return noopShutdown
	}

	if os.Getenv("OTEL_SERVICE_NAME") == "" {
		log.Warn().Msg("OTEL_SERVICE_NAME not set; spans will be tagged with the SDK default 'unknown_service:<binary>'")
	}

	exporter, err := newExporter(context.Background())
	if err != nil {
		log.Warn().Err(err).Msg("failed to create OTLP trace exporter; tracing is a noop")
		return noopShutdown
	}

	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	shutdown := shutdownFunc(tp)
	extsignals.AddSignalHandler(extsignals.SignalHandler{
		Handler: func(_ os.Signal) {
			ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
			defer cancel()
			if err := shutdown(ctx); err != nil {
				log.Warn().Err(err).Msg("OpenTelemetry shutdown failed")
			}
		},
		Order: extsignals.OrderStopExtensionHttp + 1,
		Name:  signalHandlerName,
	})

	log.Info().Str("endpoint", endpoint).Msg("OpenTelemetry tracing initialized")
	return shutdown
}

func newExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	protocol := strings.ToLower(os.Getenv("OTEL_EXPORTER_OTLP_PROTOCOL"))
	if protocol == "" {
		protocol = "grpc"
	}
	switch protocol {
	case "http/protobuf":
		return otlptracehttp.New(ctx)
	case "grpc":
		return otlptracegrpc.New(ctx)
	default:
		return otlptracegrpc.New(ctx)
	}
}

func shutdownFunc(tp *sdktrace.TracerProvider) func(context.Context) error {
	var once sync.Once
	var err error
	return func(ctx context.Context) error {
		once.Do(func() {
			err = tp.Shutdown(ctx)
		})
		return err
	}
}

func noopShutdown(context.Context) error { return nil }
