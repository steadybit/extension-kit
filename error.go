// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

// Package extension_kit provides utilities to handle extension errors.
package extension_kit

import "github.com/steadybit/extension-kit/extutil"

// ExtensionError is a generalization over AttackKit and DiscoveryKit error types. They are structurally identical
// and can be used interchangeably.
type ExtensionError struct {
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`

	// A URI reference that identifies the specific occurrence of the problem.
	Instance *string `json:"instance,omitempty"`

	// A short, human-readable summary of the problem type.
	Title string `json:"title"`

	// A URI reference that identifies the problem type.
	Type *string `json:"type,omitempty"`
}

// ToError converts an error to an ExtensionError.
func ToError(title string, err error) ExtensionError {
	var response ExtensionError
	if err != nil {
		response = ExtensionError{Title: title, Detail: extutil.Ptr(err.Error())}
	} else {
		response = ExtensionError{Title: title}
	}
	return response
}
