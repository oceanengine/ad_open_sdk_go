/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsRubeexGetV2FilteringProjectLifecycle
type ToolsRubeexGetV2FilteringProjectLifecycle string

// List of tools_rubeex_get_v2_filtering_project_lifecycle
const (
	LAUNCHED_ToolsRubeexGetV2FilteringProjectLifecycle      ToolsRubeexGetV2FilteringProjectLifecycle = "LAUNCHED"
	RELAT_PLAN_ToolsRubeexGetV2FilteringProjectLifecycle    ToolsRubeexGetV2FilteringProjectLifecycle = "RELAT_PLAN"
	AUDIT_SUCCESS_ToolsRubeexGetV2FilteringProjectLifecycle ToolsRubeexGetV2FilteringProjectLifecycle = "AUDIT_SUCCESS"
	SYNC_AD_ToolsRubeexGetV2FilteringProjectLifecycle       ToolsRubeexGetV2FilteringProjectLifecycle = "SYNC_AD"
	EDITING_ToolsRubeexGetV2FilteringProjectLifecycle       ToolsRubeexGetV2FilteringProjectLifecycle = "EDITING"
)

// Ptr returns reference to tools_rubeex_get_v2_filtering_project_lifecycle value
func (v ToolsRubeexGetV2FilteringProjectLifecycle) Ptr() *ToolsRubeexGetV2FilteringProjectLifecycle {
	return &v
}
