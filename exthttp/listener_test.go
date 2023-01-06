// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package exthttp

import (
	"github.com/stretchr/testify/require"
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
