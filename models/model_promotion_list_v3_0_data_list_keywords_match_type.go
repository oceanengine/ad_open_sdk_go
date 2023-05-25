/*
API Title

巨量引擎开放平台

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListKeywordsMatchType
type PromotionListV30DataListKeywordsMatchType string

// List of promotion_list_v3.0_data_list_keywords_match_type
const (
	EXTENSIVE_PromotionListV30DataListKeywordsMatchType PromotionListV30DataListKeywordsMatchType = "EXTENSIVE"
	PHRASE_PromotionListV30DataListKeywordsMatchType    PromotionListV30DataListKeywordsMatchType = "PHRASE"
	PRECISION_PromotionListV30DataListKeywordsMatchType PromotionListV30DataListKeywordsMatchType = "PRECISION"
)

// All allowed values of PromotionListV30DataListKeywordsMatchType enum
var AllowedPromotionListV30DataListKeywordsMatchTypeEnumValues = []PromotionListV30DataListKeywordsMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewPromotionListV30DataListKeywordsMatchTypeFromValue returns a pointer to a valid PromotionListV30DataListKeywordsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListKeywordsMatchTypeFromValue(v string) (*PromotionListV30DataListKeywordsMatchType, error) {
	ev := PromotionListV30DataListKeywordsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListKeywordsMatchType: valid values are %v", v, AllowedPromotionListV30DataListKeywordsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListKeywordsMatchType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListKeywordsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_keywords_match_type value
func (v PromotionListV30DataListKeywordsMatchType) Ptr() *PromotionListV30DataListKeywordsMatchType {
	return &v
}
