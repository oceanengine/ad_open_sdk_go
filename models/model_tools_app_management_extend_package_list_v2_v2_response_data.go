/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementExtendPackageListV2V2ResponseData
type ToolsAppManagementExtendPackageListV2V2ResponseData struct {
	// 分包数据列表
	List     []*ToolsAppManagementExtendPackageListV2V2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ToolsAppManagementAndroidAppListV2ResponseDataPageInfo         `json:"page_info,omitempty"`
}
