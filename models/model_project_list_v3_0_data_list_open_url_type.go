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

// ProjectListV30DataListOpenUrlType
type ProjectListV30DataListOpenUrlType string

// List of project_list_v3.0_data_list_open_url_type
const (
	CUSTOM_ProjectListV30DataListOpenUrlType ProjectListV30DataListOpenUrlType = "CUSTOM"
	DPA_ProjectListV30DataListOpenUrlType    ProjectListV30DataListOpenUrlType = "DPA"
	NONE_ProjectListV30DataListOpenUrlType   ProjectListV30DataListOpenUrlType = "NONE"
)

// All allowed values of ProjectListV30DataListOpenUrlType enum
var AllowedProjectListV30DataListOpenUrlTypeEnumValues = []ProjectListV30DataListOpenUrlType{
	"CUSTOM",
	"DPA",
	"NONE",
}

// NewProjectListV30DataListOpenUrlTypeFromValue returns a pointer to a valid ProjectListV30DataListOpenUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListOpenUrlTypeFromValue(v string) (*ProjectListV30DataListOpenUrlType, error) {
	ev := ProjectListV30DataListOpenUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListOpenUrlType: valid values are %v", v, AllowedProjectListV30DataListOpenUrlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListOpenUrlType) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListOpenUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_open_url_type value
func (v ProjectListV30DataListOpenUrlType) Ptr() *ProjectListV30DataListOpenUrlType {
	return &v
}
