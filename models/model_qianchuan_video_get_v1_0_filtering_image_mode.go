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

// QianchuanVideoGetV10FilteringImageMode
type QianchuanVideoGetV10FilteringImageMode string

// List of qianchuan_video_get_v1.0_filtering_image_mode
const (
	VIDEO_LARGE_QianchuanVideoGetV10FilteringImageMode    QianchuanVideoGetV10FilteringImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanVideoGetV10FilteringImageMode QianchuanVideoGetV10FilteringImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanVideoGetV10FilteringImageMode enum
var AllowedQianchuanVideoGetV10FilteringImageModeEnumValues = []QianchuanVideoGetV10FilteringImageMode{
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanVideoGetV10FilteringImageModeFromValue returns a pointer to a valid QianchuanVideoGetV10FilteringImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanVideoGetV10FilteringImageModeFromValue(v string) (*QianchuanVideoGetV10FilteringImageMode, error) {
	ev := QianchuanVideoGetV10FilteringImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanVideoGetV10FilteringImageMode: valid values are %v", v, AllowedQianchuanVideoGetV10FilteringImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanVideoGetV10FilteringImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanVideoGetV10FilteringImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_video_get_v1.0_filtering_image_mode value
func (v QianchuanVideoGetV10FilteringImageMode) Ptr() *QianchuanVideoGetV10FilteringImageMode {
	return &v
}
