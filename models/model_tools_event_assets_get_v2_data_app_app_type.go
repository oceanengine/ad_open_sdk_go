/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEventAssetsGetV2DataAppAppType
type ToolsEventAssetsGetV2DataAppAppType string

// List of tools_event_assets_get_v2_data_app_app_type
const (
	ANDROID_ToolsEventAssetsGetV2DataAppAppType ToolsEventAssetsGetV2DataAppAppType = "ANDROID"
	IOS_ToolsEventAssetsGetV2DataAppAppType     ToolsEventAssetsGetV2DataAppAppType = "IOS"
)

// All allowed values of ToolsEventAssetsGetV2DataAppAppType enum
var AllowedToolsEventAssetsGetV2DataAppAppTypeEnumValues = []ToolsEventAssetsGetV2DataAppAppType{
	"ANDROID",
	"IOS",
}

// NewToolsEventAssetsGetV2DataAppAppTypeFromValue returns a pointer to a valid ToolsEventAssetsGetV2DataAppAppType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventAssetsGetV2DataAppAppTypeFromValue(v string) (*ToolsEventAssetsGetV2DataAppAppType, error) {
	ev := ToolsEventAssetsGetV2DataAppAppType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventAssetsGetV2DataAppAppType: valid values are %v", v, AllowedToolsEventAssetsGetV2DataAppAppTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventAssetsGetV2DataAppAppType) IsValid() bool {
	for _, existing := range AllowedToolsEventAssetsGetV2DataAppAppTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_assets_get_v2_data_app_app_type value
func (v ToolsEventAssetsGetV2DataAppAppType) Ptr() *ToolsEventAssetsGetV2DataAppAppType {
	return &v
}
