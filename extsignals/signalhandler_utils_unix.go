//go:build !windows
// +build !windows

package extsignals

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sys/unix"
)

func NotifyPlatformIndependant(c chan<- os.Signal, sig ...os.Signal) {
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
}

func SignalNamePlatformIndependant(s syscall.Signal) string {
	return unix.SignalName(s)
}
