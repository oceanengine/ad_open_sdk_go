/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30DataListKeywordsMatchType
type PromotionListV30DataListKeywordsMatchType string

// List of promotion_list_v3.0_data_list_keywords_match_type
const (
	EXTENSIVE_PromotionListV30DataListKeywordsMatchType PromotionListV30DataListKeywordsMatchType = "EXTENSIVE"
	PHRASE_PromotionListV30DataListKeywordsMatchType    PromotionListV30DataListKeywordsMatchType = "PHRASE"
	PRECISION_PromotionListV30DataListKeywordsMatchType PromotionListV30DataListKeywordsMatchType = "PRECISION"
)

// Ptr returns reference to promotion_list_v3.0_data_list_keywords_match_type value
func (v PromotionListV30DataListKeywordsMatchType) Ptr() *PromotionListV30DataListKeywordsMatchType {
	return &v
}
