// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2024 Steadybit GmbH

package extension_kit

import (
	"errors"
	"github.com/steadybit/extension-kit/extutil"
	"reflect"
	"testing"
)

func TestWrapError(t *testing.T) {
	tests := []struct {
		name string
		arg  error
		want *ExtensionError
	}{
		{
			name: "nil error",
		},
		{
			name: "error",
			arg:  errors.New("some error"),
			want: &ExtensionError{Title: "some error"},
		},
		{
			name: "ExtensionError",
			arg:  &ExtensionError{Title: "ext error"},
			want: &ExtensionError{Title: "ext error"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WrapError(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToError(t *testing.T) {
	tests := []struct {
		name  string
		title string
		err   error
		want  ExtensionError
	}{
		{
			name:  "nil error",
			title: "some title",
			want:  ExtensionError{Title: "some title"},
		},
		{
			name:  "some error",
			title: "some title",
			err:   errors.New("some error"),
			want:  ExtensionError{Title: "some title", Detail: extutil.Ptr("some error")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToError(tt.title, tt.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToError() = %v, want %v", got, tt.want)
			}
		})
	}
}
