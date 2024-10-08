/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBpAssetManagementShareCancelV30ShareMode
type ToolsBpAssetManagementShareCancelV30ShareMode string

// List of tools_bp_asset_management_share_cancel_v3.0_share_mode
const (
	BP_ALL_ACCOUNTS_ToolsBpAssetManagementShareCancelV30ShareMode      ToolsBpAssetManagementShareCancelV30ShareMode = "BP_ALL_ACCOUNTS"
	COMPANY_ALL_ACCOUNTS_ToolsBpAssetManagementShareCancelV30ShareMode ToolsBpAssetManagementShareCancelV30ShareMode = "COMPANY_ALL_ACCOUNTS"
	PART_ToolsBpAssetManagementShareCancelV30ShareMode                 ToolsBpAssetManagementShareCancelV30ShareMode = "PART"
)

// Ptr returns reference to tools_bp_asset_management_share_cancel_v3.0_share_mode value
func (v ToolsBpAssetManagementShareCancelV30ShareMode) Ptr() *ToolsBpAssetManagementShareCancelV30ShareMode {
	return &v
}
