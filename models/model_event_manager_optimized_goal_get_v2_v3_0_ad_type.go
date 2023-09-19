/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerOptimizedGoalGetV2V30AdType
type EventManagerOptimizedGoalGetV2V30AdType string

// List of event_manager_optimized_goal_get_v2_v3.0_ad_type
const (
	ALL_EventManagerOptimizedGoalGetV2V30AdType    EventManagerOptimizedGoalGetV2V30AdType = "ALL"
	SEARCH_EventManagerOptimizedGoalGetV2V30AdType EventManagerOptimizedGoalGetV2V30AdType = "SEARCH"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30AdType enum
var AllowedEventManagerOptimizedGoalGetV2V30AdTypeEnumValues = []EventManagerOptimizedGoalGetV2V30AdType{
	"ALL",
	"SEARCH",
}

// NewEventManagerOptimizedGoalGetV2V30AdTypeFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30AdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30AdTypeFromValue(v string) (*EventManagerOptimizedGoalGetV2V30AdType, error) {
	ev := EventManagerOptimizedGoalGetV2V30AdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30AdType: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30AdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30AdType) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30AdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_ad_type value
func (v EventManagerOptimizedGoalGetV2V30AdType) Ptr() *EventManagerOptimizedGoalGetV2V30AdType {
	return &v
}
