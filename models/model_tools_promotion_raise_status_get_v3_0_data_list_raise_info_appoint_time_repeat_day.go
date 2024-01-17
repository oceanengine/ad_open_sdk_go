/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay
type ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay string

// List of tools_promotion_raise_status_get_v3.0_data_list_raise_info_appoint_time_repeat_day
const (
	EVERY_DAY_ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay     ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay = "EVERY_DAY"
	PER_FRIDAY_ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay    ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay = "PER_FRIDAY"
	PER_MONDAY_ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay    ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay = "PER_MONDAY"
	PER_SATURDAY_ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay  ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay = "PER_SATURDAY"
	PER_SUNDAY_ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay    ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay = "PER_SUNDAY"
	PER_THURSDAY_ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay  ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay = "PER_THURSDAY"
	PER_TUESDAY_ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay   ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay = "PER_TUESDAY"
	PER_WEDNESDAY_ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay = "PER_WEDNESDAY"
)

// All allowed values of ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay enum
var AllowedToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDayEnumValues = []ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay{
	"EVERY_DAY",
	"PER_FRIDAY",
	"PER_MONDAY",
	"PER_SATURDAY",
	"PER_SUNDAY",
	"PER_THURSDAY",
	"PER_TUESDAY",
	"PER_WEDNESDAY",
}

// NewToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDayFromValue returns a pointer to a valid ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDayFromValue(v string) (*ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay, error) {
	ev := ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay: valid values are %v", v, AllowedToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDayEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay) IsValid() bool {
	for _, existing := range AllowedToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_promotion_raise_status_get_v3.0_data_list_raise_info_appoint_time_repeat_day value
func (v ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay) Ptr() *ToolsPromotionRaiseStatusGetV30DataListRaiseInfoAppointTimeRepeatDay {
	return &v
}
