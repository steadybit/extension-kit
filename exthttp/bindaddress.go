// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

package exthttp

import (
	"errors"
	"fmt"
	"net"
	"net/netip"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

// BindAddressLocalhost binds to the IPv4 and IPv6 loopback addresses. Go would
// listen on only one address for the host name "localhost", so a client that
// resolves it to the other one could not connect.
const BindAddressLocalhost = "localhost"

var defaultBindAddress string

// SetDefaultBindAddress sets the address that the extension and health HTTP servers bind to, unless overridden through
// the environment variable STEADYBIT_EXTENSION_BIND_ADDRESS. Empty (the default) binds to all interfaces. Call it
// before exthealth.StartProbes and Listen.
func SetDefaultBindAddress(bindAddress string) {
	defaultBindAddress = bindAddress
}

// BindAddress returns the address that the extension and health HTTP servers bind to.
func BindAddress() string {
	spec := ListenSpecification{}
	spec.parseConfigurationFromEnvironment()
	return spec.bindAddress()
}

func (spec *ListenSpecification) bindAddress() string {
	if spec.BindAddress != "" {
		return spec.BindAddress
	}
	return defaultBindAddress
}

// hasIPv6Loopback reports whether ::1 is configured on this host. Variable for tests.
var hasIPv6Loopback = func() bool {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Debug().Err(err).Msg("unable to list interface addresses, assuming IPv6 loopback is available")
		return true
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.Equal(net.IPv6loopback) {
			return true
		}
	}
	return false
}

// resolveBindAddresses returns the host:port pairs to listen on for a bind address:
//   - "" binds to all interfaces (":port")
//   - "localhost" binds to 127.0.0.1 and, when configured on the host, ::1
//   - an IP literal binds to exactly that IP; 0.0.0.0 and :: bind to all interfaces
//
// Host names are rejected: Go would listen on only one of the resolved addresses.
func resolveBindAddresses(bindAddress string, port int) ([]string, error) {
	portStr := strconv.Itoa(port)
	bindAddress = strings.TrimSpace(bindAddress)

	if bindAddress == "" {
		return []string{net.JoinHostPort("", portStr)}, nil
	}

	if strings.EqualFold(bindAddress, BindAddressLocalhost) {
		addresses := []string{net.JoinHostPort("127.0.0.1", portStr)}
		if hasIPv6Loopback() {
			addresses = append(addresses, net.JoinHostPort("::1", portStr))
		}
		return addresses, nil
	}

	ip, err := netip.ParseAddr(bindAddress)
	if err != nil {
		return nil, fmt.Errorf("invalid bind address %q: must be an IP address or %q", bindAddress, BindAddressLocalhost)
	}
	return []string{net.JoinHostPort(ip.String(), portStr)}, nil
}

// ListenTCP opens TCP listeners on the given port for the bind address. See
// STEADYBIT_EXTENSION_BIND_ADDRESS in the README for the accepted values. It
// fails if any of the listeners cannot be opened.
func ListenTCP(bindAddress string, port int) ([]net.Listener, error) {
	addresses, err := resolveBindAddresses(bindAddress, port)
	if err != nil {
		return nil, err
	}

	listeners := make([]net.Listener, 0, len(addresses))
	for _, address := range addresses {
		listener, err := net.Listen("tcp", address)
		if err != nil {
			closeListeners(listeners)
			return nil, err
		}
		listeners = append(listeners, listener)
	}
	return listeners, nil
}

// ServeListeners runs serve for every listener and returns the first error. With
// http.Server, closing or shutting down the server ends all of them.
func ServeListeners(listeners []net.Listener, serve func(net.Listener) error) error {
	if len(listeners) == 0 {
		return errors.New("no listeners to serve")
	}
	errs := make(chan error, len(listeners))
	for _, listener := range listeners {
		go func() {
			errs <- serve(listener)
		}()
	}
	return <-errs
}

// ListenerAddresses joins the listeners' addresses for log messages.
func ListenerAddresses(listeners []net.Listener) string {
	addresses := make([]string, 0, len(listeners))
	for _, listener := range listeners {
		addresses = append(addresses, listener.Addr().String())
	}
	return strings.Join(addresses, ", ")
}

func closeListeners(listeners []net.Listener) {
	var errs error
	for _, listener := range listeners {
		errs = errors.Join(errs, listener.Close())
	}
	if errs != nil {
		log.Debug().Err(errs).Msg("failed to close listeners")
	}
}
