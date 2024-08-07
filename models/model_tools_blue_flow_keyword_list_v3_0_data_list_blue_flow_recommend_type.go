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

// ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType
type ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType string

// List of tools_blue_flow_keyword_list_v3.0_data_list_blue_flow_recommend_type
const (
	ORDINARY_ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType  ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType = "ORDINARY"
	PREFERRED_ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType = "PREFERRED"
)

// All allowed values of ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType enum
var AllowedToolsBlueFlowKeywordListV30DataListBlueFlowRecommendTypeEnumValues = []ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType{
	"ORDINARY",
	"PREFERRED",
}

// NewToolsBlueFlowKeywordListV30DataListBlueFlowRecommendTypeFromValue returns a pointer to a valid ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBlueFlowKeywordListV30DataListBlueFlowRecommendTypeFromValue(v string) (*ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType, error) {
	ev := ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType: valid values are %v", v, AllowedToolsBlueFlowKeywordListV30DataListBlueFlowRecommendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType) IsValid() bool {
	for _, existing := range AllowedToolsBlueFlowKeywordListV30DataListBlueFlowRecommendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_blue_flow_keyword_list_v3.0_data_list_blue_flow_recommend_type value
func (v ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType) Ptr() *ToolsBlueFlowKeywordListV30DataListBlueFlowRecommendType {
	return &v
}
