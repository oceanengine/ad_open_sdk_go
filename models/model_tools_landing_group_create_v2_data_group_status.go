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

// ToolsLandingGroupCreateV2DataGroupStatus
type ToolsLandingGroupCreateV2DataGroupStatus string

// List of tools_landing_group_create_v2_data_group_status
const (
	LANDING_GROUP_STATUS_AVAILABLE_ToolsLandingGroupCreateV2DataGroupStatus   ToolsLandingGroupCreateV2DataGroupStatus = "LANDING_GROUP_STATUS_AVAILABLE"
	LANDING_GROUP_STATUS_UNAVAILABLE_ToolsLandingGroupCreateV2DataGroupStatus ToolsLandingGroupCreateV2DataGroupStatus = "LANDING_GROUP_STATUS_UNAVAILABLE"
)

// All allowed values of ToolsLandingGroupCreateV2DataGroupStatus enum
var AllowedToolsLandingGroupCreateV2DataGroupStatusEnumValues = []ToolsLandingGroupCreateV2DataGroupStatus{
	"LANDING_GROUP_STATUS_AVAILABLE",
	"LANDING_GROUP_STATUS_UNAVAILABLE",
}

// NewToolsLandingGroupCreateV2DataGroupStatusFromValue returns a pointer to a valid ToolsLandingGroupCreateV2DataGroupStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsLandingGroupCreateV2DataGroupStatusFromValue(v string) (*ToolsLandingGroupCreateV2DataGroupStatus, error) {
	ev := ToolsLandingGroupCreateV2DataGroupStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsLandingGroupCreateV2DataGroupStatus: valid values are %v", v, AllowedToolsLandingGroupCreateV2DataGroupStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsLandingGroupCreateV2DataGroupStatus) IsValid() bool {
	for _, existing := range AllowedToolsLandingGroupCreateV2DataGroupStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_landing_group_create_v2_data_group_status value
func (v ToolsLandingGroupCreateV2DataGroupStatus) Ptr() *ToolsLandingGroupCreateV2DataGroupStatus {
	return &v
}
