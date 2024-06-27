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

// PromotionUpdateV30KeywordsMatchType
type PromotionUpdateV30KeywordsMatchType string

// List of promotion_update_v3.0_keywords_match_type
const (
	EXTENSIVE_PromotionUpdateV30KeywordsMatchType PromotionUpdateV30KeywordsMatchType = "EXTENSIVE"
	PHRASE_PromotionUpdateV30KeywordsMatchType    PromotionUpdateV30KeywordsMatchType = "PHRASE"
	PRECISION_PromotionUpdateV30KeywordsMatchType PromotionUpdateV30KeywordsMatchType = "PRECISION"
)

// All allowed values of PromotionUpdateV30KeywordsMatchType enum
var AllowedPromotionUpdateV30KeywordsMatchTypeEnumValues = []PromotionUpdateV30KeywordsMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewPromotionUpdateV30KeywordsMatchTypeFromValue returns a pointer to a valid PromotionUpdateV30KeywordsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30KeywordsMatchTypeFromValue(v string) (*PromotionUpdateV30KeywordsMatchType, error) {
	ev := PromotionUpdateV30KeywordsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30KeywordsMatchType: valid values are %v", v, AllowedPromotionUpdateV30KeywordsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30KeywordsMatchType) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30KeywordsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_keywords_match_type value
func (v PromotionUpdateV30KeywordsMatchType) Ptr() *PromotionUpdateV30KeywordsMatchType {
	return &v
}
