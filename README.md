<img src="./logo.png" height="70" align="right" alt="ExtensionKit logo depicting a wrench within a rounded rectangle on the background">

# ExtensionKit

Through kits like ActionKit and DiscoveryKit, Steadybit can be extended with new capabilities. Such *Kit usages are
called extensions. ExtensionKit
contains helpful utilities and best practices for extension authors leveraging the Go programming language.

## Installation

Add the following to your `go.mod` file:

```
go get github.com/steadybit/extension-kit
```

## Environment Variables

Extension using this extension kit can be configured through environment variables. All official extension
Helm charts set some of them from dedicated values; the rest can be set through the chart's `extraEnv`.
The following environment variables are supported:

| Environment Variable                  | Meaning                                                                                                                                                                                    | Helm value                                                               | Default |
|---------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------|---------|
| `STEADYBIT_EXTENSION_PORT`            | Overwrite the extensions default port number that the HTTP server should bind to.                                                                                                          | `extraEnv`                                                               |         |
| `STEADYBIT_EXTENSION_HEALTH_PORT`     | Overwrite the extensions default port number that the HTTP server for the health endpoints should bind to.                                                                                 | `extraEnv`                                                               |         |
| `STEADYBIT_EXTENSION_TLS_SERVER_CERT` | Optional absolute path to a TLS certificate that will be used to open an **HTTPS** server.                                                                                                 | `tls.server.certificate.fromSecret` or `tls.server.certificate.path`     |         |
| `STEADYBIT_EXTENSION_TLS_SERVER_KEY`  | Optional absolute path to a file containing the key to the server certificate.                                                                                                             | `tls.server.certificate.fromSecret` or `tls.server.certificate.key.path` |         |
| `STEADYBIT_EXTENSION_TLS_CLIENT_CAS`  | Optional comma-separated list of absolute paths to files containing TLS certificates. When specified, the server will expect clients to authenticate using mutual TLS.                     | `tls.client.certificates.fromSecrets` or `tls.client.certificates.paths` |         |
| `STEADYBIT_EXTENSION_UNIX_SOCKET`     | If set the extension will listen using a unix domain socket instead of tcp.                                                                                                                | `extraEnv`                                                               |         |
| `STEADYBIT_LOG_FORMAT`                | Defines the log format that the extension will use. Possible values are `text` and `json`.                                                                                                 | `logging.format`                                                         | text    |
| `STEADYBIT_LOG_LEVEL`                 | Defines the active log level. Possible values are `debug`, `info`, `warn` and `error`.                                                                                                     | `logging.level`                                                          | info    |
| `STEADYBIT_LOG_COLOR`                 | Defines colorization of log output. Possible values are `true`, `false` and unset. If unset will use color only if stderr is a terminal.                                                   | `extraEnv`                                                               |         |
| `STEADYBIT_EXTENSION_ENABLE_PPROF`    | Enables the `/debug/pprof/` handlers for debugging                                                                                                                                         | `extraEnv`                                                               | false   |

## OpenTelemetry Tracing

Extensions can export OpenTelemetry traces so that an operator debugging a slow or
timing-out action can see what happened *inside* the extension, not just the agent's
side of the call. Every handler registered through `exthttp.RegisterHttpHandler`
produces a `METHOD /path` server span, and incoming trace context is honoured, so an
extension's spans join the caller's trace.

Tracing is **off until an OTLP endpoint is configured**. Extensions that never enable
it pay nothing: the instrumentation short-circuits before it does any work.

### Enabling it in an extension

Call `extotel.InitOpenTelemetry()` from `main()`, before registering handlers:

```go
import "github.com/steadybit/extension-kit/extotel"

func main() {
	extlogging.InitZeroLog()
	extotel.InitOpenTelemetry()
	// ... register actions and discoveries
}
```

It returns a shutdown function that flushes buffered spans. You do not need to call it:
`InitOpenTelemetry` registers an `extsignals` handler that runs the flush on SIGTERM
after the HTTP server has stopped.

An extension that configures the OpenTelemetry SDK itself, rather than through
`extotel`, must also call `exthttp.SetTracingEnabled(true)`, or its handlers will not
be traced.

### Configuration

Configured entirely through the standard `OTEL_*` variables, so collectors and backends
can be swapped without code changes. The signal-specific variable wins over the generic
one where both exist.

| Environment Variable                 | Meaning                                                                                                                  | Helm value | Default         |
|--------------------------------------|--------------------------------------------------------------------------------------------------------------------------|------------|-----------------|
| `OTEL_EXPORTER_OTLP_ENDPOINT`        | OTLP endpoint to export spans to. **Tracing stays off while this and the traces-specific variant are both unset.**        | `extraEnv` |                 |
| `OTEL_EXPORTER_OTLP_TRACES_ENDPOINT` | Traces-only endpoint. Takes precedence over the generic one.                                                              | `extraEnv` |                 |
| `OTEL_EXPORTER_OTLP_PROTOCOL`        | `grpc` or `http/protobuf`. An unsupported value warns and falls back to the default rather than failing startup.          | `extraEnv` | `grpc`          |
| `OTEL_EXPORTER_OTLP_TRACES_PROTOCOL` | Traces-only protocol. Takes precedence over the generic one.                                                              | `extraEnv` | `grpc`          |
| `OTEL_SERVICE_NAME`                  | Service name on the exported spans. A warning is logged when unset, since spans would be tagged `unknown_service:<binary>`. | `extraEnv` |                 |
| `OTEL_SDK_DISABLED`                  | `true` forces tracing off even when an endpoint is configured.                                                            | `extraEnv` | false           |

Sampling and batching are left to the standard SDK variables — `OTEL_TRACES_SAMPLER`,
`OTEL_TRACES_SAMPLER_ARG`, `OTEL_BSP_*` — rather than being tuned per extension.

> **Match the protocol to the port.** The default is `grpc`, which means port
> **4317**. This departs from the OpenTelemetry specification, whose default is
> `http/protobuf`: the Steadybit agent's Java SDK defaults to gRPC and extensions are
> configured alongside it, so an endpoint copied from the agent works as written.
> If you point `OTEL_EXPORTER_OTLP_ENDPOINT` at an OTLP/HTTP collector on **4318**,
> set `OTEL_EXPORTER_OTLP_PROTOCOL=http/protobuf` as well, or the exporter will talk
> gRPC to a port that will not answer and no spans will arrive. Export failures are
> logged through the extension's normal logger, so check the extension log if a
> configured endpoint stays silent.

Example, exporting to a collector over gRPC (the default, so no protocol needed):

```yaml
extraEnv:
  - name: OTEL_EXPORTER_OTLP_ENDPOINT
    value: "http://otel-collector.observability:4317"
  - name: OTEL_SERVICE_NAME
    value: "steadybit-extension-http"
```
