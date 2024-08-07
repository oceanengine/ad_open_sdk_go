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

// ToolsBidSuggestV2RetargetingType
type ToolsBidSuggestV2RetargetingType string

// List of tools_bid_suggest_v2_retargeting_type
const (
	NONE_ToolsBidSuggestV2RetargetingType                ToolsBidSuggestV2RetargetingType = "NONE"
	RETARGETING_EXCLUDE_ToolsBidSuggestV2RetargetingType ToolsBidSuggestV2RetargetingType = "RETARGETING_EXCLUDE"
	RETARGETING_INCLUDE_ToolsBidSuggestV2RetargetingType ToolsBidSuggestV2RetargetingType = "RETARGETING_INCLUDE"
	RETARGETING_NONE_ToolsBidSuggestV2RetargetingType    ToolsBidSuggestV2RetargetingType = "RETARGETING_NONE"
)

// All allowed values of ToolsBidSuggestV2RetargetingType enum
var AllowedToolsBidSuggestV2RetargetingTypeEnumValues = []ToolsBidSuggestV2RetargetingType{
	"NONE",
	"RETARGETING_EXCLUDE",
	"RETARGETING_INCLUDE",
	"RETARGETING_NONE",
}

// NewToolsBidSuggestV2RetargetingTypeFromValue returns a pointer to a valid ToolsBidSuggestV2RetargetingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2RetargetingTypeFromValue(v string) (*ToolsBidSuggestV2RetargetingType, error) {
	ev := ToolsBidSuggestV2RetargetingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2RetargetingType: valid values are %v", v, AllowedToolsBidSuggestV2RetargetingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2RetargetingType) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2RetargetingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_retargeting_type value
func (v ToolsBidSuggestV2RetargetingType) Ptr() *ToolsBidSuggestV2RetargetingType {
	return &v
}
