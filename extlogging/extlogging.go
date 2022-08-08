// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

// Package extlogging contains general utilities for extension logging. We
// recommend that extensions leverage zerolog.
package extlogging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// InitZeroLog configures the zerolog logging output in a standardized way. More specifically, it configures the output to be sent to stderr,
// a human-readable output format and the time format.
func InitZeroLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
