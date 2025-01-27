package extsignals

import (
	"context"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"
)

var (
	handlers = sync.Map{}
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
	handlers.Store(signalHandler.Name, signalHandler)
}

func ClearSignalHandlers() {
	handlers.Range(func(key, value interface{}) bool {
		handlers.Delete(key)
		return true
	})
}

// RemoveSignalHandlersByName removes signal handlers by name. This is mainly used for testing.
func RemoveSignalHandlersByName(names ...string) {
	for _, name := range names {
		handlers.Delete(name)
	}
}

func createSignalChannel(context context.Context) {
	signalChannel := make(chan os.Signal, 1)
	Notify(signalChannel)
	go func(signals <-chan os.Signal) {
		for {
			select {
			case <-context.Done():
				signal.Stop(signalChannel)
				return
			case s := <-signals:
				handlerList := make([]SignalHandler, 0)
				handlers.Range(func(key, value interface{}) bool {
					handlerList = append(handlerList, value.(SignalHandler))
					return true
				})
				sort.Sort(ByOrder(handlerList))
				signalName := GetSignalName(s.(syscall.Signal))
				for _, handler := range handlerList {
					log.Debug().Str("signal", signalName).Str("handler", handler.Name).Int("order", handler.Order).Msg("received signal - call handler")
					handler.Handler(s)
				}
			}
		}
	}(signalChannel)
}

func ActivateSignalHandlers() {
	ActivateSignalHandlerWithContext(context.Background())
}

func ActivateSignalHandlerWithContext(context context.Context) {
	AddSignalHandler(SignalHandler{
		Handler: func(signal os.Signal) {
			switch signal {
			case syscall.SIGINT:
				os.Exit(128 + int(signal.(syscall.Signal)))

			case syscall.SIGTERM:
				os.Exit(128 + int(signal.(syscall.Signal)))
			}
		},
		Order: OrderTermination,
		Name:  "Termination",
	})

	createSignalChannel(context)
}
