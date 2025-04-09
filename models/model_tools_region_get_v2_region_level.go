/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsRegionGetV2RegionLevel
type ToolsRegionGetV2RegionLevel string

// List of tools_region_get_v2_region_level
const (
	REGION_LEVEL_CITY_ToolsRegionGetV2RegionLevel              ToolsRegionGetV2RegionLevel = "REGION_LEVEL_CITY"
	REGION_LEVEL_DISTRICT_ToolsRegionGetV2RegionLevel          ToolsRegionGetV2RegionLevel = "REGION_LEVEL_DISTRICT"
	REGION_LEVEL_BUSINESS_DISTRICT_ToolsRegionGetV2RegionLevel ToolsRegionGetV2RegionLevel = "REGION_LEVEL_BUSINESS_DISTRICT"
	REGION_LEVEL_PROVINCE_ToolsRegionGetV2RegionLevel          ToolsRegionGetV2RegionLevel = "REGION_LEVEL_PROVINCE"
)

// Ptr returns reference to tools_region_get_v2_region_level value
func (v ToolsRegionGetV2RegionLevel) Ptr() *ToolsRegionGetV2RegionLevel {
	return &v
}
