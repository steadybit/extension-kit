package extsignals

import (
	"os"
	"os/signal"
	"syscall"
)

func NotifyPlatformIndependant(c chan<- os.Signal, sig ...os.Signal) {
	signal.Notify(c, syscall.SIGTERM, os.Interrupt)
}

func SignalNamePlatformIndependant(s syscall.Signal) string {
	return SignalName(s)
}
