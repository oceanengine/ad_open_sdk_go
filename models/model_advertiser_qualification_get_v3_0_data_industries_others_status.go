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

// AdvertiserQualificationGetV30DataIndustriesOthersStatus
type AdvertiserQualificationGetV30DataIndustriesOthersStatus string

// List of advertiser_qualification_get_v3.0_data_industries_others_status
const (
	STATUS_CONFIRM_AdvertiserQualificationGetV30DataIndustriesOthersStatus         AdvertiserQualificationGetV30DataIndustriesOthersStatus = "STATUS_CONFIRM"
	STATUS_CONFIRM_FAIL_AdvertiserQualificationGetV30DataIndustriesOthersStatus    AdvertiserQualificationGetV30DataIndustriesOthersStatus = "STATUS_CONFIRM_FAIL"
	STATUS_NOT_SUBMIT_AdvertiserQualificationGetV30DataIndustriesOthersStatus      AdvertiserQualificationGetV30DataIndustriesOthersStatus = "STATUS_NOT_SUBMIT"
	STATUS_PENDING_CONFIRM_AdvertiserQualificationGetV30DataIndustriesOthersStatus AdvertiserQualificationGetV30DataIndustriesOthersStatus = "STATUS_PENDING_CONFIRM"
	STATUS_WAIT_CONFIRM_AdvertiserQualificationGetV30DataIndustriesOthersStatus    AdvertiserQualificationGetV30DataIndustriesOthersStatus = "STATUS_WAIT_CONFIRM"
)

// All allowed values of AdvertiserQualificationGetV30DataIndustriesOthersStatus enum
var AllowedAdvertiserQualificationGetV30DataIndustriesOthersStatusEnumValues = []AdvertiserQualificationGetV30DataIndustriesOthersStatus{
	"STATUS_CONFIRM",
	"STATUS_CONFIRM_FAIL",
	"STATUS_NOT_SUBMIT",
	"STATUS_PENDING_CONFIRM",
	"STATUS_WAIT_CONFIRM",
}

// NewAdvertiserQualificationGetV30DataIndustriesOthersStatusFromValue returns a pointer to a valid AdvertiserQualificationGetV30DataIndustriesOthersStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserQualificationGetV30DataIndustriesOthersStatusFromValue(v string) (*AdvertiserQualificationGetV30DataIndustriesOthersStatus, error) {
	ev := AdvertiserQualificationGetV30DataIndustriesOthersStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserQualificationGetV30DataIndustriesOthersStatus: valid values are %v", v, AllowedAdvertiserQualificationGetV30DataIndustriesOthersStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserQualificationGetV30DataIndustriesOthersStatus) IsValid() bool {
	for _, existing := range AllowedAdvertiserQualificationGetV30DataIndustriesOthersStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_qualification_get_v3.0_data_industries_others_status value
func (v AdvertiserQualificationGetV30DataIndustriesOthersStatus) Ptr() *AdvertiserQualificationGetV30DataIndustriesOthersStatus {
	return &v
}
