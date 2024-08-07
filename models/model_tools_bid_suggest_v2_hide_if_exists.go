/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2HideIfExists
type ToolsBidSuggestV2HideIfExists int64

// List of tools_bid_suggest_v2_hide_if_exists
const (
	Enum_0_ToolsBidSuggestV2HideIfExists ToolsBidSuggestV2HideIfExists = 0
	Enum_1_ToolsBidSuggestV2HideIfExists ToolsBidSuggestV2HideIfExists = 1
	Enum_2_ToolsBidSuggestV2HideIfExists ToolsBidSuggestV2HideIfExists = 2
)

// All allowed values of ToolsBidSuggestV2HideIfExists enum
var AllowedToolsBidSuggestV2HideIfExistsEnumValues = []ToolsBidSuggestV2HideIfExists{
	0,
	1,
	2,
}

// NewToolsBidSuggestV2HideIfExistsFromValue returns a pointer to a valid ToolsBidSuggestV2HideIfExists
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2HideIfExistsFromValue(v int64) (*ToolsBidSuggestV2HideIfExists, error) {
	ev := ToolsBidSuggestV2HideIfExists(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2HideIfExists: valid values are %v", v, AllowedToolsBidSuggestV2HideIfExistsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2HideIfExists) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2HideIfExistsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_hide_if_exists value
func (v ToolsBidSuggestV2HideIfExists) Ptr() *ToolsBidSuggestV2HideIfExists {
	return &v
}
