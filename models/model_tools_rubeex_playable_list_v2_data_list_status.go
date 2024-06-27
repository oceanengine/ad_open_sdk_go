/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsRubeexPlayableListV2DataListStatus
type ToolsRubeexPlayableListV2DataListStatus string

// List of tools_rubeex_playable_list_v2_data_list_status
const (
	AUDIT_FAIL_ToolsRubeexPlayableListV2DataListStatus       ToolsRubeexPlayableListV2DataListStatus = "AUDIT_FAIL"
	AUDIT_SUCCESS_ToolsRubeexPlayableListV2DataListStatus    ToolsRubeexPlayableListV2DataListStatus = "AUDIT_SUCCESS"
	VALIDATE_FAIL_ToolsRubeexPlayableListV2DataListStatus    ToolsRubeexPlayableListV2DataListStatus = "VALIDATE_FAIL"
	VALIDATE_SUCCESS_ToolsRubeexPlayableListV2DataListStatus ToolsRubeexPlayableListV2DataListStatus = "VALIDATE_SUCCESS"
	VALIDATING_ToolsRubeexPlayableListV2DataListStatus       ToolsRubeexPlayableListV2DataListStatus = "VALIDATING"
)

// All allowed values of ToolsRubeexPlayableListV2DataListStatus enum
var AllowedToolsRubeexPlayableListV2DataListStatusEnumValues = []ToolsRubeexPlayableListV2DataListStatus{
	"AUDIT_FAIL",
	"AUDIT_SUCCESS",
	"VALIDATE_FAIL",
	"VALIDATE_SUCCESS",
	"VALIDATING",
}

// NewToolsRubeexPlayableListV2DataListStatusFromValue returns a pointer to a valid ToolsRubeexPlayableListV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRubeexPlayableListV2DataListStatusFromValue(v string) (*ToolsRubeexPlayableListV2DataListStatus, error) {
	ev := ToolsRubeexPlayableListV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRubeexPlayableListV2DataListStatus: valid values are %v", v, AllowedToolsRubeexPlayableListV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRubeexPlayableListV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedToolsRubeexPlayableListV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rubeex_playable_list_v2_data_list_status value
func (v ToolsRubeexPlayableListV2DataListStatus) Ptr() *ToolsRubeexPlayableListV2DataListStatus {
	return &v
}
