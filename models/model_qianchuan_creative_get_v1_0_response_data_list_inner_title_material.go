/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCreativeGetV10ResponseDataListInnerTitleMaterial
type QianchuanCreativeGetV10ResponseDataListInnerTitleMaterial struct {
	//
	AwemeCarouselId *int64 `json:"aweme_carousel_id,omitempty"`
	//
	DynamicWords []*QianchuanCreativeGetV10ResponseDataListInnerTitleMaterialDynamicWordsInner `json:"dynamic_words,omitempty"`
	//
	Title     *string                                                `json:"title,omitempty"`
	TitleType *QianchuanCreativeGetV10DataListTitleMaterialTitleType `json:"title_type,omitempty"`
}
