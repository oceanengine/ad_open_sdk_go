/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEventConvertOptimizedGoalGetV30ResponseDataGoalsInner struct for ToolsEventConvertOptimizedGoalGetV30ResponseDataGoalsInner
type ToolsEventConvertOptimizedGoalGetV30ResponseDataGoalsInner struct {
	//
	AssetTypes []*ToolsEventConvertOptimizedGoalGetV30DataGoalsAssetTypes `json:"asset_types,omitempty"`
	//
	AssetsId *int64 `json:"assets_id,omitempty"`
	//
	DeepGoals      []*ToolsEventConvertOptimizedGoalGetV30ResponseDataGoalsInnerDeepGoalsInner `json:"deepGoals,omitempty"`
	ExternalAction *ToolsEventConvertOptimizedGoalGetV30DataGoalsExternalAction                `json:"external_action,omitempty"`
	//
	HistoryBack *bool `json:"history_back,omitempty"`
	//
	OptimizationName *string `json:"optimization_name,omitempty"`
	//
	SpareAssetsId *int64 `json:"spare_assets_id,omitempty"`
	//
	TwentyFourHourBack *bool                                                   `json:"twenty_four_hour_back,omitempty"`
	ValueType          *ToolsEventConvertOptimizedGoalGetV30DataGoalsValueType `json:"value_type,omitempty"`
}
