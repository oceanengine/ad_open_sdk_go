/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode
type AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode string

// List of adlab_group_create_v3.0_creative_info_video_materials_image_mode
const (
	CREATIVE_IMAGE_MODE_VIDEO_AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode          AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO"
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL_AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"
)

// All allowed values of AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode enum
var AllowedAdlabGroupCreateV30CreativeInfoVideoMaterialsImageModeEnumValues = []AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode{
	"CREATIVE_IMAGE_MODE_VIDEO",
	"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL",
}

// NewAdlabGroupCreateV30CreativeInfoVideoMaterialsImageModeFromValue returns a pointer to a valid AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30CreativeInfoVideoMaterialsImageModeFromValue(v string) (*AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode, error) {
	ev := AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode: valid values are %v", v, AllowedAdlabGroupCreateV30CreativeInfoVideoMaterialsImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30CreativeInfoVideoMaterialsImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_creative_info_video_materials_image_mode value
func (v AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode) Ptr() *AdlabGroupCreateV30CreativeInfoVideoMaterialsImageMode {
	return &v
}
