/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignGetV2DataListLandingType
type CampaignGetV2DataListLandingType string

// List of campaign_get_v2_data_list_landing_type
const (
	APP_CampaignGetV2DataListLandingType            CampaignGetV2DataListLandingType = "APP"
	DPA_CampaignGetV2DataListLandingType            CampaignGetV2DataListLandingType = "DPA"
	BRAND_EXTERNAL_CampaignGetV2DataListLandingType CampaignGetV2DataListLandingType = "BRAND_EXTERNAL"
	LIVE_CampaignGetV2DataListLandingType           CampaignGetV2DataListLandingType = "LIVE"
	AWEME_CampaignGetV2DataListLandingType          CampaignGetV2DataListLandingType = "AWEME"
	SHOP_CampaignGetV2DataListLandingType           CampaignGetV2DataListLandingType = "SHOP"
	ARTICLE_CampaignGetV2DataListLandingType        CampaignGetV2DataListLandingType = "ARTICLE"
	LINK_CampaignGetV2DataListLandingType           CampaignGetV2DataListLandingType = "LINK"
	QUICK_APP_CampaignGetV2DataListLandingType      CampaignGetV2DataListLandingType = "QUICK_APP"
	STORE_CampaignGetV2DataListLandingType          CampaignGetV2DataListLandingType = "STORE"
	GOODS_CampaignGetV2DataListLandingType          CampaignGetV2DataListLandingType = "GOODS"
)

// Ptr returns reference to campaign_get_v2_data_list_landing_type value
func (v CampaignGetV2DataListLandingType) Ptr() *CampaignGetV2DataListLandingType {
	return &v
}
