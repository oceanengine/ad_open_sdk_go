/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileMaterialBindV2Request struct for FileMaterialBindV2Request
type FileMaterialBindV2Request struct {
	// 素材归属广告主
	AdvertiserId int64 `json:"advertiser_id"`
	// 图集ID，数量限制：<=50 注意：跟video_ids必须三选一
	CarouselIds []int64 `json:"carousel_ids,omitempty"`
	// 图片ID，数量限制：<=50 注意：跟video_ids必须三选一
	ImageIds []string `json:"image_ids,omitempty"`
	// 待推送的广告主，数量限制：<=50
	TargetAdvertiserIds []int64 `json:"target_advertiser_ids"`
	// 视频ID，数量限制：<=50 注意：跟image_ids必须三选一、组织共享视频不可推送
	VideoIds []string `json:"video_ids,omitempty"`
}
