/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementBpShareCancelV2ResponseDataSuccessListInner struct for ToolsAppManagementBpShareCancelV2ResponseDataSuccessListInner
type ToolsAppManagementBpShareCancelV2ResponseDataSuccessListInner struct {
	AccountInfo         *ToolsAppManagementBpShareCancelV2ResponseDataSuccessListInnerAccountInfo         `json:"account_info,omitempty"`
	AllAccount          *ToolsAppManagementBpShareCancelV2ResponseDataSuccessListInnerAllAccount          `json:"all_account,omitempty"`
	AllAccountByCompany *ToolsAppManagementBpShareCancelV2ResponseDataSuccessListInnerAllAccountByCompany `json:"all_account_by_company,omitempty"`
	ShareMode           *ToolsAppManagementBpShareCancelV2DataSuccessListShareMode                        `json:"share_mode,omitempty"`
}
