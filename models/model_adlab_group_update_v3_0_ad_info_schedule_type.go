/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupUpdateV30AdInfoScheduleType
type AdlabGroupUpdateV30AdInfoScheduleType string

// List of adlab_group_update_v3.0_ad_info_schedule_type
const (
	SCHEDULE_FROM_NOW_AdlabGroupUpdateV30AdInfoScheduleType  AdlabGroupUpdateV30AdInfoScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_AdlabGroupUpdateV30AdInfoScheduleType AdlabGroupUpdateV30AdInfoScheduleType = "SCHEDULE_START_END"
)

// All allowed values of AdlabGroupUpdateV30AdInfoScheduleType enum
var AllowedAdlabGroupUpdateV30AdInfoScheduleTypeEnumValues = []AdlabGroupUpdateV30AdInfoScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewAdlabGroupUpdateV30AdInfoScheduleTypeFromValue returns a pointer to a valid AdlabGroupUpdateV30AdInfoScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateV30AdInfoScheduleTypeFromValue(v string) (*AdlabGroupUpdateV30AdInfoScheduleType, error) {
	ev := AdlabGroupUpdateV30AdInfoScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateV30AdInfoScheduleType: valid values are %v", v, AllowedAdlabGroupUpdateV30AdInfoScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateV30AdInfoScheduleType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateV30AdInfoScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_v3.0_ad_info_schedule_type value
func (v AdlabGroupUpdateV30AdInfoScheduleType) Ptr() *AdlabGroupUpdateV30AdInfoScheduleType {
	return &v
}
