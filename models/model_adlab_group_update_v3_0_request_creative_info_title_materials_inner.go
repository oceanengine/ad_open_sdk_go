/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupUpdateV30RequestCreativeInfoTitleMaterialsInner struct for AdlabGroupUpdateV30RequestCreativeInfoTitleMaterialsInner
type AdlabGroupUpdateV30RequestCreativeInfoTitleMaterialsInner struct {
	// 标题信息
	Title string `json:"title"`
	// 词包信息
	WordList []*AdlabGroupUpdateV30RequestCreativeInfoTitleMaterialsInnerWordListInner `json:"word_list,omitempty"`
}
