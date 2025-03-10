/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignGetV2DataListSmartBidType
type CampaignGetV2DataListSmartBidType string

// List of campaign_get_v2_data_list_smart_bid_type
const (
	BAOSHOU_CampaignGetV2DataListSmartBidType           CampaignGetV2DataListSmartBidType = "BAOSHOU"
	LITE_PACING_CampaignGetV2DataListSmartBidType       CampaignGetV2DataListSmartBidType = "LITE_PACING"
	SMART_BID_NO_BID_CampaignGetV2DataListSmartBidType  CampaignGetV2DataListSmartBidType = "SMART_BID_NO_BID"
	JIJIN_CampaignGetV2DataListSmartBidType             CampaignGetV2DataListSmartBidType = "JIJIN"
	GD_PROMOTE_CampaignGetV2DataListSmartBidType        CampaignGetV2DataListSmartBidType = "GD_PROMOTE"
	MAXCV_CampaignGetV2DataListSmartBidType             CampaignGetV2DataListSmartBidType = "MAXCV"
	CUSTOM_CampaignGetV2DataListSmartBidType            CampaignGetV2DataListSmartBidType = "CUSTOM"
	AWEME_LITE_PACING_CampaignGetV2DataListSmartBidType CampaignGetV2DataListSmartBidType = "AWEME_LITE_PACING"
	GUARANTEED_SHOW_CampaignGetV2DataListSmartBidType   CampaignGetV2DataListSmartBidType = "GUARANTEED_SHOW"
	MAX_CONVERSION_CampaignGetV2DataListSmartBidType    CampaignGetV2DataListSmartBidType = "MAX_CONVERSION"
)

// Ptr returns reference to campaign_get_v2_data_list_smart_bid_type value
func (v CampaignGetV2DataListSmartBidType) Ptr() *CampaignGetV2DataListSmartBidType {
	return &v
}
