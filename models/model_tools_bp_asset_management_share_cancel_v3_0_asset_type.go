/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBpAssetManagementShareCancelV30AssetType
type ToolsBpAssetManagementShareCancelV30AssetType string

// List of tools_bp_asset_management_share_cancel_v3.0_asset_type
const (
	APPLETS_ToolsBpAssetManagementShareCancelV30AssetType       ToolsBpAssetManagementShareCancelV30AssetType = "APPLETS"
	BYTED_APPLETS_ToolsBpAssetManagementShareCancelV30AssetType ToolsBpAssetManagementShareCancelV30AssetType = "BYTED_APPLETS"
	BYTED_GAME_ToolsBpAssetManagementShareCancelV30AssetType    ToolsBpAssetManagementShareCancelV30AssetType = "BYTED_GAME"
	WECHAT_GAME_ToolsBpAssetManagementShareCancelV30AssetType   ToolsBpAssetManagementShareCancelV30AssetType = "WECHAT_GAME"
)

// All allowed values of ToolsBpAssetManagementShareCancelV30AssetType enum
var AllowedToolsBpAssetManagementShareCancelV30AssetTypeEnumValues = []ToolsBpAssetManagementShareCancelV30AssetType{
	"APPLETS",
	"BYTED_APPLETS",
	"BYTED_GAME",
	"WECHAT_GAME",
}

// NewToolsBpAssetManagementShareCancelV30AssetTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareCancelV30AssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareCancelV30AssetTypeFromValue(v string) (*ToolsBpAssetManagementShareCancelV30AssetType, error) {
	ev := ToolsBpAssetManagementShareCancelV30AssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareCancelV30AssetType: valid values are %v", v, AllowedToolsBpAssetManagementShareCancelV30AssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareCancelV30AssetType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareCancelV30AssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_cancel_v3.0_asset_type value
func (v ToolsBpAssetManagementShareCancelV30AssetType) Ptr() *ToolsBpAssetManagementShareCancelV30AssetType {
	return &v
}