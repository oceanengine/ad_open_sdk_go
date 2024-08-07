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

// QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource
type QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource string

// List of qianchuan_uni_promotion_ad_material_get_v1.0_data_ad_material_infos_material_info_video_material_source
const (
	AD_QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource          QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource = "AD"
	AWEME_VIDEO_QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource = "AWEME_VIDEO"
)

// All allowed values of QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource enum
var AllowedQianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSourceEnumValues = []QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource{
	"AD",
	"AWEME_VIDEO",
}

// NewQianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSourceFromValue returns a pointer to a valid QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSourceFromValue(v string) (*QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource, error) {
	ev := QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource: valid values are %v", v, AllowedQianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_ad_material_get_v1.0_data_ad_material_infos_material_info_video_material_source value
func (v QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource) Ptr() *QianchuanUniPromotionAdMaterialGetV10DataAdMaterialInfosMaterialInfoVideoMaterialSource {
	return &v
}
