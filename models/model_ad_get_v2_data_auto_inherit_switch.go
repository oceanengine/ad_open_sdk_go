/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAutoInheritSwitch
type AdGetV2DataAutoInheritSwitch string

// List of ad_get_v2_data_auto_inherit_switch
const (
	ON_AdGetV2DataAutoInheritSwitch  AdGetV2DataAutoInheritSwitch = "ON"
	OFF_AdGetV2DataAutoInheritSwitch AdGetV2DataAutoInheritSwitch = "OFF"
)

// All allowed values of AdGetV2DataAutoInheritSwitch enum
var AllowedAdGetV2DataAutoInheritSwitchEnumValues = []AdGetV2DataAutoInheritSwitch{
	"ON",
	"OFF",
}

// NewAdGetV2DataAutoInheritSwitchFromValue returns a pointer to a valid AdGetV2DataAutoInheritSwitch
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAutoInheritSwitchFromValue(v string) (*AdGetV2DataAutoInheritSwitch, error) {
	ev := AdGetV2DataAutoInheritSwitch(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAutoInheritSwitch: valid values are %v", v, AllowedAdGetV2DataAutoInheritSwitchEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAutoInheritSwitch) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAutoInheritSwitchEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_auto_inherit_switch value
func (v AdGetV2DataAutoInheritSwitch) Ptr() *AdGetV2DataAutoInheritSwitch {
	return &v
}
