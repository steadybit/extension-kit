// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

package extruntime

import (
	"fmt"
	"strings"
)

// capabilityNames are the Linux capabilities by number, named as in Kubernetes securityContexts
// (without the CAP_ prefix).
var capabilityNames = []string{
	"CHOWN", "DAC_OVERRIDE", "DAC_READ_SEARCH", "FOWNER", "FSETID", "KILL", "SETGID", "SETUID",
	"SETPCAP", "LINUX_IMMUTABLE", "NET_BIND_SERVICE", "NET_BROADCAST", "NET_ADMIN", "NET_RAW",
	"IPC_LOCK", "IPC_OWNER", "SYS_MODULE", "SYS_RAWIO", "SYS_CHROOT", "SYS_PTRACE", "SYS_PACCT",
	"SYS_ADMIN", "SYS_BOOT", "SYS_NICE", "SYS_RESOURCE", "SYS_TIME", "SYS_TTY_CONFIG", "MKNOD",
	"LEASE", "AUDIT_WRITE", "AUDIT_CONTROL", "SETFCAP", "MAC_OVERRIDE", "MAC_ADMIN", "SYSLOG",
	"WAKE_ALARM", "BLOCK_SUSPEND", "AUDIT_READ", "PERFMON", "BPF", "CHECKPOINT_RESTORE",
}

func capabilityNumber(name string) (int, bool) {
	name = strings.TrimPrefix(strings.ToUpper(name), "CAP_")
	for i, n := range capabilityNames {
		if n == name {
			return i, true
		}
	}
	return 0, false
}

// RequireCapabilities returns an error naming the capabilities the extension lacks to do what,
// or nil when it has them all. See MissingCapabilities.
func RequireCapabilities(what string, required ...string) error {
	missing := MissingCapabilities(required...)
	if len(missing) == 0 {
		return nil
	}
	return fmt.Errorf("%s needs the capabilities %s, which the extension does not have: add them to the capabilities of the extension's container securityContext",
		what, strings.Join(missing, ", "))
}
