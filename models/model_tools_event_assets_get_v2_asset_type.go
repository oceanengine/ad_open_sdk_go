/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEventAssetsGetV2AssetType
type ToolsEventAssetsGetV2AssetType string

// List of tools_event_assets_get_v2_asset_type
const (
	THIRD_EXTERNAL_ToolsEventAssetsGetV2AssetType ToolsEventAssetsGetV2AssetType = "THIRD_EXTERNAL"
	APP_ToolsEventAssetsGetV2AssetType            ToolsEventAssetsGetV2AssetType = "APP"
	QUICK_APP_ToolsEventAssetsGetV2AssetType      ToolsEventAssetsGetV2AssetType = "QUICK_APP"
	MINI_PROGRAME_ToolsEventAssetsGetV2AssetType  ToolsEventAssetsGetV2AssetType = "MINI_PROGRAME"
)

// All allowed values of ToolsEventAssetsGetV2AssetType enum
var AllowedToolsEventAssetsGetV2AssetTypeEnumValues = []ToolsEventAssetsGetV2AssetType{
	"THIRD_EXTERNAL",
	"APP",
	"QUICK_APP",
	"MINI_PROGRAME",
}

// NewToolsEventAssetsGetV2AssetTypeFromValue returns a pointer to a valid ToolsEventAssetsGetV2AssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventAssetsGetV2AssetTypeFromValue(v string) (*ToolsEventAssetsGetV2AssetType, error) {
	ev := ToolsEventAssetsGetV2AssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventAssetsGetV2AssetType: valid values are %v", v, AllowedToolsEventAssetsGetV2AssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventAssetsGetV2AssetType) IsValid() bool {
	for _, existing := range AllowedToolsEventAssetsGetV2AssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_assets_get_v2_asset_type value
func (v ToolsEventAssetsGetV2AssetType) Ptr() *ToolsEventAssetsGetV2AssetType {
	return &v
}
