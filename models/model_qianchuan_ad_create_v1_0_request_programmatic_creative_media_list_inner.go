/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10RequestProgrammaticCreativeMediaListInner struct for QianchuanAdCreateV10RequestProgrammaticCreativeMediaListInner
type QianchuanAdCreateV10RequestProgrammaticCreativeMediaListInner struct {
	//
	AwemeCarouselId *int64 `json:"aweme_carousel_id,omitempty"`
	//
	AwemeItemId *int64 `json:"aweme_item_id,omitempty"`
	//
	CarouselId *int64 `json:"carousel_id,omitempty"`
	//
	ImageIds  []string                                                    `json:"image_ids,omitempty"`
	ImageMode *QianchuanAdCreateV10ProgrammaticCreativeMediaListImageMode `json:"image_mode,omitempty"`
	//
	VideoCoverId *string `json:"video_cover_id,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
}
