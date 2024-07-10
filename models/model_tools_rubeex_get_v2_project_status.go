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

// ToolsRubeexGetV2ProjectStatus
type ToolsRubeexGetV2ProjectStatus string

// List of tools_rubeex_get_v2_project_status
const (
	DELETED_ToolsRubeexGetV2ProjectStatus ToolsRubeexGetV2ProjectStatus = "DELETED"
	ENABLED_ToolsRubeexGetV2ProjectStatus ToolsRubeexGetV2ProjectStatus = "ENABLED"
	INITAL_ToolsRubeexGetV2ProjectStatus  ToolsRubeexGetV2ProjectStatus = "INITAL"
)

// All allowed values of ToolsRubeexGetV2ProjectStatus enum
var AllowedToolsRubeexGetV2ProjectStatusEnumValues = []ToolsRubeexGetV2ProjectStatus{
	"DELETED",
	"ENABLED",
	"INITAL",
}

// NewToolsRubeexGetV2ProjectStatusFromValue returns a pointer to a valid ToolsRubeexGetV2ProjectStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsRubeexGetV2ProjectStatusFromValue(v string) (*ToolsRubeexGetV2ProjectStatus, error) {
	ev := ToolsRubeexGetV2ProjectStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsRubeexGetV2ProjectStatus: valid values are %v", v, AllowedToolsRubeexGetV2ProjectStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsRubeexGetV2ProjectStatus) IsValid() bool {
	for _, existing := range AllowedToolsRubeexGetV2ProjectStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_rubeex_get_v2_project_status value
func (v ToolsRubeexGetV2ProjectStatus) Ptr() *ToolsRubeexGetV2ProjectStatus {
	return &v
}
