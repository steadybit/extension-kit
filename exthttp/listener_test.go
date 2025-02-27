/*
 * Copyright 2024 steadybit GmbH. All rights reserved.
 */

// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package exthttp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/madflojo/testcerts"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestValidateSpecificationSuccessHttp(t *testing.T) {
	spec := ListenSpecification{}
	err := spec.validateSpecification()
	assert.NoError(t, err)
}

func TestValidateSpecificationSuccessTls(t *testing.T) {
	spec := ListenSpecification{
		TlsServerCert: "cert",
		TlsServerKey:  "key",
	}
	err := spec.validateSpecification()
	assert.NoError(t, err)
}

func TestValidateSpecificationSuccessMTls(t *testing.T) {
	spec := ListenSpecification{
		TlsServerCert: "cert",
		TlsServerKey:  "key",
		TlsClientCas:  []string{"ca"},
	}
	err := spec.validateSpecification()
	assert.NoError(t, err)
}

func TestValidateSpecificationMissingCert(t *testing.T) {
	spec := ListenSpecification{
		TlsClientCas: []string{"ca"},
	}
	err := spec.validateSpecification()
	assert.ErrorContains(t, err, "certificate")
}

func TestValidateSpecificationMissingKey(t *testing.T) {
	spec := ListenSpecification{
		TlsServerCert: "cert",
		TlsClientCas:  []string{"ca"},
	}
	err := spec.validateSpecification()
	assert.ErrorContains(t, err, "key")
}

func TestStartHttpServer(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	old := http.DefaultServeMux
	defer func() { http.DefaultServeMux = old }()
	go Listen(ListenOpts{Port: port})
	WaitForServe()
	defer StopListen()

	resp, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestStartHttpsServer(t *testing.T) {
	certs, err := testcerts.NewCA().NewKeyPair("localhost")
	require.NoError(t, err)

	cert, key, err := certs.ToTempFile(t.TempDir())
	require.NoError(t, err)

	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	old := http.DefaultServeMux
	defer func() { http.DefaultServeMux = old }()
	t.Setenv("STEADYBIT_EXTENSION_TLS_SERVER_KEY", key.Name())
	t.Setenv("STEADYBIT_EXTENSION_TLS_SERVER_CERT", cert.Name())
	go Listen(ListenOpts{Port: port})
	WaitForServe()
	defer StopListen()

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(certs.PublicKey())

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: pool,
			},
		},
	}
	_, err = client.Get(fmt.Sprintf("https://localhost:%d", port))
	assert.NoError(t, err)
}

func TestStartHttpsServerMustFailWhenCertificateCannotBeFound(t *testing.T) {
	_, key, err := testcerts.GenerateCertsToTempFile(t.TempDir())
	require.NoError(t, err)

	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	old := http.DefaultServeMux
	defer func() { http.DefaultServeMux = old }()
	t.Setenv("STEADYBIT_EXTENSION_TLS_SERVER_KEY", key)
	t.Setenv("STEADYBIT_EXTENSION_TLS_SERVER_CERT", filepath.Join(t.TempDir(), "unknown.pem"))

	err = listen(ListenOpts{Port: port})
	assert.ErrorContains(t, err, "no such file or directory")
}

func TestStartHttpsServerMustFailWhenKeyCannotBeFound(t *testing.T) {
	cert, _, err := testcerts.GenerateCertsToTempFile(t.TempDir())
	require.NoError(t, err)

	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	old := http.DefaultServeMux
	defer func() { http.DefaultServeMux = old }()
	t.Setenv("STEADYBIT_EXTENSION_TLS_SERVER_KEY", filepath.Join(t.TempDir(), "unknown.pem"))
	t.Setenv("STEADYBIT_EXTENSION_TLS_SERVER_CERT", cert)
	err = listen(ListenOpts{Port: port})

	assert.ErrorContains(t, err, "no such file or directory")
}

