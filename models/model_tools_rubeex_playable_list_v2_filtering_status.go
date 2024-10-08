/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsRubeexPlayableListV2FilteringStatus
type ToolsRubeexPlayableListV2FilteringStatus string

// List of tools_rubeex_playable_list_v2_filtering_status
const (
	AUDIT_FAIL_ToolsRubeexPlayableListV2FilteringStatus       ToolsRubeexPlayableListV2FilteringStatus = "AUDIT_FAIL"
	AUDIT_SUCCESS_ToolsRubeexPlayableListV2FilteringStatus    ToolsRubeexPlayableListV2FilteringStatus = "AUDIT_SUCCESS"
	VALIDATE_FAIL_ToolsRubeexPlayableListV2FilteringStatus    ToolsRubeexPlayableListV2FilteringStatus = "VALIDATE_FAIL"
	VALIDATE_SUCCESS_ToolsRubeexPlayableListV2FilteringStatus ToolsRubeexPlayableListV2FilteringStatus = "VALIDATE_SUCCESS"
	VALIDATING_ToolsRubeexPlayableListV2FilteringStatus       ToolsRubeexPlayableListV2FilteringStatus = "VALIDATING"
)

// Ptr returns reference to tools_rubeex_playable_list_v2_filtering_status value
func (v ToolsRubeexPlayableListV2FilteringStatus) Ptr() *ToolsRubeexPlayableListV2FilteringStatus {
	return &v
}
