/*
 * Copyright 2024 steadybit GmbH. All rights reserved.
 */

// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package exthealth

import (
	"errors"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"github.com/steadybit/extension-kit/exthttp"
	"net/http"
	"sync/atomic"
)

var (
	isReady int32 = 1
	isAlive int32 = 1
	server  *http.Server
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

// addLivenessProbe registers an HTTP handler for the liveness probe. The liveness probe reports an error (HTTP 503) when the SetAlive function is called with false. Default readiness state is true.
func addLivenessProbe(registerFn func(string, http.Handler)) {
	registerFn("/health/liveness", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&isAlive) == 1 {
			w.WriteHeader(http.StatusOK)
			return
		}
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
	}))
}

// addReadinessProbe registers an HTTP handler for the readiness probe. The readiness probe reports an error (HTTP 503) when the SetReady function is called with false. Default readiness state is true.
func addReadinessProbe(registerFn func(string, http.Handler)) {
	registerFn("/health/readiness", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&isReady) == 1 {
			w.WriteHeader(http.StatusOK)
			return
		}
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
	}))
}

// StartProbes will start liveness and readiness probes.
func StartProbes(port int) {
	if exthttp.IsUnixSocketEnabled() {
		addLivenessProbe(http.Handle)
		addReadinessProbe(http.Handle)
		return
	}

	spec := HealthSpecification{}
	spec.parseConfigurationFromEnvironment()

	healthPort := port
	if spec.Port != 0 {
		healthPort = spec.Port
	}

	serverMux := http.NewServeMux()
	addLivenessProbe(serverMux.Handle)
	addReadinessProbe(serverMux.Handle)
	go func() {
		log.Info().Msgf("Starting probes server on port %d, ready: %t", healthPort, atomic.LoadInt32(&isReady) == 1)
		server = &http.Server{Addr: fmt.Sprintf(":%d", healthPort), Handler: serverMux}

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msgf("Failed to start probes server")
		}
	}()
}

func StopProbes() {
	if server != nil {
		_ = server.Close()
	}
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

// SetAlive sets the liveness state of the service. If the service is not alive the liveness probe will report an error and the container should restart.
func SetAlive(alive bool) {
	isAliveInt := atomic.LoadInt32(&isAlive)
	if isAliveInt == 1 && !alive || isAliveInt == 0 && alive {
		log.Info().Msgf("Update liveness probe - alive: %t", alive)
		if alive {
			atomic.StoreInt32(&isAlive, 1)
		} else {
			atomic.StoreInt32(&isAlive, 0)
		}
	}
}
