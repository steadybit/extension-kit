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

	// protocolHTTP and protocolGRPC are the OTLP transport protocols supported
	// by the Go SDK. The OpenTelemetry specification defines http/protobuf as
	// the default, so an unset OTEL_EXPORTER_OTLP_PROTOCOL selects it.
	protocolHTTP = "http/protobuf"
	protocolGRPC = "grpc"
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

	// The signal-specific variable takes precedence over the generic one, per
	// the OpenTelemetry specification.
	endpoint := firstNonEmpty(
		os.Getenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT"),
		os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"),
	)
	if endpoint == "" {
		log.Info().Msg("neither OTEL_EXPORTER_OTLP_TRACES_ENDPOINT nor OTEL_EXPORTER_OTLP_ENDPOINT is set; tracing is a noop")
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
	if resolveProtocol() == protocolGRPC {
		return otlptracegrpc.New(ctx)
	}
	return otlptracehttp.New(ctx)
}

// resolveProtocol reports the OTLP transport to use. The signal-specific
// variable wins over the generic one, an unset value yields the
// specification default, and an unsupported value falls back to it with a
// warning rather than failing startup.
func resolveProtocol() string {
	protocol := strings.ToLower(strings.TrimSpace(firstNonEmpty(
		os.Getenv("OTEL_EXPORTER_OTLP_TRACES_PROTOCOL"),
		os.Getenv("OTEL_EXPORTER_OTLP_PROTOCOL"),
	)))
	switch protocol {
	case "":
		return protocolHTTP
	case protocolHTTP, protocolGRPC:
		return protocol
	default:
		log.Warn().Str("protocol", protocol).Msgf("unsupported OTLP protocol; falling back to %s", protocolHTTP)
		return protocolHTTP
	}
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
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
