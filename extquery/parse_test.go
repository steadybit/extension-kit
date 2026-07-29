// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

package extquery

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The behavioural corpus lives in testdata/conformance.json and is driven by TestConformance. What
// follows covers the parts of the API and the semantics that the corpus cannot express, because
// they are either not reachable through query syntax or not about matching at all.

func TestParse_rejectsWhatValidateRejects(t *testing.T) {
	for _, query := range []string{
		"a=a (b=b)",
		"key IS NOT",
		"count(key)~value",
		"k8s.deployment={{deployment}}",
		`k8s.deployment IN ("[[deployment]]")`,
	} {
		t.Run(query, func(t *testing.T) {
			predicate, err := Parse(query)
			var parseErr *ParseError
			require.ErrorAs(t, err, &parseErr)
			assert.Nil(t, predicate, "no predicate should be returned alongside an error")
			assert.Equal(t, err, Validate(query), "Parse and Validate must agree on what is invalid")
		})
	}
}

// An empty value list is unreachable through the grammar — `IN ()` does not parse — but the rule
// still has to hold, because the platform can build such a predicate structurally and the
// implementations must agree on what it means.
func TestEmptyValueList(t *testing.T) {
	assert.False(t, valuePredicate{key: "a", operator: opEquals}.Matches(MapAttributes{"a": {"x"}}),
		"a positive membership over an empty set matches nothing")
	assert.True(t, valuePredicate{key: "a", operator: opNotEquals}.Matches(MapAttributes{"a": {"x"}}),
		"a negated membership over an empty set excludes nothing")
}

// Likewise unreachable: a composite with no children. An empty query lowers to alwaysPredicate, but
// the platform's structural form is an empty AND — and an empty OR is match-all too, not a
// contradiction.
func TestEmptyComposite(t *testing.T) {
	assert.True(t, compositePredicate{and: true}.Matches(MapAttributes{}))
	assert.True(t, compositePredicate{and: false}.Matches(MapAttributes{}), "an empty OR is match-all, not false")
}

func TestParse_emptyQueryMatchesEverything(t *testing.T) {
	for _, query := range []string{"", "   ", "\n\t "} {
		predicate, err := Parse(query)
		require.NoError(t, err)
		assert.True(t, predicate.Matches(MapAttributes{}))
		assert.True(t, predicate.Matches(MapAttributes{"anything": {"at-all"}}))
	}
}

func TestAttributesAdapters(t *testing.T) {
	predicate, err := Parse(`target.type="host" AND host.hostname~"web"`)
	require.NoError(t, err)

	t.Run("MapAttributes", func(t *testing.T) {
		assert.True(t, predicate.Matches(MapAttributes{
			"target.type":   {"host"},
			"host.hostname": {"web-01"},
		}))
	})

	t.Run("MapAttributes returns nil for an absent key", func(t *testing.T) {
		assert.Nil(t, MapAttributes{}.Values("nope"))
	})

	// The lazy adapter is how a kit injects synthetic keys (target type, label) without copying the
	// discovered attribute map for every target.
	t.Run("AttributesFunc can synthesize keys", func(t *testing.T) {
		discovered := map[string][]string{"host.hostname": {"web-01"}}
		attributes := AttributesFunc(func(key string) []string {
			if key == "target.type" {
				return []string{"host"}
			}
			return discovered[key]
		})
		assert.True(t, predicate.Matches(attributes))
	})
}

func TestMustParse(t *testing.T) {
	assert.NotNil(t, MustParse(`a="b"`))
	assert.Panics(t, func() { MustParse("a=a (b=b)") })
}

// A parsed predicate is reused across every target, so it must not accumulate state.
func TestPredicateIsReusable(t *testing.T) {
	predicate, err := Parse(`k8s.namespace="prod"`)
	require.NoError(t, err)
	for range 3 {
		assert.True(t, predicate.Matches(MapAttributes{"k8s.namespace": {"prod"}}))
		assert.False(t, predicate.Matches(MapAttributes{"k8s.namespace": {"dev"}}))
	}
}

func TestPredicateIsSafeForConcurrentUse(t *testing.T) {
	predicate, err := Parse(`k8s.namespace IN (prod, staging) AND NOT k8s.label.tier="canary"`)
	require.NoError(t, err)

	done := make(chan bool, 32)
	for i := range 32 {
		go func() {
			done <- predicate.Matches(MapAttributes{
				"k8s.namespace":  {"prod"},
				"k8s.label.tier": {"stable"},
				"index":          {string(rune('a' + i%26))},
			})
		}()
	}
	for range 32 {
		assert.True(t, <-done)
	}
}

// String() is for logging a parsed filter at start up, so an operator can see what the extension
// understood. It is not a serializer and does not round-trip to the input.
func TestPredicateString(t *testing.T) {
	for query, want := range map[string]string{
		`a="b"`:            "a = (b)",
		"a IN (b, c)":      "a = (b OR c)",
		"a NOT IN (b)":     "a != (b)",
		"a IS PRESENT":     "a IS PRESENT",
		"a IS NOT PRESENT": "a IS NOT PRESENT",
		"count(a)>=2":      "count(a) >= 2",
		"NOT a=b":          "NOT (a = (b))",
		"a=1 AND b=2":      "(a = (1) AND b = (2))",
		"a=1 OR b=2":       "(a = (1) OR b = (2))",
		"a~*x":             "a ~* (x)",
		"":                 "true",
	} {
		t.Run(query, func(t *testing.T) {
			predicate, err := Parse(query)
			require.NoError(t, err)
			assert.Equal(t, want, predicate.String())
		})
	}
}
