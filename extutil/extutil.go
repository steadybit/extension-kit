// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2022 Steadybit GmbH

// Package extutil contains a variety of util functions that were identified as common code duplication.
// More specialized packages exist for groups of use cases (extlogging and exthttp) for example.
package extutil

import (
	"fmt"
	"strconv"
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
