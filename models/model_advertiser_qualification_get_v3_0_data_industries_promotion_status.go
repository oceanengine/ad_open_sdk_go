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

// AdvertiserQualificationGetV30DataIndustriesPromotionStatus
type AdvertiserQualificationGetV30DataIndustriesPromotionStatus string

// List of advertiser_qualification_get_v3.0_data_industries_promotion_status
const (
	STATUS_CONFIRM_AdvertiserQualificationGetV30DataIndustriesPromotionStatus         AdvertiserQualificationGetV30DataIndustriesPromotionStatus = "STATUS_CONFIRM"
	STATUS_CONFIRM_FAIL_AdvertiserQualificationGetV30DataIndustriesPromotionStatus    AdvertiserQualificationGetV30DataIndustriesPromotionStatus = "STATUS_CONFIRM_FAIL"
	STATUS_NOT_SUBMIT_AdvertiserQualificationGetV30DataIndustriesPromotionStatus      AdvertiserQualificationGetV30DataIndustriesPromotionStatus = "STATUS_NOT_SUBMIT"
	STATUS_PENDING_CONFIRM_AdvertiserQualificationGetV30DataIndustriesPromotionStatus AdvertiserQualificationGetV30DataIndustriesPromotionStatus = "STATUS_PENDING_CONFIRM"
	STATUS_WAIT_CONFIRM_AdvertiserQualificationGetV30DataIndustriesPromotionStatus    AdvertiserQualificationGetV30DataIndustriesPromotionStatus = "STATUS_WAIT_CONFIRM"
)

// All allowed values of AdvertiserQualificationGetV30DataIndustriesPromotionStatus enum
var AllowedAdvertiserQualificationGetV30DataIndustriesPromotionStatusEnumValues = []AdvertiserQualificationGetV30DataIndustriesPromotionStatus{
	"STATUS_CONFIRM",
	"STATUS_CONFIRM_FAIL",
	"STATUS_NOT_SUBMIT",
	"STATUS_PENDING_CONFIRM",
	"STATUS_WAIT_CONFIRM",
}

// NewAdvertiserQualificationGetV30DataIndustriesPromotionStatusFromValue returns a pointer to a valid AdvertiserQualificationGetV30DataIndustriesPromotionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserQualificationGetV30DataIndustriesPromotionStatusFromValue(v string) (*AdvertiserQualificationGetV30DataIndustriesPromotionStatus, error) {
	ev := AdvertiserQualificationGetV30DataIndustriesPromotionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserQualificationGetV30DataIndustriesPromotionStatus: valid values are %v", v, AllowedAdvertiserQualificationGetV30DataIndustriesPromotionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserQualificationGetV30DataIndustriesPromotionStatus) IsValid() bool {
	for _, existing := range AllowedAdvertiserQualificationGetV30DataIndustriesPromotionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_qualification_get_v3.0_data_industries_promotion_status value
func (v AdvertiserQualificationGetV30DataIndustriesPromotionStatus) Ptr() *AdvertiserQualificationGetV30DataIndustriesPromotionStatus {
	return &v
}
