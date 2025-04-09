/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignCreateV2MarketingPurpose
type CampaignCreateV2MarketingPurpose string

// List of campaign_create_v2_marketing_purpose
const (
	UNLIMITED_CampaignCreateV2MarketingPurpose   CampaignCreateV2MarketingPurpose = "UNLIMITED"
	CONVERSION_CampaignCreateV2MarketingPurpose  CampaignCreateV2MarketingPurpose = "CONVERSION"
	ACKNOWLEDGE_CampaignCreateV2MarketingPurpose CampaignCreateV2MarketingPurpose = "ACKNOWLEDGE"
	INTENTION_CampaignCreateV2MarketingPurpose   CampaignCreateV2MarketingPurpose = "INTENTION"
)

// Ptr returns reference to campaign_create_v2_marketing_purpose value
func (v CampaignCreateV2MarketingPurpose) Ptr() *CampaignCreateV2MarketingPurpose {
	return &v
}
