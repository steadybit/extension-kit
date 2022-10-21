# Changelog

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