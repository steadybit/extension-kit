// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

package extcmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"
	"sync/atomic"
)

type CmdState struct {
	Id                  string
	Cmd                 *exec.Cmd
	exitCode            atomic.Int32
	mu                  *sync.Mutex
	out                 *bytes.Buffer
	lastPartialLineRead string
}

// Wait blocks until the command exits and records its exit code. It must be called
// exactly once — typically as `go cmdState.Wait()` right after the command is started.
// Wait is the sole reader of Cmd.ProcessState, so concurrent callers must obtain the
// exit code via ExitCode rather than reading Cmd.ProcessState directly; doing the
// latter races this method. The error returned by exec.Cmd.Wait is passed through for
// logging.
func (cs *CmdState) Wait() error {
	err := cs.Cmd.Wait()
	cs.exitCode.Store(int32(cs.Cmd.ProcessState.ExitCode()))
	return err
}

// ExitCode returns the command's exit code, or -1 while it is still running or if it
// was terminated by a signal, matching os.ProcessState.ExitCode(). It is safe to call
// concurrently with the Wait goroutine.
func (cs *CmdState) ExitCode() int {
	return int(cs.exitCode.Load())
}

func (cs *CmdState) Write(p []byte) (n int, err error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return cs.out.Write(p)
}

func (cs *CmdState) GetLines(includePartialLines bool) []string {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	var result []string
	line, err := cs.out.ReadString('\n')
	for ; err == nil; line, err = cs.out.ReadString('\n') {
		result = append(result, line)
	}

	if len(cs.lastPartialLineRead) > 0 && len(result) > 0 {
		result[0] = fmt.Sprintf("%s%s", cs.lastPartialLineRead, result[0])
		cs.lastPartialLineRead = ""
	}

	if len(line) > 0 {
		cs.lastPartialLineRead = fmt.Sprintf("%s%s", cs.lastPartialLineRead, line)
	}

	if includePartialLines && len(cs.lastPartialLineRead) > 0 {
		result = append(result, cs.lastPartialLineRead)
	}

	return result
}

// Message is API compatible with ActionKit and DiscoveryKit Message
type Message struct {
	Level   *string `json:"level,omitempty"`
	Message string  `json:"message"`
}

func (cs *CmdState) GetMessages(includePartialMessages bool) []Message {
	lines := cs.GetLines(includePartialMessages)
	messages := make([]Message, 0, len(lines))
	for _, line := range lines {
		messages = append(messages, Message{
			Level:   new("info"),
			Message: line,
		})
	}
	return messages
}
