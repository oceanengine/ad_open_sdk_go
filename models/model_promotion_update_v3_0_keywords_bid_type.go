/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionUpdateV30KeywordsBidType
type PromotionUpdateV30KeywordsBidType string

// List of promotion_update_v3.0_keywords_bid_type
const (
	CUSTOM_PromotionUpdateV30KeywordsBidType         PromotionUpdateV30KeywordsBidType = "CUSTOM"
	WITH_PROMOTION_PromotionUpdateV30KeywordsBidType PromotionUpdateV30KeywordsBidType = "WITH_PROMOTION"
)

// All allowed values of PromotionUpdateV30KeywordsBidType enum
var AllowedPromotionUpdateV30KeywordsBidTypeEnumValues = []PromotionUpdateV30KeywordsBidType{
	"CUSTOM",
	"WITH_PROMOTION",
}

// NewPromotionUpdateV30KeywordsBidTypeFromValue returns a pointer to a valid PromotionUpdateV30KeywordsBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30KeywordsBidTypeFromValue(v string) (*PromotionUpdateV30KeywordsBidType, error) {
	ev := PromotionUpdateV30KeywordsBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30KeywordsBidType: valid values are %v", v, AllowedPromotionUpdateV30KeywordsBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30KeywordsBidType) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30KeywordsBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_keywords_bid_type value
func (v PromotionUpdateV30KeywordsBidType) Ptr() *PromotionUpdateV30KeywordsBidType {
	return &v
}
