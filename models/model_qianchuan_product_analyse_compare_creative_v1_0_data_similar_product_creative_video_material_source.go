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

// QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource
type QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource string

// List of qianchuan_product_analyse_compare_creative_v1.0_data_similar_product_creative_video_material_source
const (
	AD_QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource          QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource = "AD"
	AWEME_VIDEO_QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource = "AWEME_VIDEO"
)

// All allowed values of QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource enum
var AllowedQianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSourceEnumValues = []QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource{
	"AD",
	"AWEME_VIDEO",
}

// NewQianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSourceFromValue returns a pointer to a valid QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSourceFromValue(v string) (*QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource, error) {
	ev := QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource: valid values are %v", v, AllowedQianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource) IsValid() bool {
	for _, existing := range AllowedQianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_product_analyse_compare_creative_v1.0_data_similar_product_creative_video_material_source value
func (v QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource) Ptr() *QianchuanProductAnalyseCompareCreativeV10DataSimilarProductCreativeVideoMaterialSource {
	return &v
}
