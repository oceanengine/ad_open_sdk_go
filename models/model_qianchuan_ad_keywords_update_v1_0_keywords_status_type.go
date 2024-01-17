/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdKeywordsUpdateV10KeywordsStatusType
type QianchuanAdKeywordsUpdateV10KeywordsStatusType string

// List of qianchuan_ad_keywords_update_v1.0_keywords_status_type
const (
	DELETE_QianchuanAdKeywordsUpdateV10KeywordsStatusType QianchuanAdKeywordsUpdateV10KeywordsStatusType = "DELETE"
	ENABLE_QianchuanAdKeywordsUpdateV10KeywordsStatusType QianchuanAdKeywordsUpdateV10KeywordsStatusType = "ENABLE"
	PAUSED_QianchuanAdKeywordsUpdateV10KeywordsStatusType QianchuanAdKeywordsUpdateV10KeywordsStatusType = "PAUSED"
)

// All allowed values of QianchuanAdKeywordsUpdateV10KeywordsStatusType enum
var AllowedQianchuanAdKeywordsUpdateV10KeywordsStatusTypeEnumValues = []QianchuanAdKeywordsUpdateV10KeywordsStatusType{
	"DELETE",
	"ENABLE",
	"PAUSED",
}

// NewQianchuanAdKeywordsUpdateV10KeywordsStatusTypeFromValue returns a pointer to a valid QianchuanAdKeywordsUpdateV10KeywordsStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdKeywordsUpdateV10KeywordsStatusTypeFromValue(v string) (*QianchuanAdKeywordsUpdateV10KeywordsStatusType, error) {
	ev := QianchuanAdKeywordsUpdateV10KeywordsStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdKeywordsUpdateV10KeywordsStatusType: valid values are %v", v, AllowedQianchuanAdKeywordsUpdateV10KeywordsStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdKeywordsUpdateV10KeywordsStatusType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdKeywordsUpdateV10KeywordsStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_keywords_update_v1.0_keywords_status_type value
func (v QianchuanAdKeywordsUpdateV10KeywordsStatusType) Ptr() *QianchuanAdKeywordsUpdateV10KeywordsStatusType {
	return &v
}
