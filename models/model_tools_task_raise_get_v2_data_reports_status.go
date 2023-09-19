/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsTaskRaiseGetV2DataReportsStatus
type ToolsTaskRaiseGetV2DataReportsStatus string

// List of tools_task_raise_get_v2_data_reports_status
const (
	RAISING_ToolsTaskRaiseGetV2DataReportsStatus ToolsTaskRaiseGetV2DataReportsStatus = "RAISING"
	STOP_ToolsTaskRaiseGetV2DataReportsStatus    ToolsTaskRaiseGetV2DataReportsStatus = "STOP"
)

// All allowed values of ToolsTaskRaiseGetV2DataReportsStatus enum
var AllowedToolsTaskRaiseGetV2DataReportsStatusEnumValues = []ToolsTaskRaiseGetV2DataReportsStatus{
	"RAISING",
	"STOP",
}

// NewToolsTaskRaiseGetV2DataReportsStatusFromValue returns a pointer to a valid ToolsTaskRaiseGetV2DataReportsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsTaskRaiseGetV2DataReportsStatusFromValue(v string) (*ToolsTaskRaiseGetV2DataReportsStatus, error) {
	ev := ToolsTaskRaiseGetV2DataReportsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsTaskRaiseGetV2DataReportsStatus: valid values are %v", v, AllowedToolsTaskRaiseGetV2DataReportsStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsTaskRaiseGetV2DataReportsStatus) IsValid() bool {
	for _, existing := range AllowedToolsTaskRaiseGetV2DataReportsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_task_raise_get_v2_data_reports_status value
func (v ToolsTaskRaiseGetV2DataReportsStatus) Ptr() *ToolsTaskRaiseGetV2DataReportsStatus {
	return &v
}
