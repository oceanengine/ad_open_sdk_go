/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandMaterialUpdateV30RequestMaterialListInner struct for BrandMaterialUpdateV30RequestMaterialListInner
type BrandMaterialUpdateV30RequestMaterialListInner struct {
	// 素材是否是ai生成，true为是
	CreativeIsAigc    *bool                                                           `json:"creative_is_aigc,omitempty"`
	CreativeTemplate  BrandMaterialUpdateV30MaterialListCreativeTemplate              `json:"creative_template"`
	MaterialComponent BrandMaterialUpdateV30RequestMaterialListInnerMaterialComponent `json:"material_component"`
	// 创意ID
	MaterialId *int64 `json:"material_id,omitempty"`
	// 创意名称
	MaterialName string `json:"material_name"`
}
