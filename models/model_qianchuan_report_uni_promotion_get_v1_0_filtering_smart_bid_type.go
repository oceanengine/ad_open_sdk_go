/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportUniPromotionGetV10FilteringSmartBidType
type QianchuanReportUniPromotionGetV10FilteringSmartBidType string

// List of qianchuan_report_uni_promotion_get_v1.0_filtering_smart_bid_type
const (
	SMART_BID_CONSERVATIVE_QianchuanReportUniPromotionGetV10FilteringSmartBidType QianchuanReportUniPromotionGetV10FilteringSmartBidType = "SMART_BID_CONSERVATIVE"
	SMART_BID_CUSTOM_QianchuanReportUniPromotionGetV10FilteringSmartBidType       QianchuanReportUniPromotionGetV10FilteringSmartBidType = "SMART_BID_CUSTOM"
)

// All allowed values of QianchuanReportUniPromotionGetV10FilteringSmartBidType enum
var AllowedQianchuanReportUniPromotionGetV10FilteringSmartBidTypeEnumValues = []QianchuanReportUniPromotionGetV10FilteringSmartBidType{
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_CUSTOM",
}

// NewQianchuanReportUniPromotionGetV10FilteringSmartBidTypeFromValue returns a pointer to a valid QianchuanReportUniPromotionGetV10FilteringSmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportUniPromotionGetV10FilteringSmartBidTypeFromValue(v string) (*QianchuanReportUniPromotionGetV10FilteringSmartBidType, error) {
	ev := QianchuanReportUniPromotionGetV10FilteringSmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportUniPromotionGetV10FilteringSmartBidType: valid values are %v", v, AllowedQianchuanReportUniPromotionGetV10FilteringSmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportUniPromotionGetV10FilteringSmartBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportUniPromotionGetV10FilteringSmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_uni_promotion_get_v1.0_filtering_smart_bid_type value
func (v QianchuanReportUniPromotionGetV10FilteringSmartBidType) Ptr() *QianchuanReportUniPromotionGetV10FilteringSmartBidType {
	return &v
}