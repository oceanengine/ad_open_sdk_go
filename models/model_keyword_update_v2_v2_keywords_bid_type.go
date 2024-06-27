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

// KeywordUpdateV2V2KeywordsBidType
type KeywordUpdateV2V2KeywordsBidType string

// List of keyword_update_v2_v2_keywords_bid_type
const (
	CUSTOM_KeywordUpdateV2V2KeywordsBidType         KeywordUpdateV2V2KeywordsBidType = "CUSTOM"
	SUGGEST_KeywordUpdateV2V2KeywordsBidType        KeywordUpdateV2V2KeywordsBidType = "SUGGEST"
	FEED_TO_SEARCH_KeywordUpdateV2V2KeywordsBidType KeywordUpdateV2V2KeywordsBidType = "FEED_TO_SEARCH"
	WITH_AD_KeywordUpdateV2V2KeywordsBidType        KeywordUpdateV2V2KeywordsBidType = "WITH_AD"
	BRAND_AD_KeywordUpdateV2V2KeywordsBidType       KeywordUpdateV2V2KeywordsBidType = "BRAND_AD"
)

// All allowed values of KeywordUpdateV2V2KeywordsBidType enum
var AllowedKeywordUpdateV2V2KeywordsBidTypeEnumValues = []KeywordUpdateV2V2KeywordsBidType{
	"CUSTOM",
	"SUGGEST",
	"FEED_TO_SEARCH",
	"WITH_AD",
	"BRAND_AD",
}

// NewKeywordUpdateV2V2KeywordsBidTypeFromValue returns a pointer to a valid KeywordUpdateV2V2KeywordsBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordUpdateV2V2KeywordsBidTypeFromValue(v string) (*KeywordUpdateV2V2KeywordsBidType, error) {
	ev := KeywordUpdateV2V2KeywordsBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordUpdateV2V2KeywordsBidType: valid values are %v", v, AllowedKeywordUpdateV2V2KeywordsBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordUpdateV2V2KeywordsBidType) IsValid() bool {
	for _, existing := range AllowedKeywordUpdateV2V2KeywordsBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_update_v2_v2_keywords_bid_type value
func (v KeywordUpdateV2V2KeywordsBidType) Ptr() *KeywordUpdateV2V2KeywordsBidType {
	return &v
}
