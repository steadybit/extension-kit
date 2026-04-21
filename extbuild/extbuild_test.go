// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package extbuild

import (
	"testing"
)

func TestGetSemverVersionStringOrUnknown(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  string
	}{
		{
			name:  "branch to unknown",
			given: "main",
			want:  "unknown",
		},
		{
			name:  "version string",
			given: "1.2.3",
			want:  "1.2.3",
		},
		{
			name:  "version string with leading v",
			given: "v11.22.33",
			want:  "11.22.33",
		},
		{
			name:  "version string without leading v",
			given: "11.22.33",
			want:  "11.22.33",
		},
		{
			name:  "semver with pre-release-identifier",
			given: "1.0.20-1776424247-next",
			want:  "1.0.20-1776424247-next",
		},
		{
			name:  "semver with pre-release-identifier (variant)",
			given: "1.0.20-next.1776424247",
			want:  "1.0.20-next.1776424247",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Version = tt.given
			if got := GetSemverVersionStringOrUnknown(); got != tt.want {
				t.Errorf("GetSemverVersionStringOrUnknown() = %v, want %v", got, tt.want)
			}
		})
	}
}
