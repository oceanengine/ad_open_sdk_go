/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2ScheduleType
type ToolsBidSuggestV2ScheduleType string

// List of tools_bid_suggest_v2_schedule_type
const (
	SCHEDULE_FROM_NOW_ToolsBidSuggestV2ScheduleType  ToolsBidSuggestV2ScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_ToolsBidSuggestV2ScheduleType ToolsBidSuggestV2ScheduleType = "SCHEDULE_START_END"
)

// All allowed values of ToolsBidSuggestV2ScheduleType enum
var AllowedToolsBidSuggestV2ScheduleTypeEnumValues = []ToolsBidSuggestV2ScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewToolsBidSuggestV2ScheduleTypeFromValue returns a pointer to a valid ToolsBidSuggestV2ScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2ScheduleTypeFromValue(v string) (*ToolsBidSuggestV2ScheduleType, error) {
	ev := ToolsBidSuggestV2ScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2ScheduleType: valid values are %v", v, AllowedToolsBidSuggestV2ScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2ScheduleType) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2ScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_schedule_type value
func (v ToolsBidSuggestV2ScheduleType) Ptr() *ToolsBidSuggestV2ScheduleType {
	return &v
}
