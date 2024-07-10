/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementBpShareCancelV2DataErrorListShareMode
type ToolsAppManagementBpShareCancelV2DataErrorListShareMode string

// List of tools_app_management_bp_share_cancel_v2_data_error_list_share_mode
const (
	COMPANY_ToolsAppManagementBpShareCancelV2DataErrorListShareMode ToolsAppManagementBpShareCancelV2DataErrorListShareMode = "COMPANY"
	PART_ToolsAppManagementBpShareCancelV2DataErrorListShareMode    ToolsAppManagementBpShareCancelV2DataErrorListShareMode = "PART"
	ALL_ToolsAppManagementBpShareCancelV2DataErrorListShareMode     ToolsAppManagementBpShareCancelV2DataErrorListShareMode = "ALL"
)

// All allowed values of ToolsAppManagementBpShareCancelV2DataErrorListShareMode enum
var AllowedToolsAppManagementBpShareCancelV2DataErrorListShareModeEnumValues = []ToolsAppManagementBpShareCancelV2DataErrorListShareMode{
	"COMPANY",
	"PART",
	"ALL",
}

// NewToolsAppManagementBpShareCancelV2DataErrorListShareModeFromValue returns a pointer to a valid ToolsAppManagementBpShareCancelV2DataErrorListShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareCancelV2DataErrorListShareModeFromValue(v string) (*ToolsAppManagementBpShareCancelV2DataErrorListShareMode, error) {
	ev := ToolsAppManagementBpShareCancelV2DataErrorListShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareCancelV2DataErrorListShareMode: valid values are %v", v, AllowedToolsAppManagementBpShareCancelV2DataErrorListShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareCancelV2DataErrorListShareMode) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareCancelV2DataErrorListShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_cancel_v2_data_error_list_share_mode value
func (v ToolsAppManagementBpShareCancelV2DataErrorListShareMode) Ptr() *ToolsAppManagementBpShareCancelV2DataErrorListShareMode {
	return &v
}
