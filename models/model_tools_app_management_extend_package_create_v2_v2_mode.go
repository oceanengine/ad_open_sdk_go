/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementExtendPackageCreateV2V2Mode
type ToolsAppManagementExtendPackageCreateV2V2Mode string

// List of tools_app_management_extend_package_create_v2_v2_mode
const (
	AUTO_ToolsAppManagementExtendPackageCreateV2V2Mode      ToolsAppManagementExtendPackageCreateV2V2Mode = "Auto"
	CUSTOMIZE_ToolsAppManagementExtendPackageCreateV2V2Mode ToolsAppManagementExtendPackageCreateV2V2Mode = "Customize"
	MANUAL_ToolsAppManagementExtendPackageCreateV2V2Mode    ToolsAppManagementExtendPackageCreateV2V2Mode = "Manual"
)

// Ptr returns reference to tools_app_management_extend_package_create_v2_v2_mode value
func (v ToolsAppManagementExtendPackageCreateV2V2Mode) Ptr() *ToolsAppManagementExtendPackageCreateV2V2Mode {
	return &v
}
