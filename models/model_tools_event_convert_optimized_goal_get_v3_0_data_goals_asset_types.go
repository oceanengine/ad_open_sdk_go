/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes
type ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes string

// List of tools_event_convert_optimized_goal_get_v3.0_data_goals_asset_types
const (
	MINI_PROGRAME_ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes   ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes = "MINI_PROGRAME"
	TETRIS_EXTERNAL_ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes = "TETRIS_EXTERNAL"
)

// All allowed values of ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes enum
var AllowedToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypesEnumValues = []ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes{
	"MINI_PROGRAME",
	"TETRIS_EXTERNAL",
}

// NewToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypesFromValue returns a pointer to a valid ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypesFromValue(v string) (*ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes, error) {
	ev := ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes: valid values are %v", v, AllowedToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes) IsValid() bool {
	for _, existing := range AllowedToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_convert_optimized_goal_get_v3.0_data_goals_asset_types value
func (v ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes) Ptr() *ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes {
	return &v
}
