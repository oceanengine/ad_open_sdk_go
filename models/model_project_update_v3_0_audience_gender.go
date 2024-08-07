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

// ProjectUpdateV30AudienceGender
type ProjectUpdateV30AudienceGender string

// List of project_update_v3.0_audience_gender
const (
	GENDER_FEMALE_ProjectUpdateV30AudienceGender    ProjectUpdateV30AudienceGender = "GENDER_FEMALE"
	GENDER_MALE_ProjectUpdateV30AudienceGender      ProjectUpdateV30AudienceGender = "GENDER_MALE"
	GENDER_UNLIMITED_ProjectUpdateV30AudienceGender ProjectUpdateV30AudienceGender = "GENDER_UNLIMITED"
	NONE_ProjectUpdateV30AudienceGender             ProjectUpdateV30AudienceGender = "NONE"
)

// All allowed values of ProjectUpdateV30AudienceGender enum
var AllowedProjectUpdateV30AudienceGenderEnumValues = []ProjectUpdateV30AudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
	"GENDER_UNLIMITED",
	"NONE",
}

// NewProjectUpdateV30AudienceGenderFromValue returns a pointer to a valid ProjectUpdateV30AudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectUpdateV30AudienceGenderFromValue(v string) (*ProjectUpdateV30AudienceGender, error) {
	ev := ProjectUpdateV30AudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectUpdateV30AudienceGender: valid values are %v", v, AllowedProjectUpdateV30AudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectUpdateV30AudienceGender) IsValid() bool {
	for _, existing := range AllowedProjectUpdateV30AudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_update_v3.0_audience_gender value
func (v ProjectUpdateV30AudienceGender) Ptr() *ProjectUpdateV30AudienceGender {
	return &v
}
