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

// QianchuanAdKeywordsUpdateV10KeywordsMatchType
type QianchuanAdKeywordsUpdateV10KeywordsMatchType string

// List of qianchuan_ad_keywords_update_v1.0_keywords_match_type
const (
	EXTENSIVE_QianchuanAdKeywordsUpdateV10KeywordsMatchType QianchuanAdKeywordsUpdateV10KeywordsMatchType = "EXTENSIVE"
	PHRASE_QianchuanAdKeywordsUpdateV10KeywordsMatchType    QianchuanAdKeywordsUpdateV10KeywordsMatchType = "PHRASE"
	PRECISION_QianchuanAdKeywordsUpdateV10KeywordsMatchType QianchuanAdKeywordsUpdateV10KeywordsMatchType = "PRECISION"
)

// All allowed values of QianchuanAdKeywordsUpdateV10KeywordsMatchType enum
var AllowedQianchuanAdKeywordsUpdateV10KeywordsMatchTypeEnumValues = []QianchuanAdKeywordsUpdateV10KeywordsMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewQianchuanAdKeywordsUpdateV10KeywordsMatchTypeFromValue returns a pointer to a valid QianchuanAdKeywordsUpdateV10KeywordsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdKeywordsUpdateV10KeywordsMatchTypeFromValue(v string) (*QianchuanAdKeywordsUpdateV10KeywordsMatchType, error) {
	ev := QianchuanAdKeywordsUpdateV10KeywordsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdKeywordsUpdateV10KeywordsMatchType: valid values are %v", v, AllowedQianchuanAdKeywordsUpdateV10KeywordsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdKeywordsUpdateV10KeywordsMatchType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdKeywordsUpdateV10KeywordsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_keywords_update_v1.0_keywords_match_type value
func (v QianchuanAdKeywordsUpdateV10KeywordsMatchType) Ptr() *QianchuanAdKeywordsUpdateV10KeywordsMatchType {
	return &v
}
