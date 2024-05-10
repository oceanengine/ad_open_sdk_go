/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAbTestListGetV2DataTestListTestType
type ToolsAbTestListGetV2DataTestListTestType string

// List of tools_ab_test_list_get_v2_data_test_list_test_type
const (
	AD_ToolsAbTestListGetV2DataTestListTestType       ToolsAbTestListGetV2DataTestListTestType = "AD"
	CAMPAIGN_ToolsAbTestListGetV2DataTestListTestType ToolsAbTestListGetV2DataTestListTestType = "CAMPAIGN"
)

// All allowed values of ToolsAbTestListGetV2DataTestListTestType enum
var AllowedToolsAbTestListGetV2DataTestListTestTypeEnumValues = []ToolsAbTestListGetV2DataTestListTestType{
	"AD",
	"CAMPAIGN",
}

// NewToolsAbTestListGetV2DataTestListTestTypeFromValue returns a pointer to a valid ToolsAbTestListGetV2DataTestListTestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAbTestListGetV2DataTestListTestTypeFromValue(v string) (*ToolsAbTestListGetV2DataTestListTestType, error) {
	ev := ToolsAbTestListGetV2DataTestListTestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAbTestListGetV2DataTestListTestType: valid values are %v", v, AllowedToolsAbTestListGetV2DataTestListTestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAbTestListGetV2DataTestListTestType) IsValid() bool {
	for _, existing := range AllowedToolsAbTestListGetV2DataTestListTestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ab_test_list_get_v2_data_test_list_test_type value
func (v ToolsAbTestListGetV2DataTestListTestType) Ptr() *ToolsAbTestListGetV2DataTestListTestType {
	return &v
}
