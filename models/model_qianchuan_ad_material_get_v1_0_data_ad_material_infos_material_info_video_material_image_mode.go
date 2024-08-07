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

// QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode
type QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode string

// List of qianchuan_ad_material_get_v1.0_data_ad_material_infos_material_info_video_material_image_mode
const (
	AWEME_LIVE_ROOM_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode = "AWEME_LIVE_ROOM"
	CAROUSEL_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode        QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode = "CAROUSEL"
	LARGE_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode           QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode  QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode           QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode = "SMALL"
	SQUARE_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode          QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode = "SQUARE"
	UNION_SPLASH_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode    QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode     QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode  QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode enum
var AllowedQianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageModeEnumValues = []QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode{
	"AWEME_LIVE_ROOM",
	"CAROUSEL",
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageModeFromValue returns a pointer to a valid QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageModeFromValue(v string) (*QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode, error) {
	ev := QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode: valid values are %v", v, AllowedQianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_material_get_v1.0_data_ad_material_infos_material_info_video_material_image_mode value
func (v QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode) Ptr() *QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialImageMode {
	return &v
}
