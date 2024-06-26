/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselCreateV2Request struct for CarouselCreateV2Request
type CarouselCreateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	// 音频 id，可以同过音频上传接口获取（优先使用 audio_id）
	AudioId      *string                      `json:"audio_id,omitempty"`
	CarouselType CarouselCreateV2CarouselType `json:"carousel_type"`
	// 千川图文素材描述
	Description *string `json:"description,omitempty"`
	//
	FileName *string `json:"file_name,omitempty"`
	//
	Images []*CarouselCreateV2RequestImagesInner `json:"images"`
	// 音频vid，可以同过视频上传接口获取（优先使用 audio_id）
	VideoId *string `json:"video_id,omitempty"`
}
