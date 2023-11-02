# Changelog

## 1.8.10

- provide http handler to deal with ETags

## 1.8.9

- log gomaxprocs in debug information

## 1.8.8

- log properly formatted timestamps in json logs
- use structured logging for http request logging

## 1.8.7

- remove broken seccomp stuff from extruntime.LogRuntimeInformation

## 1.8.6

- add LogRuntimeInformation

## 1.8.5

- log tls errors on debug

## 1.8.4

- Support loading client certificates from directories

## 1.8.3

- reload tls server key pair if the file changes

## 1.8.2

- only log response body in trace level

## 1.8.1

- add seconds and milliseconds to log output

## 1.8.0

- add support for listening on unix domain sockets instead of TCP ports
- log output is now colorized by default if the output is a terminal; can be overriden using `STEADYBIT_LOG_COLOR`

## 1.7.17

- only log caller when using debug level

## 1.7.16

- Add Caller to logging
- Added extfile with some file helpers

## 1.7.15

- HTTP-Request-Logger: do not read multipart bodies

## 1.7.14

- added util methods mainly for tests (JsonMangle)

## 1.7.13


## 1.7.12

- reduces logging for http requests

## 1.7.11

- added util methods for type conversions

## 1.7.10

- logging for readiness state

## v1.7.9

- allow overriding port for probes to be overrinden via STEADYBIT_EXTENSION_HEALTH_PORT

## v1.7.8

- start liveness and readiness probe on different port

## v1.7.7

## v1.7.6

- add exthttp.LogRequestWithLevel

## v1.7.6

- default readiness probe is true
- use atomic variable for readiness switch

## v1.7.5

- don't use generics for the conversion helper
- added first draft of readiness and liveness probe helpers

## v1.7.4

- Let ExtensionError implement the error interface
- Set timeout for handlers if Request-Timeout header is set.

## v1.7.3

- Always log errors written as http response

## v1.7.2

- Don't print "0 bytes" when there is no request body

## v1.7.1

- Trim leading `v` character in `extbuild.GetSemverVersionStringOrUnknown()` for platform compatibility.

## v1.7.0

- Support `extbuild.ExtensionName`, `extbuild.Version` and `extbuild.Revision` to retrieve build information. You have to fill these fields at build time using:

     ```
     go build -ldflags="-X 'github.com/steadybit/extension-kit/extbuild.ExtensionName=extension-prometheus' -X 'github.com/steadybit/extension-kit/extbuild.Version=v1.0.0' -X 'github.com/steadybit/extension-kit/extbuild.Revision=e3f9616ba2e838d0d3a4472cd0d0cb2e39a06e8f'"
     ```
- Extensions can now call `extbuild.PrintBuildInformation()` within their `main()` function to generate useful debugging information.
- Extensions can now call `extbuild.GetSemverVersionStringOrUnknown()` to get a fitting version number for action and type definitions.

## v1.6.0

- Added a new utility function `Listen` to `exthttp` package to listen on a port and serve HTTP requests. The function also takes care of establishing an HTTPS
  server with mutual TLS when instructed to through environment variables.

## v1.5.0

- Support for the `STEADYBIT_LOG_FORMAT` env variable. When set to `json`, extensions will log JSON lines to stderr.

## v1.4.0

- adds conversion helper to `extconversion`. This is helpful to encode and decode ActionKit's state.

## v1.3.2

- debug messages in `exthttp` are missing the request ID.

## v1.3.1

- stack overflow error in logging HTTP writer

## v1.3.0

- `exthttp` will now log request and response bodies on debug level.

## v1.2.0

- the active log level can now be configured through the `STEADYBIT_LOG_LEVEL` environment variable.

## v1.1.0

- add utilities to work with child processes across incoming HTTP calls, e.g., for ActionKit users

## v1.0.1

- fix missing conditional when logging errors

## v1.0.0

- Initial release
