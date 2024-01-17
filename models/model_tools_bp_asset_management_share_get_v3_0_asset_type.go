/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBpAssetManagementShareGetV30AssetType
type ToolsBpAssetManagementShareGetV30AssetType string

// List of tools_bp_asset_management_share_get_v3.0_asset_type
const (
	APPLETS_ToolsBpAssetManagementShareGetV30AssetType       ToolsBpAssetManagementShareGetV30AssetType = "APPLETS"
	BYTED_APPLETS_ToolsBpAssetManagementShareGetV30AssetType ToolsBpAssetManagementShareGetV30AssetType = "BYTED_APPLETS"
	BYTED_GAME_ToolsBpAssetManagementShareGetV30AssetType    ToolsBpAssetManagementShareGetV30AssetType = "BYTED_GAME"
	WECHAT_GAME_ToolsBpAssetManagementShareGetV30AssetType   ToolsBpAssetManagementShareGetV30AssetType = "WECHAT_GAME"
)

// All allowed values of ToolsBpAssetManagementShareGetV30AssetType enum
var AllowedToolsBpAssetManagementShareGetV30AssetTypeEnumValues = []ToolsBpAssetManagementShareGetV30AssetType{
	"APPLETS",
	"BYTED_APPLETS",
	"BYTED_GAME",
	"WECHAT_GAME",
}

// NewToolsBpAssetManagementShareGetV30AssetTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareGetV30AssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareGetV30AssetTypeFromValue(v string) (*ToolsBpAssetManagementShareGetV30AssetType, error) {
	ev := ToolsBpAssetManagementShareGetV30AssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareGetV30AssetType: valid values are %v", v, AllowedToolsBpAssetManagementShareGetV30AssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareGetV30AssetType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareGetV30AssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_get_v3.0_asset_type value
func (v ToolsBpAssetManagementShareGetV30AssetType) Ptr() *ToolsBpAssetManagementShareGetV30AssetType {
	return &v
}
