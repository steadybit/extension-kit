/*
 * Copyright 2023 steadybit GmbH. All rights reserved.
 */

// Package exthttp supports setup of HTTP servers to implement the *Kit contracts. To keep the resulting binary small
// the net/http server is used.
package exthttp

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
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

// RegisterHttpHandler registers a handler for the given path. Also adds panic recovery and request logging around the handler.
func RegisterHttpHandler(path string, handler func(w http.ResponseWriter, r *http.Request, body []byte)) {
	http.Handle(path, PanicRecovery(LogRequest(handler)))
}

// GetterAsHandler turns a getter function into a handler function. Typically used in combination with the RegisterHttpHandler function.
func GetterAsHandler[T any](handler func() T) func(w http.ResponseWriter, r *http.Request, body []byte) {
	return func(w http.ResponseWriter, r *http.Request, body []byte) {
		WriteBody(w, handler())
	}
}

func PanicRecovery(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Error().Msgf("Panic: %v\n %s", err, string(debug.Stack()))
				WriteError(w, extension_kit.ToError("Internal Server Error", nil))
			}
		}()
		next(w, r)
	}
}

type LoggingHttpResponseWriter struct {
	delegate http.ResponseWriter
	reqId    string
}

func (w *LoggingHttpResponseWriter) Header() http.Header {
	return w.delegate.Header()
}

func (w *LoggingHttpResponseWriter) Write(bytes []byte) (int, error) {
	log.Debug().Msgf("Req %s response body length: %d bytes", w.reqId, len(bytes))
	log.Trace().Msgf("Req %s response body: %s", w.reqId, bytes)
	return w.delegate.Write(bytes)
}

func (w *LoggingHttpResponseWriter) WriteHeader(statusCode int) {
	log.Debug().Msgf("Req %s response status code: %d", w.reqId, statusCode)
	w.delegate.WriteHeader(statusCode)
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

func LogRequest(next func(w http.ResponseWriter, r *http.Request, body []byte)) http.HandlerFunc {
	return LogRequestWithLevelFunc(next, func(r *http.Request) zerolog.Level {
		if r.Method == "GET" {
			return zerolog.DebugLevel
		} else {
			return zerolog.InfoLevel
		}
	})
}

func LogRequestWithLevelFunc(next func(w http.ResponseWriter, r *http.Request, body []byte), level func(r *http.Request) zerolog.Level) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqId := uuid.New().String()
		if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
			log.WithLevel(level(r)).Msgf("Req %s: %s %s (Multipart)", reqId, r.Method, r.URL)
			next(extutil.Ptr(LoggingHttpResponseWriter{
				delegate: w,
				reqId:    reqId,
			}), r, nil)
			return
		}

		body, bodyReadErr := io.ReadAll(r.Body)
		if bodyReadErr != nil {
			http.Error(w, bodyReadErr.Error(), http.StatusBadRequest)
			return
		}

		bodyLength := len(body)
		if bodyLength == 0 {
			log.WithLevel(level(r)).Msgf("Req %s: %s %s", reqId, r.Method, r.URL)
		} else {
			log.WithLevel(level(r)).Msgf("Req %s: %s %s with %d byte body", reqId, r.Method, r.URL, bodyLength)
		}
		log.Debug().Msgf("Req %s body: %s", reqId, body)

		next(extutil.Ptr(LoggingHttpResponseWriter{
			delegate: w,
			reqId:    reqId,
		}), r, body)
	}
}

func LogRequestWithLevel(next func(w http.ResponseWriter, r *http.Request, body []byte), level zerolog.Level) http.HandlerFunc {
	return LogRequestWithLevelFunc(next, func(r *http.Request) zerolog.Level { return level })
}

// WriteError writes the error as the HTTP response body with status code 500.
func WriteError(w http.ResponseWriter, err extension_kit.ExtensionError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	logEvent := log.Error()
	if err.Detail != nil {
		logEvent.Str("details", *err.Detail)
	}
	logEvent.Msgf(err.Title)

	encodeErr := json.NewEncoder(w).Encode(err)
	if encodeErr != nil {
		log.Err(encodeErr).Msgf("Failed to write ExtensionError as response body")
	}
}

// WriteBody writes the given value as the HTTP response body as JSON with status code 200.
func WriteBody(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	encodeErr := json.NewEncoder(w).Encode(response)
	if encodeErr != nil {
		log.Err(encodeErr).Msgf("Failed to response body")
	}
}
