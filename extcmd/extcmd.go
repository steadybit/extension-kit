// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

// Package extcmd supports use cases in which a command is supposed be started non-blocking as a result of an incoming
// HTTP requests, e.g., ActionKit's start call. It allows retrieval of log messages for streaming to the agent as well
// as stopping and status access.
//
// This package maintains state through a global variable that is a mapping from a random ID to CmdState. The random ID
// is contained within CmdState.
package extcmd

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os/exec"
	"sync"
)

var states = make(map[string]*CmdState)

// NewCmdState create a new CmdState and registers it as a global state. The expected call pattern
// is that NewCmdState is called immediately after the exec.Cmd is created, but before the Cmd
// is started.
func NewCmdState(cmd *exec.Cmd) *CmdState {
	state := new(CmdState)
	state.Id = uuid.NewString()
	state.Cmd = cmd
	state.out = new(bytes.Buffer)
	state.mu = new(sync.Mutex)

	cmd.Stdout = state
	cmd.Stderr = state

	states[state.Id] = state

	return state
}

// GetCmdState returns the state stored under the given ID or an error in case there is no persisted
// state under the ID.
func GetCmdState(id string) (*CmdState, error) {
	state, ok := states[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Failed to find a command state with ID '%s'", id))
	}
	return state, nil
}

// RemoveCmdState removes the state with the given ID. A no-op in case there is no state with this ID.
// It is the caller's responsibility to ensure that the exec.Cmd itself is stopped.
func RemoveCmdState(id string) {
	delete(states, id)
}
