/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsTaskRaiseGetV2DataReportsRaiseMode
type ToolsTaskRaiseGetV2DataReportsRaiseMode string

// List of tools_task_raise_get_v2_data_reports_raise_mode
const (
	CUSTOM_ToolsTaskRaiseGetV2DataReportsRaiseMode ToolsTaskRaiseGetV2DataReportsRaiseMode = "CUSTOM"
	STRONG_ToolsTaskRaiseGetV2DataReportsRaiseMode ToolsTaskRaiseGetV2DataReportsRaiseMode = "STRONG"
)

// All allowed values of ToolsTaskRaiseGetV2DataReportsRaiseMode enum
var AllowedToolsTaskRaiseGetV2DataReportsRaiseModeEnumValues = []ToolsTaskRaiseGetV2DataReportsRaiseMode{
	"CUSTOM",
	"STRONG",
}

// NewToolsTaskRaiseGetV2DataReportsRaiseModeFromValue returns a pointer to a valid ToolsTaskRaiseGetV2DataReportsRaiseMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsTaskRaiseGetV2DataReportsRaiseModeFromValue(v string) (*ToolsTaskRaiseGetV2DataReportsRaiseMode, error) {
	ev := ToolsTaskRaiseGetV2DataReportsRaiseMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsTaskRaiseGetV2DataReportsRaiseMode: valid values are %v", v, AllowedToolsTaskRaiseGetV2DataReportsRaiseModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsTaskRaiseGetV2DataReportsRaiseMode) IsValid() bool {
	for _, existing := range AllowedToolsTaskRaiseGetV2DataReportsRaiseModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_task_raise_get_v2_data_reports_raise_mode value
func (v ToolsTaskRaiseGetV2DataReportsRaiseMode) Ptr() *ToolsTaskRaiseGetV2DataReportsRaiseMode {
	return &v
}
