/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordCreateV2V2KeywordsBidType
type KeywordCreateV2V2KeywordsBidType string

// List of keyword_create_v2_v2_keywords_bid_type
const (
	WITH_AD_KeywordCreateV2V2KeywordsBidType        KeywordCreateV2V2KeywordsBidType = "WITH_AD"
	CUSTOM_KeywordCreateV2V2KeywordsBidType         KeywordCreateV2V2KeywordsBidType = "CUSTOM"
	BRAND_AD_KeywordCreateV2V2KeywordsBidType       KeywordCreateV2V2KeywordsBidType = "BRAND_AD"
	SUGGEST_KeywordCreateV2V2KeywordsBidType        KeywordCreateV2V2KeywordsBidType = "SUGGEST"
	FEED_TO_SEARCH_KeywordCreateV2V2KeywordsBidType KeywordCreateV2V2KeywordsBidType = "FEED_TO_SEARCH"
)

// Ptr returns reference to keyword_create_v2_v2_keywords_bid_type value
func (v KeywordCreateV2V2KeywordsBidType) Ptr() *KeywordCreateV2V2KeywordsBidType {
	return &v
}
