/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerOptimizedGoalGetV2V30AppPromotionType
type EventManagerOptimizedGoalGetV2V30AppPromotionType string

// List of event_manager_optimized_goal_get_v2_v3.0_app_promotion_type
const (
	DOWNLOAD_EventManagerOptimizedGoalGetV2V30AppPromotionType EventManagerOptimizedGoalGetV2V30AppPromotionType = "DOWNLOAD"
	LAUNCH_EventManagerOptimizedGoalGetV2V30AppPromotionType   EventManagerOptimizedGoalGetV2V30AppPromotionType = "LAUNCH"
	RESERVE_EventManagerOptimizedGoalGetV2V30AppPromotionType  EventManagerOptimizedGoalGetV2V30AppPromotionType = "RESERVE"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30AppPromotionType enum
var AllowedEventManagerOptimizedGoalGetV2V30AppPromotionTypeEnumValues = []EventManagerOptimizedGoalGetV2V30AppPromotionType{
	"DOWNLOAD",
	"LAUNCH",
	"RESERVE",
}

// NewEventManagerOptimizedGoalGetV2V30AppPromotionTypeFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30AppPromotionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30AppPromotionTypeFromValue(v string) (*EventManagerOptimizedGoalGetV2V30AppPromotionType, error) {
	ev := EventManagerOptimizedGoalGetV2V30AppPromotionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30AppPromotionType: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30AppPromotionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30AppPromotionType) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30AppPromotionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_app_promotion_type value
func (v EventManagerOptimizedGoalGetV2V30AppPromotionType) Ptr() *EventManagerOptimizedGoalGetV2V30AppPromotionType {
	return &v
}
