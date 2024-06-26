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

// AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose
type AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose string

// List of adlab_group_detail_v3.0_data_data_campaign_info_marketing_purpose
const (
	ACKNOWLEDGE_AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose = "ACKNOWLEDGE"
	CONVERSION_AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose  AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose = "CONVERSION"
	INTENTION_AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose   AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose = "INTENTION"
	UNLIMITED_AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose   AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose = "UNLIMITED"
)

// All allowed values of AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose enum
var AllowedAdlabGroupDetailV30DataDataCampaignInfoMarketingPurposeEnumValues = []AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose{
	"ACKNOWLEDGE",
	"CONVERSION",
	"INTENTION",
	"UNLIMITED",
}

// NewAdlabGroupDetailV30DataDataCampaignInfoMarketingPurposeFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataCampaignInfoMarketingPurposeFromValue(v string) (*AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose, error) {
	ev := AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataCampaignInfoMarketingPurposeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataCampaignInfoMarketingPurposeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_campaign_info_marketing_purpose value
func (v AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose) Ptr() *AdlabGroupDetailV30DataDataCampaignInfoMarketingPurpose {
	return &v
}
