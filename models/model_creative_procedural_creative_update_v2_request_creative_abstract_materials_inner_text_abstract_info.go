/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeUpdateV2RequestCreativeAbstractMaterialsInnerTextAbstractInfo
type CreativeProceduralCreativeUpdateV2RequestCreativeAbstractMaterialsInnerTextAbstractInfo struct {
	//
	AbstractText string `json:"abstract_text"`
	//
	BidwordList []*CreativeProceduralCreativeUpdateV2RequestCreativeAbstractMaterialsInnerTextAbstractInfoBidwordListInner `json:"bidword_list,omitempty"`
	//
	WordList []*CreativeProceduralCreativeUpdateV2RequestCreativeAbstractMaterialsInnerTextAbstractInfoWordListInner `json:"word_list,omitempty"`
}
