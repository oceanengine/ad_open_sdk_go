/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum
type AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum string

// List of adlab_group_detail_v3.0_data_data_campaign_info_delivery_related_num
const (
	CAMPAIGN_DPA_DEFAULT_NOT_AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum     AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum = "CAMPAIGN_DPA_DEFAULT_NOT"
	CAMPAIGN_DPA_SINGLE_DELIVERY_AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum = "CAMPAIGN_DPA_SINGLE_DELIVERY"
)

// All allowed values of AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum enum
var AllowedAdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNumEnumValues = []AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum{
	"CAMPAIGN_DPA_DEFAULT_NOT",
	"CAMPAIGN_DPA_SINGLE_DELIVERY",
}

// NewAdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNumFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNumFromValue(v string) (*AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum, error) {
	ev := AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_campaign_info_delivery_related_num value
func (v AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum) Ptr() *AdlabGroupDetailV30DataDataCampaignInfoDeliveryRelatedNum {
	return &v
}