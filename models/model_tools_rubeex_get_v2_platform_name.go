/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsRubeexGetV2PlatformName
type ToolsRubeexGetV2PlatformName string

// List of tools_rubeex_get_v2_platform_name
const (
	RUBIX_ToolsRubeexGetV2PlatformName             ToolsRubeexGetV2PlatformName = "RUBIX"
	CLAB_EDITOR_ToolsRubeexGetV2PlatformName       ToolsRubeexGetV2PlatformName = "CLAB_EDITOR"
	POWERPLAYABLE_ToolsRubeexGetV2PlatformName     ToolsRubeexGetV2PlatformName = "POWERPLAYABLE"
	PLAYABLE_TEMPLATE_ToolsRubeexGetV2PlatformName ToolsRubeexGetV2PlatformName = "PLAYABLE_TEMPLATE"
	LIGHT_PLAYABLE_ToolsRubeexGetV2PlatformName    ToolsRubeexGetV2PlatformName = "LIGHT_PLAYABLE"
	AD_PLATFORM_ToolsRubeexGetV2PlatformName       ToolsRubeexGetV2PlatformName = "AD_PLATFORM"
)

// All allowed values of ToolsRubeexGetV2PlatformName enum
var AllowedToolsRubeexGetV2PlatformNameEnumValues = []ToolsRubeexGetV2PlatformName{
	"RUBIX",
	"CLAB_EDITOR",
	"POWERPLAYABLE",
	"PLAYABLE_TEMPLATE",
	"LIGHT_PLAYABLE",
	"AD_PLATFORM",
}

// NewToolsRubeexGetV2PlatformNameFromValue returns a pointer to a valid ToolsRubeexGetV2PlatformName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRubeexGetV2PlatformNameFromValue(v string) (*ToolsRubeexGetV2PlatformName, error) {
	ev := ToolsRubeexGetV2PlatformName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRubeexGetV2PlatformName: valid values are %v", v, AllowedToolsRubeexGetV2PlatformNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRubeexGetV2PlatformName) IsValid() bool {
	for _, existing := range AllowedToolsRubeexGetV2PlatformNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rubeex_get_v2_platform_name value
func (v ToolsRubeexGetV2PlatformName) Ptr() *ToolsRubeexGetV2PlatformName {
	return &v
}
