/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementExtendPackageCreateV2Mode
type ToolsAppManagementExtendPackageCreateV2Mode string

// List of tools_app_management_extend_package_create_v2_mode
const (
	MANUAL_ToolsAppManagementExtendPackageCreateV2Mode    ToolsAppManagementExtendPackageCreateV2Mode = "Manual"
	AUTO_ToolsAppManagementExtendPackageCreateV2Mode      ToolsAppManagementExtendPackageCreateV2Mode = "Auto"
	CUSTOMIZE_ToolsAppManagementExtendPackageCreateV2Mode ToolsAppManagementExtendPackageCreateV2Mode = "Customize"
)

// Ptr returns reference to tools_app_management_extend_package_create_v2_mode value
func (v ToolsAppManagementExtendPackageCreateV2Mode) Ptr() *ToolsAppManagementExtendPackageCreateV2Mode {
	return &v
}
