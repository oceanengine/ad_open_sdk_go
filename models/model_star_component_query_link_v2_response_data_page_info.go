/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarComponentQueryLinkV2ResponseDataPageInfo 分页
type StarComponentQueryLinkV2ResponseDataPageInfo struct {
	// 数量
	Limit *int64 `json:"limit,omitempty"`
	// 页码
	Page *int64 `json:"page,omitempty"`
	// 总量
	TotalCount *int64 `json:"total_count,omitempty"`
}
