package exthealth

import (
	"fmt"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestShouldServeReadinessAndLiveness(t *testing.T) {
	SetReady(false)
	addLivenessProbe(http.Handle)
	addReadinessProbe(http.Handle)

	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	go http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	res, err := http.Get(fmt.Sprintf("http://localhost:%d/health/liveness", port))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)

	res, err = http.Get(fmt.Sprintf("http://localhost:%d/health/readiness", port))
	require.NoError(t, err)
	require.Equal(t, http.StatusServiceUnavailable, res.StatusCode)

	SetReady(true)

	res, err = http.Get(fmt.Sprintf("http://localhost:%d/health/readiness", port))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
}
