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
	DPA_CampaignCreateV2LandingType            CampaignCreateV2LandingType = "DPA"
	APP_CampaignCreateV2LandingType            CampaignCreateV2LandingType = "APP"
	LINK_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "LINK"
	SHOP_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "SHOP"
	BRAND_EXTERNAL_CampaignCreateV2LandingType CampaignCreateV2LandingType = "BRAND_EXTERNAL"
	ARTICLE_CampaignCreateV2LandingType        CampaignCreateV2LandingType = "ARTICLE"
	AWEME_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "AWEME"
	GOODS_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "GOODS"
	LIVE_CampaignCreateV2LandingType           CampaignCreateV2LandingType = "LIVE"
	STORE_CampaignCreateV2LandingType          CampaignCreateV2LandingType = "STORE"
	QUICK_APP_CampaignCreateV2LandingType      CampaignCreateV2LandingType = "QUICK_APP"
)

// Ptr returns reference to campaign_create_v2_landing_type value
func (v CampaignCreateV2LandingType) Ptr() *CampaignCreateV2LandingType {
	return &v
}
