/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus
type ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus string

// List of tool_quick_app_management_quick_app_get_v2_data_quick_app_info_status
const (
	AUDIT_ACCEPTED_ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus      ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus = "AUDIT_ACCEPTED"
	AUDIT_DOING_ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus         ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus = "AUDIT_DOING"
	AUDIT_REJECTED_ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus      ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus = "AUDIT_REJECTED"
	AUDIT_SEND_REJECTED_ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus = "AUDIT_SEND_REJECTED"
	REMOVED_ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus             ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus = "REMOVED"
)

// All allowed values of ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus enum
var AllowedToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatusEnumValues = []ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus{
	"AUDIT_ACCEPTED",
	"AUDIT_DOING",
	"AUDIT_REJECTED",
	"AUDIT_SEND_REJECTED",
	"REMOVED",
}

// NewToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatusFromValue returns a pointer to a valid ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatusFromValue(v string) (*ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus, error) {
	ev := ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus: valid values are %v", v, AllowedToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus) IsValid() bool {
	for _, existing := range AllowedToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tool_quick_app_management_quick_app_get_v2_data_quick_app_info_status value
func (v ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus) Ptr() *ToolQuickAppManagementQuickAppGetV2DataQuickAppInfoStatus {
	return &v
}
