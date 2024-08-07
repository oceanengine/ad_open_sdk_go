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

// ToolsUnionFlowPackageReportV2OrderType
type ToolsUnionFlowPackageReportV2OrderType string

// List of tools_union_flow_package_report_v2_order_type
const (
	ASC_ToolsUnionFlowPackageReportV2OrderType  ToolsUnionFlowPackageReportV2OrderType = "ASC"
	DESC_ToolsUnionFlowPackageReportV2OrderType ToolsUnionFlowPackageReportV2OrderType = "DESC"
)

// All allowed values of ToolsUnionFlowPackageReportV2OrderType enum
var AllowedToolsUnionFlowPackageReportV2OrderTypeEnumValues = []ToolsUnionFlowPackageReportV2OrderType{
	"ASC",
	"DESC",
}

// NewToolsUnionFlowPackageReportV2OrderTypeFromValue returns a pointer to a valid ToolsUnionFlowPackageReportV2OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsUnionFlowPackageReportV2OrderTypeFromValue(v string) (*ToolsUnionFlowPackageReportV2OrderType, error) {
	ev := ToolsUnionFlowPackageReportV2OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsUnionFlowPackageReportV2OrderType: valid values are %v", v, AllowedToolsUnionFlowPackageReportV2OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsUnionFlowPackageReportV2OrderType) IsValid() bool {
	for _, existing := range AllowedToolsUnionFlowPackageReportV2OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_union_flow_package_report_v2_order_type value
func (v ToolsUnionFlowPackageReportV2OrderType) Ptr() *ToolsUnionFlowPackageReportV2OrderType {
	return &v
}
