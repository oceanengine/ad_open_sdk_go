/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCarouselGetV10DataCarouselsImageMode
type QianchuanCarouselGetV10DataCarouselsImageMode string

// List of qianchuan_carousel_get_v1.0_data_carousels_image_mode
const (
	CAROUSEL_QianchuanCarouselGetV10DataCarouselsImageMode QianchuanCarouselGetV10DataCarouselsImageMode = "CAROUSEL"
)

// All allowed values of QianchuanCarouselGetV10DataCarouselsImageMode enum
var AllowedQianchuanCarouselGetV10DataCarouselsImageModeEnumValues = []QianchuanCarouselGetV10DataCarouselsImageMode{
	"CAROUSEL",
}

// NewQianchuanCarouselGetV10DataCarouselsImageModeFromValue returns a pointer to a valid QianchuanCarouselGetV10DataCarouselsImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCarouselGetV10DataCarouselsImageModeFromValue(v string) (*QianchuanCarouselGetV10DataCarouselsImageMode, error) {
	ev := QianchuanCarouselGetV10DataCarouselsImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCarouselGetV10DataCarouselsImageMode: valid values are %v", v, AllowedQianchuanCarouselGetV10DataCarouselsImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCarouselGetV10DataCarouselsImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanCarouselGetV10DataCarouselsImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_carousel_get_v1.0_data_carousels_image_mode value
func (v QianchuanCarouselGetV10DataCarouselsImageMode) Ptr() *QianchuanCarouselGetV10DataCarouselsImageMode {
	return &v
}
