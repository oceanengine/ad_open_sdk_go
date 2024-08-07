/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportAdvertiserGetV10FilteringSmartBidType
type QianchuanReportAdvertiserGetV10FilteringSmartBidType string

// List of qianchuan_report_advertiser_get_v1.0_filtering_smart_bid_type
const (
	SMART_BID_CONSERVATIVE_QianchuanReportAdvertiserGetV10FilteringSmartBidType QianchuanReportAdvertiserGetV10FilteringSmartBidType = "SMART_BID_CONSERVATIVE"
	SMART_BID_CUSTOM_QianchuanReportAdvertiserGetV10FilteringSmartBidType       QianchuanReportAdvertiserGetV10FilteringSmartBidType = "SMART_BID_CUSTOM"
)

// All allowed values of QianchuanReportAdvertiserGetV10FilteringSmartBidType enum
var AllowedQianchuanReportAdvertiserGetV10FilteringSmartBidTypeEnumValues = []QianchuanReportAdvertiserGetV10FilteringSmartBidType{
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_CUSTOM",
}

// NewQianchuanReportAdvertiserGetV10FilteringSmartBidTypeFromValue returns a pointer to a valid QianchuanReportAdvertiserGetV10FilteringSmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdvertiserGetV10FilteringSmartBidTypeFromValue(v string) (*QianchuanReportAdvertiserGetV10FilteringSmartBidType, error) {
	ev := QianchuanReportAdvertiserGetV10FilteringSmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdvertiserGetV10FilteringSmartBidType: valid values are %v", v, AllowedQianchuanReportAdvertiserGetV10FilteringSmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdvertiserGetV10FilteringSmartBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdvertiserGetV10FilteringSmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_advertiser_get_v1.0_filtering_smart_bid_type value
func (v QianchuanReportAdvertiserGetV10FilteringSmartBidType) Ptr() *QianchuanReportAdvertiserGetV10FilteringSmartBidType {
	return &v
}
