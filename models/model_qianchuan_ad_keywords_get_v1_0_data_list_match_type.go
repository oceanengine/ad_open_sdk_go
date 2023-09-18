/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdKeywordsGetV10DataListMatchType
type QianchuanAdKeywordsGetV10DataListMatchType string

// List of qianchuan_ad_keywords_get_v1.0_data_list_match_type
const (
	EXTENSIVE_QianchuanAdKeywordsGetV10DataListMatchType QianchuanAdKeywordsGetV10DataListMatchType = "EXTENSIVE"
	PHRASE_QianchuanAdKeywordsGetV10DataListMatchType    QianchuanAdKeywordsGetV10DataListMatchType = "PHRASE"
	PRECISION_QianchuanAdKeywordsGetV10DataListMatchType QianchuanAdKeywordsGetV10DataListMatchType = "PRECISION"
)

// All allowed values of QianchuanAdKeywordsGetV10DataListMatchType enum
var AllowedQianchuanAdKeywordsGetV10DataListMatchTypeEnumValues = []QianchuanAdKeywordsGetV10DataListMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewQianchuanAdKeywordsGetV10DataListMatchTypeFromValue returns a pointer to a valid QianchuanAdKeywordsGetV10DataListMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdKeywordsGetV10DataListMatchTypeFromValue(v string) (*QianchuanAdKeywordsGetV10DataListMatchType, error) {
	ev := QianchuanAdKeywordsGetV10DataListMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdKeywordsGetV10DataListMatchType: valid values are %v", v, AllowedQianchuanAdKeywordsGetV10DataListMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdKeywordsGetV10DataListMatchType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdKeywordsGetV10DataListMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_keywords_get_v1.0_data_list_match_type value
func (v QianchuanAdKeywordsGetV10DataListMatchType) Ptr() *QianchuanAdKeywordsGetV10DataListMatchType {
	return &v
}