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

// ProjectUpdateV30AudienceHideIfExists
type ProjectUpdateV30AudienceHideIfExists string

// List of project_update_v3.0_audience_hide_if_exists
const (
	FILTER_ProjectUpdateV30AudienceHideIfExists    ProjectUpdateV30AudienceHideIfExists = "FILTER"
	TARGETING_ProjectUpdateV30AudienceHideIfExists ProjectUpdateV30AudienceHideIfExists = "TARGETING"
	UNLIMITED_ProjectUpdateV30AudienceHideIfExists ProjectUpdateV30AudienceHideIfExists = "UNLIMITED"
)

// All allowed values of ProjectUpdateV30AudienceHideIfExists enum
var AllowedProjectUpdateV30AudienceHideIfExistsEnumValues = []ProjectUpdateV30AudienceHideIfExists{
	"FILTER",
	"TARGETING",
	"UNLIMITED",
}

// NewProjectUpdateV30AudienceHideIfExistsFromValue returns a pointer to a valid ProjectUpdateV30AudienceHideIfExists
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceHideIfExistsFromValue(v string) (*ProjectUpdateV30AudienceHideIfExists, error) {
	ev := ProjectUpdateV30AudienceHideIfExists(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceHideIfExists: valid values are %v", v, AllowedProjectUpdateV30AudienceHideIfExistsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceHideIfExists) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceHideIfExistsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_hide_if_exists value
func (v ProjectUpdateV30AudienceHideIfExists) Ptr() *ProjectUpdateV30AudienceHideIfExists {
	return &v
}
