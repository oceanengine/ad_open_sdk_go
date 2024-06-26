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

// ToolsBidSuggestV2District
type ToolsBidSuggestV2District string

// List of tools_bid_suggest_v2_district
const (
	COUNTY_ToolsBidSuggestV2District            ToolsBidSuggestV2District = "COUNTY"
	NONE_ToolsBidSuggestV2District              ToolsBidSuggestV2District = "NONE"
	REGION_ToolsBidSuggestV2District            ToolsBidSuggestV2District = "REGION"
	CITY_ToolsBidSuggestV2District              ToolsBidSuggestV2District = "CITY"
	BUSINESS_DISTRICT_ToolsBidSuggestV2District ToolsBidSuggestV2District = "BUSINESS_DISTRICT"
	OVERSEA_ToolsBidSuggestV2District           ToolsBidSuggestV2District = "OVERSEA"
)

// All allowed values of ToolsBidSuggestV2District enum
var AllowedToolsBidSuggestV2DistrictEnumValues = []ToolsBidSuggestV2District{
	"COUNTY",
	"NONE",
	"REGION",
	"CITY",
	"BUSINESS_DISTRICT",
	"OVERSEA",
}

// NewToolsBidSuggestV2DistrictFromValue returns a pointer to a valid ToolsBidSuggestV2District
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2DistrictFromValue(v string) (*ToolsBidSuggestV2District, error) {
	ev := ToolsBidSuggestV2District(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2District: valid values are %v", v, AllowedToolsBidSuggestV2DistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2District) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2DistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_district value
func (v ToolsBidSuggestV2District) Ptr() *ToolsBidSuggestV2District {
	return &v
}
