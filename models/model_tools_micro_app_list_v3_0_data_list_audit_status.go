/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsMicroAppListV30DataListAuditStatus
type ToolsMicroAppListV30DataListAuditStatus string

// List of tools_micro_app_list_v3.0_data_list_audit_status
const (
	AUDIT_ACCEPTED_ToolsMicroAppListV30DataListAuditStatus ToolsMicroAppListV30DataListAuditStatus = "AUDIT_ACCEPTED"
	AUDITING_ToolsMicroAppListV30DataListAuditStatus       ToolsMicroAppListV30DataListAuditStatus = "AUDITING"
	AUDIT_REJECTED_ToolsMicroAppListV30DataListAuditStatus ToolsMicroAppListV30DataListAuditStatus = "AUDIT_REJECTED"
)

// All allowed values of ToolsMicroAppListV30DataListAuditStatus enum
var AllowedToolsMicroAppListV30DataListAuditStatusEnumValues = []ToolsMicroAppListV30DataListAuditStatus{
	"AUDIT_ACCEPTED",
	"AUDITING",
	"AUDIT_REJECTED",
}

// NewToolsMicroAppListV30DataListAuditStatusFromValue returns a pointer to a valid ToolsMicroAppListV30DataListAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsMicroAppListV30DataListAuditStatusFromValue(v string) (*ToolsMicroAppListV30DataListAuditStatus, error) {
	ev := ToolsMicroAppListV30DataListAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsMicroAppListV30DataListAuditStatus: valid values are %v", v, AllowedToolsMicroAppListV30DataListAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsMicroAppListV30DataListAuditStatus) IsValid() bool {
	for _, existing := range AllowedToolsMicroAppListV30DataListAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_micro_app_list_v3.0_data_list_audit_status value
func (v ToolsMicroAppListV30DataListAuditStatus) Ptr() *ToolsMicroAppListV30DataListAuditStatus {
	return &v
}
