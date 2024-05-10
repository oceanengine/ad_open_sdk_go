/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerTrackUrlUpdateV2RequestTrackUrlGroup 监测链接组信息
type EventManagerTrackUrlUpdateV2RequestTrackUrlGroup struct {
	// 点击（监测链接）只允许传入1个
	ActionTrackUrl string `json:"action_track_url"`
	// 激活监测
	ActiveTrackUrl *string `json:"active_track_url,omitempty"`
	// 展示（监测链接）
	TrackUrl *string `json:"track_url,omitempty"`
	// 监测链接组ID
	TrackUrlGroupId int64 `json:"track_url_group_id"`
	// 监测链接组名称 应用资产必填
	TrackUrlGroupName *string `json:"track_url_group_name,omitempty"`
	// 视频播完（监测链接），只允许传入1个
	VideoPlayDoneTrackUrl *string `json:"video_play_done_track_url,omitempty"`
	// 视频有效播放（监测链接），只允许传入1个
	VideoPlayEffectiveTrackUrl *string `json:"video_play_effective_track_url,omitempty"`
	// 视频播放（监测链接），只允许传入1个
	VideoPlayTrackUrl *string `json:"video_play_track_url,omitempty"`
}
