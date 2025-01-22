// 'syscall_unix' package provides similar abstractions for UNIX but not for Windows.
package extsignals

import (
	"sort"
	"syscall"
)

func SignalName(s syscall.Signal) string {
	i := sort.Search(len(signalList), func(i int) bool {
		return signalList[i].num >= s
	})
	if i < len(signalList) && signalList[i].num == s {
		return signalList[i].name
	}
	return ""
}

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
