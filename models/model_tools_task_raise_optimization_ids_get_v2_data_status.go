/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsTaskRaiseOptimizationIdsGetV2DataStatus
type ToolsTaskRaiseOptimizationIdsGetV2DataStatus string

// List of tools_task_raise_optimization_ids_get_v2_data_status
const (
	DISABLERAISE_ToolsTaskRaiseOptimizationIdsGetV2DataStatus ToolsTaskRaiseOptimizationIdsGetV2DataStatus = "DISABLERAISE"
	ENABLERAISE_ToolsTaskRaiseOptimizationIdsGetV2DataStatus  ToolsTaskRaiseOptimizationIdsGetV2DataStatus = "ENABLERAISE"
	RAISING_ToolsTaskRaiseOptimizationIdsGetV2DataStatus      ToolsTaskRaiseOptimizationIdsGetV2DataStatus = "RAISING"
)

// All allowed values of ToolsTaskRaiseOptimizationIdsGetV2DataStatus enum
var AllowedToolsTaskRaiseOptimizationIdsGetV2DataStatusEnumValues = []ToolsTaskRaiseOptimizationIdsGetV2DataStatus{
	"DISABLERAISE",
	"ENABLERAISE",
	"RAISING",
}

// NewToolsTaskRaiseOptimizationIdsGetV2DataStatusFromValue returns a pointer to a valid ToolsTaskRaiseOptimizationIdsGetV2DataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsTaskRaiseOptimizationIdsGetV2DataStatusFromValue(v string) (*ToolsTaskRaiseOptimizationIdsGetV2DataStatus, error) {
	ev := ToolsTaskRaiseOptimizationIdsGetV2DataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsTaskRaiseOptimizationIdsGetV2DataStatus: valid values are %v", v, AllowedToolsTaskRaiseOptimizationIdsGetV2DataStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsTaskRaiseOptimizationIdsGetV2DataStatus) IsValid() bool {
	for _, existing := range AllowedToolsTaskRaiseOptimizationIdsGetV2DataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_task_raise_optimization_ids_get_v2_data_status value
func (v ToolsTaskRaiseOptimizationIdsGetV2DataStatus) Ptr() *ToolsTaskRaiseOptimizationIdsGetV2DataStatus {
	return &v
}
