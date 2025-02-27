<img src="./logo.png" height="130" align="right" alt="ExtensionKit logo depicting a wrench within a rounded rectangle on the background">

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

Extension using this extension kit can be configured through environment variables. The following environment variables
are supported:

| Environment Variable                  | Meaning                                                                                                                                                                | Default |
|---------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|
| `STEADYBIT_EXTENSION_PORT`            | Overwrite the extensions default port number that the HTTP server should bind to.                                                                                      |         |
| `STEADYBIT_EXTENSION_HEALTH_PORT`     | Overwrite the extensions default port number that the HTTP server for the health endpoints should bind to.                                                             |         |
| `STEADYBIT_EXTENSION_TLS_SERVER_CERT` | Optional absolute path to a TLS certificate that will be used to open an **HTTPS** server.                                                                             |         |
| `STEADYBIT_EXTENSION_TLS_SERVER_KEY`  | Optional absolute path to a file containing the key to the server certificate.                                                                                         |         |
| `STEADYBIT_EXTENSION_TLS_CLIENT_CAS`  | Optional comma-separated list of absolute paths to files containing TLS certificates. When specified, the server will expect clients to authenticate using mutual TLS. |         |
| `STEADYBIT_EXTENSION_UNIX_SOCKET`     | If set the extension will listen using a unix domain socket instead of tcp.                                                                                            |         |
| `STEADYBIT_LOG_FORMAT`                | Defines the log format that the extension will use. Possible values are `text` and `json`.                                                                             | text    |
| `STEADYBIT_LOG_LEVEL`                 | Defines the active log level. Possible values are `debug`, `info`, `warn` and `error`.                                                                                 | info    |
| `STEADYBIT_LOG_COLOR`                 | Defines colorization of log output. Possible values are `true`, `false` and unset. If unset will use color only if stderr is a terminal.                               |         |
| `STEADYBIT_EXTENSION_ENABLE_PPROF`    | Enables the `/debug/pprof/` handlers for debugging                                                                                                                     | false   |
