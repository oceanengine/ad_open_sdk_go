/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppListV30ResponseData
type ToolsMicroAppListV30ResponseData struct {
	//
	List     []*ToolsMicroAppListV30ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ToolsMicroAppListV30ResponseDataPageInfo    `json:"page_info,omitempty"`
}
