/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupCreateV30RequestAdInfoTrackUrlSetting 监测链接设置
type AdlabGroupCreateV30RequestAdInfoTrackUrlSetting struct {
	// 点击（监测链接），只允许传入1个
	ActionTrackUrl []string `json:"action_track_url,omitempty"`
	// 展示（监测链接），只允许传入1个
	TrackUrl []string `json:"track_url,omitempty"`
	// 监测链接组id 当 track_url_type=EXISTED时必填
	TrackUrlGroupId  *int64                                                    `json:"track_url_group_id,omitempty"`
	TrackUrlSendType *AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlSendType `json:"track_url_send_type,omitempty"`
	TrackUrlType     *AdlabGroupCreateV30AdInfoTrackUrlSettingTrackUrlType     `json:"track_url_type,omitempty"`
	// 视频播完（监测链接），只允许传入1个 投放范围为穿山甲时暂不支持设置此链接
	VideoPlayDoneTrackUrl []string `json:"video_play_done_track_url,omitempty"`
	// 视频有效播放（监测链接），只允许传入1个 投放范围为穿山甲时暂不支持设置此链接
	VideoPlayEffectiveTrackUrl []string `json:"video_play_effective_track_url,omitempty"`
	// 视频开始播放（监测链接），只允许传入1个 投放范围为穿山甲时暂不支持设置此链接
	VideoPlayTrackUrl []string `json:"video_play_track_url,omitempty"`
}
