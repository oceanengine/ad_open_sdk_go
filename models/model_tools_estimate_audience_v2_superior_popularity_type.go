/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2SuperiorPopularityType
type ToolsEstimateAudienceV2SuperiorPopularityType string

// List of tools_estimate_audience_v2_superior_popularity_type
const (
	NONE_ToolsEstimateAudienceV2SuperiorPopularityType ToolsEstimateAudienceV2SuperiorPopularityType = "NONE"
	GAME_ToolsEstimateAudienceV2SuperiorPopularityType ToolsEstimateAudienceV2SuperiorPopularityType = "GAME"
)

// All allowed values of ToolsEstimateAudienceV2SuperiorPopularityType enum
var AllowedToolsEstimateAudienceV2SuperiorPopularityTypeEnumValues = []ToolsEstimateAudienceV2SuperiorPopularityType{
	"NONE",
	"GAME",
}

// NewToolsEstimateAudienceV2SuperiorPopularityTypeFromValue returns a pointer to a valid ToolsEstimateAudienceV2SuperiorPopularityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2SuperiorPopularityTypeFromValue(v string) (*ToolsEstimateAudienceV2SuperiorPopularityType, error) {
	ev := ToolsEstimateAudienceV2SuperiorPopularityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2SuperiorPopularityType: valid values are %v", v, AllowedToolsEstimateAudienceV2SuperiorPopularityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2SuperiorPopularityType) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2SuperiorPopularityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_superior_popularity_type value
func (v ToolsEstimateAudienceV2SuperiorPopularityType) Ptr() *ToolsEstimateAudienceV2SuperiorPopularityType {
	return &v
}
