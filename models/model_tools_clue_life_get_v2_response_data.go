/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueLifeGetV2ResponseData
type ToolsClueLifeGetV2ResponseData struct {
	// 线索列表
	List     []*ToolsClueLifeGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo ToolsClueLifeGetV2ResponseDataPageInfo     `json:"page_info"`
}
