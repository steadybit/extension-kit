// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

//go:build linux

package extruntime

import (
	"errors"
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"github.com/rs/zerolog/log"
	"golang.org/x/sys/unix"
)

// RaiseCapabilities moves every capability of the permitted set into the effective set, on all
// threads of the process.
//
// Extensions run as a non-root user and get their capabilities as file capabilities of their
// binary. With the effective bit set on the file (setcap …+eip), the kernel refuses to exec the
// binary when one of them is not granted to the container, and the extension crash-loops. Without
// it (setcap …+ip), the binary starts with whatever the container grants in its permitted set, and
// this function makes them effective. What is missing can then be reported per action, see
// MissingCapabilities.
//
// Capabilities are per thread, so the change is applied to every thread at once; this needs a
// binary built without cgo. Raising them on a single thread would not help: the Go scheduler moves
// goroutines between threads. So a binary built with cgo gets an error and no capability is raised.
func RaiseCapabilities() error {
	hdr := unix.CapUserHeader{Version: unix.LINUX_CAPABILITY_VERSION_3}
	data := [2]unix.CapUserData{}
	if err := unix.Capget(&hdr, &data[0]); err != nil {
		return fmt.Errorf("reading the capabilities: %w", err)
	}
	if data[0].Effective == data[0].Permitted && data[1].Effective == data[1].Permitted {
		return nil
	}
	data[0].Effective, data[1].Effective = data[0].Permitted, data[1].Permitted

	hdr.Pid = 0
	_, _, errno := syscall.AllThreadsSyscall(unix.SYS_CAPSET, uintptr(unsafe.Pointer(&hdr)), uintptr(unsafe.Pointer(&data[0])), 0)
	if errno == 0 {
		return nil
	}
	if errors.Is(errno, syscall.ENOTSUP) {
		return errors.New("raising the capabilities on all threads needs a binary built without cgo (CGO_ENABLED=0); none were raised")
	}
	return fmt.Errorf("raising the capabilities: %w", errno)
}

// MissingCapabilities returns the capabilities among required that the extension cannot use,
// named like the input. A capability is usable when it is in the bounding set, which the
// container's securityContext decides, and either:
//   - the extension holds it (permitted set, from its file capabilities), or
//   - the extension can run root helpers (runc, nsenter, tc, …) that get it: it holds SETUID and
//     SETGID, and no_new_privs is off. With allowPrivilegeEscalation: false, no_new_privs is on: the
//     kernel caps what an exec gains at what the parent held, so the root helpers get no more than
//     the extension holds itself.
//
// Capabilities newer than the kernel (BPF and PERFMON before 5.8, CHECKPOINT_RESTORE before 5.9)
// were part of SYS_ADMIN there, so SYS_ADMIN stands in for them.
func MissingCapabilities(required ...string) []string {
	return missingCapabilities(currentProcess(), required)
}

// processCapabilities is what decides whether a capability is usable, read from the process.
type processCapabilities struct {
	permitted uint64
	// inBounding reports whether the capability is in the bounding set; known is false when the
	// kernel does not know the capability.
	inBounding func(n int) (in bool, known bool)
	noNewPrivs bool
}

func currentProcess() processCapabilities {
	p := processCapabilities{
		inBounding: func(n int) (bool, bool) {
			in, err := unix.PrctlRetInt(unix.PR_CAPBSET_READ, uintptr(n), 0, 0, 0)
			if errors.Is(err, unix.EINVAL) {
				return false, false
			}
			return err == nil && in == 1, true
		},
	}
	hdr := unix.CapUserHeader{Version: unix.LINUX_CAPABILITY_VERSION_3}
	data := [2]unix.CapUserData{}
	if err := unix.Capget(&hdr, &data[0]); err == nil {
		p.permitted = uint64(data[0].Permitted) | uint64(data[1].Permitted)<<32
	}
	if nnp, err := unix.PrctlRetInt(unix.PR_GET_NO_NEW_PRIVS, 0, 0, 0, 0); err != nil || nnp == 1 {
		p.noNewPrivs = true
	}
	return p
}

// legacyCapabilities are the capabilities split out of SYS_ADMIN by newer kernels.
var legacyCapabilities = map[string]string{"BPF": "SYS_ADMIN", "PERFMON": "SYS_ADMIN", "CHECKPOINT_RESTORE": "SYS_ADMIN"}

func missingCapabilities(p processCapabilities, required []string) []string {
	holds := func(n int) bool { return p.permitted&(1<<uint(n)) != 0 }
	setuid, _ := capabilityNumber("SETUID")
	setgid, _ := capabilityNumber("SETGID")
	canRunRootHelpers := !p.noNewPrivs && holds(setuid) && holds(setgid)

	var missing []string
	for _, name := range required {
		n, ok := capabilityNumber(name)
		if !ok {
			missing = append(missing, name)
			continue
		}
		in, known := p.inBounding(n)
		if !known {
			legacy, ok := legacyCapabilities[strings.TrimPrefix(strings.ToUpper(name), "CAP_")]
			if !ok {
				missing = append(missing, name)
				continue
			}
			n, _ = capabilityNumber(legacy)
			in, known = p.inBounding(n)
		}
		if !known || !in || !(holds(n) || canRunRootHelpers) {
			missing = append(missing, name)
		}
	}
	return missing
}

// LogMissingCapabilities logs a warning for the expected capabilities the extension lacks: the
// actions needing them are unavailable, or run degraded for an optional one.
func LogMissingCapabilities(expected ...string) {
	if missing := MissingCapabilities(expected...); len(missing) > 0 {
		log.Warn().Strs("missing", missing).Msg("The extension runs without some capabilities; the actions needing them are unavailable or degraded")
	}
}
