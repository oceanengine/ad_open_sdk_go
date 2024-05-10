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

// KeywordUpdateV2V2DataSuccessListBidType
type KeywordUpdateV2V2DataSuccessListBidType string

// List of keyword_update_v2_v2_data_success_list_bid_type
const (
	SUGGEST_KeywordUpdateV2V2DataSuccessListBidType        KeywordUpdateV2V2DataSuccessListBidType = "SUGGEST"
	WITH_AD_KeywordUpdateV2V2DataSuccessListBidType        KeywordUpdateV2V2DataSuccessListBidType = "WITH_AD"
	CUSTOM_KeywordUpdateV2V2DataSuccessListBidType         KeywordUpdateV2V2DataSuccessListBidType = "CUSTOM"
	FEED_TO_SEARCH_KeywordUpdateV2V2DataSuccessListBidType KeywordUpdateV2V2DataSuccessListBidType = "FEED_TO_SEARCH"
	BRAND_AD_KeywordUpdateV2V2DataSuccessListBidType       KeywordUpdateV2V2DataSuccessListBidType = "BRAND_AD"
)

// All allowed values of KeywordUpdateV2V2DataSuccessListBidType enum
var AllowedKeywordUpdateV2V2DataSuccessListBidTypeEnumValues = []KeywordUpdateV2V2DataSuccessListBidType{
	"SUGGEST",
	"WITH_AD",
	"CUSTOM",
	"FEED_TO_SEARCH",
	"BRAND_AD",
}

// NewKeywordUpdateV2V2DataSuccessListBidTypeFromValue returns a pointer to a valid KeywordUpdateV2V2DataSuccessListBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordUpdateV2V2DataSuccessListBidTypeFromValue(v string) (*KeywordUpdateV2V2DataSuccessListBidType, error) {
	ev := KeywordUpdateV2V2DataSuccessListBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordUpdateV2V2DataSuccessListBidType: valid values are %v", v, AllowedKeywordUpdateV2V2DataSuccessListBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordUpdateV2V2DataSuccessListBidType) IsValid() bool {
	for _, existing := range AllowedKeywordUpdateV2V2DataSuccessListBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_update_v2_v2_data_success_list_bid_type value
func (v KeywordUpdateV2V2DataSuccessListBidType) Ptr() *KeywordUpdateV2V2DataSuccessListBidType {
	return &v
}
