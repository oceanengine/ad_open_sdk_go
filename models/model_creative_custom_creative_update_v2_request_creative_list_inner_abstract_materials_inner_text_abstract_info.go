/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeUpdateV2RequestCreativeListInnerAbstractMaterialsInnerTextAbstractInfo
type CreativeCustomCreativeUpdateV2RequestCreativeListInnerAbstractMaterialsInnerTextAbstractInfo struct {
	//
	AbstractText string `json:"abstract_text"`
	//
	BidwordList []*CreativeCustomCreativeUpdateV2RequestCreativeListInnerAbstractMaterialsInnerTextAbstractInfoBidwordListInner `json:"bidword_list,omitempty"`
	//
	WordList []*CreativeCustomCreativeUpdateV2RequestCreativeListInnerAbstractMaterialsInnerTextAbstractInfoWordListInner `json:"word_list,omitempty"`
}