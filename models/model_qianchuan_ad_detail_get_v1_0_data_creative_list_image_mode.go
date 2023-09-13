/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataCreativeListImageMode
type QianchuanAdDetailGetV10DataCreativeListImageMode string

// List of qianchuan_ad_detail_get_v1.0_data_creative_list_image_mode
const (
	AWEME_LIVE_ROOM_QianchuanAdDetailGetV10DataCreativeListImageMode QianchuanAdDetailGetV10DataCreativeListImageMode = "AWEME_LIVE_ROOM"
	LARGE_QianchuanAdDetailGetV10DataCreativeListImageMode           QianchuanAdDetailGetV10DataCreativeListImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanAdDetailGetV10DataCreativeListImageMode  QianchuanAdDetailGetV10DataCreativeListImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanAdDetailGetV10DataCreativeListImageMode           QianchuanAdDetailGetV10DataCreativeListImageMode = "SMALL"
	SQUARE_QianchuanAdDetailGetV10DataCreativeListImageMode          QianchuanAdDetailGetV10DataCreativeListImageMode = "SQUARE"
	UNION_SPLASH_QianchuanAdDetailGetV10DataCreativeListImageMode    QianchuanAdDetailGetV10DataCreativeListImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanAdDetailGetV10DataCreativeListImageMode     QianchuanAdDetailGetV10DataCreativeListImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanAdDetailGetV10DataCreativeListImageMode  QianchuanAdDetailGetV10DataCreativeListImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanAdDetailGetV10DataCreativeListImageMode enum
var AllowedQianchuanAdDetailGetV10DataCreativeListImageModeEnumValues = []QianchuanAdDetailGetV10DataCreativeListImageMode{
	"AWEME_LIVE_ROOM",
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanAdDetailGetV10DataCreativeListImageModeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataCreativeListImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataCreativeListImageModeFromValue(v string) (*QianchuanAdDetailGetV10DataCreativeListImageMode, error) {
	ev := QianchuanAdDetailGetV10DataCreativeListImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataCreativeListImageMode: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataCreativeListImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataCreativeListImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataCreativeListImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_creative_list_image_mode value
func (v QianchuanAdDetailGetV10DataCreativeListImageMode) Ptr() *QianchuanAdDetailGetV10DataCreativeListImageMode {
	return &v
}
