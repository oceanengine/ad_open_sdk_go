/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason
type QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason string

// List of qianchuan_ad_recommend_keywords_get_v1.0_filtering_suggest_reason
const (
	CLICK_QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason      QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason = "CLICK"
	DARKHORSE_QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason  QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason = "DARKHORSE"
	LOW_COST_QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason   QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason = "LOW_COST"
	POTENTIAL_QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason  QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason = "POTENTIAL"
	PREFERENCE_QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason = "PREFERENCE"
)

// Ptr returns reference to qianchuan_ad_recommend_keywords_get_v1.0_filtering_suggest_reason value
func (v QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason) Ptr() *QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason {
	return &v
}
