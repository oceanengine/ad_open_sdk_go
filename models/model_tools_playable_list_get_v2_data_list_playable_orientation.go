/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableListGetV2DataListPlayableOrientation
type ToolsPlayableListGetV2DataListPlayableOrientation string

// List of tools_playable_list_get_v2_data_list_playable_orientation
const (
	BOTH_ToolsPlayableListGetV2DataListPlayableOrientation      ToolsPlayableListGetV2DataListPlayableOrientation = "BOTH"
	PORTRAIT_ToolsPlayableListGetV2DataListPlayableOrientation  ToolsPlayableListGetV2DataListPlayableOrientation = "PORTRAIT"
	LANDSCAPE_ToolsPlayableListGetV2DataListPlayableOrientation ToolsPlayableListGetV2DataListPlayableOrientation = "LANDSCAPE"
)

// Ptr returns reference to tools_playable_list_get_v2_data_list_playable_orientation value
func (v ToolsPlayableListGetV2DataListPlayableOrientation) Ptr() *ToolsPlayableListGetV2DataListPlayableOrientation {
	return &v
}
