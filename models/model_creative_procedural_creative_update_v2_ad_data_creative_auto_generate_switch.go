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

// CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch
type CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch int64

// List of creative_procedural_creative_update_v2_ad_data_creative_auto_generate_switch
const (
	Enum_0_CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch = 0
	Enum_1_CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch = 1
)

// All allowed values of CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch enum
var AllowedCreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitchEnumValues = []CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch{
	0,
	1,
}

// NewCreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitchFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitchFromValue(v int64) (*CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch, error) {
	ev := CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_ad_data_creative_auto_generate_switch value
func (v CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch) Ptr() *CreativeProceduralCreativeUpdateV2AdDataCreativeAutoGenerateSwitch {
	return &v
}
