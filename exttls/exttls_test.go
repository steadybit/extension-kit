// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

package exttls

import (
	"encoding/pem"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newSelfSignedServer(t *testing.T) *httptest.Server {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)
	return server
}

func get(url string) error {
	resp, err := (&http.Client{Transport: Transport()}).Get(url)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

func TestClientConfig_verifies_certificates_by_default(t *testing.T) {
	t.Setenv("STEADYBIT_EXTENSION_INSECURE_SKIP_VERIFY", "")
	server := newSelfSignedServer(t)

	assert.False(t, ClientConfig().InsecureSkipVerify)
	assert.Nil(t, ClientConfig().RootCAs, "RootCAs must stay nil so SSL_CERT_DIR/SSL_CERT_FILE are honoured")
	assert.ErrorContains(t, get(server.URL), "certificate")
}

func TestClientConfig_skips_verification_when_enabled(t *testing.T) {
	t.Setenv("STEADYBIT_EXTENSION_INSECURE_SKIP_VERIFY", "true")
	server := newSelfSignedServer(t)

	assert.True(t, ClientConfig().InsecureSkipVerify)
	assert.NoError(t, get(server.URL))
}

func TestTransport_does_not_modify_default_transport(t *testing.T) {
	t.Setenv("STEADYBIT_EXTENSION_INSECURE_SKIP_VERIFY", "true")

	_ = Transport()

	defaultTLS := http.DefaultTransport.(*http.Transport).TLSClientConfig
	assert.True(t, defaultTLS == nil || !defaultTLS.InsecureSkipVerify)
}

// Go loads the system certificate pool once per process, so trusting a CA through SSL_CERT_FILE is
// verified in a child process that starts with the variable already set.
func TestClientConfig_trusts_ca_from_ssl_cert_file(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("SSL_CERT_FILE is only honoured by Go's certificate loader on Linux, which is what extensions run on")
	}
	server := newSelfSignedServer(t)
	caFile := filepath.Join(t.TempDir(), "ca.crt")
	require.NoError(t, os.WriteFile(caFile, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}), 0o600))

	cmd := exec.Command(os.Args[0], "-test.run=^TestHelperGet$")
	cmd.Env = append(os.Environ(), "EXTTLS_HELPER_URL="+server.URL, "SSL_CERT_FILE="+caFile, "STEADYBIT_EXTENSION_INSECURE_SKIP_VERIFY=false")
	out, err := cmd.CombinedOutput()
	assert.NoError(t, err, string(out))
}

func TestHelperGet(t *testing.T) {
	url := os.Getenv("EXTTLS_HELPER_URL")
	if url == "" {
		t.Skip("only runs as a child process of TestClientConfig_trusts_ca_from_ssl_cert_file")
	}
	require.NoError(t, get(url))
}

func TestIsInsecureSkipVerify_parses_leniently(t *testing.T) {
	for value, expected := range map[string]bool{"": false, "false": false, "nonsense": false, "true": true, " TRUE ": true, "1": true} {
		t.Setenv("STEADYBIT_EXTENSION_INSECURE_SKIP_VERIFY", value)
		assert.Equal(t, expected, IsInsecureSkipVerify(), "value %q", value)
	}
}
