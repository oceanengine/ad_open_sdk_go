/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListDeliverySettingProjectCustom
type ProjectListV30DataListDeliverySettingProjectCustom string

// List of project_list_v3.0_data_list_delivery_setting_project_custom
const (
	OFF_ProjectListV30DataListDeliverySettingProjectCustom ProjectListV30DataListDeliverySettingProjectCustom = "OFF"
	ON_ProjectListV30DataListDeliverySettingProjectCustom  ProjectListV30DataListDeliverySettingProjectCustom = "ON"
)

// All allowed values of ProjectListV30DataListDeliverySettingProjectCustom enum
var AllowedProjectListV30DataListDeliverySettingProjectCustomEnumValues = []ProjectListV30DataListDeliverySettingProjectCustom{
	"OFF",
	"ON",
}

// NewProjectListV30DataListDeliverySettingProjectCustomFromValue returns a pointer to a valid ProjectListV30DataListDeliverySettingProjectCustom
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListDeliverySettingProjectCustomFromValue(v string) (*ProjectListV30DataListDeliverySettingProjectCustom, error) {
	ev := ProjectListV30DataListDeliverySettingProjectCustom(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListDeliverySettingProjectCustom: valid values are %v", v, AllowedProjectListV30DataListDeliverySettingProjectCustomEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListDeliverySettingProjectCustom) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListDeliverySettingProjectCustomEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_delivery_setting_project_custom value
func (v ProjectListV30DataListDeliverySettingProjectCustom) Ptr() *ProjectListV30DataListDeliverySettingProjectCustom {
	return &v
}
