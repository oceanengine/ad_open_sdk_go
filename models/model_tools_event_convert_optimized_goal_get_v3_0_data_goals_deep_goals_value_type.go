/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType
type ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType string

// List of tools_event_convert_optimized_goal_get_v3.0_data_goals_deepGoals_value_type
const (
	DISABLED_ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType              ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType = "Disabled"
	DISCRIMINATE_BY_GROUP_ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType = "DiscriminateByGroup"
	DYNAMIC_VALUE_ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType         ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType = "DynamicValue"
	FIXED_ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType                 ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType = "Fixed"
)

// Ptr returns reference to tools_event_convert_optimized_goal_get_v3.0_data_goals_deepGoals_value_type value
func (v ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType) Ptr() *ToolsEventConvertOptimizedGoalGetV30DataGoalsDeepGoalsValueType {
	return &v
}
