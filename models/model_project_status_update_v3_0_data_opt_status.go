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

// ProjectStatusUpdateV30DataOptStatus
type ProjectStatusUpdateV30DataOptStatus string

// List of project_status_update_v3.0_data_opt_status
const (
	DISABLE_ProjectStatusUpdateV30DataOptStatus ProjectStatusUpdateV30DataOptStatus = "DISABLE"
	ENABLE_ProjectStatusUpdateV30DataOptStatus  ProjectStatusUpdateV30DataOptStatus = "ENABLE"
)

// All allowed values of ProjectStatusUpdateV30DataOptStatus enum
var AllowedProjectStatusUpdateV30DataOptStatusEnumValues = []ProjectStatusUpdateV30DataOptStatus{
	"DISABLE",
	"ENABLE",
}

// NewProjectStatusUpdateV30DataOptStatusFromValue returns a pointer to a valid ProjectStatusUpdateV30DataOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectStatusUpdateV30DataOptStatusFromValue(v string) (*ProjectStatusUpdateV30DataOptStatus, error) {
	ev := ProjectStatusUpdateV30DataOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectStatusUpdateV30DataOptStatus: valid values are %v", v, AllowedProjectStatusUpdateV30DataOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectStatusUpdateV30DataOptStatus) IsValid() bool {
	for _, existing := range AllowedProjectStatusUpdateV30DataOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_status_update_v3.0_data_opt_status value
func (v ProjectStatusUpdateV30DataOptStatus) Ptr() *ProjectStatusUpdateV30DataOptStatus {
	return &v
}
