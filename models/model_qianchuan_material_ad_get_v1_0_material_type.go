/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanMaterialAdGetV10MaterialType
type QianchuanMaterialAdGetV10MaterialType string

// List of qianchuan_material_ad_get_v1.0_material_type
const (
	IMAGE_QianchuanMaterialAdGetV10MaterialType     QianchuanMaterialAdGetV10MaterialType = "IMAGE"
	LIVE_ROOM_QianchuanMaterialAdGetV10MaterialType QianchuanMaterialAdGetV10MaterialType = "LIVE_ROOM"
	TITLE_QianchuanMaterialAdGetV10MaterialType     QianchuanMaterialAdGetV10MaterialType = "TITLE"
	VIDEO_QianchuanMaterialAdGetV10MaterialType     QianchuanMaterialAdGetV10MaterialType = "VIDEO"
)

// Ptr returns reference to qianchuan_material_ad_get_v1.0_material_type value
func (v QianchuanMaterialAdGetV10MaterialType) Ptr() *QianchuanMaterialAdGetV10MaterialType {
	return &v
}
