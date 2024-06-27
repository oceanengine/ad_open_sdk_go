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

// ToolsWechatAppletListV30FilteringAuditStatus
type ToolsWechatAppletListV30FilteringAuditStatus string

// List of tools_wechat_applet_list_v3.0_filtering_audit_status
const (
	AUDIT_ACCEPTED_ToolsWechatAppletListV30FilteringAuditStatus ToolsWechatAppletListV30FilteringAuditStatus = "AUDIT_ACCEPTED"
	AUDITING_ToolsWechatAppletListV30FilteringAuditStatus       ToolsWechatAppletListV30FilteringAuditStatus = "AUDITING"
	AUDIT_REJECTED_ToolsWechatAppletListV30FilteringAuditStatus ToolsWechatAppletListV30FilteringAuditStatus = "AUDIT_REJECTED"
	ALL_ToolsWechatAppletListV30FilteringAuditStatus            ToolsWechatAppletListV30FilteringAuditStatus = "ALL"
)

// All allowed values of ToolsWechatAppletListV30FilteringAuditStatus enum
var AllowedToolsWechatAppletListV30FilteringAuditStatusEnumValues = []ToolsWechatAppletListV30FilteringAuditStatus{
	"AUDIT_ACCEPTED",
	"AUDITING",
	"AUDIT_REJECTED",
	"ALL",
}

// NewToolsWechatAppletListV30FilteringAuditStatusFromValue returns a pointer to a valid ToolsWechatAppletListV30FilteringAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsWechatAppletListV30FilteringAuditStatusFromValue(v string) (*ToolsWechatAppletListV30FilteringAuditStatus, error) {
	ev := ToolsWechatAppletListV30FilteringAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsWechatAppletListV30FilteringAuditStatus: valid values are %v", v, AllowedToolsWechatAppletListV30FilteringAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsWechatAppletListV30FilteringAuditStatus) IsValid() bool {
	for _, existing := range AllowedToolsWechatAppletListV30FilteringAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_wechat_applet_list_v3.0_filtering_audit_status value
func (v ToolsWechatAppletListV30FilteringAuditStatus) Ptr() *ToolsWechatAppletListV30FilteringAuditStatus {
	return &v
}
