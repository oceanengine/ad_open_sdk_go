/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteHandselV2ResponseData
type ToolsSiteHandselV2ResponseData struct {
	// 转赠失败列表，整体成功不返回该列表
	ErrorList []*ToolsSiteHandselV2ResponseDataErrorListInner `json:"error_list,omitempty"`
	// 转赠成功列表，整体失败不返回该列表
	SuccessList []*ToolsSiteHandselV2ResponseDataSuccessListInner `json:"success_list,omitempty"`
}