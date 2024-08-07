/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsClueLifeGetV2DataListEffectiveStateName
type ToolsClueLifeGetV2DataListEffectiveStateName string

// List of tools_clue_life_get_v2_data_list_effective_state_name
const (
	NOT_MARKED_ToolsClueLifeGetV2DataListEffectiveStateName    ToolsClueLifeGetV2DataListEffectiveStateName = "NOT_MARKED"
	NEED_CALLBACK_ToolsClueLifeGetV2DataListEffectiveStateName ToolsClueLifeGetV2DataListEffectiveStateName = "NEED_CALLBACK"
	CONVERTED_ToolsClueLifeGetV2DataListEffectiveStateName     ToolsClueLifeGetV2DataListEffectiveStateName = "CONVERTED"
	INVALID_ToolsClueLifeGetV2DataListEffectiveStateName       ToolsClueLifeGetV2DataListEffectiveStateName = "INVALID"
	ADD_WEIXIN_ToolsClueLifeGetV2DataListEffectiveStateName    ToolsClueLifeGetV2DataListEffectiveStateName = "ADD_WEIXIN"
	WAIT_CONNECT_ToolsClueLifeGetV2DataListEffectiveStateName  ToolsClueLifeGetV2DataListEffectiveStateName = "WAIT_CONNECT"
	ARRIVE_STORE_ToolsClueLifeGetV2DataListEffectiveStateName  ToolsClueLifeGetV2DataListEffectiveStateName = "ARRIVE_STORE"
)

// All allowed values of ToolsClueLifeGetV2DataListEffectiveStateName enum
var AllowedToolsClueLifeGetV2DataListEffectiveStateNameEnumValues = []ToolsClueLifeGetV2DataListEffectiveStateName{
	"NOT_MARKED",
	"NEED_CALLBACK",
	"CONVERTED",
	"INVALID",
	"ADD_WEIXIN",
	"WAIT_CONNECT",
	"ARRIVE_STORE",
}

// NewToolsClueLifeGetV2DataListEffectiveStateNameFromValue returns a pointer to a valid ToolsClueLifeGetV2DataListEffectiveStateName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueLifeGetV2DataListEffectiveStateNameFromValue(v string) (*ToolsClueLifeGetV2DataListEffectiveStateName, error) {
	ev := ToolsClueLifeGetV2DataListEffectiveStateName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueLifeGetV2DataListEffectiveStateName: valid values are %v", v, AllowedToolsClueLifeGetV2DataListEffectiveStateNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueLifeGetV2DataListEffectiveStateName) IsValid() bool {
	for _, existing := range AllowedToolsClueLifeGetV2DataListEffectiveStateNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_life_get_v2_data_list_effective_state_name value
func (v ToolsClueLifeGetV2DataListEffectiveStateName) Ptr() *ToolsClueLifeGetV2DataListEffectiveStateName {
	return &v
}
