/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30FilteringStatusFirst
type ProjectListV30FilteringStatusFirst string

// List of project_list_v3.0_filtering_status_first
const (
	PROJECT_STATUS_DELETE_ProjectListV30FilteringStatusFirst  ProjectListV30FilteringStatusFirst = "PROJECT_STATUS_DELETE"
	PROJECT_STATUS_DISABLE_ProjectListV30FilteringStatusFirst ProjectListV30FilteringStatusFirst = "PROJECT_STATUS_DISABLE"
	PROJECT_STATUS_DONE_ProjectListV30FilteringStatusFirst    ProjectListV30FilteringStatusFirst = "PROJECT_STATUS_DONE"
	PROJECT_STATUS_ENABLE_ProjectListV30FilteringStatusFirst  ProjectListV30FilteringStatusFirst = "PROJECT_STATUS_ENABLE"
)

// All allowed values of ProjectListV30FilteringStatusFirst enum
var AllowedProjectListV30FilteringStatusFirstEnumValues = []ProjectListV30FilteringStatusFirst{
	"PROJECT_STATUS_DELETE",
	"PROJECT_STATUS_DISABLE",
	"PROJECT_STATUS_DONE",
	"PROJECT_STATUS_ENABLE",
}

// NewProjectListV30FilteringStatusFirstFromValue returns a pointer to a valid ProjectListV30FilteringStatusFirst
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringStatusFirstFromValue(v string) (*ProjectListV30FilteringStatusFirst, error) {
	ev := ProjectListV30FilteringStatusFirst(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringStatusFirst: valid values are %v", v, AllowedProjectListV30FilteringStatusFirstEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringStatusFirst) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringStatusFirstEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_status_first value
func (v ProjectListV30FilteringStatusFirst) Ptr() *ProjectListV30FilteringStatusFirst {
	return &v
}
