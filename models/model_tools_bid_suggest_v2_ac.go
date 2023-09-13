/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2Ac
type ToolsBidSuggestV2Ac string

// List of tools_bid_suggest_v2_ac
const (
	Enum_3_G_ToolsBidSuggestV2Ac ToolsBidSuggestV2Ac = "3G"
	Enum_4_G_ToolsBidSuggestV2Ac ToolsBidSuggestV2Ac = "4G"
	Enum_2_G_ToolsBidSuggestV2Ac ToolsBidSuggestV2Ac = "2G"
	WIFI_ToolsBidSuggestV2Ac     ToolsBidSuggestV2Ac = "WIFI"
	Enum_5_G_ToolsBidSuggestV2Ac ToolsBidSuggestV2Ac = "5G"
)

// All allowed values of ToolsBidSuggestV2Ac enum
var AllowedToolsBidSuggestV2AcEnumValues = []ToolsBidSuggestV2Ac{
	"3G",
	"4G",
	"2G",
	"WIFI",
	"5G",
}

// NewToolsBidSuggestV2AcFromValue returns a pointer to a valid ToolsBidSuggestV2Ac
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2AcFromValue(v string) (*ToolsBidSuggestV2Ac, error) {
	ev := ToolsBidSuggestV2Ac(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2Ac: valid values are %v", v, AllowedToolsBidSuggestV2AcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2Ac) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2AcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_ac value
func (v ToolsBidSuggestV2Ac) Ptr() *ToolsBidSuggestV2Ac {
	return &v
}
