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
	AWEME_LITE_PACING_CampaignCreateV2SmartBidType CampaignCreateV2SmartBidType = "AWEME_LITE_PACING"
	JIJIN_CampaignCreateV2SmartBidType             CampaignCreateV2SmartBidType = "JIJIN"
	SMART_BID_NO_BID_CampaignCreateV2SmartBidType  CampaignCreateV2SmartBidType = "SMART_BID_NO_BID"
	MAX_CONVERSION_CampaignCreateV2SmartBidType    CampaignCreateV2SmartBidType = "MAX_CONVERSION"
	GD_PROMOTE_CampaignCreateV2SmartBidType        CampaignCreateV2SmartBidType = "GD_PROMOTE"
	CUSTOM_CampaignCreateV2SmartBidType            CampaignCreateV2SmartBidType = "CUSTOM"
	MAXCV_CampaignCreateV2SmartBidType             CampaignCreateV2SmartBidType = "MAXCV"
	GUARANTEED_SHOW_CampaignCreateV2SmartBidType   CampaignCreateV2SmartBidType = "GUARANTEED_SHOW"
	BAOSHOU_CampaignCreateV2SmartBidType           CampaignCreateV2SmartBidType = "BAOSHOU"
	LITE_PACING_CampaignCreateV2SmartBidType       CampaignCreateV2SmartBidType = "LITE_PACING"
)

// Ptr returns reference to campaign_create_v2_smart_bid_type value
func (v CampaignCreateV2SmartBidType) Ptr() *CampaignCreateV2SmartBidType {
	return &v
}
