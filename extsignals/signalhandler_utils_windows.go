package extsignals

import (
	"os"
	"os/signal"
	"syscall"
)

func Notify(c chan<- os.Signal, sig ...os.Signal) {
	signal.Notify(c, syscall.SIGTERM, os.Interrupt)
}

func GetSignalName(s syscall.Signal) string {
	return SignalName(s)
}

func Kill(pid int) (e error) {
	dll, err := syscall.LoadDLL("kernel32.dll")

	if err != nil {
		return err
	}

	process, err := dll.FindProc("GenerateConsoleCtrlEvent")

	if err != nil {
		return err
	}

	// Process call error is not checked through err variable, but instead through result code (err contains result message which can be success as well).
	// https://go.dev/src/os/signal/signal_windows_test.go
	result, _, err := process.Call(syscall.CTRL_BREAK_EVENT, uintptr(pid))

	if result == 0 {
		return err
	}

	return nil
}
