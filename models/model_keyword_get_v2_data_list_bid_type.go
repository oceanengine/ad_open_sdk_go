/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordGetV2DataListBidType
type KeywordGetV2DataListBidType string

// List of keyword_get_v2_data_list_bid_type
const (
	SUGGEST_KeywordGetV2DataListBidType        KeywordGetV2DataListBidType = "SUGGEST"
	FEED_TO_SEARCH_KeywordGetV2DataListBidType KeywordGetV2DataListBidType = "FEED_TO_SEARCH"
	CUSTOM_KeywordGetV2DataListBidType         KeywordGetV2DataListBidType = "CUSTOM"
	WITH_AD_KeywordGetV2DataListBidType        KeywordGetV2DataListBidType = "WITH_AD"
	BRAND_AD_KeywordGetV2DataListBidType       KeywordGetV2DataListBidType = "BRAND_AD"
)

// Ptr returns reference to keyword_get_v2_data_list_bid_type value
func (v KeywordGetV2DataListBidType) Ptr() *KeywordGetV2DataListBidType {
	return &v
}
