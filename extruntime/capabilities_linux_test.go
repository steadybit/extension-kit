// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

//go:build linux

package extruntime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// process builds the capabilities of a process holding the named permitted capabilities, with the
// named ones in its bounding set; newer is the lowest capability number the kernel does not know.
func process(permitted, bounding []string, newer int, noNewPrivs bool) processCapabilities {
	p := processCapabilities{noNewPrivs: noNewPrivs}
	for _, name := range permitted {
		n, _ := capabilityNumber(name)
		p.permitted |= 1 << uint(n)
	}
	inBounding := map[int]bool{}
	for _, name := range bounding {
		n, _ := capabilityNumber(name)
		inBounding[n] = true
	}
	p.inBounding = func(n int) (bool, bool) {
		if n >= newer {
			return false, false
		}
		return inBounding[n], true
	}
	return p
}

// The extension's file capabilities (what it holds) and the container's capabilities (bounding set).
var (
	fileCaps      = []string{"SETUID", "SETGID", "SYS_ADMIN", "SYS_CHROOT", "SYS_PTRACE", "DAC_OVERRIDE", "NET_ADMIN"}
	containerCaps = []string{"SETUID", "SETGID", "SYS_ADMIN", "SYS_CHROOT", "SYS_PTRACE", "DAC_OVERRIDE", "NET_ADMIN", "NET_RAW", "BPF"}
)

func TestMissingCapabilities_ReportsWhatTheContainerDoesNotGrant(t *testing.T) {
	p := process(fileCaps, []string{"SETUID", "SETGID", "SYS_ADMIN"}, 41, false)

	assert.Equal(t, []string{"NET_ADMIN", "BPF"}, missingCapabilities(p, []string{"SYS_ADMIN", "NET_ADMIN", "BPF"}))
}

func TestMissingCapabilities_RootHelpersGetTheBoundingSet(t *testing.T) {
	// NET_RAW and BPF are not file capabilities, but the root helpers (tc, iptables) get them.
	p := process(fileCaps, containerCaps, 41, false)

	assert.Empty(t, missingCapabilities(p, []string{"NET_ADMIN", "NET_RAW", "BPF"}))
}

func TestMissingCapabilities_NoNewPrivsLimitsRootHelpersToWhatTheExtensionHolds(t *testing.T) {
	// allowPrivilegeEscalation: false sets no_new_privs. The file capabilities the container grants
	// are still held, but a root helper gains nothing beyond them: NET_RAW and BPF, not file
	// capabilities, are not usable anymore. (Checked on a real binary in Docker with
	// --security-opt no-new-privileges.)
	p := process(fileCaps, containerCaps, 41, true)

	assert.Equal(t, []string{"NET_RAW", "BPF"}, missingCapabilities(p, []string{"SETUID", "NET_ADMIN", "NET_RAW", "BPF"}))
}

func TestMissingCapabilities_NoRootHelpersWithoutSetuid(t *testing.T) {
	// Only what the extension holds itself is usable.
	p := process([]string{"SYS_ADMIN", "NET_ADMIN"}, containerCaps, 41, false)

	assert.Equal(t, []string{"NET_RAW"}, missingCapabilities(p, []string{"SYS_ADMIN", "NET_ADMIN", "NET_RAW"}))
}

func TestMissingCapabilities_OldKernelsCoverBPFWithSysAdmin(t *testing.T) {
	// Kernel 5.4 knows capabilities up to AUDIT_READ (37): BPF (39) was part of SYS_ADMIN.
	withSysAdmin := process(fileCaps, []string{"SETUID", "SETGID", "SYS_ADMIN"}, 38, false)
	withoutSysAdmin := process(fileCaps, []string{"SETUID", "SETGID"}, 38, false)

	assert.Empty(t, missingCapabilities(withSysAdmin, []string{"BPF", "PERFMON", "CHECKPOINT_RESTORE"}))
	assert.Equal(t, []string{"BPF"}, missingCapabilities(withoutSysAdmin, []string{"BPF"}))
}

func TestMissingCapabilities_UnknownNames(t *testing.T) {
	p := process(fileCaps, containerCaps, 41, false)

	assert.Equal(t, []string{"NOT_A_CAPABILITY"}, missingCapabilities(p, []string{"NOT_A_CAPABILITY"}))
}
