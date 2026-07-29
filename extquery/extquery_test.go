// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

package extquery

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The accept/reject cases below mirror com.steadybit.ql.QueryLanguageTest in the platform, which is
// the behavioural ground truth for the grammar. Keep them in sync when that test changes: the two
// parsers are generated from the same grammar and must agree on what parses.
func TestValidate_accepts(t *testing.T) {
	for _, query := range []string{
		// empty query — means "match everything"
		"",
		"   ",
		"\n\t ",

		// compare clauses, quoted and unquoted
		`key="value"`,
		"key=value",
		"name=*peter",
		"name!=peter",
		"name!=*peter",
		"key~value",
		"key~*value",
		"key!~value",
		"key!~*value",

		// quoted keys and escaping
		`"foo=bar"=true`,
		`key="\"escaped value\""`,
		`question~"\o/ what is going on my \"dude\"?"`,
		`question~*"\o/ what is going on my \"dude\"?"`,

		// presence
		"key IS PRESENT",
		"key IS NOT PRESENT",

		// count
		"count(foo)=11",
		"count(foo)!=11",
		"count(foo)>11",
		"count(foo)>=11",
		"count(foo)<11",
		"count(foo)<=11",

		// membership
		`foo IN (echo, "bar", baz)`,

		// boolean composition, precedence and parentheses
		`name=peter AND age=42 OR type="premium" AND foo=bar`,
		`name=peter AND (age=42 OR type="premium" AND foo=bar)`,
		"a=a AND b=b OR c=c OR d=d AND e=e",
		"a=a AND b=b OR (c=c OR d=d) AND e=e",
		"(a=a)",
		"name~peter AND age=42",
		"name~peter OR age=42",
		"NOT kong.instance.name=staging",
		"NOT (k8s.cluster=prod OR k8s.cluster=stag)",

		// lower-case and symbolic keyword spellings
		"a=a and b=b",
		"a=a && b=b",
		"a=a or b=b",
		"a=a || b=b",
		"not a=a",
		"!a=a",
		"key is present",
		"key is not present",
		"foo in (a, b)",
		"foo not in (a, b)",
		"COUNT(foo)=1",
	} {
		t.Run(query, func(t *testing.T) {
			assert.NoError(t, Validate(query))
		})
	}
}

func TestValidate_rejects(t *testing.T) {
	for _, query := range []string{
		// count only accepts the ordering operators, and only on count()
		"count(key) IS PRESENT NOT",
		"count(key)~value",
		"count(key=value)=1",
		"key>1",
		"key>=1",
		"key<1",
		"key<=1",
		"key<=value",

		// presence keyword misuse
		"key IS PRESENT NOT",
		"key NOT IS PRESENT",
		"key PRESENT",
		"key NOT PRESENT",
		"key IS NOT",
		"IS NOT (k8s.cluster=prod OR k8s.cluster=stag)",

		// structural
		"a=a (b=b)",
		"(a=a",
		"a=a AND",
		"AND a=a",
		"=value",
		`key="unterminated`,
	} {
		t.Run(query, func(t *testing.T) {
			var parseErr *ParseError
			require.ErrorAs(t, Validate(query), &parseErr, "expected a *ParseError")
			assert.NotEmpty(t, parseErr.Message)
		})
	}
}

// Markers parse fine — the grammar accepts them in value position — but the platform resolves them
// against state an extension cannot see, so a query shipped inside an extension must not use them.
// Both the bare and the quoted spelling are rejected; they are the same authoring mistake.
func TestValidate_rejectsPlatformMarkers(t *testing.T) {
	for _, query := range []string{
		"k8s.deployment={{deployment}}",
		"k8s.deployment IN ({{deployment}})",
		`k8s.deployment IN ("{{deployment}}")`,
		"k8s.deployment=[[deployment]]",
		"k8s.deployment IN ([[deployment]])",
		`k8s.deployment IN ("[[deployment]]")`,
		"foo IN (gateway, {{deployment}}, [[service]])",
		`prefix="value-{{suffix}}"`,
	} {
		t.Run(query, func(t *testing.T) {
			var parseErr *ParseError
			require.ErrorAs(t, Validate(query), &parseErr)
			assert.Contains(t, parseErr.Error(), "resolved by the platform")
		})
	}
}

// A value that merely looks marker-ish is not a marker — the key charset comes from the grammar's
// MARKER_CHAR fragment, so brackets around anything else stay ordinary text.
func TestValidate_allowsNonMarkerBraces(t *testing.T) {
	assert.NoError(t, Validate(`json~"{{}"`))
	assert.NoError(t, Validate(`json~"{{ spaced }}"`))
}

func TestValidate_reportsPosition(t *testing.T) {
	t.Run("single line", func(t *testing.T) {
		var parseErr *ParseError
		require.ErrorAs(t, Validate("a=a (b=b)"), &parseErr)
		assert.Equal(t, 1, parseErr.Line)
		assert.Equal(t, 4, parseErr.Column, "column is 0-based and points at the unexpected '('")
	})

	t.Run("marker on a later line", func(t *testing.T) {
		var parseErr *ParseError
		require.ErrorAs(t, Validate("a=a\nAND b={{c}}"), &parseErr)
		assert.Equal(t, 2, parseErr.Line)
		assert.Equal(t, 6, parseErr.Column)
	})

	t.Run("error message carries the position", func(t *testing.T) {
		err := Validate("a=a (b=b)")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "line 1 column 4")
	})
}

func TestValidateAll(t *testing.T) {
	t.Run("reports every failure", func(t *testing.T) {
		err := ValidateAll("a=a", "b>>b", "", "c={{d}}")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "line 1 column 2") // b>>b
		assert.Contains(t, err.Error(), "resolved by the platform")
	})

	t.Run("no queries and only valid queries pass", func(t *testing.T) {
		assert.NoError(t, ValidateAll())
		assert.NoError(t, ValidateAll("a=a", "", "b IS PRESENT"))
	})
}

// Queries taken verbatim from shipped extension definitions — the artifacts this package exists to
// validate. `target.type` is a platform pseudo-attribute and lexes as an ordinary field name.
func TestValidate_realExtensionQueries(t *testing.T) {
	for name, query := range map[string]string{
		"k8s advice applicable": `(target.type="com.steadybit.extension_kubernetes.kubernetes-daemonset" OR target.type="com.steadybit.extension_kubernetes.kubernetes-deployment" OR target.type="com.steadybit.extension_kubernetes.kubernetes-statefulset") AND k8s.specification.probes.summary IS PRESENT`,
		"k8s advice quoted key": `(target.type="com.steadybit.extension_kubernetes.kubernetes-deployment") AND "k8s.label.topology.kubernetes.io/zone" IS PRESENT`,
		"k8s advice action":     `k8s.specification.probes.summary!="OK"`,
		"k8s advice boolean":    `k8s.specification.has-host-podantiaffinity="false"`,
		"k8s presence only":     "k8s.container.spec.request.memory.not-set IS PRESENT",
		"selection template":    `k8s.cluster-name="production"`,
		"aws style":             `aws.account="123456789012" AND aws.zone.id="euc1-az1"`,
	} {
		t.Run(name, func(t *testing.T) {
			assert.NoError(t, Validate(query))
		})
	}
}
