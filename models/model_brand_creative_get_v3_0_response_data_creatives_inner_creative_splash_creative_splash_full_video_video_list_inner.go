/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashFullVideoVideoListInner struct for BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashFullVideoVideoListInner
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeSplashCreativeSplashFullVideoVideoListInner struct {
	// 视频时长，单位秒
	Duration *float64 `json:"duration,omitempty"`
	// 文件大小，单位字节
	FileSize *string `json:"file_size,omitempty"`
	// 视频高,单位px
	Height *int64 `json:"height,omitempty"`
	// 播放地址
	OriginPlayUrl *string `json:"origin_play_url,omitempty"`
	// 封面图高度
	ThumbHeight *int64 `json:"thumb_height,omitempty"`
	// 封面图uri
	ThumbUri *string `json:"thumb_uri,omitempty"`
	// 封面图宽度
	ThumbWidth *int64 `json:"thumb_width,omitempty"`
	// 视频ID
	VideoId *string `json:"video_id,omitempty"`
	// 视频url
	VideoUrl *string `json:"video_url,omitempty"`
	// 视频宽,单位px
	Width *int64 `json:"width,omitempty"`
}
