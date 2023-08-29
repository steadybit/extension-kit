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

const RFC3339Micro = "2006-01-02T15:04:05.999Z07:00"

// InitZeroLog configures the zerolog logging output in a standardized way. More specifically, it configures the output to be sent to stderr,
// a human-readable output format, the time format and the global log level.
func InitZeroLog() {
	zerolog.TimeFieldFormat = RFC3339Micro

	var logger zerolog.Logger
	if strings.ToLower(os.Getenv("STEADYBIT_LOG_FORMAT")) != "json" {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: getNoColor(), TimeFormat: RFC3339Micro})
	} else {
		logger = zerolog.New(os.Stderr)
	}

	level := getLogLevel()
	zerolog.SetGlobalLevel(level)

	c := logger.With().Timestamp()
	if level == zerolog.DebugLevel {
		c = c.Caller()
	}
	log.Logger = c.Logger()
}

func getLogLevel() zerolog.Level {
	logLevel := os.Getenv("STEADYBIT_LOG_LEVEL")
	if len(logLevel) == 0 {
		logLevel = "info"
	}
	level, err := zerolog.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		log.Panic().Msgf("Unsupported log level defined via environment variable: %s\n", logLevel)
	}
	return level
}

func getNoColor() bool {
	switch strings.ToLower(os.Getenv("STEADYBIT_LOG_COLOR")) {
	case "true":
		return true
	case "false":
		return false
	default:
		if stat, err := os.Stderr.Stat(); err == nil {
			return (stat.Mode() & os.ModeCharDevice) != os.ModeCharDevice // check if stderr is not a terminal
		}
		return false
	}
}
