/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerOptimizedGoalGetV2V30ResponseDataGoalsInner struct for EventManagerOptimizedGoalGetV2V30ResponseDataGoalsInner
type EventManagerOptimizedGoalGetV2V30ResponseDataGoalsInner struct {
	//
	AssetTypes []*EventManagerOptimizedGoalGetV2V30DataGoalsAssetTypes `json:"asset_types,omitempty"`
	//
	DeepGoals      []*EventManagerOptimizedGoalGetV2V30ResponseDataGoalsInnerDeepGoalsInner `json:"deep_goals,omitempty"`
	ExternalAction *EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                `json:"external_action,omitempty"`
	//
	HistoryBack bool `json:"history_back"`
	//
	OptimizationName string `json:"optimization_name"`
	//
	TwentyFourHourBack bool                                                 `json:"twenty_four_hour_back"`
	ValueType          *EventManagerOptimizedGoalGetV2V30DataGoalsValueType `json:"value_type,omitempty"`
}
