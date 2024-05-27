/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10KeywordsMatchType
type QianchuanAdCreateV10KeywordsMatchType string

// List of qianchuan_ad_create_v1.0_keywords_match_type
const (
	EXTENSIVE_QianchuanAdCreateV10KeywordsMatchType QianchuanAdCreateV10KeywordsMatchType = "EXTENSIVE"
	PHRASE_QianchuanAdCreateV10KeywordsMatchType    QianchuanAdCreateV10KeywordsMatchType = "PHRASE"
	PRECISION_QianchuanAdCreateV10KeywordsMatchType QianchuanAdCreateV10KeywordsMatchType = "PRECISION"
)

// All allowed values of QianchuanAdCreateV10KeywordsMatchType enum
var AllowedQianchuanAdCreateV10KeywordsMatchTypeEnumValues = []QianchuanAdCreateV10KeywordsMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewQianchuanAdCreateV10KeywordsMatchTypeFromValue returns a pointer to a valid QianchuanAdCreateV10KeywordsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10KeywordsMatchTypeFromValue(v string) (*QianchuanAdCreateV10KeywordsMatchType, error) {
	ev := QianchuanAdCreateV10KeywordsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10KeywordsMatchType: valid values are %v", v, AllowedQianchuanAdCreateV10KeywordsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10KeywordsMatchType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10KeywordsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_keywords_match_type value
func (v QianchuanAdCreateV10KeywordsMatchType) Ptr() *QianchuanAdCreateV10KeywordsMatchType {
	return &v
}
