/*
 * Copyright 2023 steadybit GmbH. All rights reserved.
 */

// Package exthttp supports setup of HTTP servers to implement the *Kit contracts. To keep the resulting binary small
// the net/http server is used.
package exthttp

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"github.com/steadybit/extension-kit"
	"github.com/steadybit/extension-kit/extutil"
	"io"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

type Handler func(w http.ResponseWriter, r *http.Request, body []byte)

// RegisterHttpHandler registers a handler for the given path. Also adds panic recovery and request logging around the handler.
func RegisterHttpHandler(path string, handler Handler) {
	http.Handle(path, PanicRecovery(LogRequest(handler)))
}

// RegisterHttpHandlerWithLogLevel registers a handler for the given path. Also adds panic recovery and request logging with a given log level around the handler.
func RegisterHttpHandlerWithLogLevel(path string, handler Handler, defaultLevel zerolog.Level) {
	http.Handle(path, PanicRecovery(LogRequestWithDefaultLogLevel(handler, defaultLevel)))
}

// GetterAsHandler turns a getter function into a handler function. Typically used in combination with the RegisterHttpHandler function.
func GetterAsHandler[T any](handler func() T) Handler {
	return func(w http.ResponseWriter, r *http.Request, body []byte) {
		WriteBody(w, handler())
	}
}

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Error().Msgf("Panic: %v\n %s", err, string(debug.Stack()))
				response := extension_kit.ToError("Internal Server Error", nil)
				response.Detail = extutil.Ptr(fmt.Sprintf("Panic: %v", err))
				WriteError(w, response)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func RequestTimeoutHeaderAware(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timeout := r.Header.Get("Request-Timeout")
		if timeout == "" {
			timeout = r.Header.Get("X-Request-Timeout")
		}
		if timeout != "" {
			timeoutValue, err := strconv.ParseFloat(timeout, 32)
			if err == nil {
				log.Trace().Msgf("Using handler timeout %.1fs", timeoutValue)
				http.TimeoutHandler(http.HandlerFunc(next), time.Duration(timeoutValue*1000)*time.Millisecond, "Request timed out.").ServeHTTP(w, r)
				return
			}
		}
		next(w, r)
	}
}

func LogRequest(next Handler) http.Handler {
	return LogRequestWithDefaultLogLevel(next, zerolog.InfoLevel)
}

func LogRequestWithDefaultLogLevel(next Handler, defaultLevel zerolog.Level) http.Handler {
	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		level := defaultLevel
		if r.Method == "GET" {
			level = zerolog.DebugLevel
		}

		var reqBody []byte = nil
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
			if bytes, err := io.ReadAll(r.Body); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				reqBody = bytes
			}
		}

		hlog.FromRequest(r).Debug().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("req_size", len(reqBody)).
			Bytes("body", reqBody).
			Msg("Request received")

		hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).WithLevel(level).
				Str("method", r.Method).
				Stringer("url", r.URL).
				Int("res_size", size).
				Int("req_size", len(reqBody)).
				Dur("duration", duration).
				Int("status", status).
				Msg("")
		})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next(w, r, reqBody)
		})).ServeHTTP(w, r)
	})

	handler = hlog.RequestIDHandler("req_id", "Request-Id")(handler)
	handler = hlog.NewHandler(log.Logger)(handler)
	return handler
}

// WriteError writes the error as the HTTP response body with status code 500.
func WriteError(w http.ResponseWriter, err extension_kit.ExtensionError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	logEvent := log.Error()
	if err.Detail != nil {
		logEvent.Str("details", *err.Detail)
	}
	logEvent.Msg(err.Title)

	encodeErr := json.NewEncoder(w).Encode(err)
	if encodeErr != nil {
		log.Err(encodeErr).Msgf("Failed to write ExtensionError as response body")
	}
}

// WriteBody writes the given value as the HTTP response body as JSON with status code 200.
func WriteBody(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Err(err).Msgf("Failed to response body")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func IfNoneMatchHandler(etagFn func() string, delegate Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request, body []byte) {
		etag := etagFn()
		if len(etag) > 0 {
			if ifNoneMatch := r.Header.Get("If-None-Match"); etag == ifNoneMatch {
				w.WriteHeader(http.StatusNotModified)
				return
			}
			w.Header().Set("ETag", etag)
		}
		delegate(w, r, body)
	}
}
