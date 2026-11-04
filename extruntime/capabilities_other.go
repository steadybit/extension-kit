// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

//go:build !linux

package extruntime

// RaiseCapabilities is a no-op outside Linux.
func RaiseCapabilities() error { return nil }

// MissingCapabilities reports nothing missing outside Linux, where capabilities do not exist.
func MissingCapabilities(...string) []string { return nil }

// MissingHeldCapabilities reports nothing missing outside Linux, where capabilities do not exist.
func MissingHeldCapabilities(...string) []string { return nil }

// LogMissingCapabilities is a no-op outside Linux.
func LogMissingCapabilities(...string) {}
