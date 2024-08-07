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

// ToolsOrangeSiteGetV30Status
type ToolsOrangeSiteGetV30Status string

// List of tools_orange_site_get_v3.0_status
const (
	SITE_ALL_ToolsOrangeSiteGetV30Status      ToolsOrangeSiteGetV30Status = "SITE_ALL"
	SITE_OFFLINE_ToolsOrangeSiteGetV30Status  ToolsOrangeSiteGetV30Status = "SITE_OFFLINE"
	SITE_ONLINE_ToolsOrangeSiteGetV30Status   ToolsOrangeSiteGetV30Status = "SITE_ONLINE"
	SITE_REJECTED_ToolsOrangeSiteGetV30Status ToolsOrangeSiteGetV30Status = "SITE_REJECTED"
	SITE_TRASH_ToolsOrangeSiteGetV30Status    ToolsOrangeSiteGetV30Status = "SITE_TRASH"
)

// All allowed values of ToolsOrangeSiteGetV30Status enum
var AllowedToolsOrangeSiteGetV30StatusEnumValues = []ToolsOrangeSiteGetV30Status{
	"SITE_ALL",
	"SITE_OFFLINE",
	"SITE_ONLINE",
	"SITE_REJECTED",
	"SITE_TRASH",
}

// NewToolsOrangeSiteGetV30StatusFromValue returns a pointer to a valid ToolsOrangeSiteGetV30Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsOrangeSiteGetV30StatusFromValue(v string) (*ToolsOrangeSiteGetV30Status, error) {
	ev := ToolsOrangeSiteGetV30Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsOrangeSiteGetV30Status: valid values are %v", v, AllowedToolsOrangeSiteGetV30StatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsOrangeSiteGetV30Status) IsValid() bool {
	for _, existing := range AllowedToolsOrangeSiteGetV30StatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_orange_site_get_v3.0_status value
func (v ToolsOrangeSiteGetV30Status) Ptr() *ToolsOrangeSiteGetV30Status {
	return &v
}
