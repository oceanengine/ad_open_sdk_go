/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCreativeGetV10ResponseDataListInnerVideoMaterialListInner struct for QianchuanCreativeGetV10ResponseDataListInnerVideoMaterialListInner
type QianchuanCreativeGetV10ResponseDataListInnerVideoMaterialListInner struct {
	//
	AwemeItemId *int64                                                     `json:"aweme_item_id,omitempty"`
	ImageMode   *QianchuanCreativeGetV10DataListVideoMaterialListImageMode `json:"image_mode,omitempty"`
	//
	IsAutoGenerate *int64 `json:"is_auto_generate,omitempty"`
	//
	VideoCoverId *string `json:"video_cover_id,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
}
