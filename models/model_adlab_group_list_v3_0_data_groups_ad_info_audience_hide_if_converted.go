/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted
type AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted string

// List of adlab_group_list_v3.0_data_groups_ad_info_audience_hide_if_converted
const (
	AD_AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted           AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted = "AD"
	ADVERTISER_AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted   AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted = "ADVERTISER"
	APP_AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted          AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted = "APP"
	CAMPAIGN_AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted     AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted = "CAMPAIGN"
	CUSTOMER_AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted     AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted = "CUSTOMER"
	NO_EXCLUDE_AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted   AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted = "NO_EXCLUDE"
	ORGANIZATION_AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted = "ORGANIZATION"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted enum
var AllowedAdlabGroupListV30DataGroupsAdInfoAudienceHideIfConvertedEnumValues = []AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted{
	"AD",
	"ADVERTISER",
	"APP",
	"CAMPAIGN",
	"CUSTOMER",
	"NO_EXCLUDE",
	"ORGANIZATION",
}

// NewAdlabGroupListV30DataGroupsAdInfoAudienceHideIfConvertedFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoAudienceHideIfConvertedFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoAudienceHideIfConvertedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoAudienceHideIfConvertedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_audience_hide_if_converted value
func (v AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted) Ptr() *AdlabGroupListV30DataGroupsAdInfoAudienceHideIfConverted {
	return &v
}
