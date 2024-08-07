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

// ToolsBpAssetManagementShareV30DataErrorListShareMode
type ToolsBpAssetManagementShareV30DataErrorListShareMode string

// List of tools_bp_asset_management_share_v3.0_data_error_list_share_mode
const (
	BP_ALL_ACCOUNTS_ToolsBpAssetManagementShareV30DataErrorListShareMode      ToolsBpAssetManagementShareV30DataErrorListShareMode = "BP_ALL_ACCOUNTS"
	COMPANY_ALL_ACCOUNTS_ToolsBpAssetManagementShareV30DataErrorListShareMode ToolsBpAssetManagementShareV30DataErrorListShareMode = "COMPANY_ALL_ACCOUNTS"
	PART_ToolsBpAssetManagementShareV30DataErrorListShareMode                 ToolsBpAssetManagementShareV30DataErrorListShareMode = "PART"
)

// All allowed values of ToolsBpAssetManagementShareV30DataErrorListShareMode enum
var AllowedToolsBpAssetManagementShareV30DataErrorListShareModeEnumValues = []ToolsBpAssetManagementShareV30DataErrorListShareMode{
	"BP_ALL_ACCOUNTS",
	"COMPANY_ALL_ACCOUNTS",
	"PART",
}

// NewToolsBpAssetManagementShareV30DataErrorListShareModeFromValue returns a pointer to a valid ToolsBpAssetManagementShareV30DataErrorListShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareV30DataErrorListShareModeFromValue(v string) (*ToolsBpAssetManagementShareV30DataErrorListShareMode, error) {
	ev := ToolsBpAssetManagementShareV30DataErrorListShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareV30DataErrorListShareMode: valid values are %v", v, AllowedToolsBpAssetManagementShareV30DataErrorListShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareV30DataErrorListShareMode) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareV30DataErrorListShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_v3.0_data_error_list_share_mode value
func (v ToolsBpAssetManagementShareV30DataErrorListShareMode) Ptr() *ToolsBpAssetManagementShareV30DataErrorListShareMode {
	return &v
}
