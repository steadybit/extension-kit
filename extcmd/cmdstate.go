// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

package extcmd

import (
	"bytes"
	"fmt"
	"github.com/steadybit/extension-kit/extutil"
	"os/exec"
	"sync"
)

type CmdState struct {
	Id                  string
	Cmd                 *exec.Cmd
	mu                  *sync.Mutex
	out                 *bytes.Buffer
	lastPartialLineRead string
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
			Level:   extutil.Ptr("info"),
			Message: line,
		})
	}
	return messages
}
