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

// KeywordCreateV2V2DataSuccessListBidType
type KeywordCreateV2V2DataSuccessListBidType string

// List of keyword_create_v2_v2_data_success_list_bid_type
const (
	WITH_AD_KeywordCreateV2V2DataSuccessListBidType        KeywordCreateV2V2DataSuccessListBidType = "WITH_AD"
	CUSTOM_KeywordCreateV2V2DataSuccessListBidType         KeywordCreateV2V2DataSuccessListBidType = "CUSTOM"
	BRAND_AD_KeywordCreateV2V2DataSuccessListBidType       KeywordCreateV2V2DataSuccessListBidType = "BRAND_AD"
	SUGGEST_KeywordCreateV2V2DataSuccessListBidType        KeywordCreateV2V2DataSuccessListBidType = "SUGGEST"
	FEED_TO_SEARCH_KeywordCreateV2V2DataSuccessListBidType KeywordCreateV2V2DataSuccessListBidType = "FEED_TO_SEARCH"
)

// All allowed values of KeywordCreateV2V2DataSuccessListBidType enum
var AllowedKeywordCreateV2V2DataSuccessListBidTypeEnumValues = []KeywordCreateV2V2DataSuccessListBidType{
	"WITH_AD",
	"CUSTOM",
	"BRAND_AD",
	"SUGGEST",
	"FEED_TO_SEARCH",
}

// NewKeywordCreateV2V2DataSuccessListBidTypeFromValue returns a pointer to a valid KeywordCreateV2V2DataSuccessListBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordCreateV2V2DataSuccessListBidTypeFromValue(v string) (*KeywordCreateV2V2DataSuccessListBidType, error) {
	ev := KeywordCreateV2V2DataSuccessListBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordCreateV2V2DataSuccessListBidType: valid values are %v", v, AllowedKeywordCreateV2V2DataSuccessListBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordCreateV2V2DataSuccessListBidType) IsValid() bool {
	for _, existing := range AllowedKeywordCreateV2V2DataSuccessListBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_create_v2_v2_data_success_list_bid_type value
func (v KeywordCreateV2V2DataSuccessListBidType) Ptr() *KeywordCreateV2V2DataSuccessListBidType {
	return &v
}