func TestStartHttpsServerWithMutualTlsMustRefuseConnectionsWithoutMutualTls(t *testing.T) {
	ca := testcerts.NewCA()
	caCerts, _, err := ca.ToTempFile(t.TempDir())
	require.NoError(t, err)

	serverPair, err := ca.NewKeyPair("localhost")
	require.NoError(t, err)
	serverCert, serverKey, err := serverPair.ToTempFile(t.TempDir())
	require.NoError(t, err)

	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	old := http.DefaultServeMux
	defer func() { http.DefaultServeMux = old }()
	t.Setenv("STEADYBIT_EXTENSION_TLS_SERVER_KEY", serverKey.Name())
	t.Setenv("STEADYBIT_EXTENSION_TLS_SERVER_CERT", serverCert.Name())
	t.Setenv("STEADYBIT_EXTENSION_TLS_CLIENT_CAS", caCerts.Name())
	go Listen(ListenOpts{Port: port})
	WaitForServe()
	defer StopListen()

	_, err = http.Get(fmt.Sprintf("https://localhost:%d", port))
	assert.ErrorContains(t, err, "failed to verify certificate")
}

func TestStartHttpsServerEnforcingMutualTls(t *testing.T) {
	ca := testcerts.NewCA()

	clientPair, err := ca.NewKeyPair()
	require.NoError(t, err)
	clientCertDir := t.TempDir()
	err = os.WriteFile(filepath.Join(clientCertDir, "client.crt"), clientPair.PublicKey(), 0644)
	require.NoError(t, err)

	serverPair, err := ca.NewKeyPair("localhost")
	require.NoError(t, err)
	serverCert, serverKey, err := serverPair.ToTempFile(t.TempDir())
	require.NoError(t, err)

	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	old := http.DefaultServeMux
	defer func() { http.DefaultServeMux = old }()
	t.Setenv("STEADYBIT_EXTENSION_TLS_SERVER_KEY", serverKey.Name())
	t.Setenv("STEADYBIT_EXTENSION_TLS_SERVER_CERT", serverCert.Name())
	t.Setenv("STEADYBIT_EXTENSION_TLS_CLIENT_CAS", clientCertDir)
	go Listen(ListenOpts{Port: port})
	WaitForServe()
	defer StopListen()

	clientCertificate, err := tls.X509KeyPair(clientPair.PublicKey(), clientPair.PrivateKey())
	require.NoError(t, err)

	clientPool := x509.NewCertPool()
	clientPool.AppendCertsFromPEM(serverPair.PublicKey())

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      clientPool,
				Certificates: []tls.Certificate{clientCertificate},
			},
		},
	}

	r, err := client.Get(fmt.Sprintf("https://localhost:%d", port))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, r.StatusCode)
}

func TestStartHttpServerUsingUnixSocket(t *testing.T) {
	sock := filepath.Join(t.TempDir(), "sock")
	old := http.DefaultServeMux
	defer func() { http.DefaultServeMux = old }()

	t.Setenv("STEADYBIT_EXTENSION_UNIX_SOCKET", sock)
	go Listen(ListenOpts{})
	WaitForServe()
	defer StopListen()

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", sock)
			},
		},
	}

	resp, err := client.Get("http://localhost")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func Test_hidePprofHandlers(t *testing.T) {
	tests := []struct {
		name         string
		enabled      bool
		wantedStatus int
	}{
		{
			name:         "pprof handlers are hidden",
			enabled:      false,
			wantedStatus: http.StatusNotFound,
		},
		{
			name:         "pprof handlers are not hidden",
			enabled:      true,
			wantedStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			old := http.DefaultServeMux
			defer func() { http.DefaultServeMux = old }()

			r, _ := http.NewRequest("GET", "/debug/pprof/", nil)
			w := httptest.NewRecorder()

			hidePprofHandlers(ListenSpecification{EnablePprof: tt.enabled})

			http.DefaultServeMux.ServeHTTP(w, r)

			assert.Equal(t, tt.wantedStatus, w.Result().StatusCode)
		})
	}
}
