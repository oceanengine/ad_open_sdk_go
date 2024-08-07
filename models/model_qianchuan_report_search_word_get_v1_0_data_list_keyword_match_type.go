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

// QianchuanReportSearchWordGetV10DataListKeywordMatchType
type QianchuanReportSearchWordGetV10DataListKeywordMatchType string

// List of qianchuan_report_search_word_get_v1.0_data_list_keyword_match_type
const (
	ALL_QianchuanReportSearchWordGetV10DataListKeywordMatchType       QianchuanReportSearchWordGetV10DataListKeywordMatchType = "ALL"
	EXTENSIVE_QianchuanReportSearchWordGetV10DataListKeywordMatchType QianchuanReportSearchWordGetV10DataListKeywordMatchType = "EXTENSIVE"
	PHRASE_QianchuanReportSearchWordGetV10DataListKeywordMatchType    QianchuanReportSearchWordGetV10DataListKeywordMatchType = "PHRASE"
	PRECISION_QianchuanReportSearchWordGetV10DataListKeywordMatchType QianchuanReportSearchWordGetV10DataListKeywordMatchType = "PRECISION"
)

// All allowed values of QianchuanReportSearchWordGetV10DataListKeywordMatchType enum
var AllowedQianchuanReportSearchWordGetV10DataListKeywordMatchTypeEnumValues = []QianchuanReportSearchWordGetV10DataListKeywordMatchType{
	"ALL",
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewQianchuanReportSearchWordGetV10DataListKeywordMatchTypeFromValue returns a pointer to a valid QianchuanReportSearchWordGetV10DataListKeywordMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportSearchWordGetV10DataListKeywordMatchTypeFromValue(v string) (*QianchuanReportSearchWordGetV10DataListKeywordMatchType, error) {
	ev := QianchuanReportSearchWordGetV10DataListKeywordMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportSearchWordGetV10DataListKeywordMatchType: valid values are %v", v, AllowedQianchuanReportSearchWordGetV10DataListKeywordMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportSearchWordGetV10DataListKeywordMatchType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportSearchWordGetV10DataListKeywordMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_search_word_get_v1.0_data_list_keyword_match_type value
func (v QianchuanReportSearchWordGetV10DataListKeywordMatchType) Ptr() *QianchuanReportSearchWordGetV10DataListKeywordMatchType {
	return &v
}
