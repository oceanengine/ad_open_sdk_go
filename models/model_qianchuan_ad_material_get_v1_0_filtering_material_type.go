/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdMaterialGetV10FilteringMaterialType
type QianchuanAdMaterialGetV10FilteringMaterialType string

// List of qianchuan_ad_material_get_v1.0_filtering_material_type
const (
	IMAGE_QianchuanAdMaterialGetV10FilteringMaterialType     QianchuanAdMaterialGetV10FilteringMaterialType = "IMAGE"
	TITLE_QianchuanAdMaterialGetV10FilteringMaterialType     QianchuanAdMaterialGetV10FilteringMaterialType = "TITLE"
	LIVE_ROOM_QianchuanAdMaterialGetV10FilteringMaterialType QianchuanAdMaterialGetV10FilteringMaterialType = "LIVE_ROOM"
	VIDEO_QianchuanAdMaterialGetV10FilteringMaterialType     QianchuanAdMaterialGetV10FilteringMaterialType = "VIDEO"
)

// Ptr returns reference to qianchuan_ad_material_get_v1.0_filtering_material_type value
func (v QianchuanAdMaterialGetV10FilteringMaterialType) Ptr() *QianchuanAdMaterialGetV10FilteringMaterialType {
	return &v
}
