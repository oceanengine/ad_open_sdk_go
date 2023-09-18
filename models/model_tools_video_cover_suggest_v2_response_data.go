/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsVideoCoverSuggestV2ResponseData
type ToolsVideoCoverSuggestV2ResponseData struct {
	// 视频封面列表
	List []*ToolsVideoCoverSuggestV2ResponseDataListInner `json:"list,omitempty"`
	// 封面生成的状态枚举值：RUNNING（生成中）、SUCCESS（成功）、FAILED（失败）注意：视频封面不是实时生成，需要根据status字段判断封面生成的状态，大概需要6，7s中生成，所以当status=RUNNING时请等待后重试
	Status *string `json:"status,omitempty"`
}