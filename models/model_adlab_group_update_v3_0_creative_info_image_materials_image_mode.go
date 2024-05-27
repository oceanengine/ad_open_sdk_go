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

// AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode
type AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode string

// List of adlab_group_update_v3.0_creative_info_image_materials_image_mode
const (
	CREATIVE_IMAGE_MODE_LARGE_AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode          AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE"
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL_AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"
)

// All allowed values of AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode enum
var AllowedAdlabGroupUpdateV30CreativeInfoImageMaterialsImageModeEnumValues = []AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode{
	"CREATIVE_IMAGE_MODE_LARGE",
	"CREATIVE_IMAGE_MODE_LARGE_VERTICAL",
}

// NewAdlabGroupUpdateV30CreativeInfoImageMaterialsImageModeFromValue returns a pointer to a valid AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateV30CreativeInfoImageMaterialsImageModeFromValue(v string) (*AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode, error) {
	ev := AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode: valid values are %v", v, AllowedAdlabGroupUpdateV30CreativeInfoImageMaterialsImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateV30CreativeInfoImageMaterialsImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_v3.0_creative_info_image_materials_image_mode value
func (v AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode) Ptr() *AdlabGroupUpdateV30CreativeInfoImageMaterialsImageMode {
	return &v
}
