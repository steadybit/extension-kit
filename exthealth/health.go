// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package exthealth

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"net/http"
	"sync/atomic"
)

var (
	isReady int32 = 1
)

type HealthSpecification struct {
	Port int `json:"port" split_words:"true" required:"false"`
}

func (spec *HealthSpecification) parseConfigurationFromEnvironment() {
	err := envconfig.Process("steadybit_extension_health", spec)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to parse health HTTP server configuration from environment.")
	}
}

// AddLivenessProbe registers an HTTP handler for the liveness probe. The liveness probe reports HTTP 200 as soon as the HTTP server is up and running.
func addLivenessProbe(serverMux *http.ServeMux) {
	serverMux.Handle("/health/liveness", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
}

// AddReadinessProbe registers an HTTP handler for the readiness probe. The readiness probe reports an error (HTTP 503) when the SetReady function is called with false. Default readiness state is true.
func addReadinessProbe(serverMux *http.ServeMux) {
	serverMux.Handle("/health/readiness", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&isReady) == 1 {
			w.WriteHeader(http.StatusOK)
			return
		}
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
	}))
}

// StartProbes will start liveness and readiness probes.
func StartProbes(port int) {
	spec := HealthSpecification{}
	spec.parseConfigurationFromEnvironment()

	healthPort := port
	if spec.Port != 0 {
		healthPort = spec.Port
	}

	serverMuxProbes := http.NewServeMux()
	addLivenessProbe(serverMuxProbes)
	addReadinessProbe(serverMuxProbes)
	go func() {
		log.Info().Msgf("Starting probes server on port %d, ready: %t", healthPort, atomic.LoadInt32(&isReady) == 1)
		err := http.ListenAndServe(fmt.Sprintf(":%d", healthPort), serverMuxProbes)
		if err != nil {
			log.Fatal().Err(err).Msgf("Failed to start probes server")
		}
	}()
}

// SetReady sets the readiness state of the service. If the service is not ready the readiness probe will report an error.
func SetReady(ready bool) {
	log.Info().Msgf("Update readiness probe - ready: %t", ready)
	if ready {
		atomic.StoreInt32(&isReady, 1)
	} else {
		atomic.StoreInt32(&isReady, 0)
	}
}
