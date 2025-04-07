/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementHarmonyAppListV2FilteringStatus
type ToolsAppManagementHarmonyAppListV2FilteringStatus string

// List of tools_app_management_harmony_app_list_v2_filtering_status
const (
	ALL_ToolsAppManagementHarmonyAppListV2FilteringStatus            ToolsAppManagementHarmonyAppListV2FilteringStatus = "ALL"
	AUDIT_DOING_ToolsAppManagementHarmonyAppListV2FilteringStatus    ToolsAppManagementHarmonyAppListV2FilteringStatus = "AUDIT_DOING"
	AUDIT_REJECTED_ToolsAppManagementHarmonyAppListV2FilteringStatus ToolsAppManagementHarmonyAppListV2FilteringStatus = "AUDIT_REJECTED"
	AUDIT_ACCEPTED_ToolsAppManagementHarmonyAppListV2FilteringStatus ToolsAppManagementHarmonyAppListV2FilteringStatus = "AUDIT_ACCEPTED"
	ENABLE_ToolsAppManagementHarmonyAppListV2FilteringStatus         ToolsAppManagementHarmonyAppListV2FilteringStatus = "ENABLE"
)

// Ptr returns reference to tools_app_management_harmony_app_list_v2_filtering_status value
func (v ToolsAppManagementHarmonyAppListV2FilteringStatus) Ptr() *ToolsAppManagementHarmonyAppListV2FilteringStatus {
	return &v
}
