/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEventAllAssetsDetailV2DataAssetListAssetType
type ToolsEventAllAssetsDetailV2DataAssetListAssetType string

// List of tools_event_all_assets_detail_v2_data_asset_list_asset_type
const (
	THIRD_EXTERNAL_ToolsEventAllAssetsDetailV2DataAssetListAssetType  ToolsEventAllAssetsDetailV2DataAssetListAssetType = "THIRD_EXTERNAL"
	TETRIS_EXTERNAL_ToolsEventAllAssetsDetailV2DataAssetListAssetType ToolsEventAllAssetsDetailV2DataAssetListAssetType = "TETRIS_EXTERNAL"
	APP_ToolsEventAllAssetsDetailV2DataAssetListAssetType             ToolsEventAllAssetsDetailV2DataAssetListAssetType = "APP"
	QUICK_APP_ToolsEventAllAssetsDetailV2DataAssetListAssetType       ToolsEventAllAssetsDetailV2DataAssetListAssetType = "QUICK_APP"
	MINI_PROGRAME_ToolsEventAllAssetsDetailV2DataAssetListAssetType   ToolsEventAllAssetsDetailV2DataAssetListAssetType = "MINI_PROGRAME"
)

// Ptr returns reference to tools_event_all_assets_detail_v2_data_asset_list_asset_type value
func (v ToolsEventAllAssetsDetailV2DataAssetListAssetType) Ptr() *ToolsEventAllAssetsDetailV2DataAssetListAssetType {
	return &v
}
