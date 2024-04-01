/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppIosListV2FilteringStatus
type ToolsAppIosListV2FilteringStatus string

// List of tools_app_ios_list_v2_filtering_status
const (
	AUDIT_REJECTED_ToolsAppIosListV2FilteringStatus ToolsAppIosListV2FilteringStatus = "AUDIT_REJECTED"
	ALL_ToolsAppIosListV2FilteringStatus            ToolsAppIosListV2FilteringStatus = "ALL"
	AUDIT_ACCEPTED_ToolsAppIosListV2FilteringStatus ToolsAppIosListV2FilteringStatus = "AUDIT_ACCEPTED"
	ENABLE_ToolsAppIosListV2FilteringStatus         ToolsAppIosListV2FilteringStatus = "ENABLE"
	AUDIT_DOING_ToolsAppIosListV2FilteringStatus    ToolsAppIosListV2FilteringStatus = "AUDIT_DOING"
)

// All allowed values of ToolsAppIosListV2FilteringStatus enum
var AllowedToolsAppIosListV2FilteringStatusEnumValues = []ToolsAppIosListV2FilteringStatus{
	"AUDIT_REJECTED",
	"ALL",
	"AUDIT_ACCEPTED",
	"ENABLE",
	"AUDIT_DOING",
}

// NewToolsAppIosListV2FilteringStatusFromValue returns a pointer to a valid ToolsAppIosListV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppIosListV2FilteringStatusFromValue(v string) (*ToolsAppIosListV2FilteringStatus, error) {
	ev := ToolsAppIosListV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppIosListV2FilteringStatus: valid values are %v", v, AllowedToolsAppIosListV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppIosListV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedToolsAppIosListV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_ios_list_v2_filtering_status value
func (v ToolsAppIosListV2FilteringStatus) Ptr() *ToolsAppIosListV2FilteringStatus {
	return &v
}
