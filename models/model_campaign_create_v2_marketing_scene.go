/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignCreateV2MarketingScene
type CampaignCreateV2MarketingScene string

// List of campaign_create_v2_marketing_scene
const (
	SOCIAL_CampaignCreateV2MarketingScene            CampaignCreateV2MarketingScene = "SOCIAL"
	MERCHANTS_CampaignCreateV2MarketingScene         CampaignCreateV2MarketingScene = "MERCHANTS"
	NOVEL_CampaignCreateV2MarketingScene             CampaignCreateV2MarketingScene = "NOVEL"
	PROMOTION_PURPOSE_CampaignCreateV2MarketingScene CampaignCreateV2MarketingScene = "PROMOTION_PURPOSE"
	GAME_PROMOTION_CampaignCreateV2MarketingScene    CampaignCreateV2MarketingScene = "GAME_PROMOTION"
	ECOMMERCE_CampaignCreateV2MarketingScene         CampaignCreateV2MarketingScene = "ECOMMERCE"
	VIDEO_INFO_CampaignCreateV2MarketingScene        CampaignCreateV2MarketingScene = "VIDEO_INFO"
	CAR_CampaignCreateV2MarketingScene               CampaignCreateV2MarketingScene = "CAR"
	GAME_SUBSCRIBE_CampaignCreateV2MarketingScene    CampaignCreateV2MarketingScene = "GAME_SUBSCRIBE"
	EDUCATION_CampaignCreateV2MarketingScene         CampaignCreateV2MarketingScene = "EDUCATION"
)

// Ptr returns reference to campaign_create_v2_marketing_scene value
func (v CampaignCreateV2MarketingScene) Ptr() *CampaignCreateV2MarketingScene {
	return &v
}
