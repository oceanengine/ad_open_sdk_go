/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileAudioGetV2ResponseDataPageInfo
type FileAudioGetV2ResponseDataPageInfo struct {
	// 页数
	Page *int32 `json:"page,omitempty"`
	// 页面大小
	PageSize *int32 `json:"page_size,omitempty"`
	// 总数
	TotalNumber *int32 `json:"total_number,omitempty"`
	// 总页数
	TotalPage *int32 `json:"total_page,omitempty"`
}
