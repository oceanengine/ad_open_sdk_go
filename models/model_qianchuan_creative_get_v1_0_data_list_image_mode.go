/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCreativeGetV10DataListImageMode
type QianchuanCreativeGetV10DataListImageMode string

// List of qianchuan_creative_get_v1.0_data_list_image_mode
const (
	AWEME_LIVE_ROOM_QianchuanCreativeGetV10DataListImageMode QianchuanCreativeGetV10DataListImageMode = "AWEME_LIVE_ROOM"
	LARGE_QianchuanCreativeGetV10DataListImageMode           QianchuanCreativeGetV10DataListImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanCreativeGetV10DataListImageMode  QianchuanCreativeGetV10DataListImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanCreativeGetV10DataListImageMode           QianchuanCreativeGetV10DataListImageMode = "SMALL"
	SQUARE_QianchuanCreativeGetV10DataListImageMode          QianchuanCreativeGetV10DataListImageMode = "SQUARE"
	UNION_SPLASH_QianchuanCreativeGetV10DataListImageMode    QianchuanCreativeGetV10DataListImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanCreativeGetV10DataListImageMode     QianchuanCreativeGetV10DataListImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanCreativeGetV10DataListImageMode  QianchuanCreativeGetV10DataListImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanCreativeGetV10DataListImageMode enum
var AllowedQianchuanCreativeGetV10DataListImageModeEnumValues = []QianchuanCreativeGetV10DataListImageMode{
	"AWEME_LIVE_ROOM",
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanCreativeGetV10DataListImageModeFromValue returns a pointer to a valid QianchuanCreativeGetV10DataListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10DataListImageModeFromValue(v string) (*QianchuanCreativeGetV10DataListImageMode, error) {
	ev := QianchuanCreativeGetV10DataListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10DataListImageMode: valid values are %v", v, AllowedQianchuanCreativeGetV10DataListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10DataListImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10DataListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_data_list_image_mode value
func (v QianchuanCreativeGetV10DataListImageMode) Ptr() *QianchuanCreativeGetV10DataListImageMode {
	return &v
}