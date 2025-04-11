/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2District
type ToolsBidSuggestV2District string

// List of tools_bid_suggest_v2_district
const (
	BUSINESS_DISTRICT_ToolsBidSuggestV2District ToolsBidSuggestV2District = "BUSINESS_DISTRICT"
	COUNTY_ToolsBidSuggestV2District            ToolsBidSuggestV2District = "COUNTY"
	NONE_ToolsBidSuggestV2District              ToolsBidSuggestV2District = "NONE"
	REGION_ToolsBidSuggestV2District            ToolsBidSuggestV2District = "REGION"
	OVERSEA_ToolsBidSuggestV2District           ToolsBidSuggestV2District = "OVERSEA"
	CITY_ToolsBidSuggestV2District              ToolsBidSuggestV2District = "CITY"
)

// Ptr returns reference to tools_bid_suggest_v2_district value
func (v ToolsBidSuggestV2District) Ptr() *ToolsBidSuggestV2District {
	return &v
}
