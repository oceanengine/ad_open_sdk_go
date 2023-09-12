/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueWechatInstanceDetailV2ResponseDataWechatListInner struct for ClueWechatInstanceDetailV2ResponseDataWechatListInner
type ClueWechatInstanceDetailV2ResponseDataWechatListInner struct {
	// 公众号的开发者ID
	AppId *string `json:"app_id,omitempty"`
	// 总复制次数
	ClueCount *int64 `json:"clue_count,omitempty"`
	// 创建时间，格式：2022-07-21 15:36:09
	CreateTime *string                                            `json:"create_time,omitempty"`
	HasQrCode  *ClueWechatInstanceDetailV2DataWechatListHasQrCode `json:"has_qr_code,omitempty"`
	// 微信号在当前号码包复制次数
	InstanceClueCount *int64 `json:"instance_clue_count,omitempty"`
	// 修改时间，格式：2022-07-21 15:36:09
	ModTime *string `json:"mod_time,omitempty"`
	// 微信昵称
	NickName *string `json:"nick_name,omitempty"`
	// 启用状态，0使用中、1暂停使用
	Status *int64 `json:"status,omitempty"`
	// 微信号尾缀
	Suffix     *string                                             `json:"suffix,omitempty"`
	WechatType *ClueWechatInstanceDetailV2DataWechatListWechatType `json:"wechat_type,omitempty"`
}
