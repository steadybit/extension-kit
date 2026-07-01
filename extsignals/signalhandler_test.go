// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2025 Steadybit GmbH
//go:build !windows

package extsignals

// These tests don't consistently work in IntelliJ,
// use the native go test runner instead.

import (
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func TestSignalHandlers(t *testing.T) {
	handler1Run := atomic.Bool{}
	handler2Run := atomic.Bool{}
	handlerList := atomic.Value{}

	ClearSignalHandlers()
	defer ClearSignalHandlers()
	ActivateSignalHandlers()
	RemoveSignalHandlersByName("Termination")
	AddSignalHandler(SignalHandler{
		Handler: func(signal os.Signal) {
			log.Info().Msg("Handler1")
			handler1Run.Store(true)
			handlerList.Store(handlerList.Load().(string) + "Handler1")
		},
		Order: 30,
		Name:  "Handler1",
	})
	AddSignalHandler(SignalHandler{
		Handler: func(signal os.Signal) {
			log.Info().Msg("Handler2")
			handler2Run.Store(true)
			handlerList.Store("Handler2")
		},
		Order: 10,
		Name:  "Handler2",
	})

	err := Kill(os.Getpid())
	require.NoError(t, err)

	// Wait for the signal to be processed
	<-time.After(500 * time.Millisecond)

	require.True(t, handler1Run.Load())
	require.True(t, handler2Run.Load())
	require.Equal(t, handlerList.Load(), "Handler2Handler1")
}

func TestSignalHandler_panicking_handler_does_not_abort_others(t *testing.T) {
	survivorRun := atomic.Bool{}

	ClearSignalHandlers()
	defer ClearSignalHandlers()
	ActivateSignalHandlers()
	RemoveSignalHandlersByName("Termination")
	AddSignalHandler(SignalHandler{
		Handler: func(os.Signal) { panic("boom") },
		Order:   10,
		Name:    "Panicker",
	})
	AddSignalHandler(SignalHandler{
		Handler: func(os.Signal) { survivorRun.Store(true) },
		Order:   20, // ordered after the panicking handler
		Name:    "Survivor",
	})

	err := Kill(os.Getpid())
	require.NoError(t, err)

	<-time.After(500 * time.Millisecond)

	require.True(t, survivorRun.Load(), "a panicking handler must not prevent later handlers from running")
}

func TestRemoveSignalHandlersByName(t *testing.T) {
	handler1Run := atomic.Bool{}
	handler2Run := atomic.Bool{}

	ClearSignalHandlers()
	defer ClearSignalHandlers()
	ActivateSignalHandlers()
	AddSignalHandler(SignalHandler{
		Handler: func(signal os.Signal) {
			log.Info().Msg("Handler1")
			handler1Run.Store(true)
		},
		Order: 30,
		Name:  "Handler1",
	})
	AddSignalHandler(SignalHandler{
		Handler: func(signal os.Signal) {
			log.Info().Msg("Handler2")
			handler2Run.Store(true)
		},
		Order: 10,
		Name:  "Handler2",
	})

	RemoveSignalHandlersByName("Termination", "Handler1")
	err := Kill(os.Getpid())
	require.NoError(t, err)

	// Wait for the signal to be processed
	<-time.After(500 * time.Millisecond)

	require.False(t, handler1Run.Load())
	require.True(t, handler2Run.Load())
}
