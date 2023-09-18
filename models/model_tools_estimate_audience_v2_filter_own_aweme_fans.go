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

// ToolsEstimateAudienceV2FilterOwnAwemeFans
type ToolsEstimateAudienceV2FilterOwnAwemeFans int64

// List of tools_estimate_audience_v2_filter_own_aweme_fans
const (
	Enum_0_ToolsEstimateAudienceV2FilterOwnAwemeFans ToolsEstimateAudienceV2FilterOwnAwemeFans = 0
	Enum_1_ToolsEstimateAudienceV2FilterOwnAwemeFans ToolsEstimateAudienceV2FilterOwnAwemeFans = 1
)

// All allowed values of ToolsEstimateAudienceV2FilterOwnAwemeFans enum
var AllowedToolsEstimateAudienceV2FilterOwnAwemeFansEnumValues = []ToolsEstimateAudienceV2FilterOwnAwemeFans{
	0,
	1,
}

// NewToolsEstimateAudienceV2FilterOwnAwemeFansFromValue returns a pointer to a valid ToolsEstimateAudienceV2FilterOwnAwemeFans
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2FilterOwnAwemeFansFromValue(v int64) (*ToolsEstimateAudienceV2FilterOwnAwemeFans, error) {
	ev := ToolsEstimateAudienceV2FilterOwnAwemeFans(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2FilterOwnAwemeFans: valid values are %v", v, AllowedToolsEstimateAudienceV2FilterOwnAwemeFansEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2FilterOwnAwemeFans) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2FilterOwnAwemeFansEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_filter_own_aweme_fans value
func (v ToolsEstimateAudienceV2FilterOwnAwemeFans) Ptr() *ToolsEstimateAudienceV2FilterOwnAwemeFans {
	return &v
}