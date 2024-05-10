/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBpAssetManagementShareV30ShareMode
type ToolsBpAssetManagementShareV30ShareMode string

// List of tools_bp_asset_management_share_v3.0_share_mode
const (
	BP_ALL_ACCOUNTS_ToolsBpAssetManagementShareV30ShareMode      ToolsBpAssetManagementShareV30ShareMode = "BP_ALL_ACCOUNTS"
	COMPANY_ALL_ACCOUNTS_ToolsBpAssetManagementShareV30ShareMode ToolsBpAssetManagementShareV30ShareMode = "COMPANY_ALL_ACCOUNTS"
	PART_ToolsBpAssetManagementShareV30ShareMode                 ToolsBpAssetManagementShareV30ShareMode = "PART"
)

// All allowed values of ToolsBpAssetManagementShareV30ShareMode enum
var AllowedToolsBpAssetManagementShareV30ShareModeEnumValues = []ToolsBpAssetManagementShareV30ShareMode{
	"BP_ALL_ACCOUNTS",
	"COMPANY_ALL_ACCOUNTS",
	"PART",
}

// NewToolsBpAssetManagementShareV30ShareModeFromValue returns a pointer to a valid ToolsBpAssetManagementShareV30ShareMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareV30ShareModeFromValue(v string) (*ToolsBpAssetManagementShareV30ShareMode, error) {
	ev := ToolsBpAssetManagementShareV30ShareMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareV30ShareMode: valid values are %v", v, AllowedToolsBpAssetManagementShareV30ShareModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareV30ShareMode) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareV30ShareModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_v3.0_share_mode value
func (v ToolsBpAssetManagementShareV30ShareMode) Ptr() *ToolsBpAssetManagementShareV30ShareMode {
	return &v
}
