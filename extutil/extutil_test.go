package extutil

import (
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
