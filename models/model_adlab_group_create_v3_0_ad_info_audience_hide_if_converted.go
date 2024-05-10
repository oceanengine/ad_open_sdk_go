/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupCreateV30AdInfoAudienceHideIfConverted
type AdlabGroupCreateV30AdInfoAudienceHideIfConverted string

// List of adlab_group_create_v3.0_ad_info_audience_hide_if_converted
const (
	AD_AdlabGroupCreateV30AdInfoAudienceHideIfConverted           AdlabGroupCreateV30AdInfoAudienceHideIfConverted = "AD"
	ADVERTISER_AdlabGroupCreateV30AdInfoAudienceHideIfConverted   AdlabGroupCreateV30AdInfoAudienceHideIfConverted = "ADVERTISER"
	APP_AdlabGroupCreateV30AdInfoAudienceHideIfConverted          AdlabGroupCreateV30AdInfoAudienceHideIfConverted = "APP"
	CAMPAIGN_AdlabGroupCreateV30AdInfoAudienceHideIfConverted     AdlabGroupCreateV30AdInfoAudienceHideIfConverted = "CAMPAIGN"
	CUSTOMER_AdlabGroupCreateV30AdInfoAudienceHideIfConverted     AdlabGroupCreateV30AdInfoAudienceHideIfConverted = "CUSTOMER"
	NO_EXCLUDE_AdlabGroupCreateV30AdInfoAudienceHideIfConverted   AdlabGroupCreateV30AdInfoAudienceHideIfConverted = "NO_EXCLUDE"
	ORGANIZATION_AdlabGroupCreateV30AdInfoAudienceHideIfConverted AdlabGroupCreateV30AdInfoAudienceHideIfConverted = "ORGANIZATION"
)

// All allowed values of AdlabGroupCreateV30AdInfoAudienceHideIfConverted enum
var AllowedAdlabGroupCreateV30AdInfoAudienceHideIfConvertedEnumValues = []AdlabGroupCreateV30AdInfoAudienceHideIfConverted{
	"AD",
	"ADVERTISER",
	"APP",
	"CAMPAIGN",
	"CUSTOMER",
	"NO_EXCLUDE",
	"ORGANIZATION",
}

// NewAdlabGroupCreateV30AdInfoAudienceHideIfConvertedFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoAudienceHideIfConverted
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoAudienceHideIfConvertedFromValue(v string) (*AdlabGroupCreateV30AdInfoAudienceHideIfConverted, error) {
	ev := AdlabGroupCreateV30AdInfoAudienceHideIfConverted(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoAudienceHideIfConverted: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoAudienceHideIfConvertedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoAudienceHideIfConverted) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoAudienceHideIfConvertedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_audience_hide_if_converted value
func (v AdlabGroupCreateV30AdInfoAudienceHideIfConverted) Ptr() *AdlabGroupCreateV30AdInfoAudienceHideIfConverted {
	return &v
}
