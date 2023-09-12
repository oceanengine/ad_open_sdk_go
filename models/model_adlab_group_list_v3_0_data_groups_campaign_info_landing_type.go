/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupListV30DataGroupsCampaignInfoLandingType
type AdlabGroupListV30DataGroupsCampaignInfoLandingType string

// List of adlab_group_list_v3.0_data_groups_campaign_info_landing_type
const (
	APP_AdlabGroupListV30DataGroupsCampaignInfoLandingType  AdlabGroupListV30DataGroupsCampaignInfoLandingType = "APP"
	DPA_AdlabGroupListV30DataGroupsCampaignInfoLandingType  AdlabGroupListV30DataGroupsCampaignInfoLandingType = "DPA"
	LINK_AdlabGroupListV30DataGroupsCampaignInfoLandingType AdlabGroupListV30DataGroupsCampaignInfoLandingType = "LINK"
	LIVE_AdlabGroupListV30DataGroupsCampaignInfoLandingType AdlabGroupListV30DataGroupsCampaignInfoLandingType = "LIVE"
	SHOP_AdlabGroupListV30DataGroupsCampaignInfoLandingType AdlabGroupListV30DataGroupsCampaignInfoLandingType = "SHOP"
)

// All allowed values of AdlabGroupListV30DataGroupsCampaignInfoLandingType enum
var AllowedAdlabGroupListV30DataGroupsCampaignInfoLandingTypeEnumValues = []AdlabGroupListV30DataGroupsCampaignInfoLandingType{
	"APP",
	"DPA",
	"LINK",
	"LIVE",
	"SHOP",
}

// NewAdlabGroupListV30DataGroupsCampaignInfoLandingTypeFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsCampaignInfoLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsCampaignInfoLandingTypeFromValue(v string) (*AdlabGroupListV30DataGroupsCampaignInfoLandingType, error) {
	ev := AdlabGroupListV30DataGroupsCampaignInfoLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsCampaignInfoLandingType: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsCampaignInfoLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsCampaignInfoLandingType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsCampaignInfoLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_campaign_info_landing_type value
func (v AdlabGroupListV30DataGroupsCampaignInfoLandingType) Ptr() *AdlabGroupListV30DataGroupsCampaignInfoLandingType {
	return &v
}
