/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAnalyseCompareCreativeV10ResponseDataOwnProductCreativeInnerVideoMaterialInner struct for QianchuanProductAnalyseCompareCreativeV10ResponseDataOwnProductCreativeInnerVideoMaterialInner
type QianchuanProductAnalyseCompareCreativeV10ResponseDataOwnProductCreativeInnerVideoMaterialInner struct {
	//
	AwemeItemId *int64 `json:"aweme_item_id,omitempty"`
	//
	CoverImageHeight *int64 `json:"cover_image_height,omitempty"`
	//
	CoverImageWebUrl *string `json:"cover_image_web_url,omitempty"`
	//
	CoverImageWidth *int64                                                                              `json:"cover_image_width,omitempty"`
	Source          *QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource `json:"source,omitempty"`
	//
	VideoDuration *int64 `json:"video_duration,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
	//
	VideoUrl []string `json:"video_url,omitempty"`
}
