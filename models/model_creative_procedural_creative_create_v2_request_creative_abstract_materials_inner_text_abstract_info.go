/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeCreateV2RequestCreativeAbstractMaterialsInnerTextAbstractInfo
type CreativeProceduralCreativeCreateV2RequestCreativeAbstractMaterialsInnerTextAbstractInfo struct {
	//
	AbstractText string `json:"abstract_text"`
	//
	BidwordList []*CreativeProceduralCreativeCreateV2RequestCreativeAbstractMaterialsInnerTextAbstractInfoBidwordListInner `json:"bidword_list,omitempty"`
	//
	WordList []*CreativeProceduralCreativeCreateV2RequestCreativeAbstractMaterialsInnerTextAbstractInfoWordListInner `json:"word_list,omitempty"`
}
