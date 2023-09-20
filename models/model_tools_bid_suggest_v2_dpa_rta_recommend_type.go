/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2DpaRtaRecommendType
type ToolsBidSuggestV2DpaRtaRecommendType string

// List of tools_bid_suggest_v2_dpa_rta_recommend_type
const (
	ONLY_ToolsBidSuggestV2DpaRtaRecommendType ToolsBidSuggestV2DpaRtaRecommendType = "ONLY"
	MORE_ToolsBidSuggestV2DpaRtaRecommendType ToolsBidSuggestV2DpaRtaRecommendType = "MORE"
)

// All allowed values of ToolsBidSuggestV2DpaRtaRecommendType enum
var AllowedToolsBidSuggestV2DpaRtaRecommendTypeEnumValues = []ToolsBidSuggestV2DpaRtaRecommendType{
	"ONLY",
	"MORE",
}

// NewToolsBidSuggestV2DpaRtaRecommendTypeFromValue returns a pointer to a valid ToolsBidSuggestV2DpaRtaRecommendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2DpaRtaRecommendTypeFromValue(v string) (*ToolsBidSuggestV2DpaRtaRecommendType, error) {
	ev := ToolsBidSuggestV2DpaRtaRecommendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2DpaRtaRecommendType: valid values are %v", v, AllowedToolsBidSuggestV2DpaRtaRecommendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2DpaRtaRecommendType) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2DpaRtaRecommendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_dpa_rta_recommend_type value
func (v ToolsBidSuggestV2DpaRtaRecommendType) Ptr() *ToolsBidSuggestV2DpaRtaRecommendType {
	return &v
}