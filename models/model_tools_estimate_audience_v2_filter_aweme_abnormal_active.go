/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2FilterAwemeAbnormalActive
type ToolsEstimateAudienceV2FilterAwemeAbnormalActive int64

// List of tools_estimate_audience_v2_filter_aweme_abnormal_active
const (
	Enum_0_ToolsEstimateAudienceV2FilterAwemeAbnormalActive ToolsEstimateAudienceV2FilterAwemeAbnormalActive = 0
	Enum_1_ToolsEstimateAudienceV2FilterAwemeAbnormalActive ToolsEstimateAudienceV2FilterAwemeAbnormalActive = 1
)

// All allowed values of ToolsEstimateAudienceV2FilterAwemeAbnormalActive enum
var AllowedToolsEstimateAudienceV2FilterAwemeAbnormalActiveEnumValues = []ToolsEstimateAudienceV2FilterAwemeAbnormalActive{
	0,
	1,
}

// NewToolsEstimateAudienceV2FilterAwemeAbnormalActiveFromValue returns a pointer to a valid ToolsEstimateAudienceV2FilterAwemeAbnormalActive
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2FilterAwemeAbnormalActiveFromValue(v int64) (*ToolsEstimateAudienceV2FilterAwemeAbnormalActive, error) {
	ev := ToolsEstimateAudienceV2FilterAwemeAbnormalActive(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2FilterAwemeAbnormalActive: valid values are %v", v, AllowedToolsEstimateAudienceV2FilterAwemeAbnormalActiveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2FilterAwemeAbnormalActive) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2FilterAwemeAbnormalActiveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_filter_aweme_abnormal_active value
func (v ToolsEstimateAudienceV2FilterAwemeAbnormalActive) Ptr() *ToolsEstimateAudienceV2FilterAwemeAbnormalActive {
	return &v
}
