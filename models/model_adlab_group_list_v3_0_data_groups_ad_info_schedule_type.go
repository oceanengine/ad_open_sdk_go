/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupListV30DataGroupsAdInfoScheduleType
type AdlabGroupListV30DataGroupsAdInfoScheduleType string

// List of adlab_group_list_v3.0_data_groups_ad_info_schedule_type
const (
	SCHEDULE_FROM_NOW_AdlabGroupListV30DataGroupsAdInfoScheduleType  AdlabGroupListV30DataGroupsAdInfoScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_AdlabGroupListV30DataGroupsAdInfoScheduleType AdlabGroupListV30DataGroupsAdInfoScheduleType = "SCHEDULE_START_END"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoScheduleType enum
var AllowedAdlabGroupListV30DataGroupsAdInfoScheduleTypeEnumValues = []AdlabGroupListV30DataGroupsAdInfoScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewAdlabGroupListV30DataGroupsAdInfoScheduleTypeFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoScheduleTypeFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoScheduleType, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoScheduleType: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoScheduleType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_schedule_type value
func (v AdlabGroupListV30DataGroupsAdInfoScheduleType) Ptr() *AdlabGroupListV30DataGroupsAdInfoScheduleType {
	return &v
}
