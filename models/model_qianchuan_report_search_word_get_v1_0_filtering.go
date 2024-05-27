/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportSearchWordGetV10Filtering
type QianchuanReportSearchWordGetV10Filtering struct {
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	CampaignIds []int64 `json:"campaign_ids,omitempty"`
	//
	KeywordMerge  *bool                                                 `json:"keyword_merge,omitempty"`
	MarketingGoal QianchuanReportSearchWordGetV10FilteringMarketingGoal `json:"marketing_goal"`
	Range         *QianchuanReportSearchWordGetV10FilteringRange        `json:"range,omitempty"`
	//
	Word     *string                                           `json:"word,omitempty"`
	WordType *QianchuanReportSearchWordGetV10FilteringWordType `json:"word_type,omitempty"`
}
