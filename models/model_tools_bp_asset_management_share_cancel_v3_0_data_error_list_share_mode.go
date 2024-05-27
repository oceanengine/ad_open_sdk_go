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

// ToolsBpAssetManagementShareCancelV30DataErrorListShareMode
type ToolsBpAssetManagementShareCancelV30DataErrorListShareMode string

// List of tools_bp_asset_management_share_cancel_v3.0_data_error_list_share_mode
const (
	BP_ALL_ACCOUNTS_ToolsBpAssetManagementShareCancelV30DataErrorListShareMode      ToolsBpAssetManagementShareCancelV30DataErrorListShareMode = "BP_ALL_ACCOUNTS"
	COMPANY_ALL_ACCOUNTS_ToolsBpAssetManagementShareCancelV30DataErrorListShareMode ToolsBpAssetManagementShareCancelV30DataErrorListShareMode = "COMPANY_ALL_ACCOUNTS"
	PART_ToolsBpAssetManagementShareCancelV30DataErrorListShareMode                 ToolsBpAssetManagementShareCancelV30DataErrorListShareMode = "PART"
)

// All allowed values of ToolsBpAssetManagementShareCancelV30DataErrorListShareMode enum
var AllowedToolsBpAssetManagementShareCancelV30DataErrorListShareModeEnumValues = []ToolsBpAssetManagementShareCancelV30DataErrorListShareMode{
	"BP_ALL_ACCOUNTS",
	"COMPANY_ALL_ACCOUNTS",
	"PART",
}

// NewToolsBpAssetManagementShareCancelV30DataErrorListShareModeFromValue returns a pointer to a valid ToolsBpAssetManagementShareCancelV30DataErrorListShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareCancelV30DataErrorListShareModeFromValue(v string) (*ToolsBpAssetManagementShareCancelV30DataErrorListShareMode, error) {
	ev := ToolsBpAssetManagementShareCancelV30DataErrorListShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareCancelV30DataErrorListShareMode: valid values are %v", v, AllowedToolsBpAssetManagementShareCancelV30DataErrorListShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareCancelV30DataErrorListShareMode) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareCancelV30DataErrorListShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_cancel_v3.0_data_error_list_share_mode value
func (v ToolsBpAssetManagementShareCancelV30DataErrorListShareMode) Ptr() *ToolsBpAssetManagementShareCancelV30DataErrorListShareMode {
	return &v
}
