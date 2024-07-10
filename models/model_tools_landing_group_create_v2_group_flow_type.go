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

// ToolsLandingGroupCreateV2GroupFlowType
type ToolsLandingGroupCreateV2GroupFlowType string

// List of tools_landing_group_create_v2_group_flow_type
const (
	FLOW_DISTRIBUTION_TYPE_AVERAGE_ToolsLandingGroupCreateV2GroupFlowType      ToolsLandingGroupCreateV2GroupFlowType = "FLOW_DISTRIBUTION_TYPE_AVERAGE"
	FLOW_DISTRIBUTION_TYPE_INTELLIGENCE_ToolsLandingGroupCreateV2GroupFlowType ToolsLandingGroupCreateV2GroupFlowType = "FLOW_DISTRIBUTION_TYPE_INTELLIGENCE"
)

// All allowed values of ToolsLandingGroupCreateV2GroupFlowType enum
var AllowedToolsLandingGroupCreateV2GroupFlowTypeEnumValues = []ToolsLandingGroupCreateV2GroupFlowType{
	"FLOW_DISTRIBUTION_TYPE_AVERAGE",
	"FLOW_DISTRIBUTION_TYPE_INTELLIGENCE",
}

// NewToolsLandingGroupCreateV2GroupFlowTypeFromValue returns a pointer to a valid ToolsLandingGroupCreateV2GroupFlowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsLandingGroupCreateV2GroupFlowTypeFromValue(v string) (*ToolsLandingGroupCreateV2GroupFlowType, error) {
	ev := ToolsLandingGroupCreateV2GroupFlowType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsLandingGroupCreateV2GroupFlowType: valid values are %v", v, AllowedToolsLandingGroupCreateV2GroupFlowTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsLandingGroupCreateV2GroupFlowType) IsValid() bool {
	for _, existing := range AllowedToolsLandingGroupCreateV2GroupFlowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_landing_group_create_v2_group_flow_type value
func (v ToolsLandingGroupCreateV2GroupFlowType) Ptr() *ToolsLandingGroupCreateV2GroupFlowType {
	return &v
}
