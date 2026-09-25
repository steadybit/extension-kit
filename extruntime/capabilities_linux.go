// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

//go:build linux

package extruntime

import (
	"errors"
	"fmt"
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
// Capabilities are per thread, so the change is applied to every thread; this needs a binary built
// without cgo. It returns an error when the effective set could not be raised on all threads.
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
		// Built with cgo: at least raise them on this thread.
		if err := unix.Capset(&hdr, &data[0]); err != nil {
			return fmt.Errorf("raising the capabilities: %w", err)
		}
		return errors.New("raising the capabilities on all threads needs a binary built without cgo; raised on the main thread only")
	}
	return fmt.Errorf("raising the capabilities: %w", errno)
}

// MissingCapabilities returns the capabilities among required that the extension cannot use,
// named like the input, or all of them when they cannot be read. A capability is usable when it is
// in the bounding set: the container's securityContext decides it, and the root helpers the
// extension runs (runc, nsenter, tc, …) inherit it.
func MissingCapabilities(required ...string) []string {
	var missing []string
	for _, name := range required {
		n, ok := capabilityNumber(name)
		if !ok {
			missing = append(missing, name)
			continue
		}
		if in, err := unix.PrctlRetInt(unix.PR_CAPBSET_READ, uintptr(n), 0, 0, 0); err != nil || in != 1 {
			missing = append(missing, name)
		}
	}
	return missing
}

// LogMissingCapabilities logs a warning for the expected capabilities the extension lacks. Actions
// needing them fail, the others work.
func LogMissingCapabilities(expected ...string) {
	if missing := MissingCapabilities(expected...); len(missing) > 0 {
		log.Warn().Strs("missing", missing).Msg("The extension runs without some capabilities; the actions needing them will fail")
	}
}
