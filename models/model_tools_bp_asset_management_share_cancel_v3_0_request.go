/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBpAssetManagementShareCancelV30Request struct for ToolsBpAssetManagementShareCancelV30Request
type ToolsBpAssetManagementShareCancelV30Request struct {
	//
	AccountInfos []*ToolsBpAssetManagementShareCancelV30RequestAccountInfosInner `json:"account_infos,omitempty"`
	//
	AllAccountsByBp []*ToolsBpAssetManagementShareCancelV30AllAccountsByBp `json:"all_accounts_by_bp,omitempty"`
	//
	AllAccountsByCompany []*ToolsBpAssetManagementShareCancelV30RequestAllAccountsByCompanyInner `json:"all_accounts_by_company,omitempty"`
	AssetType            ToolsBpAssetManagementShareCancelV30AssetType                           `json:"asset_type"`
	//
	InstanceId int64 `json:"instance_id"`
	//
	OrganizationId int64                                         `json:"organization_id"`
	ShareMode      ToolsBpAssetManagementShareCancelV30ShareMode `json:"share_mode"`
}
