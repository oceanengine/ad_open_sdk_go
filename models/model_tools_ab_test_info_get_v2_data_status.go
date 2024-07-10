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

// ToolsAbTestInfoGetV2DataStatus
type ToolsAbTestInfoGetV2DataStatus string

// List of tools_ab_test_info_get_v2_data_status
const (
	CREATED_ToolsAbTestInfoGetV2DataStatus    ToolsAbTestInfoGetV2DataStatus = "CREATED"
	FAILED_ToolsAbTestInfoGetV2DataStatus     ToolsAbTestInfoGetV2DataStatus = "FAILED"
	FINISH_ToolsAbTestInfoGetV2DataStatus     ToolsAbTestInfoGetV2DataStatus = "FINISH"
	PROCESSING_ToolsAbTestInfoGetV2DataStatus ToolsAbTestInfoGetV2DataStatus = "PROCESSING"
)

// All allowed values of ToolsAbTestInfoGetV2DataStatus enum
var AllowedToolsAbTestInfoGetV2DataStatusEnumValues = []ToolsAbTestInfoGetV2DataStatus{
	"CREATED",
	"FAILED",
	"FINISH",
	"PROCESSING",
}

// NewToolsAbTestInfoGetV2DataStatusFromValue returns a pointer to a valid ToolsAbTestInfoGetV2DataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAbTestInfoGetV2DataStatusFromValue(v string) (*ToolsAbTestInfoGetV2DataStatus, error) {
	ev := ToolsAbTestInfoGetV2DataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAbTestInfoGetV2DataStatus: valid values are %v", v, AllowedToolsAbTestInfoGetV2DataStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAbTestInfoGetV2DataStatus) IsValid() bool {
	for _, existing := range AllowedToolsAbTestInfoGetV2DataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ab_test_info_get_v2_data_status value
func (v ToolsAbTestInfoGetV2DataStatus) Ptr() *ToolsAbTestInfoGetV2DataStatus {
	return &v
}
