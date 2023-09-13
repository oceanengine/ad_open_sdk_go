/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportSearchWordGetV10WordType
type QianchuanReportSearchWordGetV10WordType string

// List of qianchuan_report_search_word_get_v1.0_word_type
const (
	KEY_WORD_QianchuanReportSearchWordGetV10WordType    QianchuanReportSearchWordGetV10WordType = "KEY_WORD"
	SEARCH_WORD_QianchuanReportSearchWordGetV10WordType QianchuanReportSearchWordGetV10WordType = "SEARCH_WORD"
)

// All allowed values of QianchuanReportSearchWordGetV10WordType enum
var AllowedQianchuanReportSearchWordGetV10WordTypeEnumValues = []QianchuanReportSearchWordGetV10WordType{
	"KEY_WORD",
	"SEARCH_WORD",
}

// NewQianchuanReportSearchWordGetV10WordTypeFromValue returns a pointer to a valid QianchuanReportSearchWordGetV10WordType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportSearchWordGetV10WordTypeFromValue(v string) (*QianchuanReportSearchWordGetV10WordType, error) {
	ev := QianchuanReportSearchWordGetV10WordType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportSearchWordGetV10WordType: valid values are %v", v, AllowedQianchuanReportSearchWordGetV10WordTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportSearchWordGetV10WordType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportSearchWordGetV10WordTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_search_word_get_v1.0_word_type value
func (v QianchuanReportSearchWordGetV10WordType) Ptr() *QianchuanReportSearchWordGetV10WordType {
	return &v
}
