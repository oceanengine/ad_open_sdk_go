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

// QianchuanReportCreativeGetV10OrderType
type QianchuanReportCreativeGetV10OrderType string

// List of qianchuan_report_creative_get_v1.0_order_type
const (
	ASC_QianchuanReportCreativeGetV10OrderType  QianchuanReportCreativeGetV10OrderType = "ASC"
	DESC_QianchuanReportCreativeGetV10OrderType QianchuanReportCreativeGetV10OrderType = "DESC"
)

// All allowed values of QianchuanReportCreativeGetV10OrderType enum
var AllowedQianchuanReportCreativeGetV10OrderTypeEnumValues = []QianchuanReportCreativeGetV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanReportCreativeGetV10OrderTypeFromValue returns a pointer to a valid QianchuanReportCreativeGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportCreativeGetV10OrderTypeFromValue(v string) (*QianchuanReportCreativeGetV10OrderType, error) {
	ev := QianchuanReportCreativeGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportCreativeGetV10OrderType: valid values are %v", v, AllowedQianchuanReportCreativeGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportCreativeGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportCreativeGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_creative_get_v1.0_order_type value
func (v QianchuanReportCreativeGetV10OrderType) Ptr() *QianchuanReportCreativeGetV10OrderType {
	return &v
}
