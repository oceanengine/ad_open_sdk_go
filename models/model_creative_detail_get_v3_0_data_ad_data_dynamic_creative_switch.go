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

// CreativeDetailGetV30DataAdDataDynamicCreativeSwitch
type CreativeDetailGetV30DataAdDataDynamicCreativeSwitch string

// List of creative_detail_get_v3.0_data_ad_data_dynamic_creative_switch
const (
	DYNAMIC_CREATIVE_ABSTRACT_CreativeDetailGetV30DataAdDataDynamicCreativeSwitch CreativeDetailGetV30DataAdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_ABSTRACT"
	DYNAMIC_CREATIVE_ON_CreativeDetailGetV30DataAdDataDynamicCreativeSwitch       CreativeDetailGetV30DataAdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_ON"
	DYNAMIC_CREATIVE_SUBLINK_CreativeDetailGetV30DataAdDataDynamicCreativeSwitch  CreativeDetailGetV30DataAdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_SUBLINK"
	DYNAMIC_CREATIVE_TITLE_CreativeDetailGetV30DataAdDataDynamicCreativeSwitch    CreativeDetailGetV30DataAdDataDynamicCreativeSwitch = "DYNAMIC_CREATIVE_TITLE"
)

// All allowed values of CreativeDetailGetV30DataAdDataDynamicCreativeSwitch enum
var AllowedCreativeDetailGetV30DataAdDataDynamicCreativeSwitchEnumValues = []CreativeDetailGetV30DataAdDataDynamicCreativeSwitch{
	"DYNAMIC_CREATIVE_ABSTRACT",
	"DYNAMIC_CREATIVE_ON",
	"DYNAMIC_CREATIVE_SUBLINK",
	"DYNAMIC_CREATIVE_TITLE",
}

// NewCreativeDetailGetV30DataAdDataDynamicCreativeSwitchFromValue returns a pointer to a valid CreativeDetailGetV30DataAdDataDynamicCreativeSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataAdDataDynamicCreativeSwitchFromValue(v string) (*CreativeDetailGetV30DataAdDataDynamicCreativeSwitch, error) {
	ev := CreativeDetailGetV30DataAdDataDynamicCreativeSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataAdDataDynamicCreativeSwitch: valid values are %v", v, AllowedCreativeDetailGetV30DataAdDataDynamicCreativeSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataAdDataDynamicCreativeSwitch) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataAdDataDynamicCreativeSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_ad_data_dynamic_creative_switch value
func (v CreativeDetailGetV30DataAdDataDynamicCreativeSwitch) Ptr() *CreativeDetailGetV30DataAdDataDynamicCreativeSwitch {
	return &v
}
