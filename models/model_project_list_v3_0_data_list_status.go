/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListStatus
type ProjectListV30DataListStatus string

// List of project_list_v3.0_data_list_status
const (
	PROJECT_STATUS_BUDGET_EXCEED_ProjectListV30DataListStatus             ProjectListV30DataListStatus = "PROJECT_STATUS_BUDGET_EXCEED"
	PROJECT_STATUS_BUDGET_PRE_OFFLINE_BUDGET_ProjectListV30DataListStatus ProjectListV30DataListStatus = "PROJECT_STATUS_BUDGET_PRE_OFFLINE_BUDGET"
	PROJECT_STATUS_DELETE_ProjectListV30DataListStatus                    ProjectListV30DataListStatus = "PROJECT_STATUS_DELETE"
	PROJECT_STATUS_DISABLE_ProjectListV30DataListStatus                   ProjectListV30DataListStatus = "PROJECT_STATUS_DISABLE"
	PROJECT_STATUS_DONE_ProjectListV30DataListStatus                      ProjectListV30DataListStatus = "PROJECT_STATUS_DONE"
	PROJECT_STATUS_ENABLE_ProjectListV30DataListStatus                    ProjectListV30DataListStatus = "PROJECT_STATUS_ENABLE"
	PROJECT_STATUS_NOT_DELETE_ProjectListV30DataListStatus                ProjectListV30DataListStatus = "PROJECT_STATUS_NOT_DELETE"
	PROJECT_STATUS_NOT_START_ProjectListV30DataListStatus                 ProjectListV30DataListStatus = "PROJECT_STATUS_NOT_START"
	PROJECT_STATUS_NO_SCHEDULE_ProjectListV30DataListStatus               ProjectListV30DataListStatus = "PROJECT_STATUS_NO_SCHEDULE"
)

// All allowed values of ProjectListV30DataListStatus enum
var AllowedProjectListV30DataListStatusEnumValues = []ProjectListV30DataListStatus{
	"PROJECT_STATUS_BUDGET_EXCEED",
	"PROJECT_STATUS_BUDGET_PRE_OFFLINE_BUDGET",
	"PROJECT_STATUS_DELETE",
	"PROJECT_STATUS_DISABLE",
	"PROJECT_STATUS_DONE",
	"PROJECT_STATUS_ENABLE",
	"PROJECT_STATUS_NOT_DELETE",
	"PROJECT_STATUS_NOT_START",
	"PROJECT_STATUS_NO_SCHEDULE",
}

// NewProjectListV30DataListStatusFromValue returns a pointer to a valid ProjectListV30DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListStatusFromValue(v string) (*ProjectListV30DataListStatus, error) {
	ev := ProjectListV30DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListStatus: valid values are %v", v, AllowedProjectListV30DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListStatus) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_status value
func (v ProjectListV30DataListStatus) Ptr() *ProjectListV30DataListStatus {
	return &v
}
