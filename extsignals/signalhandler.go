package extsignals

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
	"sort"
	"syscall"
)

var (
	handlers []SignalHandler
)

const (
	OrderReadinessFalse    = 0   //Set Readiness to false
	OrderStopActions       = 10  //Stop all actions
	OrderStopCustom        = 20  //Custom handler
	OrderStopProbesHttp    = 80  //Shutdown the probes HTTP server
	OrderStopExtensionHttp = 90  //Shutdown the extension HTTP server
	OrderTermination       = 100 //Fallback handler for SIGINT and SIGTERM, the extension usually stops after shutting down the server. This is a last resort if there is an issue with the server shutdown.
)

type SignalHandler struct {
	Handler func(signal os.Signal)
	Order   int
	Name    string
}

type ByOrder []SignalHandler

func (a ByOrder) Len() int           { return len(a) }
func (a ByOrder) Less(i, j int) bool { return a[i].Order < a[j].Order }
func (a ByOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func AddSignalHandler(signalHandler SignalHandler) {
	handlers = append(handlers, signalHandler)
}

// RemoveSignalHandlersByName removes signal handlers by name. This is mainly used for testing.
func RemoveSignalHandlersByName(names ...string) {
	for _, name := range names {
		for i, handler := range handlers {
			if handler.Name == name {
				handlers = append(handlers[:i], handlers[i+1:]...)
			}
		}
	}
}

func ActivateSignalHandlers() {
	AddSignalHandler(SignalHandler{
		Handler: func(signal os.Signal) {
			switch signal {
			case syscall.SIGINT:
				os.Exit(128 + int(signal.(syscall.Signal)))

			case syscall.SIGTERM:
				fmt.Printf("Terminated: %d\n", int(signal.(syscall.Signal)))
				os.Exit(128 + int(signal.(syscall.Signal)))
			}
		},
		Order: OrderTermination,
		Name:  "Termination",
	})

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	go func(signals <-chan os.Signal) {
		for s := range signals {
			signalName := unix.SignalName(s.(syscall.Signal))
			sort.Sort(ByOrder(handlers))
			for _, handler := range handlers {
				log.Debug().Str("signal", signalName).Str("handler", handler.Name).Int("order", handler.Order).Msg("received signal - call handler")
				handler.Handler(s)
			}
		}
	}(signalChannel)
}
