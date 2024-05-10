/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportAdvertiserGetV10FilteringOrderPlatform
type QianchuanReportAdvertiserGetV10FilteringOrderPlatform string

// List of qianchuan_report_advertiser_get_v1.0_filtering_order_platform
const (
	ALL_QianchuanReportAdvertiserGetV10FilteringOrderPlatform       QianchuanReportAdvertiserGetV10FilteringOrderPlatform = "ALL"
	ECP_AWEME_QianchuanReportAdvertiserGetV10FilteringOrderPlatform QianchuanReportAdvertiserGetV10FilteringOrderPlatform = "ECP_AWEME"
	QIANCHUAN_QianchuanReportAdvertiserGetV10FilteringOrderPlatform QianchuanReportAdvertiserGetV10FilteringOrderPlatform = "QIANCHUAN"
)

// All allowed values of QianchuanReportAdvertiserGetV10FilteringOrderPlatform enum
var AllowedQianchuanReportAdvertiserGetV10FilteringOrderPlatformEnumValues = []QianchuanReportAdvertiserGetV10FilteringOrderPlatform{
	"ALL",
	"ECP_AWEME",
	"QIANCHUAN",
}

// NewQianchuanReportAdvertiserGetV10FilteringOrderPlatformFromValue returns a pointer to a valid QianchuanReportAdvertiserGetV10FilteringOrderPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdvertiserGetV10FilteringOrderPlatformFromValue(v string) (*QianchuanReportAdvertiserGetV10FilteringOrderPlatform, error) {
	ev := QianchuanReportAdvertiserGetV10FilteringOrderPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdvertiserGetV10FilteringOrderPlatform: valid values are %v", v, AllowedQianchuanReportAdvertiserGetV10FilteringOrderPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdvertiserGetV10FilteringOrderPlatform) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdvertiserGetV10FilteringOrderPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_advertiser_get_v1.0_filtering_order_platform value
func (v QianchuanReportAdvertiserGetV10FilteringOrderPlatform) Ptr() *QianchuanReportAdvertiserGetV10FilteringOrderPlatform {
	return &v
}
