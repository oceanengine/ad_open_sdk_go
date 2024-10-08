/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBpAssetManagementShareV30ResponseDataErrorListInner struct for ToolsBpAssetManagementShareV30ResponseDataErrorListInner
type ToolsBpAssetManagementShareV30ResponseDataErrorListInner struct {
	AccountInfo          *ToolsBpAssetManagementShareV30ResponseDataErrorListInnerAccountInfo          `json:"account_info,omitempty"`
	AllAccountsByBp      *ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp                   `json:"all_accounts_by_bp,omitempty"`
	AllAccountsByCompany *ToolsBpAssetManagementShareV30ResponseDataErrorListInnerAllAccountsByCompany `json:"all_accounts_by_company,omitempty"`
	// 错误原因
	ErrorMessage *string                                               `json:"error_message,omitempty"`
	ShareMode    *ToolsBpAssetManagementShareV30DataErrorListShareMode `json:"share_mode,omitempty"`
}
