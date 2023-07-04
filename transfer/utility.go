// Copyright (c) 2023 RFull Development
// This source code is managed under the MIT license. See LICENSE in the project root.
package transfer

import (
	"encoding/json"
	"strings"
)

// ToJson converts a Feed to JSON.
func ToJson(feed *Feed) (string, error) {
	c, e := json.Marshal(feed)
	if e != nil {
		return "", e
	}
	converted := string(c)
	return converted, nil
}

// XmlValue returns the value without namespace from the XML value.
func XmlValue(value string) (string, error) {
	if value == "" {
		return "", nil
	}
	v := strings.TrimSpace(value)
	i := strings.Index(v, " ")
	if i < 0 {
		return v, nil
	}
	v = v[i:]
	result := strings.TrimSpace(v)
	return result, nil
}
