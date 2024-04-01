/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30RequestTrackUrlSetting
type ProjectCreateV30RequestTrackUrlSetting struct {
	//
	ActionTrackUrl []string `json:"action_track_url,omitempty"`
	//
	ActiveTrackUrl []string                                 `json:"active_track_url,omitempty"`
	SendType       *ProjectCreateV30TrackUrlSettingSendType `json:"send_type,omitempty"`
	//
	TrackUrl []string `json:"track_url,omitempty"`
	//
	TrackUrlGroupId *int64                                       `json:"track_url_group_id,omitempty"`
	TrackUrlType    *ProjectCreateV30TrackUrlSettingTrackUrlType `json:"track_url_type,omitempty"`
	//
	VideoPlayDoneTrackUrl []string `json:"video_play_done_track_url,omitempty"`
	//
	VideoPlayEffectiveTrackUrl []string `json:"video_play_effective_track_url,omitempty"`
	//
	VideoPlayFirstTrackUrl []string `json:"video_play_first_track_url,omitempty"`
}
