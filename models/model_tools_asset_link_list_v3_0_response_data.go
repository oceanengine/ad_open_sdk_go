/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAssetLinkListV30ResponseData
type ToolsAssetLinkListV30ResponseData struct {
	//
	List     []*ToolsAssetLinkListV30ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ToolsAssetLinkListV30ResponseDataPageInfo    `json:"page_info,omitempty"`
}