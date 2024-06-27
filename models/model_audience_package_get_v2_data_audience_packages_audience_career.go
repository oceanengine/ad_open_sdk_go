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

// AudiencePackageGetV2DataAudiencePackagesAudienceCareer
type AudiencePackageGetV2DataAudiencePackagesAudienceCareer string

// List of audience_package_get_v2_data_audience_packages_audience_career
const (
	CIVIL_SERVANTS_AudiencePackageGetV2DataAudiencePackagesAudienceCareer  AudiencePackageGetV2DataAudiencePackagesAudienceCareer = "CIVIL_SERVANTS"
	COLLEGE_STUDENT_AudiencePackageGetV2DataAudiencePackagesAudienceCareer AudiencePackageGetV2DataAudiencePackagesAudienceCareer = "COLLEGE_STUDENT"
	FINANCIAL_AudiencePackageGetV2DataAudiencePackagesAudienceCareer       AudiencePackageGetV2DataAudiencePackagesAudienceCareer = "FINANCIAL"
	IT_AudiencePackageGetV2DataAudiencePackagesAudienceCareer              AudiencePackageGetV2DataAudiencePackagesAudienceCareer = "IT"
	MEDICAL_STAFF_AudiencePackageGetV2DataAudiencePackagesAudienceCareer   AudiencePackageGetV2DataAudiencePackagesAudienceCareer = "MEDICAL_STAFF"
	TEACHER_AudiencePackageGetV2DataAudiencePackagesAudienceCareer         AudiencePackageGetV2DataAudiencePackagesAudienceCareer = "TEACHER"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudienceCareer enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudienceCareerEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudienceCareer{
	"CIVIL_SERVANTS",
	"COLLEGE_STUDENT",
	"FINANCIAL",
	"IT",
	"MEDICAL_STAFF",
	"TEACHER",
}

// NewAudiencePackageGetV2DataAudiencePackagesAudienceCareerFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudienceCareer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudienceCareerFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudienceCareer, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudienceCareer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudienceCareer: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudienceCareerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudienceCareer) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudienceCareerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_career value
func (v AudiencePackageGetV2DataAudiencePackagesAudienceCareer) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudienceCareer {
	return &v
}
