// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

package extquery

import (
	"fmt"
	"strconv"
	"strings"
)

// Attributes is the view a Predicate needs of the thing it matches: a multi-valued attribute map.
// Values returns the values for a key, or nil when the attribute is absent — the two are treated
// the same by every operator.
//
// Kits adapt their own types to this. extension-kit deliberately does not import discovery-kit
// (the dependency runs the other way), so a Target adapter belongs there, not here.
type Attributes interface {
	Values(key string) []string
}

// MapAttributes adapts a plain attribute map, which is the shape discovery data already has.
type MapAttributes map[string][]string

func (m MapAttributes) Values(key string) []string { return m[key] }

// AttributesFunc adapts a lookup function, for callers that resolve attributes lazily or need to
// inject synthetic keys (e.g. target type or label) without copying the map.
type AttributesFunc func(key string) []string

func (f AttributesFunc) Values(key string) []string { return f(key) }

// Predicate is a parsed query, ready to match. Predicates are immutable and safe for concurrent
// use, so parse once at start up and reuse per target.
type Predicate interface {
	Matches(a Attributes) bool
	String() string
}

// compareOperator is the operator of a compare or membership clause.
type compareOperator int

const (
	opEquals compareOperator = iota
	opNotEquals
	opContains
	opNotContains
	opEqualsIgnoreCase
	opNotEqualsIgnoreCase
	opContainsIgnoreCase
	opNotContainsIgnoreCase
)

// negated reports whether this is a NOT_* comparison. It is what gives an empty value list its
// logical meaning: a positive membership over an empty set matches nothing, a negated one excludes
// nothing and therefore matches everything.
func (o compareOperator) negated() bool {
	switch o {
	case opNotEquals, opNotContains, opNotEqualsIgnoreCase, opNotContainsIgnoreCase:
		return true
	default:
		return false
	}
}

func (o compareOperator) String() string {
	switch o {
	case opEquals:
		return "="
	case opNotEquals:
		return "!="
	case opContains:
		return "~"
	case opNotContains:
		return "!~"
	case opEqualsIgnoreCase:
		return "=*"
	case opNotEqualsIgnoreCase:
		return "!=*"
	case opContainsIgnoreCase:
		return "~*"
	case opNotContainsIgnoreCase:
		return "!~*"
	default:
		return "?"
	}
}

// match reports whether a single attribute value satisfies the positive form of this operator.
// Negation is applied once, over the whole value list, by valuePredicate.
func (o compareOperator) match(attributeValue, queryValue string) bool {
	switch o {
	case opEquals, opNotEquals:
		return attributeValue == queryValue
	case opContains, opNotContains:
		return strings.Contains(attributeValue, queryValue)
	case opEqualsIgnoreCase, opNotEqualsIgnoreCase:
		return strings.EqualFold(attributeValue, queryValue)
	case opContainsIgnoreCase, opNotContainsIgnoreCase:
		return strings.Contains(strings.ToLower(attributeValue), strings.ToLower(queryValue))
	default:
		return false
	}
}

type countOperator int

const (
	countEqual countOperator = iota
	countNotEqual
	countGreaterThan
	countGreaterThanOrEqual
	countLessThan
	countLessThanOrEqual
)

func (o countOperator) String() string {
	switch o {
	case countEqual:
		return "="
	case countNotEqual:
		return "!="
	case countGreaterThan:
		return ">"
	case countGreaterThanOrEqual:
		return ">="
	case countLessThan:
		return "<"
	case countLessThanOrEqual:
		return "<="
	default:
		return "?"
	}
}

func (o countOperator) match(count, want int) bool {
	switch o {
	case countEqual:
		return count == want
	case countNotEqual:
		return count != want
	case countGreaterThan:
		return count > want
	case countGreaterThanOrEqual:
		return count >= want
	case countLessThan:
		return count < want
	case countLessThanOrEqual:
		return count <= want
	default:
		return false
	}
}

// alwaysPredicate is what an empty query lowers to: match everything.
type alwaysPredicate struct{}

func (alwaysPredicate) Matches(Attributes) bool { return true }
func (alwaysPredicate) String() string          { return "true" }

// valuePredicate backs both `key OP value` and `key [NOT] IN (…)`; a compare clause is just the
// single-value case.
type valuePredicate struct {
	key      string
	operator compareOperator
	values   []string
}

func (p valuePredicate) Matches(a Attributes) bool {
	if len(p.values) == 0 {
		// An empty set-membership: a positive membership matches no target; a negated one
		// excludes nothing and so matches everything.
		return p.operator.negated()
	}
	// Build the positive answer — "any attribute value satisfies any query value" — then negate
	// once for the NOT_* operators. NOT_EQUALS therefore means "no value equals any listed value",
	// not "some value differs".
	attributeValues := a.Values(p.key)
	anyMatch := false
	for _, queryValue := range p.values {
		for _, attributeValue := range attributeValues {
			if p.operator.match(attributeValue, queryValue) {
				anyMatch = true
				break
			}
		}
		if anyMatch {
			break
		}
	}
	if p.operator.negated() {
		return !anyMatch
	}
	return anyMatch
}

func (p valuePredicate) String() string {
	return p.key + " " + p.operator.String() + " (" + strings.Join(p.values, " OR ") + ")"
}

type presencePredicate struct {
	key     string
	present bool
}

func (p presencePredicate) Matches(a Attributes) bool {
	return (len(a.Values(p.key)) > 0) == p.present
}

func (p presencePredicate) String() string {
	if p.present {
		return p.key + " IS PRESENT"
	}
	return p.key + " IS NOT PRESENT"
}

type countPredicate struct {
	key      string
	operator countOperator
	value    string
}

func (p countPredicate) Matches(a Attributes) bool {
	values := a.Values(p.key)
	if len(values) == 0 {
		// An absent attribute is not a count of zero — the clause is simply false, including for
		// the operators where zero would otherwise satisfy it (e.g. `count(k) < 1`).
		return false
	}
	want, err := strconv.Atoi(p.value)
	if err != nil {
		return false
	}
	return p.operator.match(len(values), want)
}

func (p countPredicate) String() string {
	return fmt.Sprintf("count(%s) %s %s", p.key, p.operator, p.value)
}

type notPredicate struct {
	inner Predicate
}

func (p notPredicate) Matches(a Attributes) bool { return !p.inner.Matches(a) }
func (p notPredicate) String() string            { return "NOT (" + p.inner.String() + ")" }

type compositePredicate struct {
	and        bool
	predicates []Predicate
}

func (p compositePredicate) Matches(a Attributes) bool {
	// No children matches everything, for AND *and* for OR. This is what makes an empty query mean
	// match-all, and an empty OR is not treated as a contradiction.
	if len(p.predicates) == 0 {
		return true
	}
	for _, predicate := range p.predicates {
		if predicate.Matches(a) != p.and {
			// AND: a false child decides it. OR: a true child decides it.
			return !p.and
		}
	}
	return p.and
}

func (p compositePredicate) String() string {
	op := " OR "
	if p.and {
		op = " AND "
	}
	parts := make([]string, len(p.predicates))
	for i, predicate := range p.predicates {
		parts[i] = predicate.String()
	}
	return "(" + strings.Join(parts, op) + ")"
}
