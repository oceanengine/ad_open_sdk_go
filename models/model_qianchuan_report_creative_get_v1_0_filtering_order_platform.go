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

// QianchuanReportCreativeGetV10FilteringOrderPlatform
type QianchuanReportCreativeGetV10FilteringOrderPlatform string

// List of qianchuan_report_creative_get_v1.0_filtering_order_platform
const (
	ALL_QianchuanReportCreativeGetV10FilteringOrderPlatform       QianchuanReportCreativeGetV10FilteringOrderPlatform = "ALL"
	ECP_AWEME_QianchuanReportCreativeGetV10FilteringOrderPlatform QianchuanReportCreativeGetV10FilteringOrderPlatform = "ECP_AWEME"
	QIANCHUAN_QianchuanReportCreativeGetV10FilteringOrderPlatform QianchuanReportCreativeGetV10FilteringOrderPlatform = "QIANCHUAN"
)

// All allowed values of QianchuanReportCreativeGetV10FilteringOrderPlatform enum
var AllowedQianchuanReportCreativeGetV10FilteringOrderPlatformEnumValues = []QianchuanReportCreativeGetV10FilteringOrderPlatform{
	"ALL",
	"ECP_AWEME",
	"QIANCHUAN",
}

// NewQianchuanReportCreativeGetV10FilteringOrderPlatformFromValue returns a pointer to a valid QianchuanReportCreativeGetV10FilteringOrderPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportCreativeGetV10FilteringOrderPlatformFromValue(v string) (*QianchuanReportCreativeGetV10FilteringOrderPlatform, error) {
	ev := QianchuanReportCreativeGetV10FilteringOrderPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportCreativeGetV10FilteringOrderPlatform: valid values are %v", v, AllowedQianchuanReportCreativeGetV10FilteringOrderPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportCreativeGetV10FilteringOrderPlatform) IsValid() bool {
	for _, existing := range AllowedQianchuanReportCreativeGetV10FilteringOrderPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_creative_get_v1.0_filtering_order_platform value
func (v QianchuanReportCreativeGetV10FilteringOrderPlatform) Ptr() *QianchuanReportCreativeGetV10FilteringOrderPlatform {
	return &v
}
