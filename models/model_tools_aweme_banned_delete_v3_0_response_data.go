/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeBannedDeleteV30ResponseData
type ToolsAwemeBannedDeleteV30ResponseData struct {
	// 删除失败的结果
	Fail []*ToolsAwemeBannedCreateV30ResponseDataFailInner `json:"fail,omitempty"`
	// 删除成功的结果
	Success []string `json:"success,omitempty"`
}
