/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsClueLifeGetV2DataListFlowType
type ToolsClueLifeGetV2DataListFlowType string

// List of tools_clue_life_get_v2_data_list_flow_type
const (
	NATURE_ToolsClueLifeGetV2DataListFlowType ToolsClueLifeGetV2DataListFlowType = "NATURE"
	AD_ToolsClueLifeGetV2DataListFlowType     ToolsClueLifeGetV2DataListFlowType = "AD"
)

// All allowed values of ToolsClueLifeGetV2DataListFlowType enum
var AllowedToolsClueLifeGetV2DataListFlowTypeEnumValues = []ToolsClueLifeGetV2DataListFlowType{
	"NATURE",
	"AD",
}

// NewToolsClueLifeGetV2DataListFlowTypeFromValue returns a pointer to a valid ToolsClueLifeGetV2DataListFlowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueLifeGetV2DataListFlowTypeFromValue(v string) (*ToolsClueLifeGetV2DataListFlowType, error) {
	ev := ToolsClueLifeGetV2DataListFlowType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueLifeGetV2DataListFlowType: valid values are %v", v, AllowedToolsClueLifeGetV2DataListFlowTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueLifeGetV2DataListFlowType) IsValid() bool {
	for _, existing := range AllowedToolsClueLifeGetV2DataListFlowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_life_get_v2_data_list_flow_type value
func (v ToolsClueLifeGetV2DataListFlowType) Ptr() *ToolsClueLifeGetV2DataListFlowType {
	return &v
}
