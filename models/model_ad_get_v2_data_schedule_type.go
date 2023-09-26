/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataScheduleType
type AdGetV2DataScheduleType string

// List of ad_get_v2_data_schedule_type
const (
	SCHEDULE_START_END_AdGetV2DataScheduleType AdGetV2DataScheduleType = "SCHEDULE_START_END"
	SCHEDULE_FROM_NOW_AdGetV2DataScheduleType  AdGetV2DataScheduleType = "SCHEDULE_FROM_NOW"
)

// All allowed values of AdGetV2DataScheduleType enum
var AllowedAdGetV2DataScheduleTypeEnumValues = []AdGetV2DataScheduleType{
	"SCHEDULE_START_END",
	"SCHEDULE_FROM_NOW",
}

// NewAdGetV2DataScheduleTypeFromValue returns a pointer to a valid AdGetV2DataScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataScheduleTypeFromValue(v string) (*AdGetV2DataScheduleType, error) {
	ev := AdGetV2DataScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataScheduleType: valid values are %v", v, AllowedAdGetV2DataScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataScheduleType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_schedule_type value
func (v AdGetV2DataScheduleType) Ptr() *AdGetV2DataScheduleType {
	return &v
}