/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableListGetV2Status
type ToolsPlayableListGetV2Status string

// List of tools_playable_list_get_v2_status
const (
	AUDIT_FAIL_ToolsPlayableListGetV2Status       ToolsPlayableListGetV2Status = "AUDIT_FAIL"
	VALIDATE_FAIL_ToolsPlayableListGetV2Status    ToolsPlayableListGetV2Status = "VALIDATE_FAIL"
	AUDIT_SUCCESS_ToolsPlayableListGetV2Status    ToolsPlayableListGetV2Status = "AUDIT_SUCCESS"
	VALIDATING_ToolsPlayableListGetV2Status       ToolsPlayableListGetV2Status = "VALIDATING"
	VALIDATE_SUCCESS_ToolsPlayableListGetV2Status ToolsPlayableListGetV2Status = "VALIDATE_SUCCESS"
)

// Ptr returns reference to tools_playable_list_get_v2_status value
func (v ToolsPlayableListGetV2Status) Ptr() *ToolsPlayableListGetV2Status {
	return &v
}
