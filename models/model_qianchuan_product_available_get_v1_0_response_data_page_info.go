/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanProductAvailableGetV10ResponseDataPageInfo
type QianchuanProductAvailableGetV10ResponseDataPageInfo struct {
	// 页码游标值
	Cursor *int64 `json:"cursor,omitempty"`
	// 是否有更多数据
	HasMore *bool `json:"has_more,omitempty"`
	//
	Page *int64 `json:"page,omitempty"`
	//
	PageSize *int64 `json:"page_size,omitempty"`
	//
	TotalNumber *int64 `json:"total_number,omitempty"`
	//
	TotalPage *int64 `json:"total_page,omitempty"`
}
