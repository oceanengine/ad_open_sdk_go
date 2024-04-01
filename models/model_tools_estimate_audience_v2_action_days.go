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

// ToolsEstimateAudienceV2ActionDays
type ToolsEstimateAudienceV2ActionDays int64

// List of tools_estimate_audience_v2_action_days
const (
	Enum_7_ToolsEstimateAudienceV2ActionDays   ToolsEstimateAudienceV2ActionDays = 7
	Enum_15_ToolsEstimateAudienceV2ActionDays  ToolsEstimateAudienceV2ActionDays = 15
	Enum_30_ToolsEstimateAudienceV2ActionDays  ToolsEstimateAudienceV2ActionDays = 30
	Enum_60_ToolsEstimateAudienceV2ActionDays  ToolsEstimateAudienceV2ActionDays = 60
	Enum_90_ToolsEstimateAudienceV2ActionDays  ToolsEstimateAudienceV2ActionDays = 90
	Enum_180_ToolsEstimateAudienceV2ActionDays ToolsEstimateAudienceV2ActionDays = 180
	Enum_365_ToolsEstimateAudienceV2ActionDays ToolsEstimateAudienceV2ActionDays = 365
)

// All allowed values of ToolsEstimateAudienceV2ActionDays enum
var AllowedToolsEstimateAudienceV2ActionDaysEnumValues = []ToolsEstimateAudienceV2ActionDays{
	7,
	15,
	30,
	60,
	90,
	180,
	365,
}

// NewToolsEstimateAudienceV2ActionDaysFromValue returns a pointer to a valid ToolsEstimateAudienceV2ActionDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2ActionDaysFromValue(v int64) (*ToolsEstimateAudienceV2ActionDays, error) {
	ev := ToolsEstimateAudienceV2ActionDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2ActionDays: valid values are %v", v, AllowedToolsEstimateAudienceV2ActionDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2ActionDays) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2ActionDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_action_days value
func (v ToolsEstimateAudienceV2ActionDays) Ptr() *ToolsEstimateAudienceV2ActionDays {
	return &v
}
