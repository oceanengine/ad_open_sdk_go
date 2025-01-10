/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableListGetV2DataListStatus
type ToolsPlayableListGetV2DataListStatus string

// List of tools_playable_list_get_v2_data_list_status
const (
	VALIDATE_FAIL_ToolsPlayableListGetV2DataListStatus    ToolsPlayableListGetV2DataListStatus = "VALIDATE_FAIL"
	AUDIT_FAIL_ToolsPlayableListGetV2DataListStatus       ToolsPlayableListGetV2DataListStatus = "AUDIT_FAIL"
	AUDIT_SUCCESS_ToolsPlayableListGetV2DataListStatus    ToolsPlayableListGetV2DataListStatus = "AUDIT_SUCCESS"
	VALIDATE_SUCCESS_ToolsPlayableListGetV2DataListStatus ToolsPlayableListGetV2DataListStatus = "VALIDATE_SUCCESS"
	VALIDATING_ToolsPlayableListGetV2DataListStatus       ToolsPlayableListGetV2DataListStatus = "VALIDATING"
)

// Ptr returns reference to tools_playable_list_get_v2_data_list_status value
func (v ToolsPlayableListGetV2DataListStatus) Ptr() *ToolsPlayableListGetV2DataListStatus {
	return &v
}
