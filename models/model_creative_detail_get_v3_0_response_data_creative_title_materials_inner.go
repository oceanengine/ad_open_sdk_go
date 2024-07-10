/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeDetailGetV30ResponseDataCreativeTitleMaterialsInner struct for CreativeDetailGetV30ResponseDataCreativeTitleMaterialsInner
type CreativeDetailGetV30ResponseDataCreativeTitleMaterialsInner struct {
	// 搜索关键词列表
	BidwordList []*CreativeDetailGetV30ResponseDataCreativeTitleMaterialsInnerBidwordListInner `json:"bidword_list,omitempty"`
	// DPA词包ID列表
	DpaWordList []*CreativeDetailGetV30ResponseDataCreativeTitleMaterialsInnerDpaWordListInner `json:"dpa_word_list,omitempty"`
	// 创意标题
	Title *string `json:"title,omitempty"`
	// 动态词包ID列表
	WordList []*CreativeDetailGetV30ResponseDataCreativeTitleMaterialsInnerWordListInner `json:"word_list,omitempty"`
}
