/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueWechatInstanceUpdateV2ResponseDataFailWechatListInner struct for ClueWechatInstanceUpdateV2ResponseDataFailWechatListInner
type ClueWechatInstanceUpdateV2ResponseDataFailWechatListInner struct {
	// 失败原因
	FailReason *string `json:"fail_reason,omitempty"`
	// 微信号码
	Wechat *string `json:"wechat,omitempty"`
}
