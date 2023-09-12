/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsUnionFlowPackageGetV2DataListFlowPackageType
type ToolsUnionFlowPackageGetV2DataListFlowPackageType string

// List of tools_union_flow_package_get_v2_data_list_flow_package_type
const (
	CUSTOMIZE_ToolsUnionFlowPackageGetV2DataListFlowPackageType ToolsUnionFlowPackageGetV2DataListFlowPackageType = "CUSTOMIZE"
	FEATURED_ToolsUnionFlowPackageGetV2DataListFlowPackageType  ToolsUnionFlowPackageGetV2DataListFlowPackageType = "FEATURED"
	SYSTEM_ToolsUnionFlowPackageGetV2DataListFlowPackageType    ToolsUnionFlowPackageGetV2DataListFlowPackageType = "SYSTEM"
)

// All allowed values of ToolsUnionFlowPackageGetV2DataListFlowPackageType enum
var AllowedToolsUnionFlowPackageGetV2DataListFlowPackageTypeEnumValues = []ToolsUnionFlowPackageGetV2DataListFlowPackageType{
	"CUSTOMIZE",
	"FEATURED",
	"SYSTEM",
}

// NewToolsUnionFlowPackageGetV2DataListFlowPackageTypeFromValue returns a pointer to a valid ToolsUnionFlowPackageGetV2DataListFlowPackageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsUnionFlowPackageGetV2DataListFlowPackageTypeFromValue(v string) (*ToolsUnionFlowPackageGetV2DataListFlowPackageType, error) {
	ev := ToolsUnionFlowPackageGetV2DataListFlowPackageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsUnionFlowPackageGetV2DataListFlowPackageType: valid values are %v", v, AllowedToolsUnionFlowPackageGetV2DataListFlowPackageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsUnionFlowPackageGetV2DataListFlowPackageType) IsValid() bool {
	for _, existing := range AllowedToolsUnionFlowPackageGetV2DataListFlowPackageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_union_flow_package_get_v2_data_list_flow_package_type value
func (v ToolsUnionFlowPackageGetV2DataListFlowPackageType) Ptr() *ToolsUnionFlowPackageGetV2DataListFlowPackageType {
	return &v
}
