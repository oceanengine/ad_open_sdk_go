/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch
type CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch string

// List of creative_procedural_creative_create_v2_ad_data_dynamic_creative_switch
const (
	DYNAMIC_CREATIVE_TITLE_CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch    CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_TITLE"
	DYNAMIC_CREATIVE_ON_CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch       CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_ON"
	DYNAMIC_CREATIVE_ABSTRACT_CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_ABSTRACT"
	DYNAMIC_CREATIVE_SUBLINK_CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch  CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_SUBLINK"
)

// All allowed values of CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch enum
var AllowedCreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitchEnumValues = []CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch{
	"DYNAMIC_CREATIVE_TITLE",
	"DYNAMIC_CREATIVE_ON",
	"DYNAMIC_CREATIVE_ABSTRACT",
	"DYNAMIC_CREATIVE_SUBLINK",
}

// NewCreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitchFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitchFromValue(v string) (*CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch, error) {
	ev := CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_ad_data_dynamic_creative_switch value
func (v CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch) Ptr() *CreativeProceduralCreativeCreateV2AdDataDynamicCreativeSwitch {
	return &v
}
