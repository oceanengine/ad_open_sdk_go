/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestPromotionMaterialsTextAbstractListInner struct for PromotionUpdateV30RequestPromotionMaterialsTextAbstractListInner
type PromotionUpdateV30RequestPromotionMaterialsTextAbstractListInner struct {
	//
	AbstractText string `json:"abstract_text"`
	//
	BidwordList []*PromotionUpdateV30RequestPromotionMaterialsTextAbstractListInnerBidwordListInner `json:"bidword_list,omitempty"`
	//
	WordList []int64 `json:"word_list,omitempty"`
}
