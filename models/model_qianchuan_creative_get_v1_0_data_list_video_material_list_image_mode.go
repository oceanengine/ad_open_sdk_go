/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCreativeGetV10DataListVideoMaterialListImageMode
type QianchuanCreativeGetV10DataListVideoMaterialListImageMode string

// List of qianchuan_creative_get_v1.0_data_list_video_material_list_image_mode
const (
	AWEME_LIVE_ROOM_QianchuanCreativeGetV10DataListVideoMaterialListImageMode QianchuanCreativeGetV10DataListVideoMaterialListImageMode = "AWEME_LIVE_ROOM"
	LARGE_QianchuanCreativeGetV10DataListVideoMaterialListImageMode           QianchuanCreativeGetV10DataListVideoMaterialListImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanCreativeGetV10DataListVideoMaterialListImageMode  QianchuanCreativeGetV10DataListVideoMaterialListImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanCreativeGetV10DataListVideoMaterialListImageMode           QianchuanCreativeGetV10DataListVideoMaterialListImageMode = "SMALL"
	UNION_SPLASH_QianchuanCreativeGetV10DataListVideoMaterialListImageMode    QianchuanCreativeGetV10DataListVideoMaterialListImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanCreativeGetV10DataListVideoMaterialListImageMode     QianchuanCreativeGetV10DataListVideoMaterialListImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanCreativeGetV10DataListVideoMaterialListImageMode  QianchuanCreativeGetV10DataListVideoMaterialListImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanCreativeGetV10DataListVideoMaterialListImageMode enum
var AllowedQianchuanCreativeGetV10DataListVideoMaterialListImageModeEnumValues = []QianchuanCreativeGetV10DataListVideoMaterialListImageMode{
	"AWEME_LIVE_ROOM",
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanCreativeGetV10DataListVideoMaterialListImageModeFromValue returns a pointer to a valid QianchuanCreativeGetV10DataListVideoMaterialListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10DataListVideoMaterialListImageModeFromValue(v string) (*QianchuanCreativeGetV10DataListVideoMaterialListImageMode, error) {
	ev := QianchuanCreativeGetV10DataListVideoMaterialListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10DataListVideoMaterialListImageMode: valid values are %v", v, AllowedQianchuanCreativeGetV10DataListVideoMaterialListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10DataListVideoMaterialListImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10DataListVideoMaterialListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_data_list_video_material_list_image_mode value
func (v QianchuanCreativeGetV10DataListVideoMaterialListImageMode) Ptr() *QianchuanCreativeGetV10DataListVideoMaterialListImageMode {
	return &v
}
