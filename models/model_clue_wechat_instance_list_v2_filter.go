/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueWechatInstanceListV2Filter
type ClueWechatInstanceListV2Filter struct {
	// 创建时间截止日期，格式：2022-07-29
	CreateTimeEnd *string `json:"create_time_end,omitempty"`
	// 创建时间起始日期，格式：2022-07-19
	CreateTimeStart *string `json:"create_time_start,omitempty"`
	// 微信号码包IDs，超过100个时报参数错误
	InstanceIds []int64 `json:"instance_ids,omitempty"`
	// 微信号码包名称，模糊匹配
	Name *string `json:"name,omitempty"`
	// 分页，>=1，默认值1
	Page *int64 `json:"page,omitempty"`
	// 页大小，1-100，默认值20
	PageSize *int64 `json:"page_size,omitempty"`
	// 微信号，精确匹配
	WechatNumber *string `json:"wechat_number,omitempty"`
}
