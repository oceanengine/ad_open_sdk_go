/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomGetV30ResponseDataPageInfo 分页信息
type ReportCustomGetV30ResponseDataPageInfo struct {
	// 页码
	Page *int64 `json:"page,omitempty"`
	// 页面大小，即每页展示的数据量
	PageSize *int64 `json:"page_size,omitempty"`
	// 总数
	TotalNumber *int64 `json:"total_number,omitempty"`
	// 总页数
	TotalPage *int64 `json:"total_page,omitempty"`
}
