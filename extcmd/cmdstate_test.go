// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

/*
 * Copyright 2022 steadybit GmbH. All rights reserved.
 */

package extcmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
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
