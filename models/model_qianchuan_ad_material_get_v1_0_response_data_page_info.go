/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdMaterialGetV10ResponseDataPageInfo 分页结果
type QianchuanAdMaterialGetV10ResponseDataPageInfo struct {
	//
	Page int64 `json:"page"`
	//
	PageSize int64 `json:"page_size"`
	//
	TotalNum *int64 `json:"total_num,omitempty"`
	//
	TotalPage *int64 `json:"total_page,omitempty"`
}
