/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdminInfoV2SubDistrict
type ToolsAdminInfoV2SubDistrict string

// List of tools_admin_info_v2_sub_district
const (
	FOUR_LEVEL_ToolsAdminInfoV2SubDistrict  ToolsAdminInfoV2SubDistrict = "FOUR_LEVEL"
	NONE_ToolsAdminInfoV2SubDistrict        ToolsAdminInfoV2SubDistrict = "NONE"
	ONE_LEVEL_ToolsAdminInfoV2SubDistrict   ToolsAdminInfoV2SubDistrict = "ONE_LEVEL"
	THREE_LEVEL_ToolsAdminInfoV2SubDistrict ToolsAdminInfoV2SubDistrict = "THREE_LEVEL"
	TWO_LEVEL_ToolsAdminInfoV2SubDistrict   ToolsAdminInfoV2SubDistrict = "TWO_LEVEL"
)

// All allowed values of ToolsAdminInfoV2SubDistrict enum
var AllowedToolsAdminInfoV2SubDistrictEnumValues = []ToolsAdminInfoV2SubDistrict{
	"FOUR_LEVEL",
	"NONE",
	"ONE_LEVEL",
	"THREE_LEVEL",
	"TWO_LEVEL",
}

// NewToolsAdminInfoV2SubDistrictFromValue returns a pointer to a valid ToolsAdminInfoV2SubDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdminInfoV2SubDistrictFromValue(v string) (*ToolsAdminInfoV2SubDistrict, error) {
	ev := ToolsAdminInfoV2SubDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdminInfoV2SubDistrict: valid values are %v", v, AllowedToolsAdminInfoV2SubDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdminInfoV2SubDistrict) IsValid() bool {
	for _, existing := range AllowedToolsAdminInfoV2SubDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_admin_info_v2_sub_district value
func (v ToolsAdminInfoV2SubDistrict) Ptr() *ToolsAdminInfoV2SubDistrict {
	return &v
}
