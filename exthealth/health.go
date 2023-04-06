// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package exthealth

import (
	"github.com/steadybit/extension-kit/exthttp"
	"net/http"
	"sync/atomic"
)

var (
	isReady int32 = 1
)

// AddLivenessProbe registers an HTTP handler for the liveness probe. The liveness probe reports HTTP 200 as soon as the HTTP server is up and running.
func AddLivenessProbe() {
	exthttp.RegisterHttpHandler("/health/liveness", func(w http.ResponseWriter, r *http.Request, body []byte) {
		w.WriteHeader(http.StatusOK)
	})
}

// AddReadinessProbe registers an HTTP handler for the readiness probe. The readiness probe reports an error (HTTP 503) when the SetReady function is called with false. Default readiness state is true.
func AddReadinessProbe() {
	exthttp.RegisterHttpHandler("/health/readiness", func(w http.ResponseWriter, r *http.Request, body []byte) {
		if atomic.LoadInt32(&isReady) == 1 {
			w.WriteHeader(http.StatusOK)
			return
		}
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
	})
}

// AddProbes registers HTTP handlers for the liveness and readiness probes.
func AddProbes() {
	AddLivenessProbe()
	AddReadinessProbe()
}

// SetReady sets the readiness state of the service. If the service is not ready the readiness probe will report an error.
func SetReady(ready bool) {
	if ready {
		atomic.StoreInt32(&isReady, 1)
	} else {
		atomic.StoreInt32(&isReady, 0)
	}
}
