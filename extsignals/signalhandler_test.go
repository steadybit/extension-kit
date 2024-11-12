package extsignals

import (
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

func TestSignalHandlers(t *testing.T) {
	//cleanup previous test
	RemoveSignalHandlersByName("Termination", "Handler1", "Handler2")

	handler1Run := atomic.Bool{}
	handler2Run := atomic.Bool{}
	handlerList := atomic.Value{}

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

	err := syscall.Kill(os.Getpid(), syscall.SIGUSR1)
	require.NoError(t, err)

	// Wait for the signal to be processed
	<-time.After(500 * time.Millisecond)

	require.True(t, handler1Run.Load())
	require.True(t, handler2Run.Load())
	require.Equal(t, handlerList.Load(), "Handler2Handler1")
}

func TestRemoveSignalHandlersByName(t *testing.T) {
	//cleanup previous test
	RemoveSignalHandlersByName("Termination", "Handler1", "Handler2")

	handler1Run := atomic.Bool{}
	handler2Run := atomic.Bool{}

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
	err := syscall.Kill(os.Getpid(), syscall.SIGUSR1)
	require.NoError(t, err)

	// Wait for the signal to be processed
	<-time.After(500 * time.Millisecond)

	require.False(t, handler1Run.Load())
	require.True(t, handler2Run.Load())
}
