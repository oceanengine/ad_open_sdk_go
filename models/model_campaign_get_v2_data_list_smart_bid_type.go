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
	GD_PROMOTE_CampaignGetV2DataListSmartBidType        CampaignGetV2DataListSmartBidType = "GD_PROMOTE"
	GUARANTEED_SHOW_CampaignGetV2DataListSmartBidType   CampaignGetV2DataListSmartBidType = "GUARANTEED_SHOW"
	SMART_BID_NO_BID_CampaignGetV2DataListSmartBidType  CampaignGetV2DataListSmartBidType = "SMART_BID_NO_BID"
	MAX_CONVERSION_CampaignGetV2DataListSmartBidType    CampaignGetV2DataListSmartBidType = "MAX_CONVERSION"
	CUSTOM_CampaignGetV2DataListSmartBidType            CampaignGetV2DataListSmartBidType = "CUSTOM"
	BAOSHOU_CampaignGetV2DataListSmartBidType           CampaignGetV2DataListSmartBidType = "BAOSHOU"
	MAXCV_CampaignGetV2DataListSmartBidType             CampaignGetV2DataListSmartBidType = "MAXCV"
	JIJIN_CampaignGetV2DataListSmartBidType             CampaignGetV2DataListSmartBidType = "JIJIN"
	AWEME_LITE_PACING_CampaignGetV2DataListSmartBidType CampaignGetV2DataListSmartBidType = "AWEME_LITE_PACING"
	LITE_PACING_CampaignGetV2DataListSmartBidType       CampaignGetV2DataListSmartBidType = "LITE_PACING"
)

// Ptr returns reference to campaign_get_v2_data_list_smart_bid_type value
func (v CampaignGetV2DataListSmartBidType) Ptr() *CampaignGetV2DataListSmartBidType {
	return &v
}
