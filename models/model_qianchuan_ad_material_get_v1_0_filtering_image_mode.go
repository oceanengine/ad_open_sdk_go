/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdMaterialGetV10FilteringImageMode
type QianchuanAdMaterialGetV10FilteringImageMode string

// List of qianchuan_ad_material_get_v1.0_filtering_image_mode
const (
	CAROUSEL_QianchuanAdMaterialGetV10FilteringImageMode       QianchuanAdMaterialGetV10FilteringImageMode = "CAROUSEL"
	LARGE_QianchuanAdMaterialGetV10FilteringImageMode          QianchuanAdMaterialGetV10FilteringImageMode = "LARGE"
	LARGE_VERTICAL_QianchuanAdMaterialGetV10FilteringImageMode QianchuanAdMaterialGetV10FilteringImageMode = "LARGE_VERTICAL"
	SMALL_QianchuanAdMaterialGetV10FilteringImageMode          QianchuanAdMaterialGetV10FilteringImageMode = "SMALL"
	VIDEO_LARGE_QianchuanAdMaterialGetV10FilteringImageMode    QianchuanAdMaterialGetV10FilteringImageMode = "VIDEO_LARGE"
	VIDEO_VERTICAL_QianchuanAdMaterialGetV10FilteringImageMode QianchuanAdMaterialGetV10FilteringImageMode = "VIDEO_VERTICAL"
)

// Ptr returns reference to qianchuan_ad_material_get_v1.0_filtering_image_mode value
func (v QianchuanAdMaterialGetV10FilteringImageMode) Ptr() *QianchuanAdMaterialGetV10FilteringImageMode {
	return &v
}
