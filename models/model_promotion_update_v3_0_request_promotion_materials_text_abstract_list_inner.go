/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestPromotionMaterialsTextAbstractListInner struct for PromotionUpdateV30RequestPromotionMaterialsTextAbstractListInner
type PromotionUpdateV30RequestPromotionMaterialsTextAbstractListInner struct {
	//
	AbstractText *string `json:"abstract_text,omitempty"`
	//
	BidwordList []*PromotionCreateV30RequestPromotionMaterialsTextAbstractListInnerBidwordListInner `json:"bidword_list,omitempty"`
	//
	WordList []int64 `json:"word_list,omitempty"`
}
