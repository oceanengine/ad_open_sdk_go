/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataCreativeListInnerCarouselMaterial
type QianchuanAdDetailGetV10ResponseDataCreativeListInnerCarouselMaterial struct {
	Audio *QianchuanAdDetailGetV10ResponseDataCreativeListInnerCarouselMaterialAudio `json:"audio,omitempty"`
	//
	AwemeCarouselId *int64 `json:"aweme_carousel_id,omitempty"`
	//
	CarouselId *int64 `json:"carousel_id,omitempty"`
	//
	Description *string `json:"description,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	Images []*QianchuanAdDetailGetV10ResponseDataCreativeListInnerCarouselMaterialImagesInner `json:"images,omitempty"`
	//
	IsAutoGenerate *int64 `json:"is_auto_generate,omitempty"`
}
