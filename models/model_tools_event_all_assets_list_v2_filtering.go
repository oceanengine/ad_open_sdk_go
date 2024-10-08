/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEventAllAssetsListV2Filtering
type ToolsEventAllAssetsListV2Filtering struct {
	// 资产id列表,最大`100`
	AssetIds  []int64                                      `json:"asset_ids,omitempty"`
	AssetType *ToolsEventAllAssetsListV2FilteringAssetType `json:"asset_type,omitempty"`
	// 按照资产修改时间筛选，结束时间`YYYY-MM-DD`，必须与开始时间搭配传入 开始时间必须≤结束时间
	ModifyEndTime *string `json:"modify_end_time,omitempty"`
	// 按照资产修改时间筛选，开始时间`YYYY-MM-DD`，必须与结束时间搭配传入 开始时间必须≤结束时间
	ModifyStartTime *string `json:"modify_start_time,omitempty"`
}
