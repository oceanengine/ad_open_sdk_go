/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataDeepBidType
type AdGetV2DataDeepBidType string

// List of ad_get_v2_data_deep_bid_type
const (
	MIN_SECOND_STAGE_AdGetV2DataDeepBidType             AdGetV2DataDeepBidType = "MIN_SECOND_STAGE"
	SMARTBID_AdGetV2DataDeepBidType                     AdGetV2DataDeepBidType = "SMARTBID"
	DEEP_BID_DEFAULT_AdGetV2DataDeepBidType             AdGetV2DataDeepBidType = "DEEP_BID_DEFAULT"
	ROI_COEFFICIENT_AdGetV2DataDeepBidType              AdGetV2DataDeepBidType = "ROI_COEFFICIENT"
	ALI_FL_AdGetV2DataDeepBidType                       AdGetV2DataDeepBidType = "ALI_FL"
	ROI_PACING_AdGetV2DataDeepBidType                   AdGetV2DataDeepBidType = "ROI_PACING"
	PACING_SECOND_STAGE_AdGetV2DataDeepBidType          AdGetV2DataDeepBidType = "PACING_SECOND_STAGE"
	SOCIAL_ROI_AdGetV2DataDeepBidType                   AdGetV2DataDeepBidType = "SOCIAL_ROI"
	DEEP_BID_MIN_AdGetV2DataDeepBidType                 AdGetV2DataDeepBidType = "DEEP_BID_MIN"
	AUTO_MIN_SECOND_STAGE_AdGetV2DataDeepBidType        AdGetV2DataDeepBidType = "AUTO_MIN_SECOND_STAGE"
	BID_PER_ACTION_AdGetV2DataDeepBidType               AdGetV2DataDeepBidType = "BID_PER_ACTION"
	ROI_DIRECT_MAIL_AdGetV2DataDeepBidType              AdGetV2DataDeepBidType = "ROI_DIRECT_MAIL"
	DEEP_BID_TYPE_RETENTION_DAYS_AdGetV2DataDeepBidType AdGetV2DataDeepBidType = "DEEP_BID_TYPE_RETENTION_DAYS"
	DEEP_BID_PACING_AdGetV2DataDeepBidType              AdGetV2DataDeepBidType = "DEEP_BID_PACING"
)

// Ptr returns reference to ad_get_v2_data_deep_bid_type value
func (v AdGetV2DataDeepBidType) Ptr() *AdGetV2DataDeepBidType {
	return &v
}
