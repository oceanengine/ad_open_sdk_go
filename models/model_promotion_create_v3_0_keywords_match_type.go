/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30KeywordsMatchType
type PromotionCreateV30KeywordsMatchType string

// List of promotion_create_v3.0_keywords_match_type
const (
	EXTENSIVE_PromotionCreateV30KeywordsMatchType PromotionCreateV30KeywordsMatchType = "EXTENSIVE"
	PHRASE_PromotionCreateV30KeywordsMatchType    PromotionCreateV30KeywordsMatchType = "PHRASE"
	PRECISION_PromotionCreateV30KeywordsMatchType PromotionCreateV30KeywordsMatchType = "PRECISION"
)

// All allowed values of PromotionCreateV30KeywordsMatchType enum
var AllowedPromotionCreateV30KeywordsMatchTypeEnumValues = []PromotionCreateV30KeywordsMatchType{
	"EXTENSIVE",
	"PHRASE",
	"PRECISION",
}

// NewPromotionCreateV30KeywordsMatchTypeFromValue returns a pointer to a valid PromotionCreateV30KeywordsMatchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30KeywordsMatchTypeFromValue(v string) (*PromotionCreateV30KeywordsMatchType, error) {
	ev := PromotionCreateV30KeywordsMatchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30KeywordsMatchType: valid values are %v", v, AllowedPromotionCreateV30KeywordsMatchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30KeywordsMatchType) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30KeywordsMatchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_keywords_match_type value
func (v PromotionCreateV30KeywordsMatchType) Ptr() *PromotionCreateV30KeywordsMatchType {
	return &v
}
