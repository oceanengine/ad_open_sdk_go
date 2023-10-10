/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBpAssetManagementShareV30Request struct for ToolsBpAssetManagementShareV30Request
type ToolsBpAssetManagementShareV30Request struct {
	// 部分共享的账户列表
	AccountInfos []*ToolsBpAssetManagementShareV30RequestAccountInfosInner `json:"account_infos,omitempty"`
	// 组织内业务线
	AllAccountsByBp []*ToolsBpAssetManagementShareV30AllAccountsByBp `json:"all_accounts_by_bp,omitempty"`
	// 主体共享信息
	AllAccountsByCompany []*ToolsBpAssetManagementShareV30RequestAllAccountsByCompanyInner `json:"all_accounts_by_company,omitempty"`
	AssetType            ToolsBpAssetManagementShareV30AssetType                           `json:"asset_type"`
	// 资产ID
	InstanceId int64 `json:"instance_id"`
	// 组织ID
	OrganizationId int64                                   `json:"organization_id"`
	ShareMode      ToolsBpAssetManagementShareV30ShareMode `json:"share_mode"`
}
