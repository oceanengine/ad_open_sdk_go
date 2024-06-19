/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2DeviceType
type ToolsBidSuggestV2DeviceType string

// List of tools_bid_suggest_v2_device_type
const (
	PAD_ToolsBidSuggestV2DeviceType    ToolsBidSuggestV2DeviceType = "PAD"
	MOBILE_ToolsBidSuggestV2DeviceType ToolsBidSuggestV2DeviceType = "MOBILE"
)

// All allowed values of ToolsBidSuggestV2DeviceType enum
var AllowedToolsBidSuggestV2DeviceTypeEnumValues = []ToolsBidSuggestV2DeviceType{
	"PAD",
	"MOBILE",
}

// NewToolsBidSuggestV2DeviceTypeFromValue returns a pointer to a valid ToolsBidSuggestV2DeviceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2DeviceTypeFromValue(v string) (*ToolsBidSuggestV2DeviceType, error) {
	ev := ToolsBidSuggestV2DeviceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2DeviceType: valid values are %v", v, AllowedToolsBidSuggestV2DeviceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2DeviceType) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2DeviceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_device_type value
func (v ToolsBidSuggestV2DeviceType) Ptr() *ToolsBidSuggestV2DeviceType {
	return &v
}
