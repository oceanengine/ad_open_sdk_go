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

// QianchuanReportAdGetV10FilteringOrderPlatform
type QianchuanReportAdGetV10FilteringOrderPlatform string

// List of qianchuan_report_ad_get_v1.0_filtering_order_platform
const (
	ALL_QianchuanReportAdGetV10FilteringOrderPlatform       QianchuanReportAdGetV10FilteringOrderPlatform = "ALL"
	ECP_AWEME_QianchuanReportAdGetV10FilteringOrderPlatform QianchuanReportAdGetV10FilteringOrderPlatform = "ECP_AWEME"
	QIANCHUAN_QianchuanReportAdGetV10FilteringOrderPlatform QianchuanReportAdGetV10FilteringOrderPlatform = "QIANCHUAN"
)

// All allowed values of QianchuanReportAdGetV10FilteringOrderPlatform enum
var AllowedQianchuanReportAdGetV10FilteringOrderPlatformEnumValues = []QianchuanReportAdGetV10FilteringOrderPlatform{
	"ALL",
	"ECP_AWEME",
	"QIANCHUAN",
}

// NewQianchuanReportAdGetV10FilteringOrderPlatformFromValue returns a pointer to a valid QianchuanReportAdGetV10FilteringOrderPlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdGetV10FilteringOrderPlatformFromValue(v string) (*QianchuanReportAdGetV10FilteringOrderPlatform, error) {
	ev := QianchuanReportAdGetV10FilteringOrderPlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdGetV10FilteringOrderPlatform: valid values are %v", v, AllowedQianchuanReportAdGetV10FilteringOrderPlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdGetV10FilteringOrderPlatform) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdGetV10FilteringOrderPlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_ad_get_v1.0_filtering_order_platform value
func (v QianchuanReportAdGetV10FilteringOrderPlatform) Ptr() *QianchuanReportAdGetV10FilteringOrderPlatform {
	return &v
}
