/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsWechatGameListV30FilteringAssetStatus
type ToolsWechatGameListV30FilteringAssetStatus string

// List of tools_wechat_game_list_v3.0_filtering_asset_status
const (
	UPGRADED_ToolsWechatGameListV30FilteringAssetStatus ToolsWechatGameListV30FilteringAssetStatus = "UPGRADED"
	ORIGINAL_ToolsWechatGameListV30FilteringAssetStatus ToolsWechatGameListV30FilteringAssetStatus = "ORIGINAL"
)

// All allowed values of ToolsWechatGameListV30FilteringAssetStatus enum
var AllowedToolsWechatGameListV30FilteringAssetStatusEnumValues = []ToolsWechatGameListV30FilteringAssetStatus{
	"UPGRADED",
	"ORIGINAL",
}

// NewToolsWechatGameListV30FilteringAssetStatusFromValue returns a pointer to a valid ToolsWechatGameListV30FilteringAssetStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsWechatGameListV30FilteringAssetStatusFromValue(v string) (*ToolsWechatGameListV30FilteringAssetStatus, error) {
	ev := ToolsWechatGameListV30FilteringAssetStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsWechatGameListV30FilteringAssetStatus: valid values are %v", v, AllowedToolsWechatGameListV30FilteringAssetStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsWechatGameListV30FilteringAssetStatus) IsValid() bool {
	for _, existing := range AllowedToolsWechatGameListV30FilteringAssetStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_wechat_game_list_v3.0_filtering_asset_status value
func (v ToolsWechatGameListV30FilteringAssetStatus) Ptr() *ToolsWechatGameListV30FilteringAssetStatus {
	return &v
}
