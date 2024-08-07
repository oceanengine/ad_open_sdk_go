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

// ProjectCreateV30AudiencePlatform
type ProjectCreateV30AudiencePlatform string

// List of project_create_v3.0_audience_platform
const (
	ANDROID_ProjectCreateV30AudiencePlatform ProjectCreateV30AudiencePlatform = "ANDROID"
	IOS_ProjectCreateV30AudiencePlatform     ProjectCreateV30AudiencePlatform = "IOS"
)

// All allowed values of ProjectCreateV30AudiencePlatform enum
var AllowedProjectCreateV30AudiencePlatformEnumValues = []ProjectCreateV30AudiencePlatform{
	"ANDROID",
	"IOS",
}

// NewProjectCreateV30AudiencePlatformFromValue returns a pointer to a valid ProjectCreateV30AudiencePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudiencePlatformFromValue(v string) (*ProjectCreateV30AudiencePlatform, error) {
	ev := ProjectCreateV30AudiencePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudiencePlatform: valid values are %v", v, AllowedProjectCreateV30AudiencePlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudiencePlatform) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudiencePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_platform value
func (v ProjectCreateV30AudiencePlatform) Ptr() *ProjectCreateV30AudiencePlatform {
	return &v
}
