/*
 * Copyright 2023 steadybit GmbH. All rights reserved.
 */

package extconversion

import (
	"encoding/json"
)

// Convert converts a value (from - typically a struct or map[string]interface{}) to another value
// (to - typically also a struct or map[string]interface{}). This is helpful in a variety of cases,
// e.g., to encode ActionKit's action state.
//
// This function leverages json.Marshal and json.Unmarshal internally whereas it previously leveraged
// the mapstructure package. It turned out that using the json package is beneficial, as
// many go internal (time.Time) and external packages (Kubernetes resource types) are compatible
// with the json package, but not with mapstructure.
func Convert(from any, to any) error {
	bytes, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, to)
}
