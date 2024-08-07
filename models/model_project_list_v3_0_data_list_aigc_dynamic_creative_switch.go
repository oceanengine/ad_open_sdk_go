/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAigcDynamicCreativeSwitch
type ProjectListV30DataListAigcDynamicCreativeSwitch string

// List of project_list_v3.0_data_list_aigc_dynamic_creative_switch
const (
	OFF_ProjectListV30DataListAigcDynamicCreativeSwitch ProjectListV30DataListAigcDynamicCreativeSwitch = "OFF"
	ON_ProjectListV30DataListAigcDynamicCreativeSwitch  ProjectListV30DataListAigcDynamicCreativeSwitch = "ON"
)

// All allowed values of ProjectListV30DataListAigcDynamicCreativeSwitch enum
var AllowedProjectListV30DataListAigcDynamicCreativeSwitchEnumValues = []ProjectListV30DataListAigcDynamicCreativeSwitch{
	"OFF",
	"ON",
}

// NewProjectListV30DataListAigcDynamicCreativeSwitchFromValue returns a pointer to a valid ProjectListV30DataListAigcDynamicCreativeSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAigcDynamicCreativeSwitchFromValue(v string) (*ProjectListV30DataListAigcDynamicCreativeSwitch, error) {
	ev := ProjectListV30DataListAigcDynamicCreativeSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAigcDynamicCreativeSwitch: valid values are %v", v, AllowedProjectListV30DataListAigcDynamicCreativeSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAigcDynamicCreativeSwitch) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAigcDynamicCreativeSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_aigc_dynamic_creative_switch value
func (v ProjectListV30DataListAigcDynamicCreativeSwitch) Ptr() *ProjectListV30DataListAigcDynamicCreativeSwitch {
	return &v
}
