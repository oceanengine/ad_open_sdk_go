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

// ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType
type ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType string

// List of tools_blue_flow_keyword_list_v3.0_filtering_blue_flow_recommend_type
const (
	ORDINARY_ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType  ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType = "ORDINARY"
	PREFERRED_ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType = "PREFERRED"
)

// All allowed values of ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType enum
var AllowedToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendTypeEnumValues = []ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType{
	"ORDINARY",
	"PREFERRED",
}

// NewToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendTypeFromValue returns a pointer to a valid ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendTypeFromValue(v string) (*ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType, error) {
	ev := ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType: valid values are %v", v, AllowedToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType) IsValid() bool {
	for _, existing := range AllowedToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_blue_flow_keyword_list_v3.0_filtering_blue_flow_recommend_type value
func (v ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType) Ptr() *ToolsBlueFlowKeywordListV30FilteringBlueFlowRecommendType {
	return &v
}
