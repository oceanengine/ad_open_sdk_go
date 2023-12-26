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

// ToolsKeywordsBidRatioCreateV30DataConflictDimension
type ToolsKeywordsBidRatioCreateV30DataConflictDimension string

// List of tools_keywords_bid_ratio_create_v3.0_data_conflict_dimension
const (
	ADVERTISER_ToolsKeywordsBidRatioCreateV30DataConflictDimension ToolsKeywordsBidRatioCreateV30DataConflictDimension = "ADVERTISER"
	PROJECT_ToolsKeywordsBidRatioCreateV30DataConflictDimension    ToolsKeywordsBidRatioCreateV30DataConflictDimension = "PROJECT"
)

// All allowed values of ToolsKeywordsBidRatioCreateV30DataConflictDimension enum
var AllowedToolsKeywordsBidRatioCreateV30DataConflictDimensionEnumValues = []ToolsKeywordsBidRatioCreateV30DataConflictDimension{
	"ADVERTISER",
	"PROJECT",
}

// NewToolsKeywordsBidRatioCreateV30DataConflictDimensionFromValue returns a pointer to a valid ToolsKeywordsBidRatioCreateV30DataConflictDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsKeywordsBidRatioCreateV30DataConflictDimensionFromValue(v string) (*ToolsKeywordsBidRatioCreateV30DataConflictDimension, error) {
	ev := ToolsKeywordsBidRatioCreateV30DataConflictDimension(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsKeywordsBidRatioCreateV30DataConflictDimension: valid values are %v", v, AllowedToolsKeywordsBidRatioCreateV30DataConflictDimensionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsKeywordsBidRatioCreateV30DataConflictDimension) IsValid() bool {
	for _, existing := range AllowedToolsKeywordsBidRatioCreateV30DataConflictDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_keywords_bid_ratio_create_v3.0_data_conflict_dimension value
func (v ToolsKeywordsBidRatioCreateV30DataConflictDimension) Ptr() *ToolsKeywordsBidRatioCreateV30DataConflictDimension {
	return &v
}
