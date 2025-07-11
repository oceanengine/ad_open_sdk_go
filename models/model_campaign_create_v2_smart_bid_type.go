/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignCreateV2SmartBidType
type CampaignCreateV2SmartBidType string

// List of campaign_create_v2_smart_bid_type
const (
	MAX_CONVERSION_CampaignCreateV2SmartBidType    CampaignCreateV2SmartBidType = "MAX_CONVERSION"
	AWEME_LITE_PACING_CampaignCreateV2SmartBidType CampaignCreateV2SmartBidType = "AWEME_LITE_PACING"
	SMART_BID_NO_BID_CampaignCreateV2SmartBidType  CampaignCreateV2SmartBidType = "SMART_BID_NO_BID"
	MAXCV_CampaignCreateV2SmartBidType             CampaignCreateV2SmartBidType = "MAXCV"
	JIJIN_CampaignCreateV2SmartBidType             CampaignCreateV2SmartBidType = "JIJIN"
	LITE_PACING_CampaignCreateV2SmartBidType       CampaignCreateV2SmartBidType = "LITE_PACING"
	BAOSHOU_CampaignCreateV2SmartBidType           CampaignCreateV2SmartBidType = "BAOSHOU"
	CUSTOM_CampaignCreateV2SmartBidType            CampaignCreateV2SmartBidType = "CUSTOM"
	GD_PROMOTE_CampaignCreateV2SmartBidType        CampaignCreateV2SmartBidType = "GD_PROMOTE"
	GUARANTEED_SHOW_CampaignCreateV2SmartBidType   CampaignCreateV2SmartBidType = "GUARANTEED_SHOW"
)

// Ptr returns reference to campaign_create_v2_smart_bid_type value
func (v CampaignCreateV2SmartBidType) Ptr() *CampaignCreateV2SmartBidType {
	return &v
}
