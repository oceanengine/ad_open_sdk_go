/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataSmartBidType
type AdGetV2DataSmartBidType string

// List of ad_get_v2_data_smart_bid_type
const (
	SMART_BID_CUSTOM_AdGetV2DataSmartBidType       AdGetV2DataSmartBidType = "SMART_BID_CUSTOM"
	SMART_BID_RADICAL_AdGetV2DataSmartBidType      AdGetV2DataSmartBidType = "SMART_BID_RADICAL"
	SMART_BID_CONSERVATIVE_AdGetV2DataSmartBidType AdGetV2DataSmartBidType = "SMART_BID_CONSERVATIVE"
	SMART_BID_NO_BID_AdGetV2DataSmartBidType       AdGetV2DataSmartBidType = "SMART_BID_NO_BID"
)

// Ptr returns reference to ad_get_v2_data_smart_bid_type value
func (v AdGetV2DataSmartBidType) Ptr() *AdGetV2DataSmartBidType {
	return &v
}
