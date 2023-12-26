/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueWechatInstanceDetailV2ResponseData
type ClueWechatInstanceDetailV2ResponseData struct {
	// 创建时间，格式：2022-07-21 15:36:09
	CreateTime *string `json:"create_time,omitempty"`
	// 微信号码包ID
	InstanceId *int64 `json:"instance_id,omitempty"`
	// 修改时间，格式：2022-07-21 15:36:09
	ModTime *string `json:"mod_time,omitempty"`
	// 微信号码包名称
	Name    *string                                `json:"name,omitempty"`
	SubType *ClueWechatInstanceDetailV2DataSubType `json:"sub_type,omitempty"`
	Suffix  *ClueWechatInstanceDetailV2DataSuffix  `json:"suffix,omitempty"`
	//
	WechatList []*ClueWechatInstanceDetailV2ResponseDataWechatListInner `json:"wechat_list,omitempty"`
}
