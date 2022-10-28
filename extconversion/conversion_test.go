// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

package extconversion

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type StructWithTime struct {
	End time.Time `json:"end"`
}

func TestConversionOfTime(t *testing.T) {
	// Given
	const shortForm = "2006-Jan-02"
	end, err := time.Parse(shortForm, "2013-Feb-03")
	require.NoError(t, err)
	input := StructWithTime{End: end}

	// When
	var intermediate map[string]interface{}
	err = Convert(input, &intermediate)
	require.NoError(t, err)
	require.Equal(t, "2013-02-03T00:00:00Z", intermediate["end"])

	var result StructWithTime
	err = Convert(intermediate, &result)
	require.NoError(t, err)

	// Then
	require.Equal(t, input.End, result.End)
}
