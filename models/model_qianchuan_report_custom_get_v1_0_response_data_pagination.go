/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportCustomGetV10ResponseDataPagination
type QianchuanReportCustomGetV10ResponseDataPagination struct {
	//
	Page int64 `json:"page"`
	//
	PageSize int64 `json:"page_size"`
	//
	TotalNum *int64 `json:"total_num,omitempty"`
	//
	TotalPage *int64 `json:"total_page,omitempty"`
}
