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

// QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource
type QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource string

// List of qianchuan_product_analyse_compare_creative_v1.0_data_own_product_creative_video_material_source
const (
	AD_QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource          QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource = "AD"
	AWEME_VIDEO_QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource = "AWEME_VIDEO"
)

// All allowed values of QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource enum
var AllowedQianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSourceEnumValues = []QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource{
	"AD",
	"AWEME_VIDEO",
}

// NewQianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSourceFromValue returns a pointer to a valid QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSourceFromValue(v string) (*QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource, error) {
	ev := QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource: valid values are %v", v, AllowedQianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource) IsValid() bool {
	for _, existing := range AllowedQianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_product_analyse_compare_creative_v1.0_data_own_product_creative_video_material_source value
func (v QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource) Ptr() *QianchuanProductAnalyseCompareCreativeV10DataOwnProductCreativeVideoMaterialSource {
	return &v
}
