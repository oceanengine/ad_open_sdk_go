/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeBannedCreateV30ResponseData
type ToolsAwemeBannedCreateV30ResponseData struct {
	// 添加失败的结果
	Fail []*ToolsAwemeBannedCreateV30ResponseDataFailInner `json:"fail,omitempty"`
	// 添加成功的结果
	Success []string `json:"success,omitempty"`
}
