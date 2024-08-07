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

// ToolsAdminInfoV2DataDistrictsLevel
type ToolsAdminInfoV2DataDistrictsLevel string

// List of tools_admin_info_v2_data_districts_level
const (
	FOUR_LEVEL_ToolsAdminInfoV2DataDistrictsLevel  ToolsAdminInfoV2DataDistrictsLevel = "FOUR_LEVEL"
	NONE_ToolsAdminInfoV2DataDistrictsLevel        ToolsAdminInfoV2DataDistrictsLevel = "NONE"
	ONE_LEVEL_ToolsAdminInfoV2DataDistrictsLevel   ToolsAdminInfoV2DataDistrictsLevel = "ONE_LEVEL"
	THREE_LEVEL_ToolsAdminInfoV2DataDistrictsLevel ToolsAdminInfoV2DataDistrictsLevel = "THREE_LEVEL"
	TWO_LEVEL_ToolsAdminInfoV2DataDistrictsLevel   ToolsAdminInfoV2DataDistrictsLevel = "TWO_LEVEL"
)

// All allowed values of ToolsAdminInfoV2DataDistrictsLevel enum
var AllowedToolsAdminInfoV2DataDistrictsLevelEnumValues = []ToolsAdminInfoV2DataDistrictsLevel{
	"FOUR_LEVEL",
	"NONE",
	"ONE_LEVEL",
	"THREE_LEVEL",
	"TWO_LEVEL",
}

// NewToolsAdminInfoV2DataDistrictsLevelFromValue returns a pointer to a valid ToolsAdminInfoV2DataDistrictsLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdminInfoV2DataDistrictsLevelFromValue(v string) (*ToolsAdminInfoV2DataDistrictsLevel, error) {
	ev := ToolsAdminInfoV2DataDistrictsLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdminInfoV2DataDistrictsLevel: valid values are %v", v, AllowedToolsAdminInfoV2DataDistrictsLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdminInfoV2DataDistrictsLevel) IsValid() bool {
	for _, existing := range AllowedToolsAdminInfoV2DataDistrictsLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_admin_info_v2_data_districts_level value
func (v ToolsAdminInfoV2DataDistrictsLevel) Ptr() *ToolsAdminInfoV2DataDistrictsLevel {
	return &v
}
