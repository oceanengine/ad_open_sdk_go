/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus
type ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus string

// List of tools_app_management_android_basic_package_get_v2_data_current_version_status
const (
	ALL_ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus            ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus = "ALL"
	AUDIT_ACCEPTED_ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus = "AUDIT_ACCEPTED"
	AUDIT_DOING_ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus    ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus = "AUDIT_DOING"
	AUDIT_REJECTED_ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus = "AUDIT_REJECTED"
	ENABLE_ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus         ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus = "ENABLE"
)

// Ptr returns reference to tools_app_management_android_basic_package_get_v2_data_current_version_status value
func (v ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus) Ptr() *ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus {
	return &v
}
