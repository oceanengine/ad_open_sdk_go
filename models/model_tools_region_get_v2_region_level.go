/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsRegionGetV2RegionLevel
type ToolsRegionGetV2RegionLevel string

// List of tools_region_get_v2_region_level
const (
	REGION_LEVEL_BUSINESS_DISTRICT_ToolsRegionGetV2RegionLevel ToolsRegionGetV2RegionLevel = "REGION_LEVEL_BUSINESS_DISTRICT"
	REGION_LEVEL_PROVINCE_ToolsRegionGetV2RegionLevel          ToolsRegionGetV2RegionLevel = "REGION_LEVEL_PROVINCE"
	REGION_LEVEL_DISTRICT_ToolsRegionGetV2RegionLevel          ToolsRegionGetV2RegionLevel = "REGION_LEVEL_DISTRICT"
	REGION_LEVEL_CITY_ToolsRegionGetV2RegionLevel              ToolsRegionGetV2RegionLevel = "REGION_LEVEL_CITY"
)

// All allowed values of ToolsRegionGetV2RegionLevel enum
var AllowedToolsRegionGetV2RegionLevelEnumValues = []ToolsRegionGetV2RegionLevel{
	"REGION_LEVEL_BUSINESS_DISTRICT",
	"REGION_LEVEL_PROVINCE",
	"REGION_LEVEL_DISTRICT",
	"REGION_LEVEL_CITY",
}

// NewToolsRegionGetV2RegionLevelFromValue returns a pointer to a valid ToolsRegionGetV2RegionLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRegionGetV2RegionLevelFromValue(v string) (*ToolsRegionGetV2RegionLevel, error) {
	ev := ToolsRegionGetV2RegionLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRegionGetV2RegionLevel: valid values are %v", v, AllowedToolsRegionGetV2RegionLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRegionGetV2RegionLevel) IsValid() bool {
	for _, existing := range AllowedToolsRegionGetV2RegionLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_region_get_v2_region_level value
func (v ToolsRegionGetV2RegionLevel) Ptr() *ToolsRegionGetV2RegionLevel {
	return &v
}
