/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BudgetGroupListV30ResponseDataPageInfo
type BudgetGroupListV30ResponseDataPageInfo struct {
	//
	Page *int64 `json:"page,omitempty"`
	//
	PageSize *int64 `json:"page_size,omitempty"`
	// 总数
	TotalNumber *int64 `json:"total_number,omitempty"`
	// 总页数
	TotalPage *int64 `json:"total_page,omitempty"`
}
