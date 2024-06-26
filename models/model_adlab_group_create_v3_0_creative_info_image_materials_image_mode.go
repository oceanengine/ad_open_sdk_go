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

// AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode
type AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode string

// List of adlab_group_create_v3.0_creative_info_image_materials_image_mode
const (
	CREATIVE_IMAGE_MODE_LARGE_AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode          AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
)

// All allowed values of AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode enum
var AllowedAdlabGroupCreateV30CreativeInfoImageMaterialsImageModeEnumValues = []AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode{
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
}

// NewAdlabGroupCreateV30CreativeInfoImageMaterialsImageModeFromValue returns a pointer to a valid AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30CreativeInfoImageMaterialsImageModeFromValue(v string) (*AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode, error) {
	ev := AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode: valid values are %v", v, AllowedAdlabGroupCreateV30CreativeInfoImageMaterialsImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30CreativeInfoImageMaterialsImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_creative_info_image_materials_image_mode value
func (v AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode) Ptr() *AdlabGroupCreateV30CreativeInfoImageMaterialsImageMode {
	return &v
}
