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

// ToolsWechatAppletListV30DataListAuditStatus
type ToolsWechatAppletListV30DataListAuditStatus string

// List of tools_wechat_applet_list_v3.0_data_list_audit_status
const (
	AUDIT_ACCEPTED_ToolsWechatAppletListV30DataListAuditStatus ToolsWechatAppletListV30DataListAuditStatus = "AUDIT_ACCEPTED"
	AUDITING_ToolsWechatAppletListV30DataListAuditStatus       ToolsWechatAppletListV30DataListAuditStatus = "AUDITING"
	AUDIT_REJECTED_ToolsWechatAppletListV30DataListAuditStatus ToolsWechatAppletListV30DataListAuditStatus = "AUDIT_REJECTED"
)

// All allowed values of ToolsWechatAppletListV30DataListAuditStatus enum
var AllowedToolsWechatAppletListV30DataListAuditStatusEnumValues = []ToolsWechatAppletListV30DataListAuditStatus{
	"AUDIT_ACCEPTED",
	"AUDITING",
	"AUDIT_REJECTED",
}

// NewToolsWechatAppletListV30DataListAuditStatusFromValue returns a pointer to a valid ToolsWechatAppletListV30DataListAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsWechatAppletListV30DataListAuditStatusFromValue(v string) (*ToolsWechatAppletListV30DataListAuditStatus, error) {
	ev := ToolsWechatAppletListV30DataListAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsWechatAppletListV30DataListAuditStatus: valid values are %v", v, AllowedToolsWechatAppletListV30DataListAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsWechatAppletListV30DataListAuditStatus) IsValid() bool {
	for _, existing := range AllowedToolsWechatAppletListV30DataListAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_wechat_applet_list_v3.0_data_list_audit_status value
func (v ToolsWechatAppletListV30DataListAuditStatus) Ptr() *ToolsWechatAppletListV30DataListAuditStatus {
	return &v
}
