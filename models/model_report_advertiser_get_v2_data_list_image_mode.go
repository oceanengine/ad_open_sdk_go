/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdvertiserGetV2DataListImageMode
type ReportAdvertiserGetV2DataListImageMode string

// List of report_advertiser_get_v2_data_list_image_mode
const (
	CREATIVE_IMAGE_MODE_LARGE_ReportAdvertiserGetV2DataListImageMode               ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_VIDEO_ReportAdvertiserGetV2DataListImageMode               ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_ReportAdvertiserGetV2DataListImageMode ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_SMALL_ReportAdvertiserGetV2DataListImageMode               ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_ReportAdvertiserGetV2DataListImageMode   ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_ReportAdvertiserGetV2DataListImageMode          ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	TOUTIAO_SEARCH_AD_IMAGE_ReportAdvertiserGetV2DataListImageMode                 ReportAdvertiserGetV2DataListImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_ReportAdvertiserGetV2DataListImageMode      ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	SEARCH_AD_SMALL_IMAGE_ReportAdvertiserGetV2DataListImageMode                   ReportAdvertiserGetV2DataListImageMode = "SEARCH_AD_SMALL_IMAGE"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_ReportAdvertiserGetV2DataListImageMode      ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	MATERIAL_IMAGE_MODE_TITLE_ReportAdvertiserGetV2DataListImageMode               ReportAdvertiserGetV2DataListImageMode = "MATERIAL_IMAGE_MODE_TITLE"
	CREATIVE_IMAGE_MODE_GIF_ReportAdvertiserGetV2DataListImageMode                 ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_ReportAdvertiserGetV2DataListImageMode      ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_ReportAdvertiserGetV2DataListImageMode   ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_ReportAdvertiserGetV2DataListImageMode  ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_ReportAdvertiserGetV2DataListImageMode        ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_GROUP_ReportAdvertiserGetV2DataListImageMode               ReportAdvertiserGetV2DataListImageMode = "CREATIVE_IMAGE_MODE_GROUP"
)

// All allowed values of ReportAdvertiserGetV2DataListImageMode enum
var AllowedReportAdvertiserGetV2DataListImageModeEnumValues = []ReportAdvertiserGetV2DataListImageMode{
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"TOUTIAO_SEARCH_AD_IMAGE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"SEARCH_AD_SMALL_IMAGE",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"MATERIAL_IMAGE_MODE_TITLE",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_GROUP",
}

// NewReportAdvertiserGetV2DataListImageModeFromValue returns a pointer to a valid ReportAdvertiserGetV2DataListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2DataListImageModeFromValue(v string) (*ReportAdvertiserGetV2DataListImageMode, error) {
	ev := ReportAdvertiserGetV2DataListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2DataListImageMode: valid values are %v", v, AllowedReportAdvertiserGetV2DataListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2DataListImageMode) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2DataListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_data_list_image_mode value
func (v ReportAdvertiserGetV2DataListImageMode) Ptr() *ReportAdvertiserGetV2DataListImageMode {
	return &v
}
