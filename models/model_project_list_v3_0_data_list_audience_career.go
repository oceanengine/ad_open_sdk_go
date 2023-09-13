/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAudienceCareer
type ProjectListV30DataListAudienceCareer string

// List of project_list_v3.0_data_list_audience_career
const (
	CIVIL_SERVANTS_ProjectListV30DataListAudienceCareer  ProjectListV30DataListAudienceCareer = "CIVIL_SERVANTS"
	COLLEGE_STUDENT_ProjectListV30DataListAudienceCareer ProjectListV30DataListAudienceCareer = "COLLEGE_STUDENT"
	FINANCIAL_ProjectListV30DataListAudienceCareer       ProjectListV30DataListAudienceCareer = "FINANCIAL"
	IT_ProjectListV30DataListAudienceCareer              ProjectListV30DataListAudienceCareer = "IT"
	MEDICAL_STAFF_ProjectListV30DataListAudienceCareer   ProjectListV30DataListAudienceCareer = "MEDICAL_STAFF"
	TEACHER_ProjectListV30DataListAudienceCareer         ProjectListV30DataListAudienceCareer = "TEACHER"
)

// All allowed values of ProjectListV30DataListAudienceCareer enum
var AllowedProjectListV30DataListAudienceCareerEnumValues = []ProjectListV30DataListAudienceCareer{
	"CIVIL_SERVANTS",
	"COLLEGE_STUDENT",
	"FINANCIAL",
	"IT",
	"MEDICAL_STAFF",
	"TEACHER",
}

// NewProjectListV30DataListAudienceCareerFromValue returns a pointer to a valid ProjectListV30DataListAudienceCareer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceCareerFromValue(v string) (*ProjectListV30DataListAudienceCareer, error) {
	ev := ProjectListV30DataListAudienceCareer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceCareer: valid values are %v", v, AllowedProjectListV30DataListAudienceCareerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceCareer) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceCareerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_career value
func (v ProjectListV30DataListAudienceCareer) Ptr() *ProjectListV30DataListAudienceCareer {
	return &v
}
