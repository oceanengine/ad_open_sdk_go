/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAbTestInfoGetV2DataTestType
type ToolsAbTestInfoGetV2DataTestType string

// List of tools_ab_test_info_get_v2_data_test_type
const (
	AD_ToolsAbTestInfoGetV2DataTestType       ToolsAbTestInfoGetV2DataTestType = "AD"
	CAMPAIGN_ToolsAbTestInfoGetV2DataTestType ToolsAbTestInfoGetV2DataTestType = "CAMPAIGN"
)

// All allowed values of ToolsAbTestInfoGetV2DataTestType enum
var AllowedToolsAbTestInfoGetV2DataTestTypeEnumValues = []ToolsAbTestInfoGetV2DataTestType{
	"AD",
	"CAMPAIGN",
}

// NewToolsAbTestInfoGetV2DataTestTypeFromValue returns a pointer to a valid ToolsAbTestInfoGetV2DataTestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAbTestInfoGetV2DataTestTypeFromValue(v string) (*ToolsAbTestInfoGetV2DataTestType, error) {
	ev := ToolsAbTestInfoGetV2DataTestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAbTestInfoGetV2DataTestType: valid values are %v", v, AllowedToolsAbTestInfoGetV2DataTestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAbTestInfoGetV2DataTestType) IsValid() bool {
	for _, existing := range AllowedToolsAbTestInfoGetV2DataTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ab_test_info_get_v2_data_test_type value
func (v ToolsAbTestInfoGetV2DataTestType) Ptr() *ToolsAbTestInfoGetV2DataTestType {
	return &v
}
