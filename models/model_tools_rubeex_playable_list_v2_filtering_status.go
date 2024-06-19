/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

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

// All allowed values of ToolsRubeexPlayableListV2FilteringStatus enum
var AllowedToolsRubeexPlayableListV2FilteringStatusEnumValues = []ToolsRubeexPlayableListV2FilteringStatus{
	"AUDIT_FAIL",
	"AUDIT_SUCCESS",
	"VALIDATE_FAIL",
	"VALIDATE_SUCCESS",
	"VALIDATING",
}

// NewToolsRubeexPlayableListV2FilteringStatusFromValue returns a pointer to a valid ToolsRubeexPlayableListV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRubeexPlayableListV2FilteringStatusFromValue(v string) (*ToolsRubeexPlayableListV2FilteringStatus, error) {
	ev := ToolsRubeexPlayableListV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRubeexPlayableListV2FilteringStatus: valid values are %v", v, AllowedToolsRubeexPlayableListV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRubeexPlayableListV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedToolsRubeexPlayableListV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rubeex_playable_list_v2_filtering_status value
func (v ToolsRubeexPlayableListV2FilteringStatus) Ptr() *ToolsRubeexPlayableListV2FilteringStatus {
	return &v
}
