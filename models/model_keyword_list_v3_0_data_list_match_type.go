/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// KeywordListV30DataListMatchType
type KeywordListV30DataListMatchType string

// List of keyword_list_v3.0_data_list_match_type
const (
	EXTENSIVE_KeywordListV30DataListMatchType KeywordListV30DataListMatchType = "EXTENSIVE"
	PHRASE_KeywordListV30DataListMatchType    KeywordListV30DataListMatchType = "PHRASE"
	PRECISION_KeywordListV30DataListMatchType KeywordListV30DataListMatchType = "PRECISION"
)

// All allowed values of KeywordListV30DataListMatchType enum
var AllowedKeywordListV30DataListMatchTypeEnumValues = []KeywordListV30DataListMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewKeywordListV30DataListMatchTypeFromValue returns a pointer to a valid KeywordListV30DataListMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordListV30DataListMatchTypeFromValue(v string) (*KeywordListV30DataListMatchType, error) {
	ev := KeywordListV30DataListMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordListV30DataListMatchType: valid values are %v", v, AllowedKeywordListV30DataListMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordListV30DataListMatchType) IsValid() bool {
	for _, existing := range AllowedKeywordListV30DataListMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_list_v3.0_data_list_match_type value
func (v KeywordListV30DataListMatchType) Ptr() *KeywordListV30DataListMatchType {
	return &v
}
