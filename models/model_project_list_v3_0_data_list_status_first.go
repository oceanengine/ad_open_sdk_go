/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListStatusFirst
type ProjectListV30DataListStatusFirst string

// List of project_list_v3.0_data_list_status_first
const (
	PROJECT_STATUS_DELETE_ProjectListV30DataListStatusFirst  ProjectListV30DataListStatusFirst = "PROJECT_STATUS_DELETE"
	PROJECT_STATUS_DISABLE_ProjectListV30DataListStatusFirst ProjectListV30DataListStatusFirst = "PROJECT_STATUS_DISABLE"
	PROJECT_STATUS_DONE_ProjectListV30DataListStatusFirst    ProjectListV30DataListStatusFirst = "PROJECT_STATUS_DONE"
	PROJECT_STATUS_ENABLE_ProjectListV30DataListStatusFirst  ProjectListV30DataListStatusFirst = "PROJECT_STATUS_ENABLE"
)

// All allowed values of ProjectListV30DataListStatusFirst enum
var AllowedProjectListV30DataListStatusFirstEnumValues = []ProjectListV30DataListStatusFirst{
	"PROJECT_STATUS_DELETE",
	"PROJECT_STATUS_DISABLE",
	"PROJECT_STATUS_DONE",
	"PROJECT_STATUS_ENABLE",
}

// NewProjectListV30DataListStatusFirstFromValue returns a pointer to a valid ProjectListV30DataListStatusFirst
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListStatusFirstFromValue(v string) (*ProjectListV30DataListStatusFirst, error) {
	ev := ProjectListV30DataListStatusFirst(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListStatusFirst: valid values are %v", v, AllowedProjectListV30DataListStatusFirstEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListStatusFirst) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListStatusFirstEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_status_first value
func (v ProjectListV30DataListStatusFirst) Ptr() *ProjectListV30DataListStatusFirst {
	return &v
}
