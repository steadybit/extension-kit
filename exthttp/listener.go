/*
 * Copyright 2024 steadybit GmbH. All rights reserved.
 */

// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package exthttp

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	stdLog "log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type ListenSpecification struct {
	Port          int      `json:"port" split_words:"true" required:"false"`
	UnixSocket    string   `json:"unixSocket" split_words:"true" required:"false"`
	TlsServerCert string   `json:"tlsServerCert" split_words:"true" required:"false"`
	TlsServerKey  string   `json:"tlsServerKey" split_words:"true" required:"false"`
	TlsClientCas  []string `json:"tlsClientCas" split_words:"true" required:"false"`
}

var (
	wrapper   *httpServerWrapper
	serveCond = sync.NewCond(&sync.Mutex{})
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

func listen(opts ListenOpts) error {
	success := false
	serveCond.L.Lock()
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

	port := opts.Port
	if spec.Port != 0 {
		port = spec.Port
	}

	var err error
	if spec.UnixSocket != "" {
		wrapper, err = prepareUnixSocketServer(spec.UnixSocket)
	} else if spec.isTlsEnabled() {
		wrapper, err = prepareHttpsServer(port, spec)
	} else {
		wrapper, err = prepareHttpServer(port)
	}
	if err != nil {
		return err
	}

	serveCond.Broadcast()
	serveCond.L.Unlock()
	success = true
	if err = wrapper.serve(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func WaitForServe() {
	serveCond.L.Lock()
	defer serveCond.L.Unlock()
	serveCond.Wait()
}

func StopListen() {
	if wrapper == nil || wrapper.server == nil {
		return
	}
	if err := wrapper.server.Close(); err != nil {
		log.Error().Err(err).Msgf("Failed to stop extension server")
	}
	wrapper = nil
}

type forwardToZeroLogWriter struct {
}

func (fw *forwardToZeroLogWriter) Write(p []byte) (n int, err error) {
	trimmed := strings.Trim(string(p), " \t\n\r")
	if strings.Contains(trimmed, "TLS handshake error") &&
		strings.Contains(trimmed, "unknown certificate") || strings.Contains(trimmed, "client didn't provide a certificate") {
		// Ignore/only log on debug TLS handshake errors when client did not provide a certificate
		log.Debug().Msg(trimmed)
	} else {
		log.Error().Msg(trimmed)
	}
	return len([]byte(trimmed)), nil
}

func prepareHttpServer(port int) (*httpServerWrapper, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		ErrorLog: stdLog.New(&forwardToZeroLogWriter{}, "", 0),
	}

	log.Info().Msgf("Starting extension http server on port %d", port)
	return &httpServerWrapper{
		serve: func() error {
			return server.Serve(listener)
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

func prepareHttpsServer(port int, spec ListenSpecification) (*httpServerWrapper, error) {
	certReloader := NewCertReloader(spec.TlsServerCert, spec.TlsServerKey)

	if _, err := certReloader.GetCertificate(nil); err != nil {
		return nil, fmt.Errorf("failed to load TLS certificate: %w", err)
	}

	clientCAs, err := loadCertPool(spec.TlsClientCas)
	if err != nil {
		return nil, fmt.Errorf("failed to load TLS client CA certificates: %w", err)
	}

	tlsConfig := tls.Config{
		GetCertificate: certReloader.GetCertificate,
		ClientAuth:     spec.getClientAuthType(),
		ClientCAs:      clientCAs,
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		TLSConfig: &tlsConfig,
		ErrorLog:  stdLog.New(&forwardToZeroLogWriter{}, "", 0),
	}
	return &httpServerWrapper{
		serve: func() error {
			log.Info().Msgf("Starting extension https server on port %d (ClientAuth: %s)", port, spec.getClientAuthType())
			return server.ServeTLS(listener, "", "")
		},
		server: server,
	}, nil
}

func loadCertPool(filePaths []string) (*x509.CertPool, error) {
	pool := x509.NewCertPool()
	for _, filePath := range filePaths {
		_ = filepath.Walk(filePath, func(path string, info os.FileInfo, _ error) error {
			if info.IsDir() {
				return nil
			}
			caCert, err := os.ReadFile(path)
			if err == nil {
				log.Debug().Msgf("Loading CA certificate from %s", path)
				pool.AppendCertsFromPEM(caCert)
			} else {
				log.Error().Err(err).Msgf("Failed to read CA certificate from %s", path)
			}
			return nil
		})
	}
	return pool, nil
}
