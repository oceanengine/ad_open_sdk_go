/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementShareAccountListV2ResponseDataListInner struct for ToolsAppManagementShareAccountListV2ResponseDataListInner
type ToolsAppManagementShareAccountListV2ResponseDataListInner struct {
	AccountInfo         *ToolsAppManagementShareAccountListV2ResponseDataListInnerAccountInfo         `json:"account_info,omitempty"`
	AllAccount          *ToolsAppManagementShareAccountListV2ResponseDataListInnerAllAccount          `json:"all_account,omitempty"`
	AllAccountByCompany *ToolsAppManagementShareAccountListV2ResponseDataListInnerAllAccountByCompany `json:"all_account_by_company,omitempty"`
	ShareMode           *ToolsAppManagementShareAccountListV2DataListShareMode                        `json:"share_mode,omitempty"`
}
