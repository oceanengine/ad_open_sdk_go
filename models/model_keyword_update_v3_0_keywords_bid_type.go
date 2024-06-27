/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// KeywordUpdateV30KeywordsBidType
type KeywordUpdateV30KeywordsBidType string

// List of keyword_update_v3.0_keywords_bid_type
const (
	CUSTOM_KeywordUpdateV30KeywordsBidType         KeywordUpdateV30KeywordsBidType = "CUSTOM"
	WITH_PROMOTION_KeywordUpdateV30KeywordsBidType KeywordUpdateV30KeywordsBidType = "WITH_PROMOTION"
)

// All allowed values of KeywordUpdateV30KeywordsBidType enum
var AllowedKeywordUpdateV30KeywordsBidTypeEnumValues = []KeywordUpdateV30KeywordsBidType{
	"CUSTOM",
	"WITH_PROMOTION",
}

// NewKeywordUpdateV30KeywordsBidTypeFromValue returns a pointer to a valid KeywordUpdateV30KeywordsBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordUpdateV30KeywordsBidTypeFromValue(v string) (*KeywordUpdateV30KeywordsBidType, error) {
	ev := KeywordUpdateV30KeywordsBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordUpdateV30KeywordsBidType: valid values are %v", v, AllowedKeywordUpdateV30KeywordsBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordUpdateV30KeywordsBidType) IsValid() bool {
	for _, existing := range AllowedKeywordUpdateV30KeywordsBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_update_v3.0_keywords_bid_type value
func (v KeywordUpdateV30KeywordsBidType) Ptr() *KeywordUpdateV30KeywordsBidType {
	return &v
}
