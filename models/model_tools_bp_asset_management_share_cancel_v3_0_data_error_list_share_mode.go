/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBpAssetManagementShareCancelV30DataErrorListShareMode
type ToolsBpAssetManagementShareCancelV30DataErrorListShareMode string

// List of tools_bp_asset_management_share_cancel_v3.0_data_error_list_share_mode
const (
	BP_ALL_ACCOUNTS_ToolsBpAssetManagementShareCancelV30DataErrorListShareMode      ToolsBpAssetManagementShareCancelV30DataErrorListShareMode = "BP_ALL_ACCOUNTS"
	COMPANY_ALL_ACCOUNTS_ToolsBpAssetManagementShareCancelV30DataErrorListShareMode ToolsBpAssetManagementShareCancelV30DataErrorListShareMode = "COMPANY_ALL_ACCOUNTS"
	PART_ToolsBpAssetManagementShareCancelV30DataErrorListShareMode                 ToolsBpAssetManagementShareCancelV30DataErrorListShareMode = "PART"
)

// Ptr returns reference to tools_bp_asset_management_share_cancel_v3.0_data_error_list_share_mode value
func (v ToolsBpAssetManagementShareCancelV30DataErrorListShareMode) Ptr() *ToolsBpAssetManagementShareCancelV30DataErrorListShareMode {
	return &v
}
