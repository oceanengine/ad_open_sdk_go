/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableCloudGameListV2FilteringAdStatus
type ToolsPlayableCloudGameListV2FilteringAdStatus string

// List of tools_playable_cloud_game_list_v2_filtering_ad_status
const (
	INUSE_ToolsPlayableCloudGameListV2FilteringAdStatus  ToolsPlayableCloudGameListV2FilteringAdStatus = "INUSE"
	DELETE_ToolsPlayableCloudGameListV2FilteringAdStatus ToolsPlayableCloudGameListV2FilteringAdStatus = "DELETE"
	UNUSED_ToolsPlayableCloudGameListV2FilteringAdStatus ToolsPlayableCloudGameListV2FilteringAdStatus = "UNUSED"
)

// Ptr returns reference to tools_playable_cloud_game_list_v2_filtering_ad_status value
func (v ToolsPlayableCloudGameListV2FilteringAdStatus) Ptr() *ToolsPlayableCloudGameListV2FilteringAdStatus {
	return &v
}
