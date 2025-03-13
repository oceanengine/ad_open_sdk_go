/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

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

// Ptr returns reference to keyword_update_v2_v2_data_success_list_bid_type value
func (v KeywordUpdateV2V2DataSuccessListBidType) Ptr() *KeywordUpdateV2V2DataSuccessListBidType {
	return &v
}
