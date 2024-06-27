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

// ToolsAbTestListGetV2DataTestListStatus
type ToolsAbTestListGetV2DataTestListStatus string

// List of tools_ab_test_list_get_v2_data_test_list_status
const (
	CREATED_ToolsAbTestListGetV2DataTestListStatus    ToolsAbTestListGetV2DataTestListStatus = "CREATED"
	FAILED_ToolsAbTestListGetV2DataTestListStatus     ToolsAbTestListGetV2DataTestListStatus = "FAILED"
	FINISH_ToolsAbTestListGetV2DataTestListStatus     ToolsAbTestListGetV2DataTestListStatus = "FINISH"
	PROCESSING_ToolsAbTestListGetV2DataTestListStatus ToolsAbTestListGetV2DataTestListStatus = "PROCESSING"
)

// All allowed values of ToolsAbTestListGetV2DataTestListStatus enum
var AllowedToolsAbTestListGetV2DataTestListStatusEnumValues = []ToolsAbTestListGetV2DataTestListStatus{
	"CREATED",
	"FAILED",
	"FINISH",
	"PROCESSING",
}

// NewToolsAbTestListGetV2DataTestListStatusFromValue returns a pointer to a valid ToolsAbTestListGetV2DataTestListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAbTestListGetV2DataTestListStatusFromValue(v string) (*ToolsAbTestListGetV2DataTestListStatus, error) {
	ev := ToolsAbTestListGetV2DataTestListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAbTestListGetV2DataTestListStatus: valid values are %v", v, AllowedToolsAbTestListGetV2DataTestListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAbTestListGetV2DataTestListStatus) IsValid() bool {
	for _, existing := range AllowedToolsAbTestListGetV2DataTestListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ab_test_list_get_v2_data_test_list_status value
func (v ToolsAbTestListGetV2DataTestListStatus) Ptr() *ToolsAbTestListGetV2DataTestListStatus {
	return &v
}
