// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package exthttp

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
	stdLog "log"
	"net/http"
	"os"
	"strings"
)

type ListenSpecification struct {
	Port int `json:"port" split_words:"true" required:"false"`

	TlsServerCert string   `json:"tlsServerCert" split_words:"true" required:"false"`
	TlsServerKey  string   `json:"tlsServerKey" split_words:"true" required:"false"`
	TlsClientCas  []string `json:"tlsClientCas" split_words:"true" required:"false"`
}

func (spec *ListenSpecification) parseConfigurationFromEnvironment() {
	err := envconfig.Process("steadybit_extension", spec)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to parse HTTP server configuration from environment.")
	}
}

func (spec *ListenSpecification) isTlsEnabled() bool {
	return spec.TlsServerCert != "" || spec.TlsServerKey != "" || len(spec.TlsClientCas) > 0
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

func Listen(opts ListenOpts) {
	_, start, err := listen(opts)

	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to start extension server")
	}

	err = start()
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to start extension server")
	}
}

func listen(opts ListenOpts) (*http.Server, func() error, error) {
	spec := ListenSpecification{}
	spec.parseConfigurationFromEnvironment()
	err := spec.validateSpecification()
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to validate HTTP server configuration.")
	}

	port := opts.Port
	if spec.Port != 0 {
		port = spec.Port
	}

	log.Info().Msgf("Starting extension server on port %d (TLS: %v)", port, spec.isTlsEnabled())
	if spec.isTlsEnabled() {
		return prepareHttpsServer(port, spec)
	} else {
		return prepareHttpServer(port)
	}
}

type forwardToZeroLogWriter struct {
}

func (fw *forwardToZeroLogWriter) Write(p []byte) (n int, err error) {
	trimmed := strings.Trim(string(p), " \t\n\r")
	log.Error().Msg(trimmed)
	return len([]byte(trimmed)), nil
}

func prepareHttpServer(port int) (*http.Server, func() error, error) {
	server := &http.Server{
		Addr:     fmt.Sprintf(":%d", port),
		ErrorLog: stdLog.New(&forwardToZeroLogWriter{}, "", 0),
	}

	return server, server.ListenAndServe, nil
}

func prepareHttpsServer(port int, spec ListenSpecification) (*http.Server, func() error, error) {
	tlsConfig := tls.Config{}

	if len(spec.TlsClientCas) > 0 {
		clientCAs, err := loadCertPool(spec.TlsClientCas)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to load TLS client CA certificates: %w", err)
		}

		tlsConfig.ClientCAs = clientCAs
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	}

	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		TLSConfig: &tlsConfig,
		ErrorLog:  stdLog.New(&forwardToZeroLogWriter{}, "", 0),
	}
	return server, func() error {
		return server.ListenAndServeTLS(spec.TlsServerCert, spec.TlsServerKey)
	}, nil
}

func loadCertPool(filePaths []string) (*x509.CertPool, error) {
	pool := x509.NewCertPool()
	for _, filePath := range filePaths {
		caCert, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		pool.AppendCertsFromPEM(caCert)
	}
	return pool, nil
}
