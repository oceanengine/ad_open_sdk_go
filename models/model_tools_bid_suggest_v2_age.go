/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2Age
type ToolsBidSuggestV2Age string

// List of tools_bid_suggest_v2_age
const (
	AGE_BELOW_18_ToolsBidSuggestV2Age      ToolsBidSuggestV2Age = "AGE_BELOW_18"
	AGE_ABOVE_50_ToolsBidSuggestV2Age      ToolsBidSuggestV2Age = "AGE_ABOVE_50"
	AGE_BETWEEN_41_49_ToolsBidSuggestV2Age ToolsBidSuggestV2Age = "AGE_BETWEEN_41_49"
	AGE_BETWEEN_18_23_ToolsBidSuggestV2Age ToolsBidSuggestV2Age = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_31_40_ToolsBidSuggestV2Age ToolsBidSuggestV2Age = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_24_30_ToolsBidSuggestV2Age ToolsBidSuggestV2Age = "AGE_BETWEEN_24_30"
)

// All allowed values of ToolsBidSuggestV2Age enum
var AllowedToolsBidSuggestV2AgeEnumValues = []ToolsBidSuggestV2Age{
	"AGE_BELOW_18",
	"AGE_ABOVE_50",
	"AGE_BETWEEN_41_49",
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_24_30",
}

// NewToolsBidSuggestV2AgeFromValue returns a pointer to a valid ToolsBidSuggestV2Age
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2AgeFromValue(v string) (*ToolsBidSuggestV2Age, error) {
	ev := ToolsBidSuggestV2Age(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2Age: valid values are %v", v, AllowedToolsBidSuggestV2AgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2Age) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2AgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_age value
func (v ToolsBidSuggestV2Age) Ptr() *ToolsBidSuggestV2Age {
	return &v
}