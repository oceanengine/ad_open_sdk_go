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

// ProjectCreateV30DeliverySettingProjectCustom
type ProjectCreateV30DeliverySettingProjectCustom string

// List of project_create_v3.0_delivery_setting_project_custom
const (
	OFF_ProjectCreateV30DeliverySettingProjectCustom ProjectCreateV30DeliverySettingProjectCustom = "OFF"
	ON_ProjectCreateV30DeliverySettingProjectCustom  ProjectCreateV30DeliverySettingProjectCustom = "ON"
)

// All allowed values of ProjectCreateV30DeliverySettingProjectCustom enum
var AllowedProjectCreateV30DeliverySettingProjectCustomEnumValues = []ProjectCreateV30DeliverySettingProjectCustom{
	"OFF",
	"ON",
}

// NewProjectCreateV30DeliverySettingProjectCustomFromValue returns a pointer to a valid ProjectCreateV30DeliverySettingProjectCustom
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliverySettingProjectCustomFromValue(v string) (*ProjectCreateV30DeliverySettingProjectCustom, error) {
	ev := ProjectCreateV30DeliverySettingProjectCustom(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliverySettingProjectCustom: valid values are %v", v, AllowedProjectCreateV30DeliverySettingProjectCustomEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliverySettingProjectCustom) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliverySettingProjectCustomEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_setting_project_custom value
func (v ProjectCreateV30DeliverySettingProjectCustom) Ptr() *ProjectCreateV30DeliverySettingProjectCustom {
	return &v
}
