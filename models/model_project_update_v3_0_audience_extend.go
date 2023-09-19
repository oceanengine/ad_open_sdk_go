/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudienceExtend
type ProjectUpdateV30AudienceExtend string

// List of project_update_v3.0_audience_extend
const (
	OFF_ProjectUpdateV30AudienceExtend ProjectUpdateV30AudienceExtend = "OFF"
	ON_ProjectUpdateV30AudienceExtend  ProjectUpdateV30AudienceExtend = "ON"
)

// All allowed values of ProjectUpdateV30AudienceExtend enum
var AllowedProjectUpdateV30AudienceExtendEnumValues = []ProjectUpdateV30AudienceExtend{
	"OFF",
	"ON",
}

// NewProjectUpdateV30AudienceExtendFromValue returns a pointer to a valid ProjectUpdateV30AudienceExtend
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceExtendFromValue(v string) (*ProjectUpdateV30AudienceExtend, error) {
	ev := ProjectUpdateV30AudienceExtend(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceExtend: valid values are %v", v, AllowedProjectUpdateV30AudienceExtendEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceExtend) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceExtendEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_extend value
func (v ProjectUpdateV30AudienceExtend) Ptr() *ProjectUpdateV30AudienceExtend {
	return &v
}
