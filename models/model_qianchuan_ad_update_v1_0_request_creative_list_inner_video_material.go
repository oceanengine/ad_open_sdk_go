/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10RequestCreativeListInnerVideoMaterial
type QianchuanAdUpdateV10RequestCreativeListInnerVideoMaterial struct {
	//
	AwemeItemId *int64 `json:"aweme_item_id,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	// 是否为派生创意标识，1：是，0：不是
	IsAutoGenerate *int64 `json:"is_auto_generate,omitempty"`
	//
	VideoCoverId *string `json:"video_cover_id,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
}
