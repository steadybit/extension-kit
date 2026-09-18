// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

// Package extotel provides opinionated initialization of the OpenTelemetry
// SDK for Steadybit extensions. It is configured entirely via standard
// OTEL_* environment variables so operators can swap collectors and
// backends without code changes.
package extotel

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-logr/logr"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/steadybit/extension-kit/exthttp"
	"github.com/steadybit/extension-kit/extsignals"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
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
	// by the Go SDK. An unset OTEL_EXPORTER_OTLP_PROTOCOL selects gRPC, which
	// deliberately departs from the specification default of http/protobuf: the
	// Steadybit agent's Java SDK defaults to gRPC, and extensions are deployed
	// and configured alongside it against an OTLP/gRPC endpoint on port 4317.
	// Defaulting to http/protobuf would make an endpoint copied from the agent
	// export nothing at all.
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

	// OTel's global default logger writes plain text to stderr, which corrupts
	// JSON log output, and export failures otherwise go unnoticed entirely: the
	// extension logs "tracing initialized" and drops every span. Point a
	// misconfigured endpoint at the wrong port and this is the only signal.
	otel.SetErrorHandler(otel.ErrorHandlerFunc(func(err error) {
		log.Warn().Err(err).Msg("OpenTelemetry error")
	}))
	otel.SetLogger(logr.New(zerologSink{}))

	exporter, err := newExporter(context.Background())
	if err != nil {
		log.Warn().Err(err).Msg("failed to create OTLP trace exporter; tracing is a noop")
		return noopShutdown
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSpanProcessor(baggageAttributeProcessor{keys: correlationBaggageKeys}),
	)
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

	exthttp.SetTracingEnabled(true)

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
// variable wins over the generic one, an unset value yields gRPC (see the
// constants for why that, not the specification default), and an unsupported
// value falls back to it with a warning rather than failing startup.
func resolveProtocol() string {
	protocol := strings.ToLower(strings.TrimSpace(firstNonEmpty(
		os.Getenv("OTEL_EXPORTER_OTLP_TRACES_PROTOCOL"),
		os.Getenv("OTEL_EXPORTER_OTLP_PROTOCOL"),
	)))
	switch protocol {
	case "":
		return protocolGRPC
	case protocolHTTP, protocolGRPC:
		return protocol
	default:
		log.Warn().Str("protocol", protocol).Msgf("unsupported OTLP protocol; falling back to %s", protocolGRPC)
		return protocolGRPC
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

// correlationBaggageKeys are the baggage entries copied onto every span this
// extension records. The Steadybit platform puts the experiment execution id
// into baggage and the agent propagates it on every call it makes, so an
// extension's spans can be found by the run that caused them — without it they
// are only reachable by opening the agent's trace, not by querying for the run.
var correlationBaggageKeys = []string{"experiment.execution.id"}

// baggageAttributeProcessor copies selected baggage entries onto spans as they
// start. Baggage rides in the context and is not otherwise recorded, so
// attributes are what make a span queryable in a tracing backend.
//
// Only the listed keys are copied, deliberately: baggage is arbitrary
// caller-supplied key/value data, and copying all of it would put whatever an
// upstream service happened to set into this extension's telemetry.
type baggageAttributeProcessor struct {
	keys []string
}

func (p baggageAttributeProcessor) OnStart(ctx context.Context, span sdktrace.ReadWriteSpan) {
	b := baggage.FromContext(ctx)
	for _, key := range p.keys {
		if value := b.Member(key).Value(); value != "" {
			span.SetAttributes(attribute.String(key, value))
		}
	}
}

func (p baggageAttributeProcessor) OnEnd(sdktrace.ReadOnlySpan) {}

func (p baggageAttributeProcessor) Shutdown(context.Context) error { return nil }

func (p baggageAttributeProcessor) ForceFlush(context.Context) error { return nil }

// zerologSink adapts OTel's logr-based internal logging onto the extension's
// zerolog output, so SDK diagnostics keep the configured log format.
type zerologSink struct {
	name   string
	values []any
}

func (s zerologSink) Init(logr.RuntimeInfo) {}

// OTel logs at V(1) for debug-level detail and V(8)/V(4) for verbose traffic;
// only the first level is worth carrying.
func (s zerologSink) Enabled(level int) bool { return level <= 1 }

func (s zerologSink) Info(_ int, msg string, kv ...any) {
	s.event(log.Debug(), kv).Msg(msg)
}

func (s zerologSink) Error(err error, msg string, kv ...any) {
	s.event(log.Warn().Err(err), kv).Msg(msg)
}

func (s zerologSink) event(e *zerolog.Event, kv []any) *zerolog.Event {
	if s.name != "" {
		e = e.Str("otel.component", s.name)
	}
	for _, pairs := range [][]any{s.values, kv} {
		for i := 0; i+1 < len(pairs); i += 2 {
			e = e.Interface(fmt.Sprint(pairs[i]), pairs[i+1])
		}
	}
	return e
}

func (s zerologSink) WithValues(kv ...any) logr.LogSink {
	return zerologSink{name: s.name, values: append(append([]any{}, s.values...), kv...)}
}

func (s zerologSink) WithName(name string) logr.LogSink {
	if s.name != "" {
		name = s.name + "/" + name
	}
	return zerologSink{name: name, values: s.values}
}
