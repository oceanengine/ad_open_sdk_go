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

// ToolsAdConvertQueryV2DedicateType
type ToolsAdConvertQueryV2DedicateType string

// List of tools_ad_convert_query_v2_dedicate_type
const (
	UNSET_ToolsAdConvertQueryV2DedicateType        ToolsAdConvertQueryV2DedicateType = "UNSET"
	INTERMEDIATE_ToolsAdConvertQueryV2DedicateType ToolsAdConvertQueryV2DedicateType = "INTERMEDIATE"
	DEDICATED_ToolsAdConvertQueryV2DedicateType    ToolsAdConvertQueryV2DedicateType = "DEDICATED"
)

// All allowed values of ToolsAdConvertQueryV2DedicateType enum
var AllowedToolsAdConvertQueryV2DedicateTypeEnumValues = []ToolsAdConvertQueryV2DedicateType{
	"UNSET",
	"INTERMEDIATE",
	"DEDICATED",
}

// NewToolsAdConvertQueryV2DedicateTypeFromValue returns a pointer to a valid ToolsAdConvertQueryV2DedicateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertQueryV2DedicateTypeFromValue(v string) (*ToolsAdConvertQueryV2DedicateType, error) {
	ev := ToolsAdConvertQueryV2DedicateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertQueryV2DedicateType: valid values are %v", v, AllowedToolsAdConvertQueryV2DedicateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertQueryV2DedicateType) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertQueryV2DedicateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_query_v2_dedicate_type value
func (v ToolsAdConvertQueryV2DedicateType) Ptr() *ToolsAdConvertQueryV2DedicateType {
	return &v
}
