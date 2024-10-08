/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdRecommendKeywordsGetV10ResponseData
type QianchuanAdRecommendKeywordsGetV10ResponseData struct {
	// 本次查询的唯一标识ID，有效期12h。 可用于后续请求，用以保证查询到数据的完整性。
	CacheId *string `json:"cache_id,omitempty"`
	// 关键词列表
	List     []*QianchuanAdRecommendKeywordsGetV10ResponseDataListInner `json:"list,omitempty"`
	PageInfo *QianchuanAdRecommendKeywordsGetV10ResponseDataPageInfo    `json:"page_info,omitempty"`
}
