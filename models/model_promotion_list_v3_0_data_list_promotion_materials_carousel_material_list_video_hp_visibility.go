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

// PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility
type PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility string

// List of promotion_list_v3.0_data_list_promotion_materials_carousel_material_list_video_hp_visibility
const (
	ALWAYS_VISIBLE_PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility         PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility = "ALWAYS_VISIBLE"
	HIDE_AFTER_END_DATE_PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility    PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility = "HIDE_AFTER_END_DATE"
	HIDE_AFTER_NO_PLAYBACK_PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility = "HIDE_AFTER_NO_PLAYBACK"
	HIDE_VIDEO_ON_HP_PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility       PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility = "HIDE_VIDEO_ON_HP"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility enum
var AllowedPromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibilityEnumValues = []PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility{
	"ALWAYS_VISIBLE",
	"HIDE_AFTER_END_DATE",
	"HIDE_AFTER_NO_PLAYBACK",
	"HIDE_VIDEO_ON_HP",
}

// NewPromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibilityFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibilityFromValue(v string) (*PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility, error) {
	ev := PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_carousel_material_list_video_hp_visibility value
func (v PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility) Ptr() *PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility {
	return &v
}
