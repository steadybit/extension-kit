//go:build windows
// +build windows

package extsignals

import (
	"os"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
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

	dll, err := syscall.LoadDLL("kernel32.dll")
	require.NoError(t, err)

	process, err := dll.FindProc("GenerateConsoleCtrlEvent")
	require.NoError(t, err)

	// Process call error is not checked through err variable, but instead through result code (err contains result message which can be success as well).
	// https://go.dev/src/os/signal/signal_windows_test.go
	result, _, _ := process.Call(syscall.CTRL_BREAK_EVENT, uintptr(os.Getpid()))
	require.NotEqual(t, result, 0)

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
	dll, err := syscall.LoadDLL("kernel32.dll")
	require.NoError(t, err)

	process, err := dll.FindProc("GenerateConsoleCtrlEvent")
	require.NoError(t, err)

	// Process call error is not checked through err variable, but instead through result code (err contains result message which can be success as well).
	// https://go.dev/src/os/signal/signal_windows_test.go
	result, _, _ := process.Call(syscall.CTRL_BREAK_EVENT, uintptr(os.Getpid()))
	require.NotEqual(t, result, 0)

	// Wait for the signal to be processed
	<-time.After(500 * time.Millisecond)

	require.False(t, handler1Run.Load())
	require.True(t, handler2Run.Load())
}
