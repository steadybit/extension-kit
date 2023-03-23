# Changelog

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

 - Added a new utility function `Listen` to `exthttp` package to listen on a port and serve HTTP requests. The function also takes care of establishing an HTTPS server with mutual TLS when instructed to through environment variables.

## v1.5.0

 - Support for the `STEADYBIT_LOG_FORMAT` env variable. When set to `json`, extensions will log JSON lines to stderr.

## v1.4.0

 - adds conversion helper to `extconversion`. This is helpful to encode and decode ActionKit's state.

## v1.3.2

 - debug messages in `exthttp` are missing the request ID.

## v1.3.1

 -  stack overflow error in logging HTTP writer

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