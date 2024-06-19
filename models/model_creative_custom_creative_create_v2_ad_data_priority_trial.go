/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeCreateV2AdDataPriorityTrial
type CreativeCustomCreativeCreateV2AdDataPriorityTrial string

// List of creative_custom_creative_create_v2_ad_data_priority_trial
const (
	OFF_CreativeCustomCreativeCreateV2AdDataPriorityTrial CreativeCustomCreativeCreateV2AdDataPriorityTrial = "OFF"
	ON_CreativeCustomCreativeCreateV2AdDataPriorityTrial  CreativeCustomCreativeCreateV2AdDataPriorityTrial = "ON"
)

// All allowed values of CreativeCustomCreativeCreateV2AdDataPriorityTrial enum
var AllowedCreativeCustomCreativeCreateV2AdDataPriorityTrialEnumValues = []CreativeCustomCreativeCreateV2AdDataPriorityTrial{
	"OFF",
	"ON",
}

// NewCreativeCustomCreativeCreateV2AdDataPriorityTrialFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2AdDataPriorityTrial
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2AdDataPriorityTrialFromValue(v string) (*CreativeCustomCreativeCreateV2AdDataPriorityTrial, error) {
	ev := CreativeCustomCreativeCreateV2AdDataPriorityTrial(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2AdDataPriorityTrial: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2AdDataPriorityTrialEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2AdDataPriorityTrial) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2AdDataPriorityTrialEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_ad_data_priority_trial value
func (v CreativeCustomCreativeCreateV2AdDataPriorityTrial) Ptr() *CreativeCustomCreativeCreateV2AdDataPriorityTrial {
	return &v
}
