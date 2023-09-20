/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdRaiseResultGetV2V2TimeDimension
type ToolsAdRaiseResultGetV2V2TimeDimension string

// List of tools_ad_raise_result_get_v2_v2_time_dimension
const (
	SUM_ToolsAdRaiseResultGetV2V2TimeDimension    ToolsAdRaiseResultGetV2V2TimeDimension = "SUM"
	HOURLY_ToolsAdRaiseResultGetV2V2TimeDimension ToolsAdRaiseResultGetV2V2TimeDimension = "HOURLY"
)

// All allowed values of ToolsAdRaiseResultGetV2V2TimeDimension enum
var AllowedToolsAdRaiseResultGetV2V2TimeDimensionEnumValues = []ToolsAdRaiseResultGetV2V2TimeDimension{
	"SUM",
	"HOURLY",
}

// NewToolsAdRaiseResultGetV2V2TimeDimensionFromValue returns a pointer to a valid ToolsAdRaiseResultGetV2V2TimeDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdRaiseResultGetV2V2TimeDimensionFromValue(v string) (*ToolsAdRaiseResultGetV2V2TimeDimension, error) {
	ev := ToolsAdRaiseResultGetV2V2TimeDimension(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdRaiseResultGetV2V2TimeDimension: valid values are %v", v, AllowedToolsAdRaiseResultGetV2V2TimeDimensionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdRaiseResultGetV2V2TimeDimension) IsValid() bool {
	for _, existing := range AllowedToolsAdRaiseResultGetV2V2TimeDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_raise_result_get_v2_v2_time_dimension value
func (v ToolsAdRaiseResultGetV2V2TimeDimension) Ptr() *ToolsAdRaiseResultGetV2V2TimeDimension {
	return &v
}