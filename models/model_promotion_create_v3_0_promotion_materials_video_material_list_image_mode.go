/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode
type PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode string

// List of promotion_create_v3.0_promotion_materials_video_material_list_image_mode
const (
	CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO_PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode = "CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode                  PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode         PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
)

// All allowed values of PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode enum
var AllowedPromotionCreateV30PromotionMaterialsVideoMaterialListImageModeEnumValues = []PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode{
	"CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
}

// NewPromotionCreateV30PromotionMaterialsVideoMaterialListImageModeFromValue returns a pointer to a valid PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30PromotionMaterialsVideoMaterialListImageModeFromValue(v string) (*PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode, error) {
	ev := PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode: valid values are %v", v, AllowedPromotionCreateV30PromotionMaterialsVideoMaterialListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30PromotionMaterialsVideoMaterialListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_promotion_materials_video_material_list_image_mode value
func (v PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode) Ptr() *PromotionCreateV30PromotionMaterialsVideoMaterialListImageMode {
	return &v
}
