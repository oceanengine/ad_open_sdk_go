/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionAdDetailV10ResponseDataProgrammaticCreativeMediaListVideoMaterialInner struct for QianchuanUniPromotionAdDetailV10ResponseDataProgrammaticCreativeMediaListVideoMaterialInner
type QianchuanUniPromotionAdDetailV10ResponseDataProgrammaticCreativeMediaListVideoMaterialInner struct {
	//
	AwemeItemId *int64                                                                                     `json:"aweme_item_id,omitempty"`
	ImageMode   *QianchuanUniPromotionAdDetailV10DataProgrammaticCreativeMediaListVideoMaterialImageMode   `json:"image_mode,omitempty"`
	StarTraffic *QianchuanUniPromotionAdDetailV10DataProgrammaticCreativeMediaListVideoMaterialStarTraffic `json:"star_traffic,omitempty"`
	//
	VideoCoverId *string `json:"video_cover_id,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
}
