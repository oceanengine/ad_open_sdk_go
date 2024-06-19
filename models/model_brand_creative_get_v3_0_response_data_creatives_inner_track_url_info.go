/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerTrackUrlInfo 监测链接
type BrandCreativeGetV30ResponseDataCreativesInnerTrackUrlInfo struct {
	// 点击监测链接
	ActionTrackUrl []string `json:"action_track_url,omitempty"`
	// 卡片展现第三方监测链接
	CardShowTrackUrl []string `json:"card_show_track_url,omitempty"`
	// 评论监测链接
	CommentTrackUrl []string `json:"comment_track_url,omitempty"`
	// 第三方上下文内容监测链接
	ContextTrackUrl []string `json:"context_track_url,omitempty"`
	// 视频有效播放监测链接
	EffectiveFrame []string `json:"effective_frame,omitempty"`
	// 视频开始播放监测链接
	FirstFrame []string `json:"first_frame,omitempty"`
	// 互动监控链接监测链接
	InteractTrackUrl []string `json:"interact_track_url,omitempty"`
	// 视频播放完成监测链接
	LastFrame []string `json:"last_frame,omitempty"`
	// 点赞取消监测链接
	LikeCancelTrackUrl []string `json:"like_cancel_track_url,omitempty"`
	// 点赞监测链接
	LikeTrackUrl []string `json:"like_track_url,omitempty"`
	// 视频主动播放监测链接
	ManualFrame []string `json:"manual_frame,omitempty"`
	// 分享监测链接
	ShareTrackUrl []string `json:"share_track_url,omitempty"`
	// 展示监测链接
	TrackUrl []string `json:"track_url,omitempty"`
}
