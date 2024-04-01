/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CampaignGetV2FilteringSmartBidType
type CampaignGetV2FilteringSmartBidType string

// List of campaign_get_v2_filtering_smart_bid_type
const (
	JIJIN_CampaignGetV2FilteringSmartBidType             CampaignGetV2FilteringSmartBidType = "JIJIN"
	LITE_PACING_CampaignGetV2FilteringSmartBidType       CampaignGetV2FilteringSmartBidType = "LITE_PACING"
	GD_PROMOTE_CampaignGetV2FilteringSmartBidType        CampaignGetV2FilteringSmartBidType = "GD_PROMOTE"
	MAXCV_CampaignGetV2FilteringSmartBidType             CampaignGetV2FilteringSmartBidType = "MAXCV"
	CUSTOM_CampaignGetV2FilteringSmartBidType            CampaignGetV2FilteringSmartBidType = "CUSTOM"
	BAOSHOU_CampaignGetV2FilteringSmartBidType           CampaignGetV2FilteringSmartBidType = "BAOSHOU"
	MAX_CONVERSION_CampaignGetV2FilteringSmartBidType    CampaignGetV2FilteringSmartBidType = "MAX_CONVERSION"
	GUARANTEED_SHOW_CampaignGetV2FilteringSmartBidType   CampaignGetV2FilteringSmartBidType = "GUARANTEED_SHOW"
	AWEME_LITE_PACING_CampaignGetV2FilteringSmartBidType CampaignGetV2FilteringSmartBidType = "AWEME_LITE_PACING"
	SMART_BID_NO_BID_CampaignGetV2FilteringSmartBidType  CampaignGetV2FilteringSmartBidType = "SMART_BID_NO_BID"
)

// All allowed values of CampaignGetV2FilteringSmartBidType enum
var AllowedCampaignGetV2FilteringSmartBidTypeEnumValues = []CampaignGetV2FilteringSmartBidType{
	"JIJIN",
	"LITE_PACING",
	"GD_PROMOTE",
	"MAXCV",
	"CUSTOM",
	"BAOSHOU",
	"MAX_CONVERSION",
	"GUARANTEED_SHOW",
	"AWEME_LITE_PACING",
	"SMART_BID_NO_BID",
}

// NewCampaignGetV2FilteringSmartBidTypeFromValue returns a pointer to a valid CampaignGetV2FilteringSmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2FilteringSmartBidTypeFromValue(v string) (*CampaignGetV2FilteringSmartBidType, error) {
	ev := CampaignGetV2FilteringSmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2FilteringSmartBidType: valid values are %v", v, AllowedCampaignGetV2FilteringSmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2FilteringSmartBidType) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2FilteringSmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_filtering_smart_bid_type value
func (v CampaignGetV2FilteringSmartBidType) Ptr() *CampaignGetV2FilteringSmartBidType {
	return &v
}
