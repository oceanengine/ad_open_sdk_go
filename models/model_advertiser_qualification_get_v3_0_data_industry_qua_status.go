/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserQualificationGetV30DataIndustryQuaStatus
type AdvertiserQualificationGetV30DataIndustryQuaStatus string

// List of advertiser_qualification_get_v3.0_data_industry_qua_status
const (
	STATUS_CONFIRM_AdvertiserQualificationGetV30DataIndustryQuaStatus         AdvertiserQualificationGetV30DataIndustryQuaStatus = "STATUS_CONFIRM"
	STATUS_CONFIRM_FAIL_AdvertiserQualificationGetV30DataIndustryQuaStatus    AdvertiserQualificationGetV30DataIndustryQuaStatus = "STATUS_CONFIRM_FAIL"
	STATUS_NOT_SUBMIT_AdvertiserQualificationGetV30DataIndustryQuaStatus      AdvertiserQualificationGetV30DataIndustryQuaStatus = "STATUS_NOT_SUBMIT"
	STATUS_PENDING_CONFIRM_AdvertiserQualificationGetV30DataIndustryQuaStatus AdvertiserQualificationGetV30DataIndustryQuaStatus = "STATUS_PENDING_CONFIRM"
	STATUS_WAIT_CONFIRM_AdvertiserQualificationGetV30DataIndustryQuaStatus    AdvertiserQualificationGetV30DataIndustryQuaStatus = "STATUS_WAIT_CONFIRM"
)

// All allowed values of AdvertiserQualificationGetV30DataIndustryQuaStatus enum
var AllowedAdvertiserQualificationGetV30DataIndustryQuaStatusEnumValues = []AdvertiserQualificationGetV30DataIndustryQuaStatus{
	"STATUS_CONFIRM",
	"STATUS_CONFIRM_FAIL",
	"STATUS_NOT_SUBMIT",
	"STATUS_PENDING_CONFIRM",
	"STATUS_WAIT_CONFIRM",
}

// NewAdvertiserQualificationGetV30DataIndustryQuaStatusFromValue returns a pointer to a valid AdvertiserQualificationGetV30DataIndustryQuaStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserQualificationGetV30DataIndustryQuaStatusFromValue(v string) (*AdvertiserQualificationGetV30DataIndustryQuaStatus, error) {
	ev := AdvertiserQualificationGetV30DataIndustryQuaStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserQualificationGetV30DataIndustryQuaStatus: valid values are %v", v, AllowedAdvertiserQualificationGetV30DataIndustryQuaStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserQualificationGetV30DataIndustryQuaStatus) IsValid() bool {
	for _, existing := range AllowedAdvertiserQualificationGetV30DataIndustryQuaStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_qualification_get_v3.0_data_industry_qua_status value
func (v AdvertiserQualificationGetV30DataIndustryQuaStatus) Ptr() *AdvertiserQualificationGetV30DataIndustryQuaStatus {
	return &v
}
