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

// ToolsClueLifeGetV2DataListEffectiveState
type ToolsClueLifeGetV2DataListEffectiveState int64

// List of tools_clue_life_get_v2_data_list_effective_state
const (
	Enum_0_ToolsClueLifeGetV2DataListEffectiveState   ToolsClueLifeGetV2DataListEffectiveState = 0
	Enum_1_ToolsClueLifeGetV2DataListEffectiveState   ToolsClueLifeGetV2DataListEffectiveState = 1
	Enum_2_ToolsClueLifeGetV2DataListEffectiveState   ToolsClueLifeGetV2DataListEffectiveState = 2
	Enum_3_ToolsClueLifeGetV2DataListEffectiveState   ToolsClueLifeGetV2DataListEffectiveState = 3
	Enum_6_ToolsClueLifeGetV2DataListEffectiveState   ToolsClueLifeGetV2DataListEffectiveState = 6
	Enum_7_ToolsClueLifeGetV2DataListEffectiveState   ToolsClueLifeGetV2DataListEffectiveState = 7
	Enum_204_ToolsClueLifeGetV2DataListEffectiveState ToolsClueLifeGetV2DataListEffectiveState = 204
)

// All allowed values of ToolsClueLifeGetV2DataListEffectiveState enum
var AllowedToolsClueLifeGetV2DataListEffectiveStateEnumValues = []ToolsClueLifeGetV2DataListEffectiveState{
	0,
	1,
	2,
	3,
	6,
	7,
	204,
}

// NewToolsClueLifeGetV2DataListEffectiveStateFromValue returns a pointer to a valid ToolsClueLifeGetV2DataListEffectiveState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueLifeGetV2DataListEffectiveStateFromValue(v int64) (*ToolsClueLifeGetV2DataListEffectiveState, error) {
	ev := ToolsClueLifeGetV2DataListEffectiveState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueLifeGetV2DataListEffectiveState: valid values are %v", v, AllowedToolsClueLifeGetV2DataListEffectiveStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueLifeGetV2DataListEffectiveState) IsValid() bool {
	for _, existing := range AllowedToolsClueLifeGetV2DataListEffectiveStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_life_get_v2_data_list_effective_state value
func (v ToolsClueLifeGetV2DataListEffectiveState) Ptr() *ToolsClueLifeGetV2DataListEffectiveState {
	return &v
}
