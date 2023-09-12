/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementBpShareCancelV2ShareMode
type ToolsAppManagementBpShareCancelV2ShareMode string

// List of tools_app_management_bp_share_cancel_v2_share_mode
const (
	PART_ToolsAppManagementBpShareCancelV2ShareMode    ToolsAppManagementBpShareCancelV2ShareMode = "PART"
	ALL_ToolsAppManagementBpShareCancelV2ShareMode     ToolsAppManagementBpShareCancelV2ShareMode = "ALL"
	COMPANY_ToolsAppManagementBpShareCancelV2ShareMode ToolsAppManagementBpShareCancelV2ShareMode = "COMPANY"
)

// All allowed values of ToolsAppManagementBpShareCancelV2ShareMode enum
var AllowedToolsAppManagementBpShareCancelV2ShareModeEnumValues = []ToolsAppManagementBpShareCancelV2ShareMode{
	"PART",
	"ALL",
	"COMPANY",
}

// NewToolsAppManagementBpShareCancelV2ShareModeFromValue returns a pointer to a valid ToolsAppManagementBpShareCancelV2ShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareCancelV2ShareModeFromValue(v string) (*ToolsAppManagementBpShareCancelV2ShareMode, error) {
	ev := ToolsAppManagementBpShareCancelV2ShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareCancelV2ShareMode: valid values are %v", v, AllowedToolsAppManagementBpShareCancelV2ShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareCancelV2ShareMode) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareCancelV2ShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_cancel_v2_share_mode value
func (v ToolsAppManagementBpShareCancelV2ShareMode) Ptr() *ToolsAppManagementBpShareCancelV2ShareMode {
	return &v
}
