/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectUpdateV30AudienceFilterOwnAwemeFans
type ProjectUpdateV30AudienceFilterOwnAwemeFans string

// List of project_update_v3.0_audience_filter_own_aweme_fans
const (
	OFF_ProjectUpdateV30AudienceFilterOwnAwemeFans ProjectUpdateV30AudienceFilterOwnAwemeFans = "OFF"
	ON_ProjectUpdateV30AudienceFilterOwnAwemeFans  ProjectUpdateV30AudienceFilterOwnAwemeFans = "ON"
)

// All allowed values of ProjectUpdateV30AudienceFilterOwnAwemeFans enum
var AllowedProjectUpdateV30AudienceFilterOwnAwemeFansEnumValues = []ProjectUpdateV30AudienceFilterOwnAwemeFans{
	"OFF",
	"ON",
}

// NewProjectUpdateV30AudienceFilterOwnAwemeFansFromValue returns a pointer to a valid ProjectUpdateV30AudienceFilterOwnAwemeFans
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceFilterOwnAwemeFansFromValue(v string) (*ProjectUpdateV30AudienceFilterOwnAwemeFans, error) {
	ev := ProjectUpdateV30AudienceFilterOwnAwemeFans(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceFilterOwnAwemeFans: valid values are %v", v, AllowedProjectUpdateV30AudienceFilterOwnAwemeFansEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceFilterOwnAwemeFans) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceFilterOwnAwemeFansEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_filter_own_aweme_fans value
func (v ProjectUpdateV30AudienceFilterOwnAwemeFans) Ptr() *ProjectUpdateV30AudienceFilterOwnAwemeFans {
	return &v
}
