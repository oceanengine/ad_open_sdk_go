/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsRubeexGetV2FilteringProjectVersionLifecycle
type ToolsRubeexGetV2FilteringProjectVersionLifecycle string

// List of tools_rubeex_get_v2_filtering_project_version_lifecycle
const (
	AUDIT_FAIL_ToolsRubeexGetV2FilteringProjectVersionLifecycle    ToolsRubeexGetV2FilteringProjectVersionLifecycle = "AUDIT_FAIL"
	EXPORTED_ToolsRubeexGetV2FilteringProjectVersionLifecycle      ToolsRubeexGetV2FilteringProjectVersionLifecycle = "EXPORTED"
	LAUNCHED_ToolsRubeexGetV2FilteringProjectVersionLifecycle      ToolsRubeexGetV2FilteringProjectVersionLifecycle = "LAUNCHED"
	AUDIT_SUCCESS_ToolsRubeexGetV2FilteringProjectVersionLifecycle ToolsRubeexGetV2FilteringProjectVersionLifecycle = "AUDIT_SUCCESS"
	SYNC_AD_ToolsRubeexGetV2FilteringProjectVersionLifecycle       ToolsRubeexGetV2FilteringProjectVersionLifecycle = "SYNC_AD"
	RELAT_PLAN_ToolsRubeexGetV2FilteringProjectVersionLifecycle    ToolsRubeexGetV2FilteringProjectVersionLifecycle = "RELAT_PLAN"
)

// All allowed values of ToolsRubeexGetV2FilteringProjectVersionLifecycle enum
var AllowedToolsRubeexGetV2FilteringProjectVersionLifecycleEnumValues = []ToolsRubeexGetV2FilteringProjectVersionLifecycle{
	"AUDIT_FAIL",
	"EXPORTED",
	"LAUNCHED",
	"AUDIT_SUCCESS",
	"SYNC_AD",
	"RELAT_PLAN",
}

// NewToolsRubeexGetV2FilteringProjectVersionLifecycleFromValue returns a pointer to a valid ToolsRubeexGetV2FilteringProjectVersionLifecycle
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRubeexGetV2FilteringProjectVersionLifecycleFromValue(v string) (*ToolsRubeexGetV2FilteringProjectVersionLifecycle, error) {
	ev := ToolsRubeexGetV2FilteringProjectVersionLifecycle(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRubeexGetV2FilteringProjectVersionLifecycle: valid values are %v", v, AllowedToolsRubeexGetV2FilteringProjectVersionLifecycleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRubeexGetV2FilteringProjectVersionLifecycle) IsValid() bool {
	for _, existing := range AllowedToolsRubeexGetV2FilteringProjectVersionLifecycleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rubeex_get_v2_filtering_project_version_lifecycle value
func (v ToolsRubeexGetV2FilteringProjectVersionLifecycle) Ptr() *ToolsRubeexGetV2FilteringProjectVersionLifecycle {
	return &v
}
