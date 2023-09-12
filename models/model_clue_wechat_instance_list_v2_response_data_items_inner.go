/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueWechatInstanceListV2ResponseDataItemsInner struct for ClueWechatInstanceListV2ResponseDataItemsInner
type ClueWechatInstanceListV2ResponseDataItemsInner struct {
	// 复制次数
	ClueCount *int64 `json:"clue_count,omitempty"`
	// 创建时间，格式：2022-07-21 15:36:09
	CreateTime *string `json:"create_time,omitempty"`
	// 微信号码包ID
	InstanceId *int64 `json:"instance_id,omitempty"`
	// 更新时间，格式：2022-07-21 15:36:09
	ModTime *string `json:"mod_time,omitempty"`
	// 微信号码包名称
	Name    *string                                   `json:"name,omitempty"`
	SubType *ClueWechatInstanceListV2DataItemsSubType `json:"sub_type,omitempty"`
}
