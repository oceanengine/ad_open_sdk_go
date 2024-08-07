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

// ReportCreativeGetV2FilteringImageMode
type ReportCreativeGetV2FilteringImageMode string

// List of report_creative_get_v2_filtering_image_mode
const (
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_ReportCreativeGetV2FilteringImageMode                ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_ReportCreativeGetV2FilteringImageMode                ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_ReportCreativeGetV2FilteringImageMode               ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_ReportCreativeGetV2FilteringImageMode ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_ReportCreativeGetV2FilteringImageMode            ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_SMALL_ReportCreativeGetV2FilteringImageMode                         ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_ReportCreativeGetV2FilteringImageMode             ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_ReportCreativeGetV2FilteringImageMode                  ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_ReportCreativeGetV2FilteringImageMode           ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_GROUP_ReportCreativeGetV2FilteringImageMode                         ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	TOUTIAO_SEARCH_AD_IMAGE_ReportCreativeGetV2FilteringImageMode                           ReportCreativeGetV2FilteringImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_ReportCreativeGetV2FilteringImageMode        ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_GIF_ReportCreativeGetV2FilteringImageMode                           ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_ReportCreativeGetV2FilteringImageMode                ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_LARGE_ReportCreativeGetV2FilteringImageMode                         ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_VIDEO_ReportCreativeGetV2FilteringImageMode                         ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_ReportCreativeGetV2FilteringImageMode             ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_ReportCreativeGetV2FilteringImageMode                    ReportCreativeGetV2FilteringImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	SEARCH_AD_SMALL_IMAGE_ReportCreativeGetV2FilteringImageMode                             ReportCreativeGetV2FilteringImageMode = "SEARCH_AD_SMALL_IMAGE"
)

// All allowed values of ReportCreativeGetV2FilteringImageMode enum
var AllowedReportCreativeGetV2FilteringImageModeEnumValues = []ReportCreativeGetV2FilteringImageMode{
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"CREATIVE_IMAGE_MODE_GROUP",
	"TOUTIAO_SEARCH_AD_IMAGE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"SEARCH_AD_SMALL_IMAGE",
}

// NewReportCreativeGetV2FilteringImageModeFromValue returns a pointer to a valid ReportCreativeGetV2FilteringImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2FilteringImageModeFromValue(v string) (*ReportCreativeGetV2FilteringImageMode, error) {
	ev := ReportCreativeGetV2FilteringImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2FilteringImageMode: valid values are %v", v, AllowedReportCreativeGetV2FilteringImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2FilteringImageMode) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2FilteringImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_filtering_image_mode value
func (v ReportCreativeGetV2FilteringImageMode) Ptr() *ReportCreativeGetV2FilteringImageMode {
	return &v
}
