/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementExtendPackageListV2FilteringStatus
type ToolsAppManagementExtendPackageListV2FilteringStatus string

// List of tools_app_management_extend_package_list_v2_filtering_status
const (
	PUBLISHED_ToolsAppManagementExtendPackageListV2FilteringStatus       ToolsAppManagementExtendPackageListV2FilteringStatus = "PUBLISHED"
	CREATION_FAILED_ToolsAppManagementExtendPackageListV2FilteringStatus ToolsAppManagementExtendPackageListV2FilteringStatus = "CREATION_FAILED"
	NOT_UPDATE_ToolsAppManagementExtendPackageListV2FilteringStatus      ToolsAppManagementExtendPackageListV2FilteringStatus = "NOT_UPDATE"
	CREATING_ToolsAppManagementExtendPackageListV2FilteringStatus        ToolsAppManagementExtendPackageListV2FilteringStatus = "CREATING"
	ALL_ToolsAppManagementExtendPackageListV2FilteringStatus             ToolsAppManagementExtendPackageListV2FilteringStatus = "ALL"
	UPDATE_FAILED_ToolsAppManagementExtendPackageListV2FilteringStatus   ToolsAppManagementExtendPackageListV2FilteringStatus = "UPDATE_FAILED"
	UPDATING_ToolsAppManagementExtendPackageListV2FilteringStatus        ToolsAppManagementExtendPackageListV2FilteringStatus = "UPDATING"
)

// Ptr returns reference to tools_app_management_extend_package_list_v2_filtering_status value
func (v ToolsAppManagementExtendPackageListV2FilteringStatus) Ptr() *ToolsAppManagementExtendPackageListV2FilteringStatus {
	return &v
}
