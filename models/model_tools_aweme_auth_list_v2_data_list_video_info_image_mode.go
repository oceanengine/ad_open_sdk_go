/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAwemeAuthListV2DataListVideoInfoImageMode
type ToolsAwemeAuthListV2DataListVideoInfoImageMode string

// List of tools_aweme_auth_list_v2_data_list_video_info_image_mode
const (
	CREATIVE_IMAGE_MODE_AWEME_LIVE_ToolsAwemeAuthListV2DataListVideoInfoImageMode          ToolsAwemeAuthListV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_ToolsAwemeAuthListV2DataListVideoInfoImageMode      ToolsAwemeAuthListV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_GIF_ToolsAwemeAuthListV2DataListVideoInfoImageMode                 ToolsAwemeAuthListV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_GROUP_ToolsAwemeAuthListV2DataListVideoInfoImageMode               ToolsAwemeAuthListV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_LARGE_ToolsAwemeAuthListV2DataListVideoInfoImageMode               ToolsAwemeAuthListV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_ToolsAwemeAuthListV2DataListVideoInfoImageMode      ToolsAwemeAuthListV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_ToolsAwemeAuthListV2DataListVideoInfoImageMode ToolsAwemeAuthListV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_ToolsAwemeAuthListV2DataListVideoInfoImageMode   ToolsAwemeAuthListV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_SMALL_ToolsAwemeAuthListV2DataListVideoInfoImageMode               ToolsAwemeAuthListV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_ToolsAwemeAuthListV2DataListVideoInfoImageMode        ToolsAwemeAuthListV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_ToolsAwemeAuthListV2DataListVideoInfoImageMode  ToolsAwemeAuthListV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_ToolsAwemeAuthListV2DataListVideoInfoImageMode               ToolsAwemeAuthListV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_ToolsAwemeAuthListV2DataListVideoInfoImageMode      ToolsAwemeAuthListV2DataListVideoInfoImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	SEARCH_AD_SMALL_IMAGE_ToolsAwemeAuthListV2DataListVideoInfoImageMode                   ToolsAwemeAuthListV2DataListVideoInfoImageMode = "SEARCH_AD_SMALL_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_ToolsAwemeAuthListV2DataListVideoInfoImageMode                 ToolsAwemeAuthListV2DataListVideoInfoImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of ToolsAwemeAuthListV2DataListVideoInfoImageMode enum
var AllowedToolsAwemeAuthListV2DataListVideoInfoImageModeEnumValues = []ToolsAwemeAuthListV2DataListVideoInfoImageMode{
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_GROUP",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"SEARCH_AD_SMALL_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
}

// NewToolsAwemeAuthListV2DataListVideoInfoImageModeFromValue returns a pointer to a valid ToolsAwemeAuthListV2DataListVideoInfoImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeAuthListV2DataListVideoInfoImageModeFromValue(v string) (*ToolsAwemeAuthListV2DataListVideoInfoImageMode, error) {
	ev := ToolsAwemeAuthListV2DataListVideoInfoImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeAuthListV2DataListVideoInfoImageMode: valid values are %v", v, AllowedToolsAwemeAuthListV2DataListVideoInfoImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeAuthListV2DataListVideoInfoImageMode) IsValid() bool {
	for _, existing := range AllowedToolsAwemeAuthListV2DataListVideoInfoImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_auth_list_v2_data_list_video_info_image_mode value
func (v ToolsAwemeAuthListV2DataListVideoInfoImageMode) Ptr() *ToolsAwemeAuthListV2DataListVideoInfoImageMode {
	return &v
}
