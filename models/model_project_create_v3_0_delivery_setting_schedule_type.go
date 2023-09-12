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

// ProjectCreateV30DeliverySettingScheduleType
type ProjectCreateV30DeliverySettingScheduleType string

// List of project_create_v3.0_delivery_setting_schedule_type
const (
	SCHEDULE_FROM_NOW_ProjectCreateV30DeliverySettingScheduleType  ProjectCreateV30DeliverySettingScheduleType = "SCHEDULE_FROM_NOW"
	SCHEDULE_START_END_ProjectCreateV30DeliverySettingScheduleType ProjectCreateV30DeliverySettingScheduleType = "SCHEDULE_START_END"
)

// All allowed values of ProjectCreateV30DeliverySettingScheduleType enum
var AllowedProjectCreateV30DeliverySettingScheduleTypeEnumValues = []ProjectCreateV30DeliverySettingScheduleType{
	"SCHEDULE_FROM_NOW",
	"SCHEDULE_START_END",
}

// NewProjectCreateV30DeliverySettingScheduleTypeFromValue returns a pointer to a valid ProjectCreateV30DeliverySettingScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliverySettingScheduleTypeFromValue(v string) (*ProjectCreateV30DeliverySettingScheduleType, error) {
	ev := ProjectCreateV30DeliverySettingScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliverySettingScheduleType: valid values are %v", v, AllowedProjectCreateV30DeliverySettingScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliverySettingScheduleType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliverySettingScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_setting_schedule_type value
func (v ProjectCreateV30DeliverySettingScheduleType) Ptr() *ProjectCreateV30DeliverySettingScheduleType {
	return &v
}
