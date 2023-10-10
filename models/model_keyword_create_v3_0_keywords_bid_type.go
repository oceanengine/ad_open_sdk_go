/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// KeywordCreateV30KeywordsBidType
type KeywordCreateV30KeywordsBidType string

// List of keyword_create_v3.0_keywords_bid_type
const (
	CUSTOM_KeywordCreateV30KeywordsBidType         KeywordCreateV30KeywordsBidType = "CUSTOM"
	WITH_PROMOTION_KeywordCreateV30KeywordsBidType KeywordCreateV30KeywordsBidType = "WITH_PROMOTION"
)

// All allowed values of KeywordCreateV30KeywordsBidType enum
var AllowedKeywordCreateV30KeywordsBidTypeEnumValues = []KeywordCreateV30KeywordsBidType{
	"CUSTOM",
	"WITH_PROMOTION",
}

// NewKeywordCreateV30KeywordsBidTypeFromValue returns a pointer to a valid KeywordCreateV30KeywordsBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordCreateV30KeywordsBidTypeFromValue(v string) (*KeywordCreateV30KeywordsBidType, error) {
	ev := KeywordCreateV30KeywordsBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordCreateV30KeywordsBidType: valid values are %v", v, AllowedKeywordCreateV30KeywordsBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordCreateV30KeywordsBidType) IsValid() bool {
	for _, existing := range AllowedKeywordCreateV30KeywordsBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_create_v3.0_keywords_bid_type value
func (v KeywordCreateV30KeywordsBidType) Ptr() *KeywordCreateV30KeywordsBidType {
	return &v
}