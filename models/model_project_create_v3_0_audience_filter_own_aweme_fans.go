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

// ProjectCreateV30AudienceFilterOwnAwemeFans
type ProjectCreateV30AudienceFilterOwnAwemeFans string

// List of project_create_v3.0_audience_filter_own_aweme_fans
const (
	OFF_ProjectCreateV30AudienceFilterOwnAwemeFans ProjectCreateV30AudienceFilterOwnAwemeFans = "OFF"
	ON_ProjectCreateV30AudienceFilterOwnAwemeFans  ProjectCreateV30AudienceFilterOwnAwemeFans = "ON"
)

// All allowed values of ProjectCreateV30AudienceFilterOwnAwemeFans enum
var AllowedProjectCreateV30AudienceFilterOwnAwemeFansEnumValues = []ProjectCreateV30AudienceFilterOwnAwemeFans{
	"OFF",
	"ON",
}

// NewProjectCreateV30AudienceFilterOwnAwemeFansFromValue returns a pointer to a valid ProjectCreateV30AudienceFilterOwnAwemeFans
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceFilterOwnAwemeFansFromValue(v string) (*ProjectCreateV30AudienceFilterOwnAwemeFans, error) {
	ev := ProjectCreateV30AudienceFilterOwnAwemeFans(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceFilterOwnAwemeFans: valid values are %v", v, AllowedProjectCreateV30AudienceFilterOwnAwemeFansEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceFilterOwnAwemeFans) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceFilterOwnAwemeFansEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_filter_own_aweme_fans value
func (v ProjectCreateV30AudienceFilterOwnAwemeFans) Ptr() *ProjectCreateV30AudienceFilterOwnAwemeFans {
	return &v
}
