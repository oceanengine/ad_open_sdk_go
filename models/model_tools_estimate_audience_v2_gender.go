/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2Gender
type ToolsEstimateAudienceV2Gender string

// List of tools_estimate_audience_v2_gender
const (
	GENDER_FEMALE_ToolsEstimateAudienceV2Gender    ToolsEstimateAudienceV2Gender = "GENDER_FEMALE"
	NONE_ToolsEstimateAudienceV2Gender             ToolsEstimateAudienceV2Gender = "NONE"
	GENDER_MALE_ToolsEstimateAudienceV2Gender      ToolsEstimateAudienceV2Gender = "GENDER_MALE"
	GENDER_UNLIMITED_ToolsEstimateAudienceV2Gender ToolsEstimateAudienceV2Gender = "GENDER_UNLIMITED"
)

// All allowed values of ToolsEstimateAudienceV2Gender enum
var AllowedToolsEstimateAudienceV2GenderEnumValues = []ToolsEstimateAudienceV2Gender{
	"GENDER_FEMALE",
	"NONE",
	"GENDER_MALE",
	"GENDER_UNLIMITED",
}

// NewToolsEstimateAudienceV2GenderFromValue returns a pointer to a valid ToolsEstimateAudienceV2Gender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2GenderFromValue(v string) (*ToolsEstimateAudienceV2Gender, error) {
	ev := ToolsEstimateAudienceV2Gender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2Gender: valid values are %v", v, AllowedToolsEstimateAudienceV2GenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2Gender) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2GenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_gender value
func (v ToolsEstimateAudienceV2Gender) Ptr() *ToolsEstimateAudienceV2Gender {
	return &v
}
