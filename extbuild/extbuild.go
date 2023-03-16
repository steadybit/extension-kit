// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package extbuild

import (
	"github.com/rs/zerolog/log"
	"regexp"
)

var semverPattern = regexp.MustCompile("^v\\d+\\.\\d+\\.\\d+$")

var ExtensionName string = "unknown"
var Version string = "unknown"
var Revision string = "unknown"

// PrintBuildInformation sends useful build information to the log.
func PrintBuildInformation() {
	log.Info().Msgf("Build information: name=%s; version=%s; revision=%s", ExtensionName, Version, Revision)
}

// GetSemverVersionStringOrUnknown returns a version string that you can use within action and type definitions.
func GetSemverVersionStringOrUnknown() string {
	if semverPattern.MatchString(Version) {
		return Version
	}
	return "unknown"
}
