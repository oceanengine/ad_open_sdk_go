/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsLiveAuthorizeListV2Status
type ToolsLiveAuthorizeListV2Status string

// List of tools_live_authorize_list_v2_status
const (
	AUTHORIZE_OVERDUE_ToolsLiveAuthorizeListV2Status ToolsLiveAuthorizeListV2Status = "AUTHORIZE_OVERDUE"
	APPLYING_ToolsLiveAuthorizeListV2Status          ToolsLiveAuthorizeListV2Status = "APPLYING"
	AUTHORIZING_ToolsLiveAuthorizeListV2Status       ToolsLiveAuthorizeListV2Status = "AUTHORIZING"
	APPLY_OVERDUE_ToolsLiveAuthorizeListV2Status     ToolsLiveAuthorizeListV2Status = "APPLY_OVERDUE"
)

// All allowed values of ToolsLiveAuthorizeListV2Status enum
var AllowedToolsLiveAuthorizeListV2StatusEnumValues = []ToolsLiveAuthorizeListV2Status{
	"AUTHORIZE_OVERDUE",
	"APPLYING",
	"AUTHORIZING",
	"APPLY_OVERDUE",
}

// NewToolsLiveAuthorizeListV2StatusFromValue returns a pointer to a valid ToolsLiveAuthorizeListV2Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsLiveAuthorizeListV2StatusFromValue(v string) (*ToolsLiveAuthorizeListV2Status, error) {
	ev := ToolsLiveAuthorizeListV2Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsLiveAuthorizeListV2Status: valid values are %v", v, AllowedToolsLiveAuthorizeListV2StatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsLiveAuthorizeListV2Status) IsValid() bool {
	for _, existing := range AllowedToolsLiveAuthorizeListV2StatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_live_authorize_list_v2_status value
func (v ToolsLiveAuthorizeListV2Status) Ptr() *ToolsLiveAuthorizeListV2Status {
	return &v
}
