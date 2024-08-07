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

// QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode
type QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode string

// List of qianchuan_uni_aweme_ad_create_v1.0_programmatic_creative_media_list_video_material_image_mode
const (
	VIDEO_LARGE_QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode    QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode enum
var AllowedQianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageModeEnumValues = []QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode{
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageModeFromValue returns a pointer to a valid QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageModeFromValue(v string) (*QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode, error) {
	ev := QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode: valid values are %v", v, AllowedQianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_aweme_ad_create_v1.0_programmatic_creative_media_list_video_material_image_mode value
func (v QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode) Ptr() *QianchuanUniAwemeAdCreateV10ProgrammaticCreativeMediaListVideoMaterialImageMode {
	return &v
}
