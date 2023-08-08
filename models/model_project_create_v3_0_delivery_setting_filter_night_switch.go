/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30DeliverySettingFilterNightSwitch
type ProjectCreateV30DeliverySettingFilterNightSwitch string

// List of project_create_v3.0_delivery_setting_filter_night_switch
const (
	OFF_ProjectCreateV30DeliverySettingFilterNightSwitch ProjectCreateV30DeliverySettingFilterNightSwitch = "OFF"
	ON_ProjectCreateV30DeliverySettingFilterNightSwitch  ProjectCreateV30DeliverySettingFilterNightSwitch = "ON"
)

// All allowed values of ProjectCreateV30DeliverySettingFilterNightSwitch enum
var AllowedProjectCreateV30DeliverySettingFilterNightSwitchEnumValues = []ProjectCreateV30DeliverySettingFilterNightSwitch{
	"OFF",
	"ON",
}

// NewProjectCreateV30DeliverySettingFilterNightSwitchFromValue returns a pointer to a valid ProjectCreateV30DeliverySettingFilterNightSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliverySettingFilterNightSwitchFromValue(v string) (*ProjectCreateV30DeliverySettingFilterNightSwitch, error) {
	ev := ProjectCreateV30DeliverySettingFilterNightSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliverySettingFilterNightSwitch: valid values are %v", v, AllowedProjectCreateV30DeliverySettingFilterNightSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliverySettingFilterNightSwitch) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliverySettingFilterNightSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_setting_filter_night_switch value
func (v ProjectCreateV30DeliverySettingFilterNightSwitch) Ptr() *ProjectCreateV30DeliverySettingFilterNightSwitch {
	return &v
}
