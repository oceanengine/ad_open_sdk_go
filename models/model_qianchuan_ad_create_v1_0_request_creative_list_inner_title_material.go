/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10RequestCreativeListInnerTitleMaterial
type QianchuanAdCreateV10RequestCreativeListInnerTitleMaterial struct {
	//
	DynamicWords []*QianchuanAdCreateV10RequestCreativeListInnerTitleMaterialDynamicWordsInner `json:"dynamic_words,omitempty"`
	//
	Title     string                                                  `json:"title"`
	TitleType *QianchuanAdCreateV10CreativeListTitleMaterialTitleType `json:"title_type,omitempty"`
}
