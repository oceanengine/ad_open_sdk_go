/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdKeywordsGetV10ResponseDataListInner struct for QianchuanAdKeywordsGetV10ResponseDataListInner
type QianchuanAdKeywordsGetV10ResponseDataListInner struct {
	// 计划ID
	AdId *int64 `json:"ad_id,omitempty"`
	// 关键词id
	Id        *int64                                      `json:"id,omitempty"`
	MatchType *QianchuanAdKeywordsGetV10DataListMatchType `json:"match_type,omitempty"`
	Status    *QianchuanAdKeywordsGetV10DataListStatus    `json:"status,omitempty"`
	// 关键词字面
	Word *string `json:"word,omitempty"`
	// 词ID
	WordId *int64 `json:"word_id,omitempty"`
}
