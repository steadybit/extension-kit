// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

package extcmd

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestNewCmdState(t *testing.T) {
	_, err := GetCmdState("I am unknown")
	assert.NotNil(t, err)

	cmd := exec.Command("echo", "hello", "world")
	cs := NewCmdState(cmd)

	persistedState, err := GetCmdState(cs.Id)
	assert.Nil(t, err)
	assert.NotNil(t, persistedState)

	err = cmd.Start()
	assert.Nil(t, err)
	err = cmd.Wait()
	assert.Nil(t, err)

	messages := cs.GetMessages(true)
	assert.Len(t, messages, 1)
	assert.Equal(t, "info", *messages[0].Level)
	assert.Equal(t, "hello world\n", messages[0].Message)

	RemoveCmdState(cs.Id)
	persistedState, err = GetCmdState(cs.Id)
	assert.NotNil(t, err)
	assert.Nil(t, persistedState)
}
