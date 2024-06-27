/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudienceDpaLbs
type ProjectUpdateV30AudienceDpaLbs string

// List of project_update_v3.0_audience_dpa_lbs
const (
	OFF_ProjectUpdateV30AudienceDpaLbs ProjectUpdateV30AudienceDpaLbs = "OFF"
	ON_ProjectUpdateV30AudienceDpaLbs  ProjectUpdateV30AudienceDpaLbs = "ON"
)

// All allowed values of ProjectUpdateV30AudienceDpaLbs enum
var AllowedProjectUpdateV30AudienceDpaLbsEnumValues = []ProjectUpdateV30AudienceDpaLbs{
	"OFF",
	"ON",
}

// NewProjectUpdateV30AudienceDpaLbsFromValue returns a pointer to a valid ProjectUpdateV30AudienceDpaLbs
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceDpaLbsFromValue(v string) (*ProjectUpdateV30AudienceDpaLbs, error) {
	ev := ProjectUpdateV30AudienceDpaLbs(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceDpaLbs: valid values are %v", v, AllowedProjectUpdateV30AudienceDpaLbsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceDpaLbs) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceDpaLbsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_dpa_lbs value
func (v ProjectUpdateV30AudienceDpaLbs) Ptr() *ProjectUpdateV30AudienceDpaLbs {
	return &v
}
