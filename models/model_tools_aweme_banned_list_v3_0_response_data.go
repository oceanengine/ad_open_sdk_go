/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeBannedListV30ResponseData
type ToolsAwemeBannedListV30ResponseData struct {
	// 屏蔽用户列表
	List     []*ToolsAwemeBannedListV30ResponseDataListInner `json:"list"`
	PageInfo ToolsAwemeBannedListV30ResponseDataPageInfo     `json:"page_info"`
}
