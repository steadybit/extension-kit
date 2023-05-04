// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

// Package extlogging contains general utilities for extension logging. We
// recommend that extensions leverage zerolog.
package extlogging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

// InitZeroLog configures the zerolog logging output in a standardized way. More specifically, it configures the output to be sent to stderr,
// a human-readable output format, the time format and the global log level.
func InitZeroLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logFormat := os.Getenv("STEADYBIT_LOG_FORMAT")
	log.Logger = log.With().Caller().Logger()
	if strings.ToLower(logFormat) != "json" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	logLevel := os.Getenv("STEADYBIT_LOG_LEVEL")
	if len(logLevel) == 0 {
		logLevel = "info"
	}
	level, err := zerolog.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		log.Panic().Msgf("Unsupported log level defined via environment variable: %s\n", logLevel)
	}
	zerolog.SetGlobalLevel(level)
}
