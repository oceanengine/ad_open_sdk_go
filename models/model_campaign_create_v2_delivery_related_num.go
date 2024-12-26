/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignCreateV2DeliveryRelatedNum
type CampaignCreateV2DeliveryRelatedNum string

// List of campaign_create_v2_delivery_related_num
const (
	CAMPAIGN_DPA_SINGLE_DELIVERY_CampaignCreateV2DeliveryRelatedNum CampaignCreateV2DeliveryRelatedNum = "CAMPAIGN_DPA_SINGLE_DELIVERY"
	CAMPAIGN_DPA_MULTI_DELIVERY_CampaignCreateV2DeliveryRelatedNum  CampaignCreateV2DeliveryRelatedNum = "CAMPAIGN_DPA_MULTI_DELIVERY"
	CAMPAIGN_DPA_DEFAULT_NOT_CampaignCreateV2DeliveryRelatedNum     CampaignCreateV2DeliveryRelatedNum = "CAMPAIGN_DPA_DEFAULT_NOT"
)

// Ptr returns reference to campaign_create_v2_delivery_related_num value
func (v CampaignCreateV2DeliveryRelatedNum) Ptr() *CampaignCreateV2DeliveryRelatedNum {
	return &v
}
