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

// KeywordUpdateV2V2DataErrorListMatchType
type KeywordUpdateV2V2DataErrorListMatchType string

// List of keyword_update_v2_v2_data_error_list_match_type
const (
	EXTENSIVE_KeywordUpdateV2V2DataErrorListMatchType KeywordUpdateV2V2DataErrorListMatchType = "EXTENSIVE"
	PHRASE_KeywordUpdateV2V2DataErrorListMatchType    KeywordUpdateV2V2DataErrorListMatchType = "PHRASE"
	PRECISION_KeywordUpdateV2V2DataErrorListMatchType KeywordUpdateV2V2DataErrorListMatchType = "PRECISION"
)

// All allowed values of KeywordUpdateV2V2DataErrorListMatchType enum
var AllowedKeywordUpdateV2V2DataErrorListMatchTypeEnumValues = []KeywordUpdateV2V2DataErrorListMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewKeywordUpdateV2V2DataErrorListMatchTypeFromValue returns a pointer to a valid KeywordUpdateV2V2DataErrorListMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordUpdateV2V2DataErrorListMatchTypeFromValue(v string) (*KeywordUpdateV2V2DataErrorListMatchType, error) {
	ev := KeywordUpdateV2V2DataErrorListMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordUpdateV2V2DataErrorListMatchType: valid values are %v", v, AllowedKeywordUpdateV2V2DataErrorListMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordUpdateV2V2DataErrorListMatchType) IsValid() bool {
	for _, existing := range AllowedKeywordUpdateV2V2DataErrorListMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_update_v2_v2_data_error_list_match_type value
func (v KeywordUpdateV2V2DataErrorListMatchType) Ptr() *KeywordUpdateV2V2DataErrorListMatchType {
	return &v
}
