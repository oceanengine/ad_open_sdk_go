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

// ToolsEventAllAssetsListV2DataAssetListAssetType
type ToolsEventAllAssetsListV2DataAssetListAssetType string

// List of tools_event_all_assets_list_v2_data_asset_list_asset_type
const (
	THIRD_EXTERNAL_ToolsEventAllAssetsListV2DataAssetListAssetType  ToolsEventAllAssetsListV2DataAssetListAssetType = "THIRD_EXTERNAL"
	TETRIS_EXTERNAL_ToolsEventAllAssetsListV2DataAssetListAssetType ToolsEventAllAssetsListV2DataAssetListAssetType = "TETRIS_EXTERNAL"
	APP_ToolsEventAllAssetsListV2DataAssetListAssetType             ToolsEventAllAssetsListV2DataAssetListAssetType = "APP"
	QUICK_APP_ToolsEventAllAssetsListV2DataAssetListAssetType       ToolsEventAllAssetsListV2DataAssetListAssetType = "QUICK_APP"
	MINI_PROGRAME_ToolsEventAllAssetsListV2DataAssetListAssetType   ToolsEventAllAssetsListV2DataAssetListAssetType = "MINI_PROGRAME"
)

// All allowed values of ToolsEventAllAssetsListV2DataAssetListAssetType enum
var AllowedToolsEventAllAssetsListV2DataAssetListAssetTypeEnumValues = []ToolsEventAllAssetsListV2DataAssetListAssetType{
	"THIRD_EXTERNAL",
	"TETRIS_EXTERNAL",
	"APP",
	"QUICK_APP",
	"MINI_PROGRAME",
}

// NewToolsEventAllAssetsListV2DataAssetListAssetTypeFromValue returns a pointer to a valid ToolsEventAllAssetsListV2DataAssetListAssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventAllAssetsListV2DataAssetListAssetTypeFromValue(v string) (*ToolsEventAllAssetsListV2DataAssetListAssetType, error) {
	ev := ToolsEventAllAssetsListV2DataAssetListAssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventAllAssetsListV2DataAssetListAssetType: valid values are %v", v, AllowedToolsEventAllAssetsListV2DataAssetListAssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventAllAssetsListV2DataAssetListAssetType) IsValid() bool {
	for _, existing := range AllowedToolsEventAllAssetsListV2DataAssetListAssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_all_assets_list_v2_data_asset_list_asset_type value
func (v ToolsEventAllAssetsListV2DataAssetListAssetType) Ptr() *ToolsEventAllAssetsListV2DataAssetListAssetType {
	return &v
}
