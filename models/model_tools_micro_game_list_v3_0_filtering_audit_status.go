/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsMicroGameListV30FilteringAuditStatus
type ToolsMicroGameListV30FilteringAuditStatus string

// List of tools_micro_game_list_v3.0_filtering_audit_status
const (
	AUDIT_ACCEPTED_ToolsMicroGameListV30FilteringAuditStatus ToolsMicroGameListV30FilteringAuditStatus = "AUDIT_ACCEPTED"
	AUDITING_ToolsMicroGameListV30FilteringAuditStatus       ToolsMicroGameListV30FilteringAuditStatus = "AUDITING"
	AUDIT_REJECTED_ToolsMicroGameListV30FilteringAuditStatus ToolsMicroGameListV30FilteringAuditStatus = "AUDIT_REJECTED"
	ALL_ToolsMicroGameListV30FilteringAuditStatus            ToolsMicroGameListV30FilteringAuditStatus = "ALL"
)

// All allowed values of ToolsMicroGameListV30FilteringAuditStatus enum
var AllowedToolsMicroGameListV30FilteringAuditStatusEnumValues = []ToolsMicroGameListV30FilteringAuditStatus{
	"AUDIT_ACCEPTED",
	"AUDITING",
	"AUDIT_REJECTED",
	"ALL",
}

// NewToolsMicroGameListV30FilteringAuditStatusFromValue returns a pointer to a valid ToolsMicroGameListV30FilteringAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsMicroGameListV30FilteringAuditStatusFromValue(v string) (*ToolsMicroGameListV30FilteringAuditStatus, error) {
	ev := ToolsMicroGameListV30FilteringAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsMicroGameListV30FilteringAuditStatus: valid values are %v", v, AllowedToolsMicroGameListV30FilteringAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsMicroGameListV30FilteringAuditStatus) IsValid() bool {
	for _, existing := range AllowedToolsMicroGameListV30FilteringAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_micro_game_list_v3.0_filtering_audit_status value
func (v ToolsMicroGameListV30FilteringAuditStatus) Ptr() *ToolsMicroGameListV30FilteringAuditStatus {
	return &v
}
