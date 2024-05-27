/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsLandingGroupGetV2DataListSitesSiteAuditStatus
type ToolsLandingGroupGetV2DataListSitesSiteAuditStatus string

// List of tools_landing_group_get_v2_data_list_sites_site_audit_status
const (
	REJECTED_ToolsLandingGroupGetV2DataListSitesSiteAuditStatus ToolsLandingGroupGetV2DataListSitesSiteAuditStatus = "REJECTED"
	AUDITING_ToolsLandingGroupGetV2DataListSitesSiteAuditStatus ToolsLandingGroupGetV2DataListSitesSiteAuditStatus = "AUDITING"
	CREATED_ToolsLandingGroupGetV2DataListSitesSiteAuditStatus  ToolsLandingGroupGetV2DataListSitesSiteAuditStatus = "CREATED"
	BANNED_ToolsLandingGroupGetV2DataListSitesSiteAuditStatus   ToolsLandingGroupGetV2DataListSitesSiteAuditStatus = "BANNED"
	ACCEPTED_ToolsLandingGroupGetV2DataListSitesSiteAuditStatus ToolsLandingGroupGetV2DataListSitesSiteAuditStatus = "ACCEPTED"
	UNKNOWN_ToolsLandingGroupGetV2DataListSitesSiteAuditStatus  ToolsLandingGroupGetV2DataListSitesSiteAuditStatus = "UNKNOWN"
)

// All allowed values of ToolsLandingGroupGetV2DataListSitesSiteAuditStatus enum
var AllowedToolsLandingGroupGetV2DataListSitesSiteAuditStatusEnumValues = []ToolsLandingGroupGetV2DataListSitesSiteAuditStatus{
	"REJECTED",
	"AUDITING",
	"CREATED",
	"BANNED",
	"ACCEPTED",
	"UNKNOWN",
}

// NewToolsLandingGroupGetV2DataListSitesSiteAuditStatusFromValue returns a pointer to a valid ToolsLandingGroupGetV2DataListSitesSiteAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsLandingGroupGetV2DataListSitesSiteAuditStatusFromValue(v string) (*ToolsLandingGroupGetV2DataListSitesSiteAuditStatus, error) {
	ev := ToolsLandingGroupGetV2DataListSitesSiteAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsLandingGroupGetV2DataListSitesSiteAuditStatus: valid values are %v", v, AllowedToolsLandingGroupGetV2DataListSitesSiteAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsLandingGroupGetV2DataListSitesSiteAuditStatus) IsValid() bool {
	for _, existing := range AllowedToolsLandingGroupGetV2DataListSitesSiteAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_landing_group_get_v2_data_list_sites_site_audit_status value
func (v ToolsLandingGroupGetV2DataListSitesSiteAuditStatus) Ptr() *ToolsLandingGroupGetV2DataListSitesSiteAuditStatus {
	return &v
}
