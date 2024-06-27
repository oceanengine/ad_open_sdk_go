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

// CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode
type CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode string

// List of creative_custom_creative_create_v2_creative_list_decoration_material_image_mode
const (
	CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode             CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL"
	CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode               CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE"
	CREATIVE_IMAGE_MODE_GROUP_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode                         CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_GROUP"
	TOUTIAO_SEARCH_AD_IMAGE_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode                           CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "TOUTIAO_SEARCH_AD_IMAGE"
	CREATIVE_IMAGE_MODE_SMALL_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode                         CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_SMALL"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode            CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO"
	CREATIVE_IMAGE_MODE_UNION_SPLASH_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode                  CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"
	SEARCH_AD_SMALL_IMAGE_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode                             CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "SEARCH_AD_SMALL_IMAGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode                CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
	CREATIVE_IMAGE_MODE_GIF_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode                           CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_GIF"
	CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode           CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL"
	CREATIVE_IMAGE_MODE_DECORATION_COUPON_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode             CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_DECORATION_COUPON"
	CREATIVE_IMAGE_MODE_AWEME_LIVE_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode                    CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode        CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode                CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"
	CREATIVE_IMAGE_MODE_LARGE_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode                         CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_VIDEO_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode                         CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode                CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE_CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE"
)

// All allowed values of CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode enum
var AllowedCreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageModeEnumValues = []CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode{
	"CREATIVE_IMAGE_MODE_PLAYABLE_VERTICAL",
	"CREATIVE_IMAGE_MODE_DIRECT_PLAYABLE",
	"CREATIVE_IMAGE_MODE_GROUP",
	"TOUTIAO_SEARCH_AD_IMAGE",
	"CREATIVE_IMAGE_MODE_SMALL",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO",
	"CREATIVE_IMAGE_MODE_UNION_SPLASH",
	"SEARCH_AD_SMALL_IMAGE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
	"CREATIVE_IMAGE_MODE_GIF",
	"CREATIVE_IMAGE_MODE_PLAYABLE_HORIZONTAL",
	"CREATIVE_IMAGE_MODE_DECORATION_COUPON",
	"CREATIVE_IMAGE_MODE_AWEME_LIVE",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO",
	"CREATIVE_IMAGE_MODE_DISPLAY_WINDOW",
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_SINGLE_IMAGE",
}

// NewCreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageModeFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageModeFromValue(v string) (*CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode, error) {
	ev := CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_creative_list_decoration_material_image_mode value
func (v CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode) Ptr() *CreativeCustomCreativeCreateV2CreativeListDecorationMaterialImageMode {
	return &v
}
