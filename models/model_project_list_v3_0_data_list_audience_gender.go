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

// ProjectListV30DataListAudienceGender
type ProjectListV30DataListAudienceGender string

// List of project_list_v3.0_data_list_audience_gender
const (
	GENDER_FEMALE_ProjectListV30DataListAudienceGender ProjectListV30DataListAudienceGender = "GENDER_FEMALE"
	GENDER_MALE_ProjectListV30DataListAudienceGender   ProjectListV30DataListAudienceGender = "GENDER_MALE"
	NONE_ProjectListV30DataListAudienceGender          ProjectListV30DataListAudienceGender = "NONE"
)

// All allowed values of ProjectListV30DataListAudienceGender enum
var AllowedProjectListV30DataListAudienceGenderEnumValues = []ProjectListV30DataListAudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
	"NONE",
}

// NewProjectListV30DataListAudienceGenderFromValue returns a pointer to a valid ProjectListV30DataListAudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceGenderFromValue(v string) (*ProjectListV30DataListAudienceGender, error) {
	ev := ProjectListV30DataListAudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceGender: valid values are %v", v, AllowedProjectListV30DataListAudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceGender) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_gender value
func (v ProjectListV30DataListAudienceGender) Ptr() *ProjectListV30DataListAudienceGender {
	return &v
}
