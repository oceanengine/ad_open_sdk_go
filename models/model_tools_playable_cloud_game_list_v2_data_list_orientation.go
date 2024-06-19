/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPlayableCloudGameListV2DataListOrientation
type ToolsPlayableCloudGameListV2DataListOrientation string

// List of tools_playable_cloud_game_list_v2_data_list_orientation
const (
	HORIZONTAL_ToolsPlayableCloudGameListV2DataListOrientation ToolsPlayableCloudGameListV2DataListOrientation = "HORIZONTAL"
	VERTICAL_ToolsPlayableCloudGameListV2DataListOrientation   ToolsPlayableCloudGameListV2DataListOrientation = "VERTICAL"
)

// All allowed values of ToolsPlayableCloudGameListV2DataListOrientation enum
var AllowedToolsPlayableCloudGameListV2DataListOrientationEnumValues = []ToolsPlayableCloudGameListV2DataListOrientation{
	"HORIZONTAL",
	"VERTICAL",
}

// NewToolsPlayableCloudGameListV2DataListOrientationFromValue returns a pointer to a valid ToolsPlayableCloudGameListV2DataListOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPlayableCloudGameListV2DataListOrientationFromValue(v string) (*ToolsPlayableCloudGameListV2DataListOrientation, error) {
	ev := ToolsPlayableCloudGameListV2DataListOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPlayableCloudGameListV2DataListOrientation: valid values are %v", v, AllowedToolsPlayableCloudGameListV2DataListOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPlayableCloudGameListV2DataListOrientation) IsValid() bool {
	for _, existing := range AllowedToolsPlayableCloudGameListV2DataListOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_playable_cloud_game_list_v2_data_list_orientation value
func (v ToolsPlayableCloudGameListV2DataListOrientation) Ptr() *ToolsPlayableCloudGameListV2DataListOrientation {
	return &v
}
