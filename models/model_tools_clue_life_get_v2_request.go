/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueLifeGetV2Request struct for ToolsClueLifeGetV2Request
type ToolsClueLifeGetV2Request struct {
	// 查询截止时间，格式：yyyy-MM-dd hh:mm:ss
	EndTime string `json:"end_time"`
	// 广告主ids，上限50
	LocalAccountIds []int64 `json:"local_account_ids"`
	// 页数 默认值: 1
	Page int32 `json:"page"`
	// 页面大小 默认值: 10，上限：100
	PageSize int32 `json:"page_size"`
	// 查询起始时间，格式：yyyy-MM-dd hh:mm:ss
	StartTime string `json:"start_time"`
}
