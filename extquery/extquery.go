// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

// Package extquery parses Steadybit's target query language — the language used for environment
// scopes, experiment blast radii, action target selections and advice assessment queries, e.g.
//
//	k8s.namespace="prod" AND NOT k8s.label.tier="canary"
//
// It is generated from the same ANTLR grammar the platform uses, so a query accepted here is
// accepted there. See grammar/ and generate.sh.
//
// Today the package only offers syntax validation. Extensions embed query strings in their action
// and advice definitions, where a malformed query is currently only discovered by the platform at
// registration time — against an already-released artifact. Validate lets an extension catch that
// in its own tests instead.
package extquery

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// markerPattern matches a variable ({{key}}) or template-placeholder ([[key]]) marker. The key
// charset mirrors the grammar's MARKER_CHAR fragment. Deliberately applied to the raw query rather
// than to VARIABLE/PLACEHOLDER tokens, so the quoted form ("{{key}}", which lexes as QUOTED) is
// caught too — it is the same authoring mistake.
var markerPattern = regexp.MustCompile(`\{\{[a-zA-Z0-9_.\-]+}}|\[\[[a-zA-Z0-9_.\-]+]]`)

// ParseError describes why a query could not be parsed. Line is 1-based and Column is 0-based,
// matching what ANTLR reports and what the platform surfaces in its own validation messages.
type ParseError struct {
	Message string
	Line    int
	Column  int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("failed to parse query at line %d column %d: %s", e.Line, e.Column, e.Message)
}

// Validate reports whether query is syntactically valid, returning a *ParseError if it is not.
//
// An empty query is valid and means "match everything", consistent with the platform.
//
// Variable ({{key}}) and template-placeholder ([[key]]) markers are rejected, quoted or not: the
// platform resolves them against experiment and environment state that an extension has no access
// to, so a query shipped inside an extension can never legitimately contain one.
//
// Only syntax is checked. Whether the referenced attributes exist, or whether the query selects
// anything, is not knowable here.
func Validate(query string) error {
	_, err := parseTree(query)
	return err
}

// rejectMarkers fails a query containing a variable or template-placeholder marker.
func rejectMarkers(query string) error {
	loc := markerPattern.FindStringIndex(query)
	if loc == nil {
		return nil
	}
	line, column := position(query, loc[0])
	return &ParseError{
		Message: fmt.Sprintf("%s is resolved by the platform and cannot be used in a query defined by an extension", query[loc[0]:loc[1]]),
		Line:    line,
		Column:  column,
	}
}

// ValidateAll validates every query and joins the failures, so a caller checking a whole
// definition reports all of its bad queries at once rather than one per run.
func ValidateAll(queries ...string) error {
	var errs []error
	for _, query := range queries {
		if err := Validate(query); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

// position converts a byte offset into the 1-based line and 0-based column ANTLR would report.
func position(s string, offset int) (line, column int) {
	line = 1 + strings.Count(s[:offset], "\n")
	if lastNewline := strings.LastIndex(s[:offset], "\n"); lastNewline >= 0 {
		return line, offset - lastNewline - 1
	}
	return line, offset
}

// firstErrorListener keeps only the first syntax error. ANTLR's default error strategy
// resynchronizes and goes on to report cascading follow-up errors for a single real mistake; the
// first one is the one that points at it.
type firstErrorListener struct {
	*antlr.DefaultErrorListener
	err *ParseError
}

func (l *firstErrorListener) SyntaxError(_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException) {
	if l.err == nil {
		l.err = &ParseError{Message: msg, Line: line, Column: column}
	}
}
