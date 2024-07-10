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

// KeywordUpdateV2V2DataSuccessListMatchType
type KeywordUpdateV2V2DataSuccessListMatchType string

// List of keyword_update_v2_v2_data_success_list_match_type
const (
	PRECISION_KeywordUpdateV2V2DataSuccessListMatchType KeywordUpdateV2V2DataSuccessListMatchType = "PRECISION"
	EXTENSIVE_KeywordUpdateV2V2DataSuccessListMatchType KeywordUpdateV2V2DataSuccessListMatchType = "EXTENSIVE"
	PHRASE_KeywordUpdateV2V2DataSuccessListMatchType    KeywordUpdateV2V2DataSuccessListMatchType = "PHRASE"
)

// All allowed values of KeywordUpdateV2V2DataSuccessListMatchType enum
var AllowedKeywordUpdateV2V2DataSuccessListMatchTypeEnumValues = []KeywordUpdateV2V2DataSuccessListMatchType{
	"PRECISION",
	"EXTENSIVE",
	"PHRASE",
}

// NewKeywordUpdateV2V2DataSuccessListMatchTypeFromValue returns a pointer to a valid KeywordUpdateV2V2DataSuccessListMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordUpdateV2V2DataSuccessListMatchTypeFromValue(v string) (*KeywordUpdateV2V2DataSuccessListMatchType, error) {
	ev := KeywordUpdateV2V2DataSuccessListMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordUpdateV2V2DataSuccessListMatchType: valid values are %v", v, AllowedKeywordUpdateV2V2DataSuccessListMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordUpdateV2V2DataSuccessListMatchType) IsValid() bool {
	for _, existing := range AllowedKeywordUpdateV2V2DataSuccessListMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_update_v2_v2_data_success_list_match_type value
func (v KeywordUpdateV2V2DataSuccessListMatchType) Ptr() *KeywordUpdateV2V2DataSuccessListMatchType {
	return &v
}
