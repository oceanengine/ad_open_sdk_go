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

// CreativeProceduralCreativeCreateV2AdDataPriorityTrial
type CreativeProceduralCreativeCreateV2AdDataPriorityTrial string

// List of creative_procedural_creative_create_v2_ad_data_priority_trial
const (
	OFF_CreativeProceduralCreativeCreateV2AdDataPriorityTrial CreativeProceduralCreativeCreateV2AdDataPriorityTrial = "OFF"
	ON_CreativeProceduralCreativeCreateV2AdDataPriorityTrial  CreativeProceduralCreativeCreateV2AdDataPriorityTrial = "ON"
)

// All allowed values of CreativeProceduralCreativeCreateV2AdDataPriorityTrial enum
var AllowedCreativeProceduralCreativeCreateV2AdDataPriorityTrialEnumValues = []CreativeProceduralCreativeCreateV2AdDataPriorityTrial{
	"OFF",
	"ON",
}

// NewCreativeProceduralCreativeCreateV2AdDataPriorityTrialFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2AdDataPriorityTrial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2AdDataPriorityTrialFromValue(v string) (*CreativeProceduralCreativeCreateV2AdDataPriorityTrial, error) {
	ev := CreativeProceduralCreativeCreateV2AdDataPriorityTrial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2AdDataPriorityTrial: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2AdDataPriorityTrialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2AdDataPriorityTrial) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2AdDataPriorityTrialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_ad_data_priority_trial value
func (v CreativeProceduralCreativeCreateV2AdDataPriorityTrial) Ptr() *CreativeProceduralCreativeCreateV2AdDataPriorityTrial {
	return &v
}
