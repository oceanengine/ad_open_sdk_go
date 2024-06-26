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

// AdlabGroupCreateV30AdInfoDeepBidType
type AdlabGroupCreateV30AdInfoDeepBidType string

// List of adlab_group_create_v3.0_ad_info_deep_bid_type
const (
	AUTO_MIN_SECOND_STAGE_AdlabGroupCreateV30AdInfoDeepBidType        AdlabGroupCreateV30AdInfoDeepBidType = "AUTO_MIN_SECOND_STAGE"
	BID_PER_ACTION_AdlabGroupCreateV30AdInfoDeepBidType               AdlabGroupCreateV30AdInfoDeepBidType = "BID_PER_ACTION"
	DEEP_BID_DEFAULT_AdlabGroupCreateV30AdInfoDeepBidType             AdlabGroupCreateV30AdInfoDeepBidType = "DEEP_BID_DEFAULT"
	DEEP_BID_MIN_AdlabGroupCreateV30AdInfoDeepBidType                 AdlabGroupCreateV30AdInfoDeepBidType = "DEEP_BID_MIN"
	DEEP_BID_PACING_AdlabGroupCreateV30AdInfoDeepBidType              AdlabGroupCreateV30AdInfoDeepBidType = "DEEP_BID_PACING"
	DEEP_BID_TYPE_RETENTION_DAYS_AdlabGroupCreateV30AdInfoDeepBidType AdlabGroupCreateV30AdInfoDeepBidType = "DEEP_BID_TYPE_RETENTION_DAYS"
	MIN_SECOND_STAGE_AdlabGroupCreateV30AdInfoDeepBidType             AdlabGroupCreateV30AdInfoDeepBidType = "MIN_SECOND_STAGE"
	PACING_SECOND_STAGE_AdlabGroupCreateV30AdInfoDeepBidType          AdlabGroupCreateV30AdInfoDeepBidType = "PACING_SECOND_STAGE"
	ROI_COEFFICIENT_AdlabGroupCreateV30AdInfoDeepBidType              AdlabGroupCreateV30AdInfoDeepBidType = "ROI_COEFFICIENT"
	ROI_PACING_AdlabGroupCreateV30AdInfoDeepBidType                   AdlabGroupCreateV30AdInfoDeepBidType = "ROI_PACING"
	SMARTBID_AdlabGroupCreateV30AdInfoDeepBidType                     AdlabGroupCreateV30AdInfoDeepBidType = "SMARTBID"
	SOCIAL_ROI_AdlabGroupCreateV30AdInfoDeepBidType                   AdlabGroupCreateV30AdInfoDeepBidType = "SOCIAL_ROI"
)

// All allowed values of AdlabGroupCreateV30AdInfoDeepBidType enum
var AllowedAdlabGroupCreateV30AdInfoDeepBidTypeEnumValues = []AdlabGroupCreateV30AdInfoDeepBidType{
	"AUTO_MIN_SECOND_STAGE",
	"BID_PER_ACTION",
	"DEEP_BID_DEFAULT",
	"DEEP_BID_MIN",
	"DEEP_BID_PACING",
	"DEEP_BID_TYPE_RETENTION_DAYS",
	"MIN_SECOND_STAGE",
	"PACING_SECOND_STAGE",
	"ROI_COEFFICIENT",
	"ROI_PACING",
	"SMARTBID",
	"SOCIAL_ROI",
}

// NewAdlabGroupCreateV30AdInfoDeepBidTypeFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoDeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoDeepBidTypeFromValue(v string) (*AdlabGroupCreateV30AdInfoDeepBidType, error) {
	ev := AdlabGroupCreateV30AdInfoDeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoDeepBidType: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoDeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoDeepBidType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoDeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_deep_bid_type value
func (v AdlabGroupCreateV30AdInfoDeepBidType) Ptr() *AdlabGroupCreateV30AdInfoDeepBidType {
	return &v
}
