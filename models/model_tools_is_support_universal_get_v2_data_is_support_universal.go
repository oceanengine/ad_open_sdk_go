/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsIsSupportUniversalGetV2DataIsSupportUniversal
type ToolsIsSupportUniversalGetV2DataIsSupportUniversal string

// List of tools_is_support_universal_get_v2_data_is_support_universal
const (
	CANNOT_UNIVERSAL_ToolsIsSupportUniversalGetV2DataIsSupportUniversal  ToolsIsSupportUniversalGetV2DataIsSupportUniversal = "CANNOT_UNIVERSAL"
	FORCE_UNIVERSAL_ToolsIsSupportUniversalGetV2DataIsSupportUniversal   ToolsIsSupportUniversalGetV2DataIsSupportUniversal = "FORCE_UNIVERSAL"
	SUPPORT_UNIVERSAL_ToolsIsSupportUniversalGetV2DataIsSupportUniversal ToolsIsSupportUniversalGetV2DataIsSupportUniversal = "SUPPORT_UNIVERSAL"
)

// All allowed values of ToolsIsSupportUniversalGetV2DataIsSupportUniversal enum
var AllowedToolsIsSupportUniversalGetV2DataIsSupportUniversalEnumValues = []ToolsIsSupportUniversalGetV2DataIsSupportUniversal{
	"CANNOT_UNIVERSAL",
	"FORCE_UNIVERSAL",
	"SUPPORT_UNIVERSAL",
}

// NewToolsIsSupportUniversalGetV2DataIsSupportUniversalFromValue returns a pointer to a valid ToolsIsSupportUniversalGetV2DataIsSupportUniversal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsIsSupportUniversalGetV2DataIsSupportUniversalFromValue(v string) (*ToolsIsSupportUniversalGetV2DataIsSupportUniversal, error) {
	ev := ToolsIsSupportUniversalGetV2DataIsSupportUniversal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsIsSupportUniversalGetV2DataIsSupportUniversal: valid values are %v", v, AllowedToolsIsSupportUniversalGetV2DataIsSupportUniversalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsIsSupportUniversalGetV2DataIsSupportUniversal) IsValid() bool {
	for _, existing := range AllowedToolsIsSupportUniversalGetV2DataIsSupportUniversalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_is_support_universal_get_v2_data_is_support_universal value
func (v ToolsIsSupportUniversalGetV2DataIsSupportUniversal) Ptr() *ToolsIsSupportUniversalGetV2DataIsSupportUniversal {
	return &v
}
