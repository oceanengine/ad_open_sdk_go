/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLiveAuthorizeListV2ResponseData
type ToolsLiveAuthorizeListV2ResponseData struct {
	//
	List     []*ToolsLiveAuthorizeListV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ToolsLiveAuthorizeListV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
