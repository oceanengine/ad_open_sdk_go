/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueLifeGetV2ResponseDataPageInfo 分页信息
type ToolsClueLifeGetV2ResponseDataPageInfo struct {
	// 页面大小
	PageNumber *int32 `json:"page_number,omitempty"`
	// 页数
	PageSize *int32 `json:"page_size,omitempty"`
	// 总页数
	PageTotal *int32 `json:"page_total,omitempty"`
	// 总数
	Total *int32 `json:"total,omitempty"`
}
