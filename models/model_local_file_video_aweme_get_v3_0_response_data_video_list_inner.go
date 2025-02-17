/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalFileVideoAwemeGetV30ResponseDataVideoListInner struct for LocalFileVideoAwemeGetV30ResponseDataVideoListInner
type LocalFileVideoAwemeGetV30ResponseDataVideoListInner struct {
	// 抖音号id
	AwemeId *string `json:"aweme_id,omitempty"`
	// 视频播放地址
	AwemeVideoUrl *string `json:"aweme_video_url,omitempty"`
	// 是否可投放
	CanDelivery *bool `json:"can_delivery,omitempty"`
	// 视频封面图片地址
	CoverImageUrl *string `json:"cover_image_url,omitempty"`
	// 时长
	Duration  string                                          `json:"duration"`
	ImageMode LocalFileVideoAwemeGetV30DataVideoListImageMode `json:"image_mode"`
	// 抖音视频ID
	ItemId string `json:"item_id"`
	//
	LegoMaterialId *int64 `json:"lego_material_id,omitempty"`
	// 不可投放原因
	NotDeliveryReason []string `json:"not_delivery_reason,omitempty"`
	// 视频标题
	Title string `json:"title"`
	//
	VideoHeigh *int64 `json:"video_heigh,omitempty"`
	// 视频ID
	VideoId string `json:"video_id"`
	//
	VideoWidth *int64 `json:"video_width,omitempty"`
}
