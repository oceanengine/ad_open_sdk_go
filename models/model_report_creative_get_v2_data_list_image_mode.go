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

// ReportCreativeGetV2DataListImageMode
type ReportCreativeGetV2DataListImageMode string

// List of report_creative_get_v2_data_list_image_mode
const (
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_ReportCreativeGetV2DataListImageMode      ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_ReportCreativeGetV2DataListImageMode      ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	TOUTIAO_SEARCH_AD_IMAGE_ReportCreativeGetV2DataListImageMode                 ReportCreativeGetV2DataListImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_VIDEO_ReportCreativeGetV2DataListImageMode               ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_ReportCreativeGetV2DataListImageMode   ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_ReportCreativeGetV2DataListImageMode  ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_ReportCreativeGetV2DataListImageMode      ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_ReportCreativeGetV2DataListImageMode ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_ReportCreativeGetV2DataListImageMode          ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_ReportCreativeGetV2DataListImageMode   ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_GIF_ReportCreativeGetV2DataListImageMode                 ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_SMALL_ReportCreativeGetV2DataListImageMode               ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	SEARCH_AD_SMALL_IMAGE_ReportCreativeGetV2DataListImageMode                   ReportCreativeGetV2DataListImageMode = "SEARCH_AD_SMALL_IMAGE"
	CREATIVE_IMAGE_MODE_GROUP_ReportCreativeGetV2DataListImageMode               ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_ReportCreativeGetV2DataListImageMode        ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_LARGE_ReportCreativeGetV2DataListImageMode               ReportCreativeGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	MATERIAL_IMAGE_MODE_TITLE_ReportCreativeGetV2DataListImageMode               ReportCreativeGetV2DataListImageMode = "MATERIAL_IMAGE_MODE_TITLE"
)

// All allowed values of ReportCreativeGetV2DataListImageMode enum
var AllowedReportCreativeGetV2DataListImageModeEnumValues = []ReportCreativeGetV2DataListImageMode{
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"TOUTIAO_SEARCH_AD_IMAGE",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_SMALL",
	"SEARCH_AD_SMALL_IMAGE",
	"CREATIVE_IMAGE_MODE_GROUP",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_LARGE",
	"MATERIAL_IMAGE_MODE_TITLE",
}

// NewReportCreativeGetV2DataListImageModeFromValue returns a pointer to a valid ReportCreativeGetV2DataListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2DataListImageModeFromValue(v string) (*ReportCreativeGetV2DataListImageMode, error) {
	ev := ReportCreativeGetV2DataListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2DataListImageMode: valid values are %v", v, AllowedReportCreativeGetV2DataListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2DataListImageMode) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2DataListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_data_list_image_mode value
func (v ReportCreativeGetV2DataListImageMode) Ptr() *ReportCreativeGetV2DataListImageMode {
	return &v
}
