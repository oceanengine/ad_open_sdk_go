/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerPromotionMaterialsTitleMaterialListInner struct for PromotionListV30ResponseDataListInnerPromotionMaterialsTitleMaterialListInner
type PromotionListV30ResponseDataListInnerPromotionMaterialsTitleMaterialListInner struct {
	//
	BidwordList []*PromotionListV30ResponseDataListInnerPromotionMaterialsTextAbstractListInnerBidwordListInner `json:"bidword_list,omitempty"`
	//
	DpaWordList []int64 `json:"dpa_word_list,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	WordList []int64 `json:"word_list,omitempty"`
}
