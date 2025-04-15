// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

package extcmd

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"runtime"
	"testing"
)

func TestNewCmdState(t *testing.T) {
	_, err := GetCmdState("I am unknown")
	assert.Error(t, err)

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "echo hello world")
	} else {
		cmd = exec.Command("echo", "hello", "world")
	}

	cs := NewCmdState(cmd)

	persistedState, err := GetCmdState(cs.Id)
	assert.NoError(t, err)
	assert.NotNil(t, persistedState)

	err = cmd.Start()
	assert.NoError(t, err)
	err = cmd.Wait()
	assert.NoError(t, err)

	messages := cs.GetMessages(true)
	assert.Len(t, messages, 1)
	assert.Equal(t, "info", *messages[0].Level)
	assert.Contains(t, messages[0].Message, "hello world")

	RemoveCmdState(cs.Id)
	persistedState, err = GetCmdState(cs.Id)
	assert.Error(t, err)
	assert.Nil(t, persistedState)
}
