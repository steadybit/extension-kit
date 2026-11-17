// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2025 Steadybit GmbH

/*
 * Copyright 2024 steadybit GmbH. All rights reserved.
 */

package exthttp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	stdLog "log"
	"net"
	"net/http"
	_ "net/http/pprof" // NOSONAR go:S4507 (pprof handlers are disabled by default; see hidePprofHandlers)
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	"github.com/steadybit/extension-kit/extsignals"
)

type ListenSpecification struct {
	Port          int      `json:"port" split_words:"true" required:"false"`
	BindAddress   string   `json:"bindAddress" split_words:"true" required:"false"`
	UnixSocket    string   `json:"unixSocket" split_words:"true" required:"false"`
	TlsServerCert string   `json:"tlsServerCert" split_words:"true" required:"false"`
	TlsServerKey  string   `json:"tlsServerKey" split_words:"true" required:"false"`
	TlsClientCas  []string `json:"tlsClientCas" split_words:"true" required:"false"`
	EnablePprof   bool     `json:"enablePprof" split_words:"true" required:"false"`
}

var (
	wrapper   *httpServerWrapper
	serveCond = sync.NewCond(&sync.Mutex{})
	// serving records whether the server has reached the point of accepting
	// connections. WaitForServe needs it as a predicate: listen() holds
	// serveCond.L from its first statement, so a waiter that has not acquired
	// the lock yet is not waiting on the condition and misses the broadcast.
	serving bool
)

func (spec *ListenSpecification) parseConfigurationFromEnvironment() {
	err := envconfig.Process("steadybit_extension", spec)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to parse HTTP server configuration from environment.")
	}
}

func (spec *ListenSpecification) isTlsEnabled() bool {
	return spec.TlsServerCert != "" || spec.TlsServerKey != "" || len(spec.TlsClientCas) > 0
}

func (spec *ListenSpecification) getClientAuthType() tls.ClientAuthType {
	if len(spec.TlsClientCas) > 0 {
		return tls.RequireAndVerifyClientCert
	} else {
		return tls.NoClientCert
	}
}

func (spec *ListenSpecification) validateSpecification() error {
	tlsEnabled := spec.isTlsEnabled()

	if tlsEnabled && spec.TlsServerCert == "" {
		return fmt.Errorf("TLS server certificate must be provided when TLS is enabled")
	}
	if tlsEnabled && spec.TlsServerKey == "" {
		return fmt.Errorf("TLS server key must be provided when TLS is enabled")
	}
	return nil
}

type ListenOpts struct {
	// Port Default port to bind to. Can be overridden through the environment variable STEADYBIT_EXTENSION_PORT.
	Port int
}
type httpServerWrapper struct {
	serve  func() error
	server *http.Server
}

func Listen(opts ListenOpts) {
	err := listen(opts)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Err(err).Msgf("Failed to start extension server")
	}
}

func IsUnixSocketEnabled() bool {
	spec := ListenSpecification{}
	spec.parseConfigurationFromEnvironment()
	return spec.UnixSocket != ""
}

func hidePprofHandlers(spec ListenSpecification) {
	if spec.EnablePprof {
		log.Info().Msg("pprof handlers enabled")
		return
	}

	log.Debug().Msg("disabling pprof handlers")

	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNotFound)
	})
	mux.HandleFunc("/debug/pprof/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusNotFound)
	})
	mux.Handle("/", http.DefaultServeMux)
	http.DefaultServeMux = mux
}

