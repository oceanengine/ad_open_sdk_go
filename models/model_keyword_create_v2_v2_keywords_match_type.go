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

// KeywordCreateV2V2KeywordsMatchType
type KeywordCreateV2V2KeywordsMatchType string

// List of keyword_create_v2_v2_keywords_match_type
const (
	PRECISION_KeywordCreateV2V2KeywordsMatchType KeywordCreateV2V2KeywordsMatchType = "PRECISION"
	EXTENSIVE_KeywordCreateV2V2KeywordsMatchType KeywordCreateV2V2KeywordsMatchType = "EXTENSIVE"
	PHRASE_KeywordCreateV2V2KeywordsMatchType    KeywordCreateV2V2KeywordsMatchType = "PHRASE"
)

// All allowed values of KeywordCreateV2V2KeywordsMatchType enum
var AllowedKeywordCreateV2V2KeywordsMatchTypeEnumValues = []KeywordCreateV2V2KeywordsMatchType{
	"PRECISION",
	"EXTENSIVE",
	"PHRASE",
}

// NewKeywordCreateV2V2KeywordsMatchTypeFromValue returns a pointer to a valid KeywordCreateV2V2KeywordsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordCreateV2V2KeywordsMatchTypeFromValue(v string) (*KeywordCreateV2V2KeywordsMatchType, error) {
	ev := KeywordCreateV2V2KeywordsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordCreateV2V2KeywordsMatchType: valid values are %v", v, AllowedKeywordCreateV2V2KeywordsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordCreateV2V2KeywordsMatchType) IsValid() bool {
	for _, existing := range AllowedKeywordCreateV2V2KeywordsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_create_v2_v2_keywords_match_type value
func (v KeywordCreateV2V2KeywordsMatchType) Ptr() *KeywordCreateV2V2KeywordsMatchType {
	return &v
}
