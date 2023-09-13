/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementBpShareCancelV2ResponseDataErrorListInner struct for ToolsAppManagementBpShareCancelV2ResponseDataErrorListInner
type ToolsAppManagementBpShareCancelV2ResponseDataErrorListInner struct {
	AccountInfo         *ToolsAppManagementBpShareCancelV2ResponseDataErrorListInnerAccountInfo         `json:"account_info,omitempty"`
	AllAccount          *ToolsAppManagementBpShareCancelV2ResponseDataErrorListInnerAllAccount          `json:"all_account,omitempty"`
	AllAccountByCompany *ToolsAppManagementBpShareCancelV2ResponseDataErrorListInnerAllAccountByCompany `json:"all_account_by_company,omitempty"`
	//
	ErrorMessage *string                                                  `json:"error_message,omitempty"`
	ShareMode    *ToolsAppManagementBpShareCancelV2DataErrorListShareMode `json:"share_mode,omitempty"`
}
