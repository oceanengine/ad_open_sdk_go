/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPlayableListGetV2Status
type ToolsPlayableListGetV2Status string

// List of tools_playable_list_get_v2_status
const (
	AUDIT_FAIL_ToolsPlayableListGetV2Status       ToolsPlayableListGetV2Status = "AUDIT_FAIL"
	VALIDATING_ToolsPlayableListGetV2Status       ToolsPlayableListGetV2Status = "VALIDATING"
	VALIDATE_SUCCESS_ToolsPlayableListGetV2Status ToolsPlayableListGetV2Status = "VALIDATE_SUCCESS"
	VALIDATE_FAIL_ToolsPlayableListGetV2Status    ToolsPlayableListGetV2Status = "VALIDATE_FAIL"
	AUDIT_SUCCESS_ToolsPlayableListGetV2Status    ToolsPlayableListGetV2Status = "AUDIT_SUCCESS"
)

// All allowed values of ToolsPlayableListGetV2Status enum
var AllowedToolsPlayableListGetV2StatusEnumValues = []ToolsPlayableListGetV2Status{
	"AUDIT_FAIL",
	"VALIDATING",
	"VALIDATE_SUCCESS",
	"VALIDATE_FAIL",
	"AUDIT_SUCCESS",
}

// NewToolsPlayableListGetV2StatusFromValue returns a pointer to a valid ToolsPlayableListGetV2Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableListGetV2StatusFromValue(v string) (*ToolsPlayableListGetV2Status, error) {
	ev := ToolsPlayableListGetV2Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableListGetV2Status: valid values are %v", v, AllowedToolsPlayableListGetV2StatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableListGetV2Status) IsValid() bool {
	for _, existing := range AllowedToolsPlayableListGetV2StatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_list_get_v2_status value
func (v ToolsPlayableListGetV2Status) Ptr() *ToolsPlayableListGetV2Status {
	return &v
}