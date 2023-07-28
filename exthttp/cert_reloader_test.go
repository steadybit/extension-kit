package exthttp

import (
	"github.com/madflojo/testcerts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ReloadingCertificateProvider(t *testing.T) {
	cert, key, err := testcerts.GenerateCertsToTempFile(t.TempDir())
	require.NoError(t, err)

	reloader := NewCertReloader(cert, key)

	certificate, err := reloader.GetCertificate(nil)
	require.NoError(t, err)
	cached, err := reloader.GetCertificate(nil)
	require.NoError(t, err)
	assert.Same(t, certificate, cached, "certificate should be cached")

	// Regenerate certificate
	err = testcerts.GenerateCertsToFile(cert, key)
	require.NoError(t, err)
	reloaded, err := reloader.GetCertificate(nil)
	require.NoError(t, err)
	assert.NotEqualf(t, certificate, reloaded, "certificate should have been reloaded")
}
