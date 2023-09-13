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

// AdvertiserQualificationSubmitV30SubjectCompanyType
type AdvertiserQualificationSubmitV30SubjectCompanyType string

// List of advertiser_qualification_submit_v3.0_subject_company_type
const (
	COMPANY_AdvertiserQualificationSubmitV30SubjectCompanyType    AdvertiserQualificationSubmitV30SubjectCompanyType = "COMPANY"
	INDIVIDUAL_AdvertiserQualificationSubmitV30SubjectCompanyType AdvertiserQualificationSubmitV30SubjectCompanyType = "INDIVIDUAL"
)

// All allowed values of AdvertiserQualificationSubmitV30SubjectCompanyType enum
var AllowedAdvertiserQualificationSubmitV30SubjectCompanyTypeEnumValues = []AdvertiserQualificationSubmitV30SubjectCompanyType{
	"COMPANY",
	"INDIVIDUAL",
}

// NewAdvertiserQualificationSubmitV30SubjectCompanyTypeFromValue returns a pointer to a valid AdvertiserQualificationSubmitV30SubjectCompanyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserQualificationSubmitV30SubjectCompanyTypeFromValue(v string) (*AdvertiserQualificationSubmitV30SubjectCompanyType, error) {
	ev := AdvertiserQualificationSubmitV30SubjectCompanyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserQualificationSubmitV30SubjectCompanyType: valid values are %v", v, AllowedAdvertiserQualificationSubmitV30SubjectCompanyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserQualificationSubmitV30SubjectCompanyType) IsValid() bool {
	for _, existing := range AllowedAdvertiserQualificationSubmitV30SubjectCompanyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_qualification_submit_v3.0_subject_company_type value
func (v AdvertiserQualificationSubmitV30SubjectCompanyType) Ptr() *AdvertiserQualificationSubmitV30SubjectCompanyType {
	return &v
}
