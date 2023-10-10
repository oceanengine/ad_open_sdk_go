/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeTemplateDetailGetV2ResponseDataTemplateTextSchemaInner struct for CreativeTemplateDetailGetV2ResponseDataTemplateTextSchemaInner
type CreativeTemplateDetailGetV2ResponseDataTemplateTextSchemaInner struct {
	// 填充内容的键
	Key *string `json:"key,omitempty"`
	// 填充内容的名称，e.g. 卖点主文本、卖点副文本
	Name *string `json:"name,omitempty"`
	// 文本的字符最长长度
	TextMaxLength *int64 `json:"text_max_length,omitempty"`
	// 文本的字符最小长度
	TextMinLength *int64 `json:"text_min_length,omitempty"`
	// 默认填充内容的值（文本类型的填充值为文字）
	Value *string `json:"value,omitempty"`
}
