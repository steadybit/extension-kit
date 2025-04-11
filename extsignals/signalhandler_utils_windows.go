//go:build windows

package extsignals

import (
	"os"
	"os/signal"
	"sort"
	"syscall"
)

// 'syscall_unix' package provides similar abstractions for UNIX but not for Windows.
// https://learn.microsoft.com/en-us/windows/console/handlerroutine
var signalList = [...]struct {
	num  syscall.Signal
	name string
	desc string
}{
	{0, "CTRL_C_EVENT", "'ctrl + c' received"},
	{1, "CTRL_BREAK_EVENT", "'ctrl + break' received"},
	{2, "CTRL_CLOSE_EVENT", "close on windows window menu/end task using task manager"},
	{5, "CTRL_LOGOFF_EVENT", "user logging off"},
	{6, "CTRL_SHUTDOWN_EVENT", "system shutting down"},
}

func GetSignalName(s syscall.Signal) string {
	i := sort.Search(len(signalList), func(i int) bool {
		return signalList[i].num >= s
	})
	if i < len(signalList) && signalList[i].num == s {
		return signalList[i].name
	}
	return ""
}

func Notify(c chan<- os.Signal, _ ...os.Signal) {
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
}

func Kill(pid int) (e error) {
	dll, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		return err
	}
	defer func(dll *syscall.DLL) {
		_ = dll.Release()
	}(dll)

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
