/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsTaskRaiseCreateV2RaiseMode
type ToolsTaskRaiseCreateV2RaiseMode string

// List of tools_task_raise_create_v2_raise_mode
const (
	CUSTOM_ToolsTaskRaiseCreateV2RaiseMode ToolsTaskRaiseCreateV2RaiseMode = "CUSTOM"
	STRONG_ToolsTaskRaiseCreateV2RaiseMode ToolsTaskRaiseCreateV2RaiseMode = "STRONG"
)

// All allowed values of ToolsTaskRaiseCreateV2RaiseMode enum
var AllowedToolsTaskRaiseCreateV2RaiseModeEnumValues = []ToolsTaskRaiseCreateV2RaiseMode{
	"CUSTOM",
	"STRONG",
}

// NewToolsTaskRaiseCreateV2RaiseModeFromValue returns a pointer to a valid ToolsTaskRaiseCreateV2RaiseMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsTaskRaiseCreateV2RaiseModeFromValue(v string) (*ToolsTaskRaiseCreateV2RaiseMode, error) {
	ev := ToolsTaskRaiseCreateV2RaiseMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsTaskRaiseCreateV2RaiseMode: valid values are %v", v, AllowedToolsTaskRaiseCreateV2RaiseModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsTaskRaiseCreateV2RaiseMode) IsValid() bool {
	for _, existing := range AllowedToolsTaskRaiseCreateV2RaiseModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_task_raise_create_v2_raise_mode value
func (v ToolsTaskRaiseCreateV2RaiseMode) Ptr() *ToolsTaskRaiseCreateV2RaiseMode {
	return &v
}
