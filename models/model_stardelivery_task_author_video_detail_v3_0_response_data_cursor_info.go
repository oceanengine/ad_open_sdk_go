/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskAuthorVideoDetailV30ResponseDataCursorInfo 分页信息
type StardeliveryTaskAuthorVideoDetailV30ResponseDataCursorInfo struct {
	// 页码游标值
	Cursor *int64 `json:"cursor,omitempty"`
	// 是否有下一页
	HasMore *bool `json:"has_more,omitempty"`
}
