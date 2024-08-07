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

// ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType
type ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType string

// List of tools_event_convert_optimized_goal_get_v3.0_data_goals_value_type
const (
	DISABLED_ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType              ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType = "Disabled"
	DISCRIMINATE_BY_GROUP_ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType = "DiscriminateByGroup"
	DYNAMIC_VALUE_ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType         ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType = "DynamicValue"
	FIXED_ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType                 ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType = "Fixed"
)

// All allowed values of ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType enum
var AllowedToolsEventConvertOptimizedGoalGetV30DataGoalsValueTypeEnumValues = []ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType{
	"Disabled",
	"DiscriminateByGroup",
	"DynamicValue",
	"Fixed",
}

// NewToolsEventConvertOptimizedGoalGetV30DataGoalsValueTypeFromValue returns a pointer to a valid ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventConvertOptimizedGoalGetV30DataGoalsValueTypeFromValue(v string) (*ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType, error) {
	ev := ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType: valid values are %v", v, AllowedToolsEventConvertOptimizedGoalGetV30DataGoalsValueTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType) IsValid() bool {
	for _, existing := range AllowedToolsEventConvertOptimizedGoalGetV30DataGoalsValueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_convert_optimized_goal_get_v3.0_data_goals_value_type value
func (v ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType) Ptr() *ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType {
	return &v
}
