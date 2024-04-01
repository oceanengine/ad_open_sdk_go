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

// EventManagerOptimizedGoalGetV2V30DeliveryType
type EventManagerOptimizedGoalGetV2V30DeliveryType string

// List of event_manager_optimized_goal_get_v2_v3.0_delivery_type
const (
	DURATION_EventManagerOptimizedGoalGetV2V30DeliveryType EventManagerOptimizedGoalGetV2V30DeliveryType = "DURATION"
	NORMAL_EventManagerOptimizedGoalGetV2V30DeliveryType   EventManagerOptimizedGoalGetV2V30DeliveryType = "NORMAL"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30DeliveryType enum
var AllowedEventManagerOptimizedGoalGetV2V30DeliveryTypeEnumValues = []EventManagerOptimizedGoalGetV2V30DeliveryType{
	"DURATION",
	"NORMAL",
}

// NewEventManagerOptimizedGoalGetV2V30DeliveryTypeFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30DeliveryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30DeliveryTypeFromValue(v string) (*EventManagerOptimizedGoalGetV2V30DeliveryType, error) {
	ev := EventManagerOptimizedGoalGetV2V30DeliveryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30DeliveryType: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30DeliveryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30DeliveryType) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30DeliveryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_delivery_type value
func (v EventManagerOptimizedGoalGetV2V30DeliveryType) Ptr() *EventManagerOptimizedGoalGetV2V30DeliveryType {
	return &v
}
