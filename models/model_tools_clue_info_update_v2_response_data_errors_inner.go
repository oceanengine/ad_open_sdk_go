/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueInfoUpdateV2ResponseDataErrorsInner struct for ToolsClueInfoUpdateV2ResponseDataErrorsInner
type ToolsClueInfoUpdateV2ResponseDataErrorsInner struct {
	// 更新失败线索id
	Key int64 `json:"key"`
	// 失败原因
	Message *string `json:"message,omitempty"`
}
