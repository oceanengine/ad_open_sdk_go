/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsClueLifeGetV2DataListFollowStateName
type ToolsClueLifeGetV2DataListFollowStateName string

// List of tools_clue_life_get_v2_data_list_follow_state_name
const (
	NOT_CALLED_ToolsClueLifeGetV2DataListFollowStateName     ToolsClueLifeGetV2DataListFollowStateName = "NOT_CALLED"
	NOT_ANSWERED_ToolsClueLifeGetV2DataListFollowStateName   ToolsClueLifeGetV2DataListFollowStateName = "NOT_ANSWERED"
	SHORT_ANSWERED_ToolsClueLifeGetV2DataListFollowStateName ToolsClueLifeGetV2DataListFollowStateName = "SHORT_ANSWERED"
	ANSWERED_ToolsClueLifeGetV2DataListFollowStateName       ToolsClueLifeGetV2DataListFollowStateName = "ANSWERED"
	DEEP_ANSWERED_ToolsClueLifeGetV2DataListFollowStateName  ToolsClueLifeGetV2DataListFollowStateName = "DEEP_ANSWERED"
)

// All allowed values of ToolsClueLifeGetV2DataListFollowStateName enum
var AllowedToolsClueLifeGetV2DataListFollowStateNameEnumValues = []ToolsClueLifeGetV2DataListFollowStateName{
	"NOT_CALLED",
	"NOT_ANSWERED",
	"SHORT_ANSWERED",
	"ANSWERED",
	"DEEP_ANSWERED",
}

// NewToolsClueLifeGetV2DataListFollowStateNameFromValue returns a pointer to a valid ToolsClueLifeGetV2DataListFollowStateName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueLifeGetV2DataListFollowStateNameFromValue(v string) (*ToolsClueLifeGetV2DataListFollowStateName, error) {
	ev := ToolsClueLifeGetV2DataListFollowStateName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueLifeGetV2DataListFollowStateName: valid values are %v", v, AllowedToolsClueLifeGetV2DataListFollowStateNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueLifeGetV2DataListFollowStateName) IsValid() bool {
	for _, existing := range AllowedToolsClueLifeGetV2DataListFollowStateNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_life_get_v2_data_list_follow_state_name value
func (v ToolsClueLifeGetV2DataListFollowStateName) Ptr() *ToolsClueLifeGetV2DataListFollowStateName {
	return &v
}
