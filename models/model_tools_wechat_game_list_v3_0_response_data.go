/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatGameListV30ResponseData
type ToolsWechatGameListV30ResponseData struct {
	//
	List     []*ToolsWechatGameListV30ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ToolsWechatGameListV30ResponseDataPageInfo    `json:"page_info,omitempty"`
}
