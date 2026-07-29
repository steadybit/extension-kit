// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

package extquery

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestConformance drives testdata/conformance.json — the corpus shared with the platform's Java and
// SQL evaluators. It is the mechanism that keeps three implementations of one language in
// agreement; see the $comment in that file before changing it, and mirror any change in the
// platform's copy.
func TestConformance(t *testing.T) {
	var corpus struct {
		Cases []struct {
			Name       string              `json:"name"`
			Query      string              `json:"query"`
			Attributes map[string][]string `json:"attributes"`
			Matches    bool                `json:"matches"`
		} `json:"cases"`
	}

	raw, err := os.ReadFile("testdata/conformance.json")
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(raw, &corpus))
	require.NotEmpty(t, corpus.Cases, "conformance corpus is empty")

	seen := make(map[string]struct{}, len(corpus.Cases))
	for _, c := range corpus.Cases {
		require.NotEmpty(t, c.Name, "every case needs a name")
		_, duplicate := seen[c.Name]
		require.False(t, duplicate, "duplicate case name %q", c.Name)
		seen[c.Name] = struct{}{}

		t.Run(c.Name, func(t *testing.T) {
			predicate, err := Parse(c.Query)
			require.NoError(t, err, "query: %s", c.Query)
			assert.Equal(t, c.Matches, predicate.Matches(MapAttributes(c.Attributes)),
				"query: %s\nattributes: %v", c.Query, c.Attributes)
		})
	}
}
