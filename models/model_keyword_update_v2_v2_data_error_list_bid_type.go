/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV2V2DataErrorListBidType
type KeywordUpdateV2V2DataErrorListBidType string

// List of keyword_update_v2_v2_data_error_list_bid_type
const (
	SUGGEST_KeywordUpdateV2V2DataErrorListBidType        KeywordUpdateV2V2DataErrorListBidType = "SUGGEST"
	CUSTOM_KeywordUpdateV2V2DataErrorListBidType         KeywordUpdateV2V2DataErrorListBidType = "CUSTOM"
	FEED_TO_SEARCH_KeywordUpdateV2V2DataErrorListBidType KeywordUpdateV2V2DataErrorListBidType = "FEED_TO_SEARCH"
	BRAND_AD_KeywordUpdateV2V2DataErrorListBidType       KeywordUpdateV2V2DataErrorListBidType = "BRAND_AD"
	WITH_AD_KeywordUpdateV2V2DataErrorListBidType        KeywordUpdateV2V2DataErrorListBidType = "WITH_AD"
)

// Ptr returns reference to keyword_update_v2_v2_data_error_list_bid_type value
func (v KeywordUpdateV2V2DataErrorListBidType) Ptr() *KeywordUpdateV2V2DataErrorListBidType {
	return &v
}
