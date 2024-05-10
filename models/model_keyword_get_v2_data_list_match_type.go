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

// KeywordGetV2DataListMatchType
type KeywordGetV2DataListMatchType string

// List of keyword_get_v2_data_list_match_type
const (
	PRECISION_KeywordGetV2DataListMatchType KeywordGetV2DataListMatchType = "PRECISION"
	PHRASE_KeywordGetV2DataListMatchType    KeywordGetV2DataListMatchType = "PHRASE"
	EXTENSIVE_KeywordGetV2DataListMatchType KeywordGetV2DataListMatchType = "EXTENSIVE"
)

// All allowed values of KeywordGetV2DataListMatchType enum
var AllowedKeywordGetV2DataListMatchTypeEnumValues = []KeywordGetV2DataListMatchType{
	"PRECISION",
	"PHRASE",
	"EXTENSIVE",
}

// NewKeywordGetV2DataListMatchTypeFromValue returns a pointer to a valid KeywordGetV2DataListMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordGetV2DataListMatchTypeFromValue(v string) (*KeywordGetV2DataListMatchType, error) {
	ev := KeywordGetV2DataListMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordGetV2DataListMatchType: valid values are %v", v, AllowedKeywordGetV2DataListMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordGetV2DataListMatchType) IsValid() bool {
	for _, existing := range AllowedKeywordGetV2DataListMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_get_v2_data_list_match_type value
func (v KeywordGetV2DataListMatchType) Ptr() *KeywordGetV2DataListMatchType {
	return &v
}
