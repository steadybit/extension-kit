// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package exthttp

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/require"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestValidateSpecificationSuccessHttp(t *testing.T) {
	spec := ListenSpecification{}
	err := spec.validateSpecification()
	require.NoError(t, err)
}

func TestValidateSpecificationSuccessTls(t *testing.T) {
	spec := ListenSpecification{
		TlsServerCert: "cert",
		TlsServerKey:  "key",
	}
	err := spec.validateSpecification()
	require.NoError(t, err)
}

func TestValidateSpecificationSuccessMTls(t *testing.T) {
	spec := ListenSpecification{
		TlsServerCert: "cert",
		TlsServerKey:  "key",
		TlsClientCas:  []string{"ca"},
	}
	err := spec.validateSpecification()
	require.NoError(t, err)
}

func TestValidateSpecificationMissingCert(t *testing.T) {
	spec := ListenSpecification{
		TlsClientCas: []string{"ca"},
	}
	err := spec.validateSpecification()
	require.ErrorContains(t, err, "certificate")
}

func TestValidateSpecificationMissingKey(t *testing.T) {
	spec := ListenSpecification{
		TlsServerCert: "cert",
		TlsClientCas:  []string{"ca"},
	}
	err := spec.validateSpecification()
	require.ErrorContains(t, err, "key")
}

func TestStartHttpServer(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	server, start, err := prepareHttpServer(port)
	require.NoError(t, err)

	go start()
	defer server.Close()
	time.Sleep(1 * time.Second)

	resp, err := http.Get(fmt.Sprintf("http://localhost:%d", port))

	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestStartHttpsServer(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	server, start, err := prepareHttpsServer(port, ListenSpecification{
		TlsServerCert: "testdata/cert.pem",
		TlsServerKey:  "testdata/key.pem",
	})
	require.NoError(t, err)

	go start()
	defer server.Close()
	time.Sleep(1 * time.Second)

	_, err = http.Get(fmt.Sprintf("https://localhost:%d", port))
	require.ErrorContains(t, err, "certificate")
}

func TestStartHttpsServerMustFailWhenCertificateCannotBeFound(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	_, start, err := prepareHttpsServer(port, ListenSpecification{
		TlsServerCert: "testdata/unknown.pem",
		TlsServerKey:  "testdata/key.pem",
	})
	require.NoError(t, err)

	err = start()
	require.ErrorContains(t, err, "no such file or directory")
}

func TestStartHttpsServerMustFailWhenKeyCannotBeFound(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	_, start, err := prepareHttpsServer(port, ListenSpecification{
		TlsServerCert: "testdata/cert.pem",
		TlsServerKey:  "testdata/unknown.pem",
	})
	require.NoError(t, err)

	err = start()
	require.ErrorContains(t, err, "no such file or directory")
}

func TestStartHttpsServerWithMutualTlsMustRefuseConnectionsWithoutMutualTls(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	server, start, err := prepareHttpsServer(port, ListenSpecification{
		TlsServerCert: "testdata/cert.pem",
		TlsServerKey:  "testdata/unknown.pem",
		TlsClientCas:  []string{"testdata/cert.pem"},
	})
	require.NoError(t, err)

	go start()
	defer server.Close()
	time.Sleep(1 * time.Second)

	_, err = http.Get(fmt.Sprintf("https://localhost:%d", port))

	require.ErrorContains(t, err, "connect: connection refused")
}

func TestStartHttpsServerWithMutualTlsMustSuccessfullyAllowMutualTlsConnections(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	server, start, err := prepareHttpsServer(port, ListenSpecification{
		TlsServerCert: "testdata/cert.pem",
		TlsServerKey:  "testdata/key.pem",
		TlsClientCas:  []string{"testdata/cert.pem"},
	})
	require.NoError(t, err)

	go start()
	defer server.Close()
	time.Sleep(1 * time.Second)

	cert, err := tls.LoadX509KeyPair("testdata/cert.pem", "testdata/key.pem")
	require.NoError(t, err)

	caCert, err := os.ReadFile("testdata/cert.pem")
	require.NoError(t, err)
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	r, err := client.Get(fmt.Sprintf("https://localhost:%d", port))
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, r.StatusCode)
}
