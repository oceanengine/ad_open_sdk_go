/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2LocationType
type ToolsBidSuggestV2LocationType string

// List of tools_bid_suggest_v2_location_type
const (
	TRAVEL_ToolsBidSuggestV2LocationType  ToolsBidSuggestV2LocationType = "TRAVEL"
	CURRENT_ToolsBidSuggestV2LocationType ToolsBidSuggestV2LocationType = "CURRENT"
	ALL_ToolsBidSuggestV2LocationType     ToolsBidSuggestV2LocationType = "ALL"
	HOME_ToolsBidSuggestV2LocationType    ToolsBidSuggestV2LocationType = "HOME"
)

// All allowed values of ToolsBidSuggestV2LocationType enum
var AllowedToolsBidSuggestV2LocationTypeEnumValues = []ToolsBidSuggestV2LocationType{
	"TRAVEL",
	"CURRENT",
	"ALL",
	"HOME",
}

// NewToolsBidSuggestV2LocationTypeFromValue returns a pointer to a valid ToolsBidSuggestV2LocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2LocationTypeFromValue(v string) (*ToolsBidSuggestV2LocationType, error) {
	ev := ToolsBidSuggestV2LocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2LocationType: valid values are %v", v, AllowedToolsBidSuggestV2LocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2LocationType) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2LocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_location_type value
func (v ToolsBidSuggestV2LocationType) Ptr() *ToolsBidSuggestV2LocationType {
	return &v
}
