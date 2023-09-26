/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// KeywordFeedadsSuggestV2DataListMatchType
type KeywordFeedadsSuggestV2DataListMatchType string

// List of keyword_feedads_suggest_v2_data_list_match_type
const (
	PRECISION_KeywordFeedadsSuggestV2DataListMatchType KeywordFeedadsSuggestV2DataListMatchType = "PRECISION"
	EXTENSIVE_KeywordFeedadsSuggestV2DataListMatchType KeywordFeedadsSuggestV2DataListMatchType = "EXTENSIVE"
	PHRASE_KeywordFeedadsSuggestV2DataListMatchType    KeywordFeedadsSuggestV2DataListMatchType = "PHRASE"
)

// All allowed values of KeywordFeedadsSuggestV2DataListMatchType enum
var AllowedKeywordFeedadsSuggestV2DataListMatchTypeEnumValues = []KeywordFeedadsSuggestV2DataListMatchType{
	"PRECISION",
	"EXTENSIVE",
	"PHRASE",
}

// NewKeywordFeedadsSuggestV2DataListMatchTypeFromValue returns a pointer to a valid KeywordFeedadsSuggestV2DataListMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordFeedadsSuggestV2DataListMatchTypeFromValue(v string) (*KeywordFeedadsSuggestV2DataListMatchType, error) {
	ev := KeywordFeedadsSuggestV2DataListMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordFeedadsSuggestV2DataListMatchType: valid values are %v", v, AllowedKeywordFeedadsSuggestV2DataListMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordFeedadsSuggestV2DataListMatchType) IsValid() bool {
	for _, existing := range AllowedKeywordFeedadsSuggestV2DataListMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_feedads_suggest_v2_data_list_match_type value
func (v KeywordFeedadsSuggestV2DataListMatchType) Ptr() *KeywordFeedadsSuggestV2DataListMatchType {
	return &v
}