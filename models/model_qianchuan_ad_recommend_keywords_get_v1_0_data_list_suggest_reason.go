/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdRecommendKeywordsGetV10DataListSuggestReason
type QianchuanAdRecommendKeywordsGetV10DataListSuggestReason string

// List of qianchuan_ad_recommend_keywords_get_v1.0_data_list_suggest_reason
const (
	CLICK_QianchuanAdRecommendKeywordsGetV10DataListSuggestReason      QianchuanAdRecommendKeywordsGetV10DataListSuggestReason = "CLICK"
	DARKHORSE_QianchuanAdRecommendKeywordsGetV10DataListSuggestReason  QianchuanAdRecommendKeywordsGetV10DataListSuggestReason = "DARKHORSE"
	LOW_COST_QianchuanAdRecommendKeywordsGetV10DataListSuggestReason   QianchuanAdRecommendKeywordsGetV10DataListSuggestReason = "LOW_COST"
	POTENTIAL_QianchuanAdRecommendKeywordsGetV10DataListSuggestReason  QianchuanAdRecommendKeywordsGetV10DataListSuggestReason = "POTENTIAL"
	PREFERENCE_QianchuanAdRecommendKeywordsGetV10DataListSuggestReason QianchuanAdRecommendKeywordsGetV10DataListSuggestReason = "PREFERENCE"
)

// Ptr returns reference to qianchuan_ad_recommend_keywords_get_v1.0_data_list_suggest_reason value
func (v QianchuanAdRecommendKeywordsGetV10DataListSuggestReason) Ptr() *QianchuanAdRecommendKeywordsGetV10DataListSuggestReason {
	return &v
}
