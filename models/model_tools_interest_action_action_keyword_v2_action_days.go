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

// ToolsInterestActionActionKeywordV2ActionDays
type ToolsInterestActionActionKeywordV2ActionDays int64

// List of tools_interest_action_action_keyword_v2_action_days
const (
	Enum_15_ToolsInterestActionActionKeywordV2ActionDays  ToolsInterestActionActionKeywordV2ActionDays = 15
	Enum_180_ToolsInterestActionActionKeywordV2ActionDays ToolsInterestActionActionKeywordV2ActionDays = 180
	Enum_30_ToolsInterestActionActionKeywordV2ActionDays  ToolsInterestActionActionKeywordV2ActionDays = 30
	Enum_365_ToolsInterestActionActionKeywordV2ActionDays ToolsInterestActionActionKeywordV2ActionDays = 365
	Enum_60_ToolsInterestActionActionKeywordV2ActionDays  ToolsInterestActionActionKeywordV2ActionDays = 60
	Enum_7_ToolsInterestActionActionKeywordV2ActionDays   ToolsInterestActionActionKeywordV2ActionDays = 7
	Enum_90_ToolsInterestActionActionKeywordV2ActionDays  ToolsInterestActionActionKeywordV2ActionDays = 90
)

// All allowed values of ToolsInterestActionActionKeywordV2ActionDays enum
var AllowedToolsInterestActionActionKeywordV2ActionDaysEnumValues = []ToolsInterestActionActionKeywordV2ActionDays{
	15,
	180,
	30,
	365,
	60,
	7,
	90,
}

// NewToolsInterestActionActionKeywordV2ActionDaysFromValue returns a pointer to a valid ToolsInterestActionActionKeywordV2ActionDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsInterestActionActionKeywordV2ActionDaysFromValue(v int64) (*ToolsInterestActionActionKeywordV2ActionDays, error) {
	ev := ToolsInterestActionActionKeywordV2ActionDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsInterestActionActionKeywordV2ActionDays: valid values are %v", v, AllowedToolsInterestActionActionKeywordV2ActionDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsInterestActionActionKeywordV2ActionDays) IsValid() bool {
	for _, existing := range AllowedToolsInterestActionActionKeywordV2ActionDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_interest_action_action_keyword_v2_action_days value
func (v ToolsInterestActionActionKeywordV2ActionDays) Ptr() *ToolsInterestActionActionKeywordV2ActionDays {
	return &v
}