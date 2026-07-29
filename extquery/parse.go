// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

package extquery

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/steadybit/extension-kit/extquery/internal/gen"
)

// Parse compiles a query into a Predicate.
//
// The accepted syntax is exactly what Validate accepts, and the same restrictions apply — see its
// documentation for empty queries and for why platform markers are rejected.
//
// Parsing is the expensive part; matching is not. Parse once at start up and reuse the Predicate
// for every target.
func Parse(query string) (Predicate, error) {
	tree, err := parseTree(query)
	if err != nil {
		return nil, err
	}
	if tree == nil || tree.Query() == nil {
		// An empty query means "match everything", consistent with the platform.
		return alwaysPredicate{}, nil
	}
	return buildQuery(tree.Query())
}

// MustParse is Parse for queries that are compile-time constants, panicking on a bad one. Use it
// only for literals in package-level variables, never for anything derived from configuration.
func MustParse(query string) Predicate {
	predicate, err := Parse(query)
	if err != nil {
		panic(err)
	}
	return predicate
}

// parseTree runs the marker check and the parser, returning nil for an empty query.
func parseTree(query string) (gen.ITopLevelQueryContext, error) {
	if strings.TrimSpace(query) == "" {
		return nil, nil
	}
	if err := rejectMarkers(query); err != nil {
		return nil, err
	}

	listener := &firstErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}

	lexer := gen.NewQueryLanguageLexer(antlr.NewInputStream(query))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	parser := gen.NewQueryLanguageParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)
	tree := parser.TopLevelQuery()

	if listener.err != nil {
		return nil, listener.err
	}
	return tree, nil
}

// buildQuery lowers a `query` node. It mirrors TargetPredicateQueryLanguageVisitor#visitQuery in
// the platform, including the order of the checks — NOT is tested before AND/OR because a
// parenthesised child is itself a `query` node.
func buildQuery(ctx gen.IQueryContext) (Predicate, error) {
	if ctx.NOT() != nil {
		inner, err := buildQuery(ctx.Query(0))
		if err != nil {
			return nil, err
		}
		return notPredicate{inner: inner}, nil
	}

	if ctx.AND() != nil || ctx.OR() != nil {
		children := ctx.AllQuery()
		predicates := make([]Predicate, 0, len(children))
		for _, child := range children {
			predicate, err := buildQuery(child)
			if err != nil {
				return nil, err
			}
			predicates = append(predicates, predicate)
		}
		return compositePredicate{and: ctx.AND() != nil, predicates: predicates}, nil
	}

	if ctx.Clause() != nil {
		return buildClause(ctx.Clause())
	}

	// A parenthesised query: `( query )`.
	if children := ctx.AllQuery(); len(children) == 1 {
		return buildQuery(children[0])
	}

	return nil, &ParseError{Message: "unsupported query structure", Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn()}
}

func buildClause(ctx gen.IClauseContext) (Predicate, error) {
	switch {
	case ctx.IsPresentClause() != nil:
		c := ctx.IsPresentClause()
		return presencePredicate{key: fieldName(c.FieldName()), present: c.NOT() == nil}, nil

	case ctx.InClause() != nil:
		c := ctx.InClause()
		operator := opEquals
		if c.NOT() != nil {
			operator = opNotEquals
		}
		return valuePredicate{key: fieldName(c.FieldName()), operator: operator, values: valueList(c.Valuelist())}, nil

	case ctx.CompareClause() != nil:
		c := ctx.CompareClause()
		operator, err := compareOperatorOf(c.GetOp())
		if err != nil {
			return nil, err
		}
		return valuePredicate{key: fieldName(c.FieldName()), operator: operator, values: []string{value(c.Value())}}, nil

	case ctx.CountClause() != nil:
		c := ctx.CountClause()
		operator, err := countOperatorOf(c.GetOp())
		if err != nil {
			return nil, err
		}
		return countPredicate{key: fieldName(c.FieldName()), operator: operator, value: c.Int_value().GetText()}, nil
	}

	return nil, &ParseError{Message: "unsupported clause", Line: ctx.GetStart().GetLine(), Column: ctx.GetStart().GetColumn()}
}

func compareOperatorOf(token antlr.Token) (compareOperator, error) {
	switch token.GetTokenType() {
	case gen.QueryLanguageParserOP_EQUAL:
		return opEquals, nil
	case gen.QueryLanguageParserOP_NOT_EQUAL:
		return opNotEquals, nil
	case gen.QueryLanguageParserOP_EQUAL_IGNORE_CASE:
		return opEqualsIgnoreCase, nil
	case gen.QueryLanguageParserOP_NOT_EQUAL_IGNORE_CASE:
		return opNotEqualsIgnoreCase, nil
	case gen.QueryLanguageParserOP_TILDE:
		return opContains, nil
	case gen.QueryLanguageParserOP_NOT_TILDE:
		return opNotContains, nil
	case gen.QueryLanguageParserOP_TILDE_IGNORE_CASE:
		return opContainsIgnoreCase, nil
	case gen.QueryLanguageParserOP_NOT_TILDE_IGNORE_CASE:
		return opNotContainsIgnoreCase, nil
	}
	return 0, unexpectedOperator(token)
}

func countOperatorOf(token antlr.Token) (countOperator, error) {
	switch token.GetTokenType() {
	case gen.QueryLanguageParserOP_EQUAL:
		return countEqual, nil
	case gen.QueryLanguageParserOP_NOT_EQUAL:
		return countNotEqual, nil
	case gen.QueryLanguageParserOP_GREATER_THAN:
		return countGreaterThan, nil
	case gen.QueryLanguageParserOP_GREATER_THAN_EQUAL:
		return countGreaterThanOrEqual, nil
	case gen.QueryLanguageParserOP_LESS_THAN:
		return countLessThan, nil
	case gen.QueryLanguageParserOP_LESS_THAN_EQUAL:
		return countLessThanOrEqual, nil
	}
	return 0, unexpectedOperator(token)
}

func unexpectedOperator(token antlr.Token) error {
	return &ParseError{
		Message: fmt.Sprintf("unexpected operator %q", token.GetText()),
		Line:    token.GetLine(),
		Column:  token.GetColumn(),
	}
}

func valueList(ctx gen.IValuelistContext) []string {
	values := ctx.AllValue()
	out := make([]string, len(values))
	for i, v := range values {
		out[i] = value(v)
	}
	return out
}

func value(ctx gen.IValueContext) string {
	return content(ctx.GetText(), ctx.QUOTED() != nil)
}

func fieldName(ctx gen.IFieldNameContext) string {
	return content(ctx.GetText(), ctx.QUOTED() != nil)
}

// content strips the quotes of a QUOTED token and unescapes the value, mirroring
// TargetPredicateQueryLanguageVisitor#getContent / #unescapeValue.
func content(text string, quoted bool) string {
	if quoted && len(text) >= 2 {
		text = text[1 : len(text)-1]
	}
	text = strings.ReplaceAll(text, `\"`, `"`)
	return strings.ReplaceAll(text, `\\`, `\`)
}
