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

// QianchuanReportSearchWordGetV10FilteringWordType
type QianchuanReportSearchWordGetV10FilteringWordType string

// List of qianchuan_report_search_word_get_v1.0_filtering_word_type
const (
	KEY_WORD_QianchuanReportSearchWordGetV10FilteringWordType    QianchuanReportSearchWordGetV10FilteringWordType = "KEY_WORD"
	SEARCH_WORD_QianchuanReportSearchWordGetV10FilteringWordType QianchuanReportSearchWordGetV10FilteringWordType = "SEARCH_WORD"
)

// All allowed values of QianchuanReportSearchWordGetV10FilteringWordType enum
var AllowedQianchuanReportSearchWordGetV10FilteringWordTypeEnumValues = []QianchuanReportSearchWordGetV10FilteringWordType{
	"KEY_WORD",
	"SEARCH_WORD",
}

// NewQianchuanReportSearchWordGetV10FilteringWordTypeFromValue returns a pointer to a valid QianchuanReportSearchWordGetV10FilteringWordType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportSearchWordGetV10FilteringWordTypeFromValue(v string) (*QianchuanReportSearchWordGetV10FilteringWordType, error) {
	ev := QianchuanReportSearchWordGetV10FilteringWordType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportSearchWordGetV10FilteringWordType: valid values are %v", v, AllowedQianchuanReportSearchWordGetV10FilteringWordTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportSearchWordGetV10FilteringWordType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportSearchWordGetV10FilteringWordTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_search_word_get_v1.0_filtering_word_type value
func (v QianchuanReportSearchWordGetV10FilteringWordType) Ptr() *QianchuanReportSearchWordGetV10FilteringWordType {
	return &v
}
