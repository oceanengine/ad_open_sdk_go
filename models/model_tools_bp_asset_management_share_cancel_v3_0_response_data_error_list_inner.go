/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBpAssetManagementShareCancelV30ResponseDataErrorListInner struct for ToolsBpAssetManagementShareCancelV30ResponseDataErrorListInner
type ToolsBpAssetManagementShareCancelV30ResponseDataErrorListInner struct {
	AccountInfo          *ToolsBpAssetManagementShareCancelV30ResponseDataErrorListInnerAccountInfo          `json:"account_info,omitempty"`
	AllAccountsByBp      *ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByBp                   `json:"all_accounts_by_bp,omitempty"`
	AllAccountsByCompany *ToolsBpAssetManagementShareCancelV30ResponseDataErrorListInnerAllAccountsByCompany `json:"all_accounts_by_company,omitempty"`
	//
	ErrorMessage *string                                                     `json:"error_message,omitempty"`
	ShareMode    *ToolsBpAssetManagementShareCancelV30DataErrorListShareMode `json:"share_mode,omitempty"`
}
