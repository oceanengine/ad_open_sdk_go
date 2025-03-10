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
	BAOSHOU_CampaignCreateV2SmartBidType           CampaignCreateV2SmartBidType = "BAOSHOU"
	LITE_PACING_CampaignCreateV2SmartBidType       CampaignCreateV2SmartBidType = "LITE_PACING"
	SMART_BID_NO_BID_CampaignCreateV2SmartBidType  CampaignCreateV2SmartBidType = "SMART_BID_NO_BID"
	JIJIN_CampaignCreateV2SmartBidType             CampaignCreateV2SmartBidType = "JIJIN"
	GD_PROMOTE_CampaignCreateV2SmartBidType        CampaignCreateV2SmartBidType = "GD_PROMOTE"
	MAXCV_CampaignCreateV2SmartBidType             CampaignCreateV2SmartBidType = "MAXCV"
	CUSTOM_CampaignCreateV2SmartBidType            CampaignCreateV2SmartBidType = "CUSTOM"
	AWEME_LITE_PACING_CampaignCreateV2SmartBidType CampaignCreateV2SmartBidType = "AWEME_LITE_PACING"
	GUARANTEED_SHOW_CampaignCreateV2SmartBidType   CampaignCreateV2SmartBidType = "GUARANTEED_SHOW"
	MAX_CONVERSION_CampaignCreateV2SmartBidType    CampaignCreateV2SmartBidType = "MAX_CONVERSION"
)

// Ptr returns reference to campaign_create_v2_smart_bid_type value
func (v CampaignCreateV2SmartBidType) Ptr() *CampaignCreateV2SmartBidType {
	return &v
}
