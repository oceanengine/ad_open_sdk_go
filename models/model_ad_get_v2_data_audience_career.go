/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAudienceCareer
type AdGetV2DataAudienceCareer string

// List of ad_get_v2_data_audience_career
const (
	TEACHER_AdGetV2DataAudienceCareer         AdGetV2DataAudienceCareer = "TEACHER"
	COLLEGE_STUDENT_AdGetV2DataAudienceCareer AdGetV2DataAudienceCareer = "COLLEGE_STUDENT"
	IT_AdGetV2DataAudienceCareer              AdGetV2DataAudienceCareer = "IT"
	FINANCIAL_AdGetV2DataAudienceCareer       AdGetV2DataAudienceCareer = "FINANCIAL"
	CIVIL_SERVANTS_AdGetV2DataAudienceCareer  AdGetV2DataAudienceCareer = "CIVIL_SERVANTS"
	MEDICAL_STAFF_AdGetV2DataAudienceCareer   AdGetV2DataAudienceCareer = "MEDICAL_STAFF"
)

// All allowed values of AdGetV2DataAudienceCareer enum
var AllowedAdGetV2DataAudienceCareerEnumValues = []AdGetV2DataAudienceCareer{
	"TEACHER",
	"COLLEGE_STUDENT",
	"IT",
	"FINANCIAL",
	"CIVIL_SERVANTS",
	"MEDICAL_STAFF",
}

// NewAdGetV2DataAudienceCareerFromValue returns a pointer to a valid AdGetV2DataAudienceCareer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceCareerFromValue(v string) (*AdGetV2DataAudienceCareer, error) {
	ev := AdGetV2DataAudienceCareer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceCareer: valid values are %v", v, AllowedAdGetV2DataAudienceCareerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceCareer) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceCareerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_career value
func (v AdGetV2DataAudienceCareer) Ptr() *AdGetV2DataAudienceCareer {
	return &v
}
