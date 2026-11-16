// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

package exthttp

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_resolveBindAddresses(t *testing.T) {
	tests := []struct {
		name            string
		bindAddress     string
		hasIPv6Loopback bool
		want            []string
		wantErr         string
	}{
		{name: "empty binds to all interfaces", bindAddress: "", want: []string{":8085"}},
		{name: "blank binds to all interfaces", bindAddress: "  ", want: []string{":8085"}},
		{name: "localhost binds to both loopbacks", bindAddress: "localhost", hasIPv6Loopback: true, want: []string{"127.0.0.1:8085", "[::1]:8085"}},
		{name: "localhost is case-insensitive", bindAddress: "LocalHost", hasIPv6Loopback: true, want: []string{"127.0.0.1:8085", "[::1]:8085"}},
		{name: "localhost without IPv6 loopback", bindAddress: "localhost", hasIPv6Loopback: false, want: []string{"127.0.0.1:8085"}},
		{name: "IPv4 address", bindAddress: "10.0.0.5", want: []string{"10.0.0.5:8085"}},
		{name: "IPv4 wildcard", bindAddress: "0.0.0.0", want: []string{"0.0.0.0:8085"}},
		{name: "IPv6 address", bindAddress: "::1", want: []string{"[::1]:8085"}},
		{name: "IPv6 wildcard", bindAddress: "::", want: []string{"[::]:8085"}},
		{name: "host names are rejected", bindAddress: "example.com", wantErr: `invalid bind address "example.com"`},
		{name: "addresses with port are rejected", bindAddress: "127.0.0.1:8085", wantErr: "invalid bind address"},
		{name: "lists are rejected", bindAddress: "127.0.0.1,::1", wantErr: "invalid bind address"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			old := hasIPv6Loopback
			defer func() { hasIPv6Loopback = old }()
			hasIPv6Loopback = func() bool { return tt.hasIPv6Loopback }

			got, err := resolveBindAddresses(tt.bindAddress, 8085)

			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestListenTCPOnLocalhostAcceptsOnlyLoopback(t *testing.T) {
	if !hasIPv6Loopback() {
		t.Skip("IPv6 loopback not available")
	}
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	listeners, err := ListenTCP("localhost", port)
	require.NoError(t, err)
	defer closeListeners(listeners)
	go acceptAndClose(listeners)

	assertReachable(t, "127.0.0.1", port)
	assertReachable(t, "::1", port)
	if ip := nonLoopbackIP(t); ip != "" {
		assertNotReachable(t, ip, port)
	}
}

func TestListenTCPOnAllInterfaces(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	listeners, err := ListenTCP("", port)
	require.NoError(t, err)
	defer closeListeners(listeners)
	go acceptAndClose(listeners)

	assertReachable(t, "127.0.0.1", port)
	if ip := nonLoopbackIP(t); ip != "" {
		assertReachable(t, ip, port)
	}
}

func TestListenTCPClosesListenersOnFailure(t *testing.T) {
	if !hasIPv6Loopback() {
		t.Skip("IPv6 loopback not available")
	}
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	// Occupy ::1 so that the second listener fails after 127.0.0.1 was opened.
	blocker, err := net.Listen("tcp", net.JoinHostPort("::1", strconv.Itoa(port)))
	require.NoError(t, err)
	defer blocker.Close()

	_, err = ListenTCP("localhost", port)
	require.Error(t, err)

	// 127.0.0.1 must have been released again.
	listener, err := net.Listen("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(port)))
	require.NoError(t, err)
	_ = listener.Close()
}

func TestStartHttpServerOnDefaultBindAddress(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	old := http.DefaultServeMux
	defer func() { http.DefaultServeMux = old }()
	SetDefaultBindAddress(BindAddressLocalhost)
	defer SetDefaultBindAddress("")
	go Listen(ListenOpts{Port: port})
	WaitForServe()
	defer StopListen()

	resp, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	if ip := nonLoopbackIP(t); ip != "" {
		assertNotReachable(t, ip, port)
	}
}

func TestStartHttpServerBindAddressFromEnvironmentOverridesDefault(t *testing.T) {
	ip := nonLoopbackIP(t)
	if ip == "" {
		t.Skip("no non-loopback IPv4 address available")
	}
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	old := http.DefaultServeMux
	defer func() { http.DefaultServeMux = old }()
	SetDefaultBindAddress(BindAddressLocalhost)
	defer SetDefaultBindAddress("")
	t.Setenv("STEADYBIT_EXTENSION_BIND_ADDRESS", "0.0.0.0")
	go Listen(ListenOpts{Port: port})
	WaitForServe()
	defer StopListen()

	assertReachable(t, ip, port)
}

func TestStartHttpServerMustFailForInvalidBindAddress(t *testing.T) {
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	old := http.DefaultServeMux
	defer func() { http.DefaultServeMux = old }()
	t.Setenv("STEADYBIT_EXTENSION_BIND_ADDRESS", "example.com")

	err = listen(ListenOpts{Port: port})

	assert.ErrorContains(t, err, `invalid bind address "example.com"`)
}

func acceptAndClose(listeners []net.Listener) {
	for _, listener := range listeners {
		go func() {
			for {
				conn, err := listener.Accept()
				if err != nil {
					return
				}
				_ = conn.Close()
			}
		}()
	}
}

func assertReachable(t *testing.T, ip string, port int) {
	t.Helper()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, strconv.Itoa(port)), time.Second)
	if assert.NoError(t, err, "expected %s to be reachable", ip) {
		_ = conn.Close()
	}
}

func assertNotReachable(t *testing.T, ip string, port int) {
	t.Helper()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, strconv.Itoa(port)), time.Second)
	if err == nil {
		_ = conn.Close()
	}
	assert.Error(t, err, "expected %s not to be reachable", ip)
}

// nonLoopbackIP returns an IPv4 address of this host that is not a loopback address, or "" if there is none.
func nonLoopbackIP(t *testing.T) string {
	t.Helper()
	addrs, err := net.InterfaceAddrs()
	require.NoError(t, err)
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.To4() != nil && !ipNet.IP.IsLoopback() && !ipNet.IP.IsLinkLocalUnicast() {
			return ipNet.IP.String()
		}
	}
	return ""
}
