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

// QianchuanToolsSmartBoostAdBoostReportGetV10OrderType
type QianchuanToolsSmartBoostAdBoostReportGetV10OrderType string

// List of qianchuan_tools_smart_boost_ad_boost_report_get_v1.0_order_type
const (
	ASC_QianchuanToolsSmartBoostAdBoostReportGetV10OrderType  QianchuanToolsSmartBoostAdBoostReportGetV10OrderType = "ASC"
	DESC_QianchuanToolsSmartBoostAdBoostReportGetV10OrderType QianchuanToolsSmartBoostAdBoostReportGetV10OrderType = "DESC"
)

// All allowed values of QianchuanToolsSmartBoostAdBoostReportGetV10OrderType enum
var AllowedQianchuanToolsSmartBoostAdBoostReportGetV10OrderTypeEnumValues = []QianchuanToolsSmartBoostAdBoostReportGetV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanToolsSmartBoostAdBoostReportGetV10OrderTypeFromValue returns a pointer to a valid QianchuanToolsSmartBoostAdBoostReportGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsSmartBoostAdBoostReportGetV10OrderTypeFromValue(v string) (*QianchuanToolsSmartBoostAdBoostReportGetV10OrderType, error) {
	ev := QianchuanToolsSmartBoostAdBoostReportGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsSmartBoostAdBoostReportGetV10OrderType: valid values are %v", v, AllowedQianchuanToolsSmartBoostAdBoostReportGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsSmartBoostAdBoostReportGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsSmartBoostAdBoostReportGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_smart_boost_ad_boost_report_get_v1.0_order_type value
func (v QianchuanToolsSmartBoostAdBoostReportGetV10OrderType) Ptr() *QianchuanToolsSmartBoostAdBoostReportGetV10OrderType {
	return &v
}
