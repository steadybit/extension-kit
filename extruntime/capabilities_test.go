// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

package extruntime

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapabilityNumber(t *testing.T) {
	for name, want := range map[string]int{"CHOWN": 0, "NET_ADMIN": 12, "cap_sys_admin": 21, "BPF": 39, "CHECKPOINT_RESTORE": 40} {
		got, ok := capabilityNumber(name)
		assert.True(t, ok, name)
		assert.Equal(t, want, got, name)
	}
	_, ok := capabilityNumber("NOT_A_CAPABILITY")
	assert.False(t, ok)
}

func TestMissingCapabilities_ReportsUnknownNames(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("capabilities are Linux only")
	}
	assert.Equal(t, []string{"NOT_A_CAPABILITY"}, MissingCapabilities("NOT_A_CAPABILITY"))
}

func TestRequireCapabilities(t *testing.T) {
	if runtime.GOOS != "linux" {
		assert.NoError(t, RequireCapabilities("network attacks", "NOT_A_CAPABILITY"))
		return
	}
	err := RequireCapabilities("network attacks", "NOT_A_CAPABILITY")
	assert.EqualError(t, err, "network attacks needs the capabilities NOT_A_CAPABILITY, which the extension does not have: "+
		"add them to the capabilities of the extension's container securityContext")
}

func TestRaiseCapabilities(t *testing.T) {
	// Without file capabilities, permitted and effective are equal already (both empty for a
	// non-root user, both full for root), so this must be a no-op that succeeds.
	assert.NoError(t, RaiseCapabilities())
}
