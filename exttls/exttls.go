// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

// Package exttls gives extensions one consistent way to configure the TLS settings of the
// clients they use to reach third-party services, which are often self-hosted behind
// self-signed certificates.
//
// Two mechanisms are supported, the same in every extension:
//
//   - STEADYBIT_EXTENSION_INSECURE_SKIP_VERIFY=true disables certificate verification.
//   - Custom CA certificates are trusted by mounting them into the container and pointing
//     SSL_CERT_DIR (or SSL_CERT_FILE) at them. Go reads these variables when it loads the
//     system certificate pool, so no extension code is involved beyond not replacing that
//     pool. The tls.Config returned here leaves RootCAs nil, which means the system pool.
package exttls

import (
	"crypto/tls"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

const envInsecureSkipVerify = "STEADYBIT_EXTENSION_INSECURE_SKIP_VERIFY"

var warnInsecureOnce sync.Once

// IsInsecureSkipVerify reports whether STEADYBIT_EXTENSION_INSECURE_SKIP_VERIFY is enabled. An empty or
// invalid value keeps verification on rather than failing startup, as charts may render the variable empty.
func IsInsecureSkipVerify() bool {
	value := strings.TrimSpace(os.Getenv(envInsecureSkipVerify))
	if value == "" {
		return false
	}
	insecure, err := strconv.ParseBool(value)
	if err != nil {
		log.Warn().Str("value", value).Msgf("Ignoring invalid value of %s, TLS certificates will be verified.", envInsecureSkipVerify)
		return false
	}
	return insecure
}

// ClientConfig returns the tls.Config to use for outgoing connections to third-party services.
// Callers may modify the returned value, e.g. to add client certificates for mutual TLS.
func ClientConfig() *tls.Config {
	insecure := IsInsecureSkipVerify()
	if insecure {
		warnInsecureOnce.Do(func() {
			log.Warn().Msg("TLS certificate verification is disabled (STEADYBIT_EXTENSION_INSECURE_SKIP_VERIFY=true). Prefer trusting a custom CA via SSL_CERT_DIR.")
		})
	}
	return &tls.Config{InsecureSkipVerify: insecure} //NOSONAR explicit user choice
}

// Transport returns a clone of http.DefaultTransport using ClientConfig. Use it as the transport of
// an http.Client, or pass it to resty via SetTransport. If http.DefaultTransport was replaced by a
// wrapper (e.g. otelhttp), a transport with the same defaults is created instead.
func Transport() *http.Transport {
	var t *http.Transport
	if defaultTransport, ok := http.DefaultTransport.(*http.Transport); ok {
		t = defaultTransport.Clone()
	} else {
		t = &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
	}
	t.TLSClientConfig = ClientConfig()
	return t
}
