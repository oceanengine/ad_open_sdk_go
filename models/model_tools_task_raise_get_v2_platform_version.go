/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsTaskRaiseGetV2PlatformVersion
type ToolsTaskRaiseGetV2PlatformVersion string

// List of tools_task_raise_get_v2_platform_version
const (
	V2_ToolsTaskRaiseGetV2PlatformVersion ToolsTaskRaiseGetV2PlatformVersion = "V2"
)

// All allowed values of ToolsTaskRaiseGetV2PlatformVersion enum
var AllowedToolsTaskRaiseGetV2PlatformVersionEnumValues = []ToolsTaskRaiseGetV2PlatformVersion{
	"V2",
}

// NewToolsTaskRaiseGetV2PlatformVersionFromValue returns a pointer to a valid ToolsTaskRaiseGetV2PlatformVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsTaskRaiseGetV2PlatformVersionFromValue(v string) (*ToolsTaskRaiseGetV2PlatformVersion, error) {
	ev := ToolsTaskRaiseGetV2PlatformVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsTaskRaiseGetV2PlatformVersion: valid values are %v", v, AllowedToolsTaskRaiseGetV2PlatformVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsTaskRaiseGetV2PlatformVersion) IsValid() bool {
	for _, existing := range AllowedToolsTaskRaiseGetV2PlatformVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_task_raise_get_v2_platform_version value
func (v ToolsTaskRaiseGetV2PlatformVersion) Ptr() *ToolsTaskRaiseGetV2PlatformVersion {
	return &v
}
