/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeAuthListGetV10ResponseDataAuthorizationInfosInner struct for QianchuanAwemeAuthListGetV10ResponseDataAuthorizationInfosInner
type QianchuanAwemeAuthListGetV10ResponseDataAuthorizationInfosInner struct {
	AuthRange  *QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthRange  `json:"auth_range,omitempty"`
	AuthSource *QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthSource `json:"auth_source,omitempty"`
	AuthStatus *QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthStatus `json:"auth_status,omitempty"`
	//
	AuthType      []*QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType                 `json:"auth_type,omitempty"`
	AuthVideoInfo *QianchuanAwemeAuthListGetV10ResponseDataAuthorizationInfosInnerAuthVideoInfo `json:"auth_video_info,omitempty"`
	//
	AwemeAvatar *string `json:"aweme_avatar,omitempty"`
	// 抖音号uid
	AwemeId *int64 `json:"aweme_id,omitempty"`
	//
	AwemeName *string `json:"aweme_name,omitempty"`
	//
	AwemeShowId *string `json:"aweme_show_id,omitempty"`
	// 申请创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// 结束时间
	EndTime *string `json:"end_time,omitempty"`
	// 是待确认解除授权
	IsCancellationProgress *bool `json:"is_cancellation_progress,omitempty"`
	// 开始时间
	StartTime *string `json:"start_time,omitempty"`
}
