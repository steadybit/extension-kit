package extsignals

import (
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestSignalHandlers(t *testing.T) {
	handler1Run := false
	handler2Run := false
	handlerList := make([]string, 0)

	ActivateSignalHandlers()
	RemoveSignalHandlersByName("Termination")
	AddSignalHandler(SignalHandler{
		Handler: func(signal os.Signal) {
			log.Info().Msg("Handler1")
			handler1Run = true
			handlerList = append(handlerList, "Handler1")
		},
		Order: 30,
		Name:  "Handler1",
	})
	AddSignalHandler(SignalHandler{
		Handler: func(signal os.Signal) {
			log.Info().Msg("Handler2")
			handler2Run = true
			handlerList = append(handlerList, "Handler2")
		},
		Order: 10,
		Name:  "Handler2",
	})

	err := syscall.Kill(os.Getpid(), syscall.SIGUSR1)
	require.NoError(t, err)

	// Wait for the signal to be processed
	<-time.After(500 * time.Millisecond)

	require.True(t, handler1Run)
	require.True(t, handler2Run)
	require.Equal(t, handlerList, []string{"Handler2", "Handler1"})
}

func TestRemoveSignalHandlersByName(t *testing.T) {
	handler1Run := false
	handler2Run := false
	handlerList := make([]string, 0)

	ActivateSignalHandlers()
	AddSignalHandler(SignalHandler{
		Handler: func(signal os.Signal) {
			log.Info().Msg("Handler1")
			handler1Run = true
			handlerList = append(handlerList, "Handler1")
		},
		Order: 30,
		Name:  "Handler1",
	})
	AddSignalHandler(SignalHandler{
		Handler: func(signal os.Signal) {
			log.Info().Msg("Handler2")
			handler2Run = true
			handlerList = append(handlerList, "Handler2")
		},
		Order: 10,
		Name:  "Handler2",
	})

	RemoveSignalHandlersByName("Termination", "Handler1")
	err := syscall.Kill(os.Getpid(), syscall.SIGUSR1)
	require.NoError(t, err)

	// Wait for the signal to be processed
	<-time.After(500 * time.Millisecond)

	require.False(t, handler1Run)
	require.True(t, handler2Run)
	require.Equal(t, handlerList, []string{"Handler2"})
}
