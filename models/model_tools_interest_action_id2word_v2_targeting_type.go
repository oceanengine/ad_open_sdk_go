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

// ToolsInterestActionId2wordV2TargetingType
type ToolsInterestActionId2wordV2TargetingType string

// List of tools_interest_action_id2word_v2_targeting_type
const (
	INTEREST_ToolsInterestActionId2wordV2TargetingType ToolsInterestActionId2wordV2TargetingType = "INTEREST"
	ACTION_ToolsInterestActionId2wordV2TargetingType   ToolsInterestActionId2wordV2TargetingType = "ACTION"
)

// All allowed values of ToolsInterestActionId2wordV2TargetingType enum
var AllowedToolsInterestActionId2wordV2TargetingTypeEnumValues = []ToolsInterestActionId2wordV2TargetingType{
	"INTEREST",
	"ACTION",
}

// NewToolsInterestActionId2wordV2TargetingTypeFromValue returns a pointer to a valid ToolsInterestActionId2wordV2TargetingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsInterestActionId2wordV2TargetingTypeFromValue(v string) (*ToolsInterestActionId2wordV2TargetingType, error) {
	ev := ToolsInterestActionId2wordV2TargetingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsInterestActionId2wordV2TargetingType: valid values are %v", v, AllowedToolsInterestActionId2wordV2TargetingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsInterestActionId2wordV2TargetingType) IsValid() bool {
	for _, existing := range AllowedToolsInterestActionId2wordV2TargetingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_interest_action_id2word_v2_targeting_type value
func (v ToolsInterestActionId2wordV2TargetingType) Ptr() *ToolsInterestActionId2wordV2TargetingType {
	return &v
}