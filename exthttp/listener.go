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
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ListenSpecification struct {
	Port          int      `json:"port" split_words:"true" required:"false"`
	UnixSocket    string   `json:"unixSocket" split_words:"true" required:"false"`
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

func IsUnixSocketEnabled() bool {
	spec := ListenSpecification{}
	spec.parseConfigurationFromEnvironment()
	return spec.UnixSocket != ""
}

func listen(opts ListenOpts) (*http.Server, func() error, error) {
	spec := ListenSpecification{}
	spec.parseConfigurationFromEnvironment()
	err := spec.validateSpecification()
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to validate HTTP server configuration.")
	}

	if spec.UnixSocket != "" {
		log.Info().Msgf("Starting extension server on unix socket %s", spec.UnixSocket)
		return prepareUnixSocketServer(spec.UnixSocket)
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

func prepareUnixSocketServer(path string) (*http.Server, func() error, error) {
	server := &http.Server{
		ErrorLog: stdLog.New(&forwardToZeroLogWriter{}, "", 0),
	}

	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create directory for unix socket: %w", err)
		}
	} else {
		_ = os.Remove(path)
	}

	unixListener, err := net.Listen("unix", path)
	if err != nil {
		return nil, nil, fmt.Errorf("failed listen on unix socket: %w", err)
	}

	return server, func() error {
		return server.Serve(unixListener)
	}, nil
}

func prepareHttpsServer(port int, spec ListenSpecification) (*http.Server, func() error, error) {
	certReloader := NewCertReloader(spec.TlsServerCert, spec.TlsServerKey)

	if _, err := certReloader.GetCertificate(nil); err != nil {
		return nil, nil, fmt.Errorf("failed to load TLS certificate: %w", err)
	}

	tlsConfig := tls.Config{
		GetCertificate: certReloader.GetCertificate,
	}

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
		return server.ListenAndServeTLS("", "")
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
