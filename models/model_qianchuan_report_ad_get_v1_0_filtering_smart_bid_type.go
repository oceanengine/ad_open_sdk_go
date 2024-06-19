/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportAdGetV10FilteringSmartBidType
type QianchuanReportAdGetV10FilteringSmartBidType string

// List of qianchuan_report_ad_get_v1.0_filtering_smart_bid_type
const (
	SMART_BID_CONSERVATIVE_QianchuanReportAdGetV10FilteringSmartBidType QianchuanReportAdGetV10FilteringSmartBidType = "SMART_BID_CONSERVATIVE"
	SMART_BID_CUSTOM_QianchuanReportAdGetV10FilteringSmartBidType       QianchuanReportAdGetV10FilteringSmartBidType = "SMART_BID_CUSTOM"
)

// All allowed values of QianchuanReportAdGetV10FilteringSmartBidType enum
var AllowedQianchuanReportAdGetV10FilteringSmartBidTypeEnumValues = []QianchuanReportAdGetV10FilteringSmartBidType{
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_CUSTOM",
}

// NewQianchuanReportAdGetV10FilteringSmartBidTypeFromValue returns a pointer to a valid QianchuanReportAdGetV10FilteringSmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdGetV10FilteringSmartBidTypeFromValue(v string) (*QianchuanReportAdGetV10FilteringSmartBidType, error) {
	ev := QianchuanReportAdGetV10FilteringSmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdGetV10FilteringSmartBidType: valid values are %v", v, AllowedQianchuanReportAdGetV10FilteringSmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdGetV10FilteringSmartBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdGetV10FilteringSmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_ad_get_v1.0_filtering_smart_bid_type value
func (v QianchuanReportAdGetV10FilteringSmartBidType) Ptr() *QianchuanReportAdGetV10FilteringSmartBidType {
	return &v
}
