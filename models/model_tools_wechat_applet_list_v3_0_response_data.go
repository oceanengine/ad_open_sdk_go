/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatAppletListV30ResponseData
type ToolsWechatAppletListV30ResponseData struct {
	//
	List     []*ToolsWechatAppletListV30ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ToolsWechatAppletListV30ResponseDataPageInfo    `json:"page_info,omitempty"`
}
