/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthV2ResponseData
type ToolsAwemeAuthV2ResponseData struct {
	// 发起授权是否成功标记
	AuthStatus *string `json:"auth_status,omitempty"`
	// 错误信息
	ErrMsg *string `json:"err_msg,omitempty"`
}
