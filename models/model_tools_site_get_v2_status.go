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

// ToolsSiteGetV2Status
type ToolsSiteGetV2Status string

// List of tools_site_get_v2_status
const (
	SITE_ALL_ToolsSiteGetV2Status      ToolsSiteGetV2Status = "SITE_ALL"
	SITE_ONLINE_ToolsSiteGetV2Status   ToolsSiteGetV2Status = "SITE_ONLINE"
	SITE_OFFLINE_ToolsSiteGetV2Status  ToolsSiteGetV2Status = "SITE_OFFLINE"
	SITE_REJECTED_ToolsSiteGetV2Status ToolsSiteGetV2Status = "SITE_REJECTED"
	SITE_TRASH_ToolsSiteGetV2Status    ToolsSiteGetV2Status = "SITE_TRASH"
)

// All allowed values of ToolsSiteGetV2Status enum
var AllowedToolsSiteGetV2StatusEnumValues = []ToolsSiteGetV2Status{
	"SITE_ALL",
	"SITE_ONLINE",
	"SITE_OFFLINE",
	"SITE_REJECTED",
	"SITE_TRASH",
}

// NewToolsSiteGetV2StatusFromValue returns a pointer to a valid ToolsSiteGetV2Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteGetV2StatusFromValue(v string) (*ToolsSiteGetV2Status, error) {
	ev := ToolsSiteGetV2Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteGetV2Status: valid values are %v", v, AllowedToolsSiteGetV2StatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteGetV2Status) IsValid() bool {
	for _, existing := range AllowedToolsSiteGetV2StatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_get_v2_status value
func (v ToolsSiteGetV2Status) Ptr() *ToolsSiteGetV2Status {
	return &v
}
