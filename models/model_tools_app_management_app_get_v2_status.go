/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementAppGetV2Status
type ToolsAppManagementAppGetV2Status string

// List of tools_app_management_app_get_v2_status
const (
	ALL_ToolsAppManagementAppGetV2Status            ToolsAppManagementAppGetV2Status = "ALL"
	AUDIT_ACCEPTED_ToolsAppManagementAppGetV2Status ToolsAppManagementAppGetV2Status = "AUDIT_ACCEPTED"
	AUDIT_DOING_ToolsAppManagementAppGetV2Status    ToolsAppManagementAppGetV2Status = "AUDIT_DOING"
	AUDIT_REJECTED_ToolsAppManagementAppGetV2Status ToolsAppManagementAppGetV2Status = "AUDIT_REJECTED"
	ENABLE_ToolsAppManagementAppGetV2Status         ToolsAppManagementAppGetV2Status = "ENABLE"
)

// Ptr returns reference to tools_app_management_app_get_v2_status value
func (v ToolsAppManagementAppGetV2Status) Ptr() *ToolsAppManagementAppGetV2Status {
	return &v
}
