//go:build !windows

package extsignals

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sys/unix"
)

func Notify(c chan<- os.Signal, sig ...os.Signal) {
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
}

func GetSignalName(s syscall.Signal) string {
	return unix.SignalName(s)
}

func Kill(pid int) (e error) {
	return syscall.Kill(os.Getpid(), syscall.SIGUSR1)
}
