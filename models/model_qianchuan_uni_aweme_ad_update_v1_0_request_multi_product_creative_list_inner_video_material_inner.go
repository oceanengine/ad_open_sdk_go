/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniAwemeAdUpdateV10RequestMultiProductCreativeListInnerVideoMaterialInner struct for QianchuanUniAwemeAdUpdateV10RequestMultiProductCreativeListInnerVideoMaterialInner
type QianchuanUniAwemeAdUpdateV10RequestMultiProductCreativeListInnerVideoMaterialInner struct {
	//
	AwemeItemId *int64                                                                      `json:"aweme_item_id,omitempty"`
	ImageMode   *QianchuanUniAwemeAdUpdateV10MultiProductCreativeListVideoMaterialImageMode `json:"image_mode,omitempty"`
	//
	VideoCoverId *string `json:"video_cover_id,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
}
