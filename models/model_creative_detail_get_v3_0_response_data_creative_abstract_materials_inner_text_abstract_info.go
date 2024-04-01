/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeDetailGetV30ResponseDataCreativeAbstractMaterialsInnerTextAbstractInfo 文本摘要信息
type CreativeDetailGetV30ResponseDataCreativeAbstractMaterialsInnerTextAbstractInfo struct {
	// 文本摘要内容
	AbstractText *string `json:"abstract_text,omitempty"`
	// 搜索关键词列表
	BidwordList []*CreativeDetailGetV30ResponseDataCreativeAbstractMaterialsInnerTextAbstractInfoBidwordListInner `json:"bidword_list,omitempty"`
	// 动态词包ID列表
	WordList []*CreativeDetailGetV30ResponseDataCreativeAbstractMaterialsInnerTextAbstractInfoWordListInner `json:"word_list,omitempty"`
}
