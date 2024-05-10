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

// ToolsPlayableSaveV2DataPlayableOrientation
type ToolsPlayableSaveV2DataPlayableOrientation string

// List of tools_playable_save_v2_data_playable_orientation
const (
	BOTH_ToolsPlayableSaveV2DataPlayableOrientation      ToolsPlayableSaveV2DataPlayableOrientation = "BOTH"
	LANDSCAPE_ToolsPlayableSaveV2DataPlayableOrientation ToolsPlayableSaveV2DataPlayableOrientation = "LANDSCAPE"
	PORTRAIT_ToolsPlayableSaveV2DataPlayableOrientation  ToolsPlayableSaveV2DataPlayableOrientation = "PORTRAIT"
)

// All allowed values of ToolsPlayableSaveV2DataPlayableOrientation enum
var AllowedToolsPlayableSaveV2DataPlayableOrientationEnumValues = []ToolsPlayableSaveV2DataPlayableOrientation{
	"BOTH",
	"LANDSCAPE",
	"PORTRAIT",
}

// NewToolsPlayableSaveV2DataPlayableOrientationFromValue returns a pointer to a valid ToolsPlayableSaveV2DataPlayableOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableSaveV2DataPlayableOrientationFromValue(v string) (*ToolsPlayableSaveV2DataPlayableOrientation, error) {
	ev := ToolsPlayableSaveV2DataPlayableOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableSaveV2DataPlayableOrientation: valid values are %v", v, AllowedToolsPlayableSaveV2DataPlayableOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableSaveV2DataPlayableOrientation) IsValid() bool {
	for _, existing := range AllowedToolsPlayableSaveV2DataPlayableOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_save_v2_data_playable_orientation value
func (v ToolsPlayableSaveV2DataPlayableOrientation) Ptr() *ToolsPlayableSaveV2DataPlayableOrientation {
	return &v
}
