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

// ProjectListV30FilteringStatusSecond
type ProjectListV30FilteringStatusSecond string

// List of project_list_v3.0_filtering_status_second
const (
	PROJECT_STATUS_BUDGET_EXCEED_ProjectListV30FilteringStatusSecond               ProjectListV30FilteringStatusSecond = "PROJECT_STATUS_BUDGET_EXCEED"
	PROJECT_STATUS_BUDGET_GROUP_OFFLINE_BUDGET_ProjectListV30FilteringStatusSecond ProjectListV30FilteringStatusSecond = "PROJECT_STATUS_BUDGET_GROUP_OFFLINE_BUDGET"
	PROJECT_STATUS_NOT_START_ProjectListV30FilteringStatusSecond                   ProjectListV30FilteringStatusSecond = "PROJECT_STATUS_NOT_START"
	PROJECT_STATUS_NO_SCHEDULE_ProjectListV30FilteringStatusSecond                 ProjectListV30FilteringStatusSecond = "PROJECT_STATUS_NO_SCHEDULE"
	PROJECT_STATUS_STOP_ProjectListV30FilteringStatusSecond                        ProjectListV30FilteringStatusSecond = "PROJECT_STATUS_STOP"
)

// All allowed values of ProjectListV30FilteringStatusSecond enum
var AllowedProjectListV30FilteringStatusSecondEnumValues = []ProjectListV30FilteringStatusSecond{
	"PROJECT_STATUS_BUDGET_EXCEED",
	"PROJECT_STATUS_BUDGET_GROUP_OFFLINE_BUDGET",
	"PROJECT_STATUS_NOT_START",
	"PROJECT_STATUS_NO_SCHEDULE",
	"PROJECT_STATUS_STOP",
}

// NewProjectListV30FilteringStatusSecondFromValue returns a pointer to a valid ProjectListV30FilteringStatusSecond
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringStatusSecondFromValue(v string) (*ProjectListV30FilteringStatusSecond, error) {
	ev := ProjectListV30FilteringStatusSecond(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringStatusSecond: valid values are %v", v, AllowedProjectListV30FilteringStatusSecondEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringStatusSecond) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringStatusSecondEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_status_second value
func (v ProjectListV30FilteringStatusSecond) Ptr() *ProjectListV30FilteringStatusSecond {
	return &v
}
