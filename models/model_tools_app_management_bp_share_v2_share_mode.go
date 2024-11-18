/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementBpShareV2ShareMode
type ToolsAppManagementBpShareV2ShareMode string

// List of tools_app_management_bp_share_v2_share_mode
const (
	COMPANY_ToolsAppManagementBpShareV2ShareMode ToolsAppManagementBpShareV2ShareMode = "COMPANY"
	PART_ToolsAppManagementBpShareV2ShareMode    ToolsAppManagementBpShareV2ShareMode = "PART"
	ALL_ToolsAppManagementBpShareV2ShareMode     ToolsAppManagementBpShareV2ShareMode = "ALL"
)

// Ptr returns reference to tools_app_management_bp_share_v2_share_mode value
func (v ToolsAppManagementBpShareV2ShareMode) Ptr() *ToolsAppManagementBpShareV2ShareMode {
	return &v
}
