/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeCreateV2RequestCreativeListInnerTitleMaterial
type CreativeCustomCreativeCreateV2RequestCreativeListInnerTitleMaterial struct {
	//
	BidwordList []*CreativeCustomCreativeCreateV2RequestCreativeListInnerTitleMaterialBidwordListInner `json:"bidword_list,omitempty"`
	//
	DpaWordList []*CreativeCustomCreativeCreateV2RequestCreativeListInnerTitleMaterialDpaWordListInner `json:"dpa_word_list,omitempty"`
	//
	Title string `json:"title"`
	//
	WordList []*CreativeCustomCreativeCreateV2RequestCreativeListInnerTitleMaterialWordListInner `json:"word_list,omitempty"`
}
