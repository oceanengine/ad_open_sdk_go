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

// CreativeProceduralCreativeUpdateV2AdDataPriorityTrial
type CreativeProceduralCreativeUpdateV2AdDataPriorityTrial string

// List of creative_procedural_creative_update_v2_ad_data_priority_trial
const (
	OFF_CreativeProceduralCreativeUpdateV2AdDataPriorityTrial CreativeProceduralCreativeUpdateV2AdDataPriorityTrial = "OFF"
	ON_CreativeProceduralCreativeUpdateV2AdDataPriorityTrial  CreativeProceduralCreativeUpdateV2AdDataPriorityTrial = "ON"
)

// All allowed values of CreativeProceduralCreativeUpdateV2AdDataPriorityTrial enum
var AllowedCreativeProceduralCreativeUpdateV2AdDataPriorityTrialEnumValues = []CreativeProceduralCreativeUpdateV2AdDataPriorityTrial{
	"OFF",
	"ON",
}

// NewCreativeProceduralCreativeUpdateV2AdDataPriorityTrialFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2AdDataPriorityTrial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2AdDataPriorityTrialFromValue(v string) (*CreativeProceduralCreativeUpdateV2AdDataPriorityTrial, error) {
	ev := CreativeProceduralCreativeUpdateV2AdDataPriorityTrial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2AdDataPriorityTrial: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2AdDataPriorityTrialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2AdDataPriorityTrial) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2AdDataPriorityTrialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_ad_data_priority_trial value
func (v CreativeProceduralCreativeUpdateV2AdDataPriorityTrial) Ptr() *CreativeProceduralCreativeUpdateV2AdDataPriorityTrial {
	return &v
}
