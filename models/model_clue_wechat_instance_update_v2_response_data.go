/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueWechatInstanceUpdateV2ResponseData
type ClueWechatInstanceUpdateV2ResponseData struct {
	// 修改失败的微信列表
	FailWechatList []*ClueWechatInstanceUpdateV2ResponseDataFailWechatListInner `json:"fail_wechat_list,omitempty"`
	// 修改成功的微信列表
	SuccessWechatList []string `json:"success_wechat_list,omitempty"`
}
