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
	"github.com/stretchr/testify/require"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"testing"
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

	resp, err := http.Get(fmt.Sprintf("http://localhost:%d", port))

	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestStartHttpsServer(t *testing.T) {
	certs, err := testcerts.NewCA().NewKeyPair("localhost")
	require.NoError(t, err)

	cert, key, err := certs.ToTempFile(t.TempDir())
	require.NoError(t, err)

	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	server, start, err := prepareHttpsServer(port, ListenSpecification{
		TlsServerCert: cert.Name(),
		TlsServerKey:  key.Name(),
	})
	require.NoError(t, err)

	go start()
	defer server.Close()

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
	require.NoError(t, err)
}

func TestStartHttpsServerMustFailWhenCertificateCannotBeFound(t *testing.T) {
	_, key, err := testcerts.GenerateCertsToTempFile(t.TempDir())
	require.NoError(t, err)

	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	_, _, err = prepareHttpsServer(port, ListenSpecification{
		TlsServerCert: filepath.Join(t.TempDir(), "unknown.pem"),
		TlsServerKey:  key,
	})
	require.ErrorContains(t, err, "no such file or directory")
}

func TestStartHttpsServerMustFailWhenKeyCannotBeFound(t *testing.T) {
	_, key, err := testcerts.GenerateCertsToTempFile(t.TempDir())
	require.NoError(t, err)

	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	_, _, err = prepareHttpsServer(port, ListenSpecification{
		TlsServerCert: key,
		TlsServerKey:  filepath.Join(t.TempDir(), "unknown.pem"),
	})
	require.ErrorContains(t, err, "no such file or directory")
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

	server, start, err := prepareHttpsServer(port, ListenSpecification{
		TlsServerCert: serverCert.Name(),
		TlsServerKey:  serverKey.Name(),
		TlsClientCas:  []string{caCerts.Name()},
	})
	require.NoError(t, err)

	go start()
	defer server.Close()

	_, err = http.Get(fmt.Sprintf("https://localhost:%d", port))
	require.ErrorContains(t, err, "failed to verify certificate")
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

	server, start, err := prepareHttpsServer(port, ListenSpecification{
		TlsServerCert: serverCert.Name(),
		TlsServerKey:  serverKey.Name(),
		TlsClientCas:  []string{clientCertDir},
	})
	require.NoError(t, err)

	go start()
	defer server.Close()

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
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, r.StatusCode)
}

func TestStartHttpServerUsingUnixSocket(t *testing.T) {
	sock := filepath.Join(t.TempDir(), "sock")

	server, start, err := prepareUnixSocketServer(sock)
	require.NoError(t, err)

	go start()
	defer server.Close()

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", sock)
			},
		},
	}

	resp, err := client.Get("http://localhost")

	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}
