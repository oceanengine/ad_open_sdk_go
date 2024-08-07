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

// QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode
type QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode string

// List of qianchuan_product_analyse_compare_creative_v1.0_data_own_product_creative_image_mode
const (
	AWEME_LIVE_ROOM_QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode = "AWEME_LIVE_ROOM"
	LARGE_QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode           QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode  QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode           QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode = "SMALL"
	SQUARE_QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode          QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode = "SQUARE"
	UNION_SPLASH_QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode    QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode = "UNION_SPLASH"
	VIDEO_LARGE_QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode     QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode  QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode = "VIDEO_VERTICAL"
)

// All allowed values of QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode enum
var AllowedQianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageModeEnumValues = []QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode{
	"AWEME_LIVE_ROOM",
	"LARGE",
	"LARGE_VERTICAL",
	"SMALL",
	"SQUARE",
	"UNION_SPLASH",
	"VIDEO_LARGE",
	"VIDEO_VERTICAL",
}

// NewQianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageModeFromValue returns a pointer to a valid QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageModeFromValue(v string) (*QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode, error) {
	ev := QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode: valid values are %v", v, AllowedQianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode) IsValid() bool {
	for _, existing := range AllowedQianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_product_analyse_compare_creative_v1.0_data_own_product_creative_image_mode value
func (v QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode) Ptr() *QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeImageMode {
	return &v
}
