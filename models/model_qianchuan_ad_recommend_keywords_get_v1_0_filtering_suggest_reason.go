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

// All allowed values of QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason enum
var AllowedQianchuanAdRecommendKeywordsGetV10FilteringSuggestReasonEnumValues = []QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason{
	"CLICK",
	"DARKHORSE",
	"LOW_COST",
	"POTENTIAL",
	"PREFERENCE",
}

// NewQianchuanAdRecommendKeywordsGetV10FilteringSuggestReasonFromValue returns a pointer to a valid QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdRecommendKeywordsGetV10FilteringSuggestReasonFromValue(v string) (*QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason, error) {
	ev := QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason: valid values are %v", v, AllowedQianchuanAdRecommendKeywordsGetV10FilteringSuggestReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason) IsValid() bool {
	for _, existing := range AllowedQianchuanAdRecommendKeywordsGetV10FilteringSuggestReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_recommend_keywords_get_v1.0_filtering_suggest_reason value
func (v QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason) Ptr() *QianchuanAdRecommendKeywordsGetV10FilteringSuggestReason {
	return &v
}
