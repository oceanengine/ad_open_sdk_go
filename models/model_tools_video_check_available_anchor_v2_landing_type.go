/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsVideoCheckAvailableAnchorV2LandingType
type ToolsVideoCheckAvailableAnchorV2LandingType string

// List of tools_video_check_available_anchor_v2_landing_type
const (
	APP_ToolsVideoCheckAvailableAnchorV2LandingType ToolsVideoCheckAvailableAnchorV2LandingType = "APP"
)

// All allowed values of ToolsVideoCheckAvailableAnchorV2LandingType enum
var AllowedToolsVideoCheckAvailableAnchorV2LandingTypeEnumValues = []ToolsVideoCheckAvailableAnchorV2LandingType{
	"APP",
}

// NewToolsVideoCheckAvailableAnchorV2LandingTypeFromValue returns a pointer to a valid ToolsVideoCheckAvailableAnchorV2LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsVideoCheckAvailableAnchorV2LandingTypeFromValue(v string) (*ToolsVideoCheckAvailableAnchorV2LandingType, error) {
	ev := ToolsVideoCheckAvailableAnchorV2LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsVideoCheckAvailableAnchorV2LandingType: valid values are %v", v, AllowedToolsVideoCheckAvailableAnchorV2LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsVideoCheckAvailableAnchorV2LandingType) IsValid() bool {
	for _, existing := range AllowedToolsVideoCheckAvailableAnchorV2LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_video_check_available_anchor_v2_landing_type value
func (v ToolsVideoCheckAvailableAnchorV2LandingType) Ptr() *ToolsVideoCheckAvailableAnchorV2LandingType {
	return &v
}
