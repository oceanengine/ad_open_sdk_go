/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2AutoExtendTargets
type ToolsEstimateAudienceV2AutoExtendTargets string

// List of tools_estimate_audience_v2_auto_extend_targets
const (
	INTEREST_TAG_ToolsEstimateAudienceV2AutoExtendTargets    ToolsEstimateAudienceV2AutoExtendTargets = "INTEREST_TAG"
	CUSTOM_AUDIENCE_ToolsEstimateAudienceV2AutoExtendTargets ToolsEstimateAudienceV2AutoExtendTargets = "CUSTOM_AUDIENCE"
	REGION_ToolsEstimateAudienceV2AutoExtendTargets          ToolsEstimateAudienceV2AutoExtendTargets = "REGION"
	AD_TAG_ToolsEstimateAudienceV2AutoExtendTargets          ToolsEstimateAudienceV2AutoExtendTargets = "AD_TAG"
	AGE_ToolsEstimateAudienceV2AutoExtendTargets             ToolsEstimateAudienceV2AutoExtendTargets = "AGE"
	GENDER_ToolsEstimateAudienceV2AutoExtendTargets          ToolsEstimateAudienceV2AutoExtendTargets = "GENDER"
)

// All allowed values of ToolsEstimateAudienceV2AutoExtendTargets enum
var AllowedToolsEstimateAudienceV2AutoExtendTargetsEnumValues = []ToolsEstimateAudienceV2AutoExtendTargets{
	"INTEREST_TAG",
	"CUSTOM_AUDIENCE",
	"REGION",
	"AD_TAG",
	"AGE",
	"GENDER",
}

// NewToolsEstimateAudienceV2AutoExtendTargetsFromValue returns a pointer to a valid ToolsEstimateAudienceV2AutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2AutoExtendTargetsFromValue(v string) (*ToolsEstimateAudienceV2AutoExtendTargets, error) {
	ev := ToolsEstimateAudienceV2AutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2AutoExtendTargets: valid values are %v", v, AllowedToolsEstimateAudienceV2AutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2AutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2AutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_auto_extend_targets value
func (v ToolsEstimateAudienceV2AutoExtendTargets) Ptr() *ToolsEstimateAudienceV2AutoExtendTargets {
	return &v
}
