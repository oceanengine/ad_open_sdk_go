/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueWechatInstanceUpdateV2Request struct for ClueWechatInstanceUpdateV2Request
type ClueWechatInstanceUpdateV2Request struct {
	Action ClueWechatInstanceUpdateV2Action `json:"action"`
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 微信号码包ID
	InstanceId int64 `json:"instance_id"`
	// 要修改的微信列表，最多不超过100个
	WechatList []string `json:"wechat_list"`
}
