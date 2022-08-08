// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

// Package extutil contains a variety of util functions that were identified as common code duplication.
// More specialized packages exist for groups of use cases (extlogging and exthttp) for example.
package extutil

// Ptr returns a pointer to the given value. You will find this helpful when desiring to pass a literal value to a function that requires a pointer.
func Ptr[T any](val T) *T {
	return &val
}
