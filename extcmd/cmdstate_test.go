// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

/*
 * Copyright 2022 steadybit GmbH. All rights reserved.
 */

package extcmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os/exec"
	"runtime"
	"sync"
	"testing"
)

func TestCmdStateReadsFullLines(t *testing.T) {
	is := new(CmdState)
	is.mu = new(sync.Mutex)
	is.out = new(bytes.Buffer)

	assert.Equal(t, 0, len(is.GetLines(false)))

	is.Write([]byte("first"))
	assert.Equal(t, 0, len(is.GetLines(false)))

	is.Write([]byte(" line"))
	assert.Equal(t, 0, len(is.GetLines(false)))

	is.Write([]byte("\nSecond line"))
	assert.Equal(t, []string{"first line\n"}, is.GetLines(false))

	assert.Equal(t, 0, len(is.GetLines(false)))

	assert.Equal(t, []string{"Second line"}, is.GetLines(true))
}

func TestCmdStateExitCode(t *testing.T) {
	cs := NewCmdState(exec.Command("sh", "-c", "exit 7"))
	defer RemoveCmdState(cs.Id)

	// -1 while the command has not yet exited, matching os.ProcessState.ExitCode().
	assert.Equal(t, -1, cs.ExitCode())

	require.NoError(t, cs.Cmd.Start())
	err := cs.Wait()

	require.Error(t, err) // a non-zero exit makes exec.Cmd.Wait return an *ExitError
	assert.Equal(t, 7, cs.ExitCode())
}

func TestCmdStateExitCodeSuccess(t *testing.T) {
	cs := NewCmdState(exec.Command("sh", "-c", "exit 0"))
	defer RemoveCmdState(cs.Id)

	require.NoError(t, cs.Cmd.Start())
	require.NoError(t, cs.Wait())
	assert.Equal(t, 0, cs.ExitCode())
}

func TestCmdStateExitCodeIsConcurrencySafe(t *testing.T) {
	cs := NewCmdState(exec.Command("sh", "-c", "exit 0"))
	defer RemoveCmdState(cs.Id)
	require.NoError(t, cs.Cmd.Start())

	// Readers (as the status/stop handlers do) must not race the Wait goroutine
	// (caught by `go test -race`).
	var wg sync.WaitGroup
	for r := 0; r < 4; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for cs.ExitCode() == -1 {
				runtime.Gosched()
			}
		}()
	}
	require.NoError(t, cs.Wait())
	wg.Wait()
	assert.Equal(t, 0, cs.ExitCode())
}
