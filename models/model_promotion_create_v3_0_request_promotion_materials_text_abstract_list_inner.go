/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30RequestPromotionMaterialsTextAbstractListInner struct for PromotionCreateV30RequestPromotionMaterialsTextAbstractListInner
type PromotionCreateV30RequestPromotionMaterialsTextAbstractListInner struct {
	//
	AbstractText string `json:"abstract_text"`
	//
	BidwordList []*PromotionCreateV30RequestPromotionMaterialsTextAbstractListInnerBidwordListInner `json:"bidword_list,omitempty"`
	//
	WordList []int64 `json:"word_list,omitempty"`
}
