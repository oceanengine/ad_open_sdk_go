/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteGetV2ResponseData
type ToolsSiteGetV2ResponseData struct {
	// 建站列表
	List     []*ToolsSiteGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ToolsSiteGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
