/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30AudienceGender
type ProjectCreateV30AudienceGender string

// List of project_create_v3.0_audience_gender
const (
	GENDER_FEMALE_ProjectCreateV30AudienceGender    ProjectCreateV30AudienceGender = "GENDER_FEMALE"
	GENDER_MALE_ProjectCreateV30AudienceGender      ProjectCreateV30AudienceGender = "GENDER_MALE"
	GENDER_UNLIMITED_ProjectCreateV30AudienceGender ProjectCreateV30AudienceGender = "GENDER_UNLIMITED"
	NONE_ProjectCreateV30AudienceGender             ProjectCreateV30AudienceGender = "NONE"
)

// All allowed values of ProjectCreateV30AudienceGender enum
var AllowedProjectCreateV30AudienceGenderEnumValues = []ProjectCreateV30AudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
	"GENDER_UNLIMITED",
	"NONE",
}

// NewProjectCreateV30AudienceGenderFromValue returns a pointer to a valid ProjectCreateV30AudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30AudienceGenderFromValue(v string) (*ProjectCreateV30AudienceGender, error) {
	ev := ProjectCreateV30AudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30AudienceGender: valid values are %v", v, AllowedProjectCreateV30AudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30AudienceGender) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30AudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_audience_gender value
func (v ProjectCreateV30AudienceGender) Ptr() *ProjectCreateV30AudienceGender {
	return &v
}
