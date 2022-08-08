// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

// Package exthttp supports setup of HTTP servers to implement the *Kit contracts. To keep the resulting binary small
// the net/http server is used.
package exthttp

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/steadybit/extension-kit"
	"io/ioutil"
	"net/http"
	"runtime/debug"
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

func LogRequest(next func(w http.ResponseWriter, r *http.Request, body []byte)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, bodyReadErr := ioutil.ReadAll(r.Body)
		if bodyReadErr != nil {
			http.Error(w, bodyReadErr.Error(), http.StatusBadRequest)
			return
		}

		if len(body) > 0 {
			log.Info().Msgf("%s %s with body %s", r.Method, r.URL, body)
		} else {
			log.Info().Msgf("%s %s", r.Method, r.URL)
		}

		next(w, r, body)
	}
}

// WriteError writes the error as the HTTP response body with status code 500.
func WriteError(w http.ResponseWriter, err extension_kit.ExtensionError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
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
