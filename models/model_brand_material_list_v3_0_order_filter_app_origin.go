/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandMaterialListV30OrderFilterAppOrigin
type BrandMaterialListV30OrderFilterAppOrigin string

// List of brand_material_list_v3.0_order_filter_app_origin
const (
	AWEME_BrandMaterialListV30OrderFilterAppOrigin   BrandMaterialListV30OrderFilterAppOrigin = "AWEME"
	TOUTIAO_BrandMaterialListV30OrderFilterAppOrigin BrandMaterialListV30OrderFilterAppOrigin = "TOUTIAO"
	XIGUA_BrandMaterialListV30OrderFilterAppOrigin   BrandMaterialListV30OrderFilterAppOrigin = "XIGUA"
)

// Ptr returns reference to brand_material_list_v3.0_order_filter_app_origin value
func (v BrandMaterialListV30OrderFilterAppOrigin) Ptr() *BrandMaterialListV30OrderFilterAppOrigin {
	return &v
}