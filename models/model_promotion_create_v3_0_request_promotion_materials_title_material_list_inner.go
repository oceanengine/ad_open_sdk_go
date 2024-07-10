/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30RequestPromotionMaterialsTitleMaterialListInner struct for PromotionCreateV30RequestPromotionMaterialsTitleMaterialListInner
type PromotionCreateV30RequestPromotionMaterialsTitleMaterialListInner struct {
	//
	BidwordList []*PromotionCreateV30RequestPromotionMaterialsTitleMaterialListInnerBidwordListInner `json:"bidword_list,omitempty"`
	// 创建DPA创意时可以选择传入DPA词包ID列表，动态词包与DPA词包总数最多为2。可通过【获取DPA词包】接口获取，结合标题中的词包格式您需要填写相同个数与顺序的词包ID，如果实际ID顺序与标题中词包名顺序不一致我们将以词包ID顺序为准
	DpaWordList []int64 `json:"dpa_word_list,omitempty"`
	//
	Title string `json:"title"`
	//
	WordList []int64 `json:"word_list,omitempty"`
}
