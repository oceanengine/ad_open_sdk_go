/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EventManagerTrackUrlGetV2ResponseDataPageInfo 分页信息
type EventManagerTrackUrlGetV2ResponseDataPageInfo struct {
	// 当前页 格式: int64
	Page *int64 `json:"page,omitempty"`
	// 页面大小 格式: int64
	PageSize *int64 `json:"page_size,omitempty"`
	// 总数 格式: int64
	TotalNumber *int64 `json:"total_number,omitempty"`
	// 总页数 格式: int64
	TotalPage *int64 `json:"total_page,omitempty"`
}
