package extruntime

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveOOMScoreAdj(t *testing.T) {
	tests := []struct {
		name     string
		envSet   bool
		envValue string
		want     int
	}{
		{name: "uses default when unset", envSet: false, want: defaultOOMScoreAdj},
		{name: "uses valid negative value", envSet: true, envValue: "-997", want: -997},
		{name: "uses valid positive value", envSet: true, envValue: "500", want: 500},
		{name: "clamps below minimum", envSet: true, envValue: "-5000", want: minOOMScoreAdj},
		{name: "clamps above maximum", envSet: true, envValue: "5000", want: maxOOMScoreAdj},
		{name: "falls back on non-numeric", envSet: true, envValue: "not-a-number", want: defaultOOMScoreAdj},
		{name: "falls back on empty value", envSet: true, envValue: "", want: defaultOOMScoreAdj},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envSet {
				t.Setenv(envOOMScoreAdj, tt.envValue)
			} else {
				assert.NoError(t, os.Unsetenv(envOOMScoreAdj))
			}

			assert.Equal(t, tt.want, resolveOOMScoreAdj())
		})
	}
}
