/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus
type ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus string

// List of tool_quick_app_management_quick_app_get_v2_data_quick_app_info_status
const (
	AUDIT_ACCEPTED_ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus      ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus = "AUDIT_ACCEPTED"
	AUDIT_DOING_ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus         ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus = "AUDIT_DOING"
	AUDIT_REJECTED_ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus      ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus = "AUDIT_REJECTED"
	AUDIT_SEND_REJECTED_ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus = "AUDIT_SEND_REJECTED"
	REMOVED_ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus             ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus = "REMOVED"
)

// Ptr returns reference to tool_quick_app_management_quick_app_get_v2_data_quick_app_info_status value
func (v ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus) Ptr() *ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus {
	return &v
}
