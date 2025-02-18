/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignCreateV2LandingType
type CampaignCreateV2LandingType string

// List of campaign_create_v2_landing_type
const (
	APP_CampaignCreateV2LandingType            CampaignCreateV2LandingType = "APP"
	DPA_CampaignCreateV2LandingType            CampaignCreateV2LandingType = "DPA"
	BRAND_EXTERNAL_CampaignCreateV2LandingType CampaignCreateV2LandingType = "BRAND_EXTERNAL"
	LIVE_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "LIVE"
	AWEME_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "AWEME"
	SHOP_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "SHOP"
	ARTICLE_CampaignCreateV2LandingType        CampaignCreateV2LandingType = "ARTICLE"
	LINK_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "LINK"
	QUICK_APP_CampaignCreateV2LandingType      CampaignCreateV2LandingType = "QUICK_APP"
	STORE_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "STORE"
	GOODS_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "GOODS"
)

// Ptr returns reference to campaign_create_v2_landing_type value
func (v CampaignCreateV2LandingType) Ptr() *CampaignCreateV2LandingType {
	return &v
}
