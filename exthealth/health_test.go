/*
 * Copyright 2024 steadybit GmbH. All rights reserved.
 */

package exthealth

import (
	"context"
	"fmt"
	"github.com/phayes/freeport"
	"github.com/steadybit/extension-kit/exthttp"
	"github.com/stretchr/testify/require"
	"net"
	"net/http"
	"path/filepath"
	"testing"
)

func TestServeProbes(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	SetReady(false)
	StartProbes(port)
	defer StopProbes()

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

func TestServerProbesUsingUnixSocket(t *testing.T) {
	sock := filepath.Join(t.TempDir(), "sock")

	t.Setenv("STEADYBIT_EXTENSION_UNIX_SOCKET", sock)
	go exthttp.Listen(exthttp.ListenOpts{})
	exthttp.WaitForServe()
	defer exthttp.StopListen()

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", sock)
			},
		},
	}

	SetReady(false)
	StartProbes(0)
	defer StopProbes()

	res, err := client.Get("http://localhost/health/liveness")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)

	res, err = client.Get("http://localhost/health/readiness")
	require.NoError(t, err)
	require.Equal(t, http.StatusServiceUnavailable, res.StatusCode)

	SetReady(true)

	res, err = client.Get("http://localhost/health/readiness")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)

	SetAlive(false)

	res, err = client.Get("http://localhost/health/liveness")
	require.NoError(t, err)
	require.Equal(t, http.StatusServiceUnavailable, res.StatusCode)

	SetAlive(true)

	res, err = client.Get("http://localhost/health/liveness")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
}
