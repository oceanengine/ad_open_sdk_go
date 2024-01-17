/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserQualificationGetV30DataSubjectQualificationType
type AdvertiserQualificationGetV30DataSubjectQualificationType string

// List of advertiser_qualification_get_v3.0_data_subject_qualification_type
const (
	ASSOCIATION_REGISTER_CODE_AdvertiserQualificationGetV30DataSubjectQualificationType  AdvertiserQualificationGetV30DataSubjectQualificationType = "ASSOCIATION_REGISTER_CODE"
	COMMERCIAL_REGISTER_NUMBER_AdvertiserQualificationGetV30DataSubjectQualificationType AdvertiserQualificationGetV30DataSubjectQualificationType = "COMMERCIAL_REGISTER_NUMBER"
	COMPANY_CREDIT_CODE_AdvertiserQualificationGetV30DataSubjectQualificationType        AdvertiserQualificationGetV30DataSubjectQualificationType = "COMPANY_CREDIT_CODE"
	COMPANY_REGISTER_CODE_AdvertiserQualificationGetV30DataSubjectQualificationType      AdvertiserQualificationGetV30DataSubjectQualificationType = "COMPANY_REGISTER_CODE"
	CREDIT_CODE_AdvertiserQualificationGetV30DataSubjectQualificationType                AdvertiserQualificationGetV30DataSubjectQualificationType = "CREDIT_CODE"
	HK_MACAO_TAIWAN_ID_AdvertiserQualificationGetV30DataSubjectQualificationType         AdvertiserQualificationGetV30DataSubjectQualificationType = "HK_MACAO_TAIWAN_ID"
	HK_REGISTER_CODE_AdvertiserQualificationGetV30DataSubjectQualificationType           AdvertiserQualificationGetV30DataSubjectQualificationType = "HK_REGISTER_CODE"
	ID_AdvertiserQualificationGetV30DataSubjectQualificationType                         AdvertiserQualificationGetV30DataSubjectQualificationType = "ID"
	INDIVIDUAL_CREDIT_CODE_AdvertiserQualificationGetV30DataSubjectQualificationType     AdvertiserQualificationGetV30DataSubjectQualificationType = "INDIVIDUAL_CREDIT_CODE"
	INDIVIDUAL_REGISTER_CODE_AdvertiserQualificationGetV30DataSubjectQualificationType   AdvertiserQualificationGetV30DataSubjectQualificationType = "INDIVIDUAL_REGISTER_CODE"
	LAWYER_CERTIFICATE_NUMBER_AdvertiserQualificationGetV30DataSubjectQualificationType  AdvertiserQualificationGetV30DataSubjectQualificationType = "LAWYER_CERTIFICATE_NUMBER"
	LAWYER_PERMIT_CODE_AdvertiserQualificationGetV30DataSubjectQualificationType         AdvertiserQualificationGetV30DataSubjectQualificationType = "LAWYER_PERMIT_CODE"
	LEGAL_PERSON_CREDIT_CODE_AdvertiserQualificationGetV30DataSubjectQualificationType   AdvertiserQualificationGetV30DataSubjectQualificationType = "LEGAL_PERSON_CREDIT_CODE"
	ORGANIZATION_REGISTER_CODE_AdvertiserQualificationGetV30DataSubjectQualificationType AdvertiserQualificationGetV30DataSubjectQualificationType = "ORGANIZATION_REGISTER_CODE"
	OTHER_AdvertiserQualificationGetV30DataSubjectQualificationType                      AdvertiserQualificationGetV30DataSubjectQualificationType = "OTHER"
	OVERSEA_REGISTER_CODE_AdvertiserQualificationGetV30DataSubjectQualificationType      AdvertiserQualificationGetV30DataSubjectQualificationType = "OVERSEA_REGISTER_CODE"
	PASSPORT_ID_AdvertiserQualificationGetV30DataSubjectQualificationType                AdvertiserQualificationGetV30DataSubjectQualificationType = "PASSPORT_ID"
	SCHOOL_PERMIT_CODE_AdvertiserQualificationGetV30DataSubjectQualificationType         AdvertiserQualificationGetV30DataSubjectQualificationType = "SCHOOL_PERMIT_CODE"
	SOCIAL_REGISTER_CODE_AdvertiserQualificationGetV30DataSubjectQualificationType       AdvertiserQualificationGetV30DataSubjectQualificationType = "SOCIAL_REGISTER_CODE"
)

// All allowed values of AdvertiserQualificationGetV30DataSubjectQualificationType enum
var AllowedAdvertiserQualificationGetV30DataSubjectQualificationTypeEnumValues = []AdvertiserQualificationGetV30DataSubjectQualificationType{
	"ASSOCIATION_REGISTER_CODE",
	"COMMERCIAL_REGISTER_NUMBER",
	"COMPANY_CREDIT_CODE",
	"COMPANY_REGISTER_CODE",
	"CREDIT_CODE",
	"HK_MACAO_TAIWAN_ID",
	"HK_REGISTER_CODE",
	"ID",
	"INDIVIDUAL_CREDIT_CODE",
	"INDIVIDUAL_REGISTER_CODE",
	"LAWYER_CERTIFICATE_NUMBER",
	"LAWYER_PERMIT_CODE",
	"LEGAL_PERSON_CREDIT_CODE",
	"ORGANIZATION_REGISTER_CODE",
	"OTHER",
	"OVERSEA_REGISTER_CODE",
	"PASSPORT_ID",
	"SCHOOL_PERMIT_CODE",
	"SOCIAL_REGISTER_CODE",
}

// NewAdvertiserQualificationGetV30DataSubjectQualificationTypeFromValue returns a pointer to a valid AdvertiserQualificationGetV30DataSubjectQualificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserQualificationGetV30DataSubjectQualificationTypeFromValue(v string) (*AdvertiserQualificationGetV30DataSubjectQualificationType, error) {
	ev := AdvertiserQualificationGetV30DataSubjectQualificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserQualificationGetV30DataSubjectQualificationType: valid values are %v", v, AllowedAdvertiserQualificationGetV30DataSubjectQualificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserQualificationGetV30DataSubjectQualificationType) IsValid() bool {
	for _, existing := range AllowedAdvertiserQualificationGetV30DataSubjectQualificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_qualification_get_v3.0_data_subject_qualification_type value
func (v AdvertiserQualificationGetV30DataSubjectQualificationType) Ptr() *AdvertiserQualificationGetV30DataSubjectQualificationType {
	return &v
}
