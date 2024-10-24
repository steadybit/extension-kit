// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2024 Steadybit GmbH

// Package extension_kit provides utilities to handle extension errors.
package extension_kit

import (
	"errors"
	"fmt"
	"github.com/steadybit/extension-kit/extutil"
)

// ExtensionError is a generalization over ActionKit and DiscoveryKit error types. They are structurally identical
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

func (e ExtensionError) Error() string {
	if e.Detail != nil {
		return fmt.Sprintf("%s: %s", e.Title, *e.Detail)
	}
	return e.Title
}

// ToError converts an error to an ExtensionError.
func ToError(title string, err error) ExtensionError {
	if err != nil {
		return ExtensionError{Title: title, Detail: extutil.Ptr(err.Error())}
	} else {
		return ExtensionError{Title: title}
	}
}

// WrapError if the error is an ExtensionError, it is returned as is. Otherwise, a new ExtensionError is with the error as title.
func WrapError(err error) *ExtensionError {
	if err == nil {
		return nil
	}
	var extErr *ExtensionError
	if errors.As(err, &extErr) {
		return extErr
	}
	return &ExtensionError{Title: err.Error()}
}
