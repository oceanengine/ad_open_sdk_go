/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolQuickAppManagementQuickAppGetV2Status
type ToolQuickAppManagementQuickAppGetV2Status string

// List of tool_quick_app_management_quick_app_get_v2_status
const (
	AUDIT_ACCEPTED_ToolQuickAppManagementQuickAppGetV2Status      ToolQuickAppManagementQuickAppGetV2Status = "AUDIT_ACCEPTED"
	AUDIT_DOING_ToolQuickAppManagementQuickAppGetV2Status         ToolQuickAppManagementQuickAppGetV2Status = "AUDIT_DOING"
	AUDIT_REJECTED_ToolQuickAppManagementQuickAppGetV2Status      ToolQuickAppManagementQuickAppGetV2Status = "AUDIT_REJECTED"
	AUDIT_SEND_REJECTED_ToolQuickAppManagementQuickAppGetV2Status ToolQuickAppManagementQuickAppGetV2Status = "AUDIT_SEND_REJECTED"
	REMOVED_ToolQuickAppManagementQuickAppGetV2Status             ToolQuickAppManagementQuickAppGetV2Status = "REMOVED"
)

// Ptr returns reference to tool_quick_app_management_quick_app_get_v2_status value
func (v ToolQuickAppManagementQuickAppGetV2Status) Ptr() *ToolQuickAppManagementQuickAppGetV2Status {
	return &v
}
