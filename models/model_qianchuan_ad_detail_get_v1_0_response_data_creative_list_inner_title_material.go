/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataCreativeListInnerTitleMaterial
type QianchuanAdDetailGetV10ResponseDataCreativeListInnerTitleMaterial struct {
	//
	DynamicWords []*QianchuanAdDetailGetV10ResponseDataCreativeListInnerTitleMaterialDynamicWordsInner `json:"dynamic_words,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	Title     *string                                                        `json:"title,omitempty"`
	TitleType *QianchuanAdDetailGetV10DataCreativeListTitleMaterialTitleType `json:"title_type,omitempty"`
}
