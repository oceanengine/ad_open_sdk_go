/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCreativeGetV10ResponseDataListInnerTitleMaterialListInner struct for QianchuanCreativeGetV10ResponseDataListInnerTitleMaterialListInner
type QianchuanCreativeGetV10ResponseDataListInnerTitleMaterialListInner struct {
	//
	AwemeCarouselId *int64 `json:"aweme_carousel_id,omitempty"`
	//
	DynamicWords []*QianchuanCreativeGetV10ResponseDataListInnerTitleMaterialListInnerDynamicWordsInner `json:"dynamic_words,omitempty"`
	//
	Title     *string                                                    `json:"title,omitempty"`
	TitleType *QianchuanCreativeGetV10DataListTitleMaterialListTitleType `json:"title_type,omitempty"`
}
