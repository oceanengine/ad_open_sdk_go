/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppIosListV2DataListStatus
type ToolsAppIosListV2DataListStatus string

// List of tools_app_ios_list_v2_data_list_status
const (
	ENABLE_ToolsAppIosListV2DataListStatus         ToolsAppIosListV2DataListStatus = "ENABLE"
	AUDIT_ACCEPTED_ToolsAppIosListV2DataListStatus ToolsAppIosListV2DataListStatus = "AUDIT_ACCEPTED"
	AUDIT_DOING_ToolsAppIosListV2DataListStatus    ToolsAppIosListV2DataListStatus = "AUDIT_DOING"
	ALL_ToolsAppIosListV2DataListStatus            ToolsAppIosListV2DataListStatus = "ALL"
	AUDIT_REJECTED_ToolsAppIosListV2DataListStatus ToolsAppIosListV2DataListStatus = "AUDIT_REJECTED"
)

// Ptr returns reference to tools_app_ios_list_v2_data_list_status value
func (v ToolsAppIosListV2DataListStatus) Ptr() *ToolsAppIosListV2DataListStatus {
	return &v
}
