/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsRubeexPlayableAdListV2ResponseData
type ToolsRubeexPlayableAdListV2ResponseData struct {
	// 计划列表
	AdIds    []int64                                          `json:"ad_ids,omitempty"`
	PageInfo *ToolsRubeexPlayableAdListV2ResponseDataPageInfo `json:"page_info,omitempty"`
}
