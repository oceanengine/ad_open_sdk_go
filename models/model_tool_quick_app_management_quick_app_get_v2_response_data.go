/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolQuickAppManagementQuickAppGetV2ResponseData
type ToolQuickAppManagementQuickAppGetV2ResponseData struct {
	PageInfo ToolQuickAppManagementQuickAppGetV2ResponseDataPageInfo `json:"page_info"`
	//
	QuickAppInfo []*ToolQuickAppManagementQuickAppGetV2ResponseDataQuickAppInfoInner `json:"quick_app_info"`
}
