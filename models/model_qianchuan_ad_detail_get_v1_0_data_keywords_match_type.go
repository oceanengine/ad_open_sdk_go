/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataKeywordsMatchType
type QianchuanAdDetailGetV10DataKeywordsMatchType string

// List of qianchuan_ad_detail_get_v1.0_data_keywords_match_type
const (
	EXTENSIVE_QianchuanAdDetailGetV10DataKeywordsMatchType QianchuanAdDetailGetV10DataKeywordsMatchType = "EXTENSIVE"
	PHRASE_QianchuanAdDetailGetV10DataKeywordsMatchType    QianchuanAdDetailGetV10DataKeywordsMatchType = "PHRASE"
	PRECISION_QianchuanAdDetailGetV10DataKeywordsMatchType QianchuanAdDetailGetV10DataKeywordsMatchType = "PRECISION"
)

// All allowed values of QianchuanAdDetailGetV10DataKeywordsMatchType enum
var AllowedQianchuanAdDetailGetV10DataKeywordsMatchTypeEnumValues = []QianchuanAdDetailGetV10DataKeywordsMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewQianchuanAdDetailGetV10DataKeywordsMatchTypeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataKeywordsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataKeywordsMatchTypeFromValue(v string) (*QianchuanAdDetailGetV10DataKeywordsMatchType, error) {
	ev := QianchuanAdDetailGetV10DataKeywordsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataKeywordsMatchType: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataKeywordsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataKeywordsMatchType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataKeywordsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_keywords_match_type value
func (v QianchuanAdDetailGetV10DataKeywordsMatchType) Ptr() *QianchuanAdDetailGetV10DataKeywordsMatchType {
	return &v
}
