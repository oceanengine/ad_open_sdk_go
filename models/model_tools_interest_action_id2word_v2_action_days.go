/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsInterestActionId2wordV2ActionDays
type ToolsInterestActionId2wordV2ActionDays int64

// List of tools_interest_action_id2word_v2_action_days
const (
	Enum_7_ToolsInterestActionId2wordV2ActionDays   ToolsInterestActionId2wordV2ActionDays = 7
	Enum_15_ToolsInterestActionId2wordV2ActionDays  ToolsInterestActionId2wordV2ActionDays = 15
	Enum_30_ToolsInterestActionId2wordV2ActionDays  ToolsInterestActionId2wordV2ActionDays = 30
	Enum_60_ToolsInterestActionId2wordV2ActionDays  ToolsInterestActionId2wordV2ActionDays = 60
	Enum_90_ToolsInterestActionId2wordV2ActionDays  ToolsInterestActionId2wordV2ActionDays = 90
	Enum_180_ToolsInterestActionId2wordV2ActionDays ToolsInterestActionId2wordV2ActionDays = 180
	Enum_365_ToolsInterestActionId2wordV2ActionDays ToolsInterestActionId2wordV2ActionDays = 365
)

// All allowed values of ToolsInterestActionId2wordV2ActionDays enum
var AllowedToolsInterestActionId2wordV2ActionDaysEnumValues = []ToolsInterestActionId2wordV2ActionDays{
	7,
	15,
	30,
	60,
	90,
	180,
	365,
}

// NewToolsInterestActionId2wordV2ActionDaysFromValue returns a pointer to a valid ToolsInterestActionId2wordV2ActionDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsInterestActionId2wordV2ActionDaysFromValue(v int64) (*ToolsInterestActionId2wordV2ActionDays, error) {
	ev := ToolsInterestActionId2wordV2ActionDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsInterestActionId2wordV2ActionDays: valid values are %v", v, AllowedToolsInterestActionId2wordV2ActionDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsInterestActionId2wordV2ActionDays) IsValid() bool {
	for _, existing := range AllowedToolsInterestActionId2wordV2ActionDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_interest_action_id2word_v2_action_days value
func (v ToolsInterestActionId2wordV2ActionDays) Ptr() *ToolsInterestActionId2wordV2ActionDays {
	return &v
}