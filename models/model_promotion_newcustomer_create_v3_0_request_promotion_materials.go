/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionNewcustomerCreateV30RequestPromotionMaterials
type PromotionNewcustomerCreateV30RequestPromotionMaterials struct {
	//
	CarouselMaterialList []*PromotionNewcustomerCreateV30RequestPromotionMaterialsCarouselMaterialListInner `json:"carousel_material_list,omitempty"`
	// 落地页链接，仅支持传入橙子落地页，落地页数量上限为1
	ExternalUrlMaterialList []*PromotionNewcustomerCreateV30RequestPromotionMaterialsExternalUrlMaterialListInner `json:"external_url_material_list,omitempty"`
	// 标题，上限10个，单个标题字数限制[5-55]
	TitleMaterialList []*PromotionNewcustomerCreateV30RequestPromotionMaterialsTitleMaterialListInner `json:"title_material_list"`
	//
	VideoMaterialList []*PromotionNewcustomerCreateV30RequestPromotionMaterialsVideoMaterialListInner `json:"video_material_list,omitempty"`
}
