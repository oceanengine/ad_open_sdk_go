/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsOrangeSiteGetV30DataListStatus
type ToolsOrangeSiteGetV30DataListStatus string

// List of tools_orange_site_get_v3.0_data_list_status
const (
	AUDIT_ACCEPTED_ToolsOrangeSiteGetV30DataListStatus ToolsOrangeSiteGetV30DataListStatus = "AUDIT_ACCEPTED"
	AUDIT_BANNED_ToolsOrangeSiteGetV30DataListStatus   ToolsOrangeSiteGetV30DataListStatus = "AUDIT_BANNED"
	AUDIT_DOING_ToolsOrangeSiteGetV30DataListStatus    ToolsOrangeSiteGetV30DataListStatus = "AUDIT_DOING"
	AUDIT_REJECTED_ToolsOrangeSiteGetV30DataListStatus ToolsOrangeSiteGetV30DataListStatus = "AUDIT_REJECTED"
	DELETED_ToolsOrangeSiteGetV30DataListStatus        ToolsOrangeSiteGetV30DataListStatus = "DELETED"
	DISABLE_ToolsOrangeSiteGetV30DataListStatus        ToolsOrangeSiteGetV30DataListStatus = "DISABLE"
	EDIT_ToolsOrangeSiteGetV30DataListStatus           ToolsOrangeSiteGetV30DataListStatus = "EDIT"
	ENABLE_ToolsOrangeSiteGetV30DataListStatus         ToolsOrangeSiteGetV30DataListStatus = "ENABLE"
	FORBIDDEN_ToolsOrangeSiteGetV30DataListStatus      ToolsOrangeSiteGetV30DataListStatus = "FORBIDDEN"
)

// All allowed values of ToolsOrangeSiteGetV30DataListStatus enum
var AllowedToolsOrangeSiteGetV30DataListStatusEnumValues = []ToolsOrangeSiteGetV30DataListStatus{
	"AUDIT_ACCEPTED",
	"AUDIT_BANNED",
	"AUDIT_DOING",
	"AUDIT_REJECTED",
	"DELETED",
	"DISABLE",
	"EDIT",
	"ENABLE",
	"FORBIDDEN",
}

// NewToolsOrangeSiteGetV30DataListStatusFromValue returns a pointer to a valid ToolsOrangeSiteGetV30DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsOrangeSiteGetV30DataListStatusFromValue(v string) (*ToolsOrangeSiteGetV30DataListStatus, error) {
	ev := ToolsOrangeSiteGetV30DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsOrangeSiteGetV30DataListStatus: valid values are %v", v, AllowedToolsOrangeSiteGetV30DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsOrangeSiteGetV30DataListStatus) IsValid() bool {
	for _, existing := range AllowedToolsOrangeSiteGetV30DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_orange_site_get_v3.0_data_list_status value
func (v ToolsOrangeSiteGetV30DataListStatus) Ptr() *ToolsOrangeSiteGetV30DataListStatus {
	return &v
}
