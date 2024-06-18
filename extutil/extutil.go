// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2024 Steadybit GmbH

// Package extutil contains a variety of util functions that were identified as common code duplication.
// More specialized packages exist for groups of use cases (extlogging and exthttp) for example.
package extutil

import (
	"fmt"
	"github.com/steadybit/extension-kit/extconversion"
	"reflect"
	"strconv"
	"strings"
)

// Ptr returns a pointer to the given value. You will find this helpful when desiring to pass a literal value to a function that requires a pointer.
func Ptr[T any](val T) *T {
	return &val
}

func ToInt64(val interface{}) int64 {
	switch val := val.(type) {
	case int:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	case string:
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0
		}
		return i
	default:
		return 0
	}
}

func ToUInt64(val interface{}) uint64 {
	switch val := val.(type) {
	case int:
		return uint64(val)
	case int32:
		return uint64(val)
	case int64:
		return uint64(val)
	case uint64:
		return val
	case float32:
		return uint64(val)
	case float64:
		return uint64(val)
	case string:
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0
		}
		return uint64(i)
	default:
		return 0
	}
}

func ToInt(val interface{}) int {
	switch val := val.(type) {
	case int:
		return val
	case int32:
		return int(val)
	case int64:
		return int(val)
	case float32:
		return int(val)
	case float64:
		return int(val)
	case string:
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0
		}
		return int(i)
	default:
		return 0
	}
}

func ToInt32(val interface{}) int32 {
	switch val := val.(type) {
	case int:
		return int32(val)
	case int32:
		return val
	case int64:
		return int32(val)
	case float32:
		return int32(val)
	case float64:
		return int32(val)
	case string:
		i, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return 0
		}
		return int32(i)
	}
	return 0

}

func ToString(val interface{}) string {
	if val == nil {
		return ""
	}
	return val.(string)
}

func ToBool(val interface{}) bool {
	if val == nil || val == "" {
		return false
	}
	// parse bool string
	if val, ok := val.(string); ok {
		return val == "true"
	}
	return val.(bool)
}

func ToKeyValue(config map[string]interface{}, configName string) (map[string]string, error) {
	kv, ok := config[configName].([]any)
	if !ok {
		return nil, fmt.Errorf("failed to interpret config value for %s as a key/value array", configName)
	}

	result := make(map[string]string, len(kv))
	for _, rawEntry := range kv {
		entry, ok := rawEntry.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("failed to interpret config value for %s as a key/value array", configName)
		}
		result[entry["key"].(string)] = entry["value"].(string)
	}

	return result, nil
}

func ToUInt(val interface{}) uint {
	switch val := val.(type) {
	case int:
		return uint(val)
	case int32:
		return uint(val)
	case int64:
		return uint(val)
	case float32:
		return uint(val)
	case float64:
		return uint(val)
	case string:
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0
		}
		return uint(i)
	default:
		return 0
	}
}

// MustHaveValue panics if the given key is not present in the map or the value is nil or empty.
func MustHaveValue[T any, K comparable](m map[K]T, key K) T {
	val, ok := m[key]
	if !ok {
		panic(fmt.Sprintf("missing value for '%v'", key))
	}

	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.Array || kind == reflect.Chan || kind == reflect.Map || kind == reflect.Slice || kind == reflect.String {
		if reflect.ValueOf(val).Len() == 0 {
			panic(fmt.Sprintf("value for '%v' is empty ", key))
		}
	} else if kind == reflect.Ptr {
		if reflect.ValueOf(val).IsNil() {
			panic(fmt.Sprintf("value for '%v' is nil ", key))
		}
	}
	return val
}

func ToStringArray(s interface{}) []string {
	if s == nil {
		return nil
	}

	tokens := make([]string, len(s.([]interface{})))
	for i, v := range s.([]interface{}) {
		tokens[i] = v.(string)
	}
	return tokens
}

func JsonMangle[T any](in T) T {
	err := extconversion.Convert(in, &in)
	if err != nil {
		panic(err)
	}
	return in
}

func MaskString(s string, search string, remaining int) string {
	searchStringIndex := strings.Index(s, search)
	if searchStringIndex == -1 {
		return s
	}

	startIndex := searchStringIndex
	stopIndex := startIndex + len(search) - remaining

	out := []rune(s)
	for i := startIndex; i < stopIndex; i++ {
		out[i] = '*'
	}

	return string(out)
}
