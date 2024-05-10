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

// ToolsEventConvertOptimizedGoalGetV30AssetType
type ToolsEventConvertOptimizedGoalGetV30AssetType string

// List of tools_event_convert_optimized_goal_get_v3.0_asset_type
const (
	MINI_PROGRAME_ToolsEventConvertOptimizedGoalGetV30AssetType   ToolsEventConvertOptimizedGoalGetV30AssetType = "MINI_PROGRAME"
	TETRIS_EXTERNAL_ToolsEventConvertOptimizedGoalGetV30AssetType ToolsEventConvertOptimizedGoalGetV30AssetType = "TETRIS_EXTERNAL"
)

// All allowed values of ToolsEventConvertOptimizedGoalGetV30AssetType enum
var AllowedToolsEventConvertOptimizedGoalGetV30AssetTypeEnumValues = []ToolsEventConvertOptimizedGoalGetV30AssetType{
	"MINI_PROGRAME",
	"TETRIS_EXTERNAL",
}

// NewToolsEventConvertOptimizedGoalGetV30AssetTypeFromValue returns a pointer to a valid ToolsEventConvertOptimizedGoalGetV30AssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventConvertOptimizedGoalGetV30AssetTypeFromValue(v string) (*ToolsEventConvertOptimizedGoalGetV30AssetType, error) {
	ev := ToolsEventConvertOptimizedGoalGetV30AssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventConvertOptimizedGoalGetV30AssetType: valid values are %v", v, AllowedToolsEventConvertOptimizedGoalGetV30AssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventConvertOptimizedGoalGetV30AssetType) IsValid() bool {
	for _, existing := range AllowedToolsEventConvertOptimizedGoalGetV30AssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_convert_optimized_goal_get_v3.0_asset_type value
func (v ToolsEventConvertOptimizedGoalGetV30AssetType) Ptr() *ToolsEventConvertOptimizedGoalGetV30AssetType {
	return &v
}