func listen(opts ListenOpts) error {

	success := false
	serveCond.L.Lock()
	serving = false
	defer func() {
		if !success {
			serveCond.L.Unlock()
		}
	}()

	spec := ListenSpecification{}
	spec.parseConfigurationFromEnvironment()
	if err := spec.validateSpecification(); err != nil {
		return fmt.Errorf("failed to validate listen specification: %w", err)

	}
	hidePprofHandlers(spec)

	port := opts.Port
	if spec.Port != 0 {
		port = spec.Port
	}

	bindAddress := spec.bindAddress()

	// Serve through the local variable: wrapper is only accessed under
	// serveCond.L, and StopListen may reset it before serve is called.
	var w *httpServerWrapper
	var err error
	if spec.UnixSocket != "" {
		w, err = prepareUnixSocketServer(spec.UnixSocket)
	} else if spec.isTlsEnabled() {
		w, err = prepareHttpsServer(bindAddress, port, spec)
	} else {
		w, err = prepareHttpServer(bindAddress, port)
	}
	if err != nil {
		return err
	}
	wrapper = w

	extsignals.AddSignalHandler(extsignals.SignalHandler{
		Handler: func(signal os.Signal) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer func() {
				cancel()
			}()
			w := takeWrapper()
			if w == nil || w.server == nil {
				return
			}
			log.Info().Msg("Stopping Extension HTTP Server")
			if err := w.server.Shutdown(ctx); err != nil {
				log.Warn().Msgf("Extension HTTP Server Shutdown Failed: %+v", err)
			}
		},
		Order: extsignals.OrderStopExtensionHttp,
		Name:  "StopExtensionHTTP",
	})

	serving = true
	serveCond.Broadcast()
	serveCond.L.Unlock()
	success = true
	if err = w.serve(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func WaitForServe() {
	serveCond.L.Lock()
	defer serveCond.L.Unlock()
	for !serving {
		serveCond.Wait()
	}
}

func StopListen() {
	serveCond.L.Lock()
	serving = false
	serveCond.L.Unlock()

	w := takeWrapper()
	if w == nil || w.server == nil {
		return
	}
	if err := w.server.Close(); err != nil {
		log.Error().Err(err).Msgf("Failed to stop extension server")
	}
}

// takeWrapper returns the current server wrapper and resets it, so that only one caller stops the server.
func takeWrapper() *httpServerWrapper {
	serveCond.L.Lock()
	defer serveCond.L.Unlock()
	w := wrapper
	wrapper = nil
	return w
}

type forwardToZeroLogWriter struct {
}

func (fw *forwardToZeroLogWriter) Write(p []byte) (n int, err error) {
	trimmed := strings.Trim(string(p), " \t\n\r")
	if strings.Contains(trimmed, "TLS handshake error") &&
		strings.Contains(trimmed, "unknown certificate") ||
		strings.Contains(trimmed, "client didn't provide a certificate") ||
		strings.Contains(trimmed, "EOF") {
		log.Debug().Msg(trimmed)
	} else {
		log.Error().Msg(trimmed)
	}
	return len([]byte(trimmed)), nil
}

func prepareHttpServer(bindAddress string, port int) (*httpServerWrapper, error) {
	listeners, err := ListenTCP(bindAddress, port)
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		ErrorLog: stdLog.New(&forwardToZeroLogWriter{}, "", 0),
	}

	log.Info().Msgf("Starting extension http server on %s", ListenerAddresses(listeners))
	return &httpServerWrapper{
		serve: func() error {
			return ServeListeners(listeners, server.Serve)
		},
		server: server,
	}, nil
}

func prepareUnixSocketServer(path string) (*httpServerWrapper, error) {
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			return nil, fmt.Errorf("failed to create directory for unix socket: %w", err)
		}
	} else {
		_ = os.Remove(path)
	}

	unixListener, err := net.Listen("unix", path)
	if err != nil {
		return nil, fmt.Errorf("failed listen on unix socket: %w", err)
	}

	server := &http.Server{
		ErrorLog: stdLog.New(&forwardToZeroLogWriter{}, "", 0),
	}

	return &httpServerWrapper{
		serve: func() error {
			log.Info().Msgf("Starting extension http server on unix domain socket (%s)", path)
			return server.Serve(unixListener)
		},
		server: server,
	}, nil
}

func prepareHttpsServer(bindAddress string, port int, spec ListenSpecification) (*httpServerWrapper, error) {
	certReloader := NewCertReloader(spec.TlsServerCert, spec.TlsServerKey)

	if _, err := certReloader.GetCertificate(nil); err != nil {
		return nil, fmt.Errorf("failed to load TLS certificate: %w", err)
	}

	clientCAs, err := loadCertPool(spec.TlsClientCas)
	if err != nil {
		log.Warn().Err(err).Msg("failed to load TLS client CA certificates")
	}

	tlsConfig := tls.Config{
		GetCertificate: certReloader.GetCertificate,
		ClientAuth:     spec.getClientAuthType(),
		ClientCAs:      clientCAs,
	}

	listeners, err := ListenTCP(bindAddress, port)
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		TLSConfig: &tlsConfig,
		ErrorLog:  stdLog.New(&forwardToZeroLogWriter{}, "", 0),
	}
	return &httpServerWrapper{
		serve: func() error {
			log.Info().Msgf("Starting extension https server on %s (ClientAuth: %s)", ListenerAddresses(listeners), spec.getClientAuthType())
			return ServeListeners(listeners, func(listener net.Listener) error {
				return server.ServeTLS(listener, "", "")
			})
		},
		server: server,
	}, nil
}

func loadCertPool(filePaths []string) (*x509.CertPool, error) {
	pool := x509.NewCertPool()

	var err error
	for _, filePath := range filePaths {
		walkErr := filepath.WalkDir(filePath, func(path string, entry os.DirEntry, err error) error {
			if err != nil {
				return err
			} else if entry.IsDir() {
				return nil
			}
			if caCert, err := os.ReadFile(path); err == nil {
				log.Debug().Msgf("loading CA certificate from %s", path)
				pool.AppendCertsFromPEM(caCert)
			} else {
				log.Error().Err(err).Msgf("failed to read CA certificate from %s", path)
			}
			return nil
		})
		err = errors.Join(err, walkErr)
	}
	return pool, err
}
