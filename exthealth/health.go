// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package exthealth

import (
	"github.com/steadybit/extension-kit/exthttp"
	"net/http"
)

var (
	isReady = false
)

// AddLivenessProbe registers a HTTP handler for the liveness probe. The liveness probe reports HTTP 200 as soon as the HTTP server is up and running.
func AddLivenessProbe() {
	exthttp.RegisterHttpHandler("/health/liveness", func(w http.ResponseWriter, r *http.Request, body []byte) {
		w.WriteHeader(http.StatusOK)
	})
}

// AddReadinessProbe registers a HTTP handler for the readiness probe. The readiness probe reports an error (HTTP 503) until the SetReady function is called with true.
func AddReadinessProbe() {
	exthttp.RegisterHttpHandler("/health/readiness", func(w http.ResponseWriter, r *http.Request, body []byte) {
		if isReady {
			w.WriteHeader(http.StatusOK)
			return
		}
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
	})
}

// SetReady sets the readiness state of the service. If the service is not ready the readiness probe will report an error.
func SetReady(ready bool) {
	isReady = ready
}
