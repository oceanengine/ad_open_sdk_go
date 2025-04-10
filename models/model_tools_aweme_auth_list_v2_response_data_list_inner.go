/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthListV2ResponseDataListInner struct for ToolsAwemeAuthListV2ResponseDataListInner
type ToolsAwemeAuthListV2ResponseDataListInner struct {
	//
	AuthAutoExpireDate *string                                                     `json:"auth_auto_expire_date,omitempty"`
	AuthStatus         *ToolsAwemeAuthListV2DataListAuthStatus                     `json:"auth_status,omitempty"`
	AuthThresholdInfo  *ToolsAwemeAuthListV2ResponseDataListInnerAuthThresholdInfo `json:"auth_threshold_info,omitempty"`
	AuthType           *ToolsAwemeAuthListV2DataListAuthType                       `json:"auth_type,omitempty"`
	// 抖音号作者发起解除授权时上传的凭证信息（选填项，抖音号作者可能不填，此时该参数返回为null）
	AwemeCancelImageList []string `json:"aweme_cancel_image_list,omitempty"`
	// 抖音号作者发起解除授权时填写的联系方式（选填项，抖音号作者可能不填，此时该参数返回为null）
	AwemeCancelNote *string `json:"aweme_cancel_note,omitempty"`
	// 抖音号作者发起解除授权的原因，仅当抖音号作者发起解除授权时有值，100字以内，可能包括： - 不知道该授权是怎么建立的，申请解除授权 - 联系不到对方，无法进行合作沟通，申请解除授权 - 与对方合作到期或者有纠纷，申请解除授权 - 其他情况（作者会填写其他文案给到）
	AwemeCancelReason *string `json:"aweme_cancel_reason,omitempty"`
	// 抖音号
	AwemeId *string `json:"aweme_id,omitempty"`
	// 抖音账号名称
	AwemeName     *string                                    `json:"aweme_name,omitempty"`
	AwemeUserType *ToolsAwemeAuthListV2DataListAwemeUserType `json:"aweme_user_type,omitempty"`
	// 授权结束时间
	EndTime *string `json:"end_time,omitempty"`
	// 发布新视频素材到该抖音号下时，视频主页可见性只能设置「仅单次展示可见」 枚举值： true：是 false：否，表示无此限制  当值返回true时，代表在创建广告时添加新视频素材到该抖音号下推广，视频的主页可见性设置只允许HIDE_VIDEO_ON_HP「仅单次展示可见」
	HasVideoHpVisibilityLimit *bool `json:"has_video_hp_visibility_limit,omitempty"`
	// 备注
	Note      *string                                `json:"note,omitempty"`
	ShareType *ToolsAwemeAuthListV2DataListShareType `json:"share_type,omitempty"`
	// 授权开始时间
	StartTime *string                                             `json:"start_time,omitempty"`
	SubStatus *ToolsAwemeAuthListV2DataListSubStatus              `json:"sub_status,omitempty"`
	VideoInfo *ToolsAwemeAuthListV2ResponseDataListInnerVideoInfo `json:"video_info,omitempty"`
	// 抖音授权关系警告信息，您可根据该信息及时处理，可能返回 - 不达门槛：表示发起授权的抖音号未达到要求，详细未达门槛信息可通过auth_threshold_info获取 - 即将解除：表示抖音号作者已发起解除授权申请，您需要及时联系作者或同意解除
	WarningContent []string `json:"warning_content,omitempty"`
	// 抖音授权关系警告类型
	WarningTypes []*ToolsAwemeAuthListV2DataListWarningTypes `json:"warning_types,omitempty"`
}
