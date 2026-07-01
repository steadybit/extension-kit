// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2024 Steadybit GmbH

package extutil

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestToInt64(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "int 1 should return 1", args: args{val: 1}, want: 1},
		{name: "int32 1 should return 1", args: args{val: int32(1)}, want: 1},
		{name: "int64 1 should return 1", args: args{val: int64(1)}, want: 1},
		{name: "float32 1.0 should return 1", args: args{val: float32(1.0)}, want: 1},
		{name: "float64 1.0 should return 1", args: args{val: float64(1.0)}, want: 1},
		{name: "string '1' should return 1", args: args{val: "1"}, want: 1},
		{name: "string 'anything' should return 0", args: args{val: "anything"}, want: 0},
		{name: "nil should return 0", args: args{val: nil}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64(tt.args.val); got != tt.want {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUInt64(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "int 1 should return 1", args: args{val: 1}, want: 1},
		{name: "int32 1 should return 1", args: args{val: int32(1)}, want: 1},
		{name: "int64 1 should return 1", args: args{val: int64(1)}, want: 1},
		{name: "float32 1.0 should return 1", args: args{val: float32(1.0)}, want: 1},
		{name: "float64 1.0 should return 1", args: args{val: float64(1.0)}, want: 1},
		{name: "string '1' should return 1", args: args{val: "1"}, want: 1},
		{name: "string 'anything' should return 0", args: args{val: "anything"}, want: 0},
		{name: "nil should return 0", args: args{val: nil}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUInt64(tt.args.val); got != tt.want {
				t.Errorf("ToUInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUInt(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{name: "int 1 should return 1", args: args{val: 1}, want: 1},
		{name: "int32 1 should return 1", args: args{val: int32(1)}, want: 1},
		{name: "int64 1 should return 1", args: args{val: int64(1)}, want: 1},
		{name: "float32 1.0 should return 1", args: args{val: float32(1.0)}, want: 1},
		{name: "float64 1.0 should return 1", args: args{val: float64(1.0)}, want: 1},
		{name: "string '1' should return 1", args: args{val: "1"}, want: 1},
		{name: "string 'anything' should return 0", args: args{val: "anything"}, want: 0},
		{name: "nil should return 0", args: args{val: nil}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUInt(tt.args.val); got != tt.want {
				t.Errorf("ToUInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "int 1 should return 1", args: args{val: 1}, want: 1},
		{name: "int32 1 should return 1", args: args{val: int32(1)}, want: 1},
		{name: "int64 1 should return 1", args: args{val: int64(1)}, want: 1},
		{name: "float32 1.0 should return 1", args: args{val: float32(1.0)}, want: 1},
		{name: "float64 1.0 should return 1", args: args{val: float64(1.0)}, want: 1},
		{name: "string '1' should return 1", args: args{val: "1"}, want: 1},
		{name: "string 'anything' should return 0", args: args{val: "anything"}, want: 0},
		{name: "nil should return 0", args: args{val: nil}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.val); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "int 1 should return 1", args: args{val: 1}, want: 1},
		{name: "int32 1 should return 1", args: args{val: int32(1)}, want: 1},
		{name: "int64 1 should return 1", args: args{val: int64(1)}, want: 1},
		{name: "float32 1.0 should return 1", args: args{val: float32(1.0)}, want: 1},
		{name: "float64 1.0 should return 1", args: args{val: float64(1.0)}, want: 1},
		{name: "string '1' should return 1", args: args{val: "1"}, want: 1},
		{name: "string 'anything' should return 0", args: args{val: "anything"}, want: 0},
		{name: "nil should return 0", args: args{val: nil}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt32(tt.args.val); got != tt.want {
				t.Errorf("ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestToString(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "string 'anything' should return 'anything'", args: args{val: "anything"}, want: "anything"},
		{name: "nil should return ''", args: args{val: nil}, want: ""},
		{name: "non-string number should return '' (not panic)", args: args{val: float64(42)}, want: ""},
		{name: "non-string bool should return '' (not panic)", args: args{val: true}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.val); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "string 'true' should return true", args: args{val: "true"}, want: true},
		{name: "boolean true should return true", args: args{val: true}, want: true},
		{name: "string 'anything' should return false", args: args{val: "anything"}, want: false},
		{name: "string '' should return false", args: args{val: ""}, want: false},
		{name: "nil should return false", args: args{val: nil}, want: false},
		{name: "non-string non-bool number should return false (not panic)", args: args{val: float64(1)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBool(tt.args.val); got != tt.want {
				t.Errorf("ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToKeyValue(t *testing.T) {
	value, err := ToKeyValue(map[string]interface{}{"headers": []any{
		map[string]any{"key": "testKey", "value": "testValue"},
	}}, "headers")
	require.NoError(t, err)
	require.Equal(t, "testValue", value["testKey"])
}

func TestToToStringArray(t *testing.T) {
	value := ToStringArray([]any{"key", "testKey", "value", "testValue"})
	require.Equal(t, []string([]string{"key", "testKey", "value", "testValue"}), value)
}

func TestToKeyValue_malformed_returns_error_not_panic(t *testing.T) {
	// value that isn't a key/value array
	_, err := ToKeyValue(map[string]interface{}{"headers": "not-an-array"}, "headers")
	require.Error(t, err)

	// entry with a non-string value must return an error, not panic
	_, err = ToKeyValue(map[string]interface{}{"headers": []any{
		map[string]any{"key": "k", "value": 42},
	}}, "headers")
	require.Error(t, err)

	// entry missing the value key must return an error, not panic
	_, err = ToKeyValue(map[string]interface{}{"headers": []any{
		map[string]any{"key": "k"},
	}}, "headers")
	require.Error(t, err)
}

func TestToStringArray_malformed_returns_nil_or_skips_not_panic(t *testing.T) {
	assert.Nil(t, ToStringArray(nil))
	assert.Nil(t, ToStringArray("not-a-slice"))
	// non-string elements are skipped rather than panicking
	assert.Equal(t, []string{"a", "b"}, ToStringArray([]any{"a", 42, "b", true}))
}

func TestMaskString(t *testing.T) {
	type args struct {
		s         string
		search    string
		remaining int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "should mask", args: args{
			s:         "command --apiKey=123456 --fast",
			search:    "123456",
			remaining: 0},
			want: "command --apiKey=****** --fast"},
		{name: "should mask with remaining runes", args: args{
			s:         "command --apiKey=123456 --fast",
			search:    "123456",
			remaining: 2},
			want: "command --apiKey=****56 --fast"},
		{name: "should not fail if remaining is too large", args: args{
			s:         "command --apiKey=123456 --fast",
			search:    "123456",
			remaining: 10},
			want: "command --apiKey=123456 --fast"},
		{name: "should ignore if search not present", args: args{
			s:         "command --fast",
			search:    "123456",
			remaining: 0},
			want: "command --fast"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaskString(tt.args.s, tt.args.search, tt.args.remaining); got != tt.want {
				t.Errorf("MaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustHaveValue(t *testing.T) {
	assert.Equal(t, 0, MustHaveValue(map[string]int{"key": 0}, "key"))
	assert.Panics(t, func() {
		MustHaveValue(map[string]int{"key": 0}, "missing")
	})

	assert.Equal(t, "value", MustHaveValue(map[string]string{"key": "value"}, "key"))
	assert.Panics(t, func() {
		MustHaveValue(map[string]string{"key": "value"}, "missing")
	})

	assert.Equal(t, Ptr("value"), MustHaveValue(map[string]*string{"key": Ptr("value")}, "key"))
	assert.Panics(t, func() {
		MustHaveValue(map[string]*string{"empty": nil}, "empty")
	})

	assert.Equal(t, []string{"value"}, MustHaveValue(map[string][]string{"key": {"value"}}, "key"))
	assert.Panics(t, func() {
		MustHaveValue(map[string][]string{"empty": {}}, "empty")
	})
}
