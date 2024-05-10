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

// QianchuanReportSearchWordGetV10OrderType
type QianchuanReportSearchWordGetV10OrderType string

// List of qianchuan_report_search_word_get_v1.0_order_type
const (
	ASC_QianchuanReportSearchWordGetV10OrderType  QianchuanReportSearchWordGetV10OrderType = "ASC"
	DESC_QianchuanReportSearchWordGetV10OrderType QianchuanReportSearchWordGetV10OrderType = "DESC"
)

// All allowed values of QianchuanReportSearchWordGetV10OrderType enum
var AllowedQianchuanReportSearchWordGetV10OrderTypeEnumValues = []QianchuanReportSearchWordGetV10OrderType{
	"ASC",
	"DESC",
}

// NewQianchuanReportSearchWordGetV10OrderTypeFromValue returns a pointer to a valid QianchuanReportSearchWordGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportSearchWordGetV10OrderTypeFromValue(v string) (*QianchuanReportSearchWordGetV10OrderType, error) {
	ev := QianchuanReportSearchWordGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportSearchWordGetV10OrderType: valid values are %v", v, AllowedQianchuanReportSearchWordGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportSearchWordGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportSearchWordGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_search_word_get_v1.0_order_type value
func (v QianchuanReportSearchWordGetV10OrderType) Ptr() *QianchuanReportSearchWordGetV10OrderType {
	return &v
}
