/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupUpdateV30AdInfoAudienceHideIfConverted
type AdlabGroupUpdateV30AdInfoAudienceHideIfConverted string

// List of adlab_group_update_v3.0_ad_info_audience_hide_if_converted
const (
	AD_AdlabGroupUpdateV30AdInfoAudienceHideIfConverted           AdlabGroupUpdateV30AdInfoAudienceHideIfConverted = "AD"
	ADVERTISER_AdlabGroupUpdateV30AdInfoAudienceHideIfConverted   AdlabGroupUpdateV30AdInfoAudienceHideIfConverted = "ADVERTISER"
	APP_AdlabGroupUpdateV30AdInfoAudienceHideIfConverted          AdlabGroupUpdateV30AdInfoAudienceHideIfConverted = "APP"
	CAMPAIGN_AdlabGroupUpdateV30AdInfoAudienceHideIfConverted     AdlabGroupUpdateV30AdInfoAudienceHideIfConverted = "CAMPAIGN"
	CUSTOMER_AdlabGroupUpdateV30AdInfoAudienceHideIfConverted     AdlabGroupUpdateV30AdInfoAudienceHideIfConverted = "CUSTOMER"
	NO_EXCLUDE_AdlabGroupUpdateV30AdInfoAudienceHideIfConverted   AdlabGroupUpdateV30AdInfoAudienceHideIfConverted = "NO_EXCLUDE"
	ORGANIZATION_AdlabGroupUpdateV30AdInfoAudienceHideIfConverted AdlabGroupUpdateV30AdInfoAudienceHideIfConverted = "ORGANIZATION"
)

// All allowed values of AdlabGroupUpdateV30AdInfoAudienceHideIfConverted enum
var AllowedAdlabGroupUpdateV30AdInfoAudienceHideIfConvertedEnumValues = []AdlabGroupUpdateV30AdInfoAudienceHideIfConverted{
	"AD",
	"ADVERTISER",
	"APP",
	"CAMPAIGN",
	"CUSTOMER",
	"NO_EXCLUDE",
	"ORGANIZATION",
}

// NewAdlabGroupUpdateV30AdInfoAudienceHideIfConvertedFromValue returns a pointer to a valid AdlabGroupUpdateV30AdInfoAudienceHideIfConverted
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateV30AdInfoAudienceHideIfConvertedFromValue(v string) (*AdlabGroupUpdateV30AdInfoAudienceHideIfConverted, error) {
	ev := AdlabGroupUpdateV30AdInfoAudienceHideIfConverted(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateV30AdInfoAudienceHideIfConverted: valid values are %v", v, AllowedAdlabGroupUpdateV30AdInfoAudienceHideIfConvertedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateV30AdInfoAudienceHideIfConverted) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateV30AdInfoAudienceHideIfConvertedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_v3.0_ad_info_audience_hide_if_converted value
func (v AdlabGroupUpdateV30AdInfoAudienceHideIfConverted) Ptr() *AdlabGroupUpdateV30AdInfoAudienceHideIfConverted {
	return &v
}
