/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsKeywordsBidRatioGetV30DataListDimension
type ToolsKeywordsBidRatioGetV30DataListDimension string

// List of tools_keywords_bid_ratio_get_v3.0_data_list_dimension
const (
	ADVERTISER_ToolsKeywordsBidRatioGetV30DataListDimension ToolsKeywordsBidRatioGetV30DataListDimension = "ADVERTISER"
	PROJECT_ToolsKeywordsBidRatioGetV30DataListDimension    ToolsKeywordsBidRatioGetV30DataListDimension = "PROJECT"
)

// All allowed values of ToolsKeywordsBidRatioGetV30DataListDimension enum
var AllowedToolsKeywordsBidRatioGetV30DataListDimensionEnumValues = []ToolsKeywordsBidRatioGetV30DataListDimension{
	"ADVERTISER",
	"PROJECT",
}

// NewToolsKeywordsBidRatioGetV30DataListDimensionFromValue returns a pointer to a valid ToolsKeywordsBidRatioGetV30DataListDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsKeywordsBidRatioGetV30DataListDimensionFromValue(v string) (*ToolsKeywordsBidRatioGetV30DataListDimension, error) {
	ev := ToolsKeywordsBidRatioGetV30DataListDimension(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsKeywordsBidRatioGetV30DataListDimension: valid values are %v", v, AllowedToolsKeywordsBidRatioGetV30DataListDimensionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsKeywordsBidRatioGetV30DataListDimension) IsValid() bool {
	for _, existing := range AllowedToolsKeywordsBidRatioGetV30DataListDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_keywords_bid_ratio_get_v3.0_data_list_dimension value
func (v ToolsKeywordsBidRatioGetV30DataListDimension) Ptr() *ToolsKeywordsBidRatioGetV30DataListDimension {
	return &v
}
