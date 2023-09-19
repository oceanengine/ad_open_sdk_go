/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPreAuditSendV2DataListType
type ToolsPreAuditSendV2DataListType string

// List of tools_pre_audit_send_v2_data_list_type
const (
	IMG_ToolsPreAuditSendV2DataListType   ToolsPreAuditSendV2DataListType = "IMG"
	SITE_ToolsPreAuditSendV2DataListType  ToolsPreAuditSendV2DataListType = "SITE"
	TITLE_ToolsPreAuditSendV2DataListType ToolsPreAuditSendV2DataListType = "TITLE"
	VIDEO_ToolsPreAuditSendV2DataListType ToolsPreAuditSendV2DataListType = "VIDEO"
)

// All allowed values of ToolsPreAuditSendV2DataListType enum
var AllowedToolsPreAuditSendV2DataListTypeEnumValues = []ToolsPreAuditSendV2DataListType{
	"IMG",
	"SITE",
	"TITLE",
	"VIDEO",
}

// NewToolsPreAuditSendV2DataListTypeFromValue returns a pointer to a valid ToolsPreAuditSendV2DataListType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPreAuditSendV2DataListTypeFromValue(v string) (*ToolsPreAuditSendV2DataListType, error) {
	ev := ToolsPreAuditSendV2DataListType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPreAuditSendV2DataListType: valid values are %v", v, AllowedToolsPreAuditSendV2DataListTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPreAuditSendV2DataListType) IsValid() bool {
	for _, existing := range AllowedToolsPreAuditSendV2DataListTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_pre_audit_send_v2_data_list_type value
func (v ToolsPreAuditSendV2DataListType) Ptr() *ToolsPreAuditSendV2DataListType {
	return &v
}
