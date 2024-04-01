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

// QianchuanAdScheduleDateUpdateV10ScheduleType
type QianchuanAdScheduleDateUpdateV10ScheduleType string

// List of qianchuan_ad_schedule_date_update_v1.0_schedule_type
const (
	SCHEDULE_FROM_NOW_QianchuanAdScheduleDateUpdateV10ScheduleType  QianchuanAdScheduleDateUpdateV10ScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_QianchuanAdScheduleDateUpdateV10ScheduleType QianchuanAdScheduleDateUpdateV10ScheduleType = "SCHEDULE_START_END"
)

// All allowed values of QianchuanAdScheduleDateUpdateV10ScheduleType enum
var AllowedQianchuanAdScheduleDateUpdateV10ScheduleTypeEnumValues = []QianchuanAdScheduleDateUpdateV10ScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewQianchuanAdScheduleDateUpdateV10ScheduleTypeFromValue returns a pointer to a valid QianchuanAdScheduleDateUpdateV10ScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdScheduleDateUpdateV10ScheduleTypeFromValue(v string) (*QianchuanAdScheduleDateUpdateV10ScheduleType, error) {
	ev := QianchuanAdScheduleDateUpdateV10ScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdScheduleDateUpdateV10ScheduleType: valid values are %v", v, AllowedQianchuanAdScheduleDateUpdateV10ScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdScheduleDateUpdateV10ScheduleType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdScheduleDateUpdateV10ScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_schedule_date_update_v1.0_schedule_type value
func (v QianchuanAdScheduleDateUpdateV10ScheduleType) Ptr() *QianchuanAdScheduleDateUpdateV10ScheduleType {
	return &v
}
