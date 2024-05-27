/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode
type AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode string

// List of adlab_group_update_v3.0_creative_info_video_materials_image_mode
const (
	CREATIVE_IMAGE_MODE_VIDEO_AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode          AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
)

// All allowed values of AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode enum
var AllowedAdlabGroupUpdateV30CreativeInfoVideoMaterialsImageModeEnumValues = []AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode{
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
}

// NewAdlabGroupUpdateV30CreativeInfoVideoMaterialsImageModeFromValue returns a pointer to a valid AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateV30CreativeInfoVideoMaterialsImageModeFromValue(v string) (*AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode, error) {
	ev := AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode: valid values are %v", v, AllowedAdlabGroupUpdateV30CreativeInfoVideoMaterialsImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateV30CreativeInfoVideoMaterialsImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_v3.0_creative_info_video_materials_image_mode value
func (v AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode) Ptr() *AdlabGroupUpdateV30CreativeInfoVideoMaterialsImageMode {
	return &v
}
