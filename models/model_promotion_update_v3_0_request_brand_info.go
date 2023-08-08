/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestBrandInfo 品牌信息
type PromotionUpdateV30RequestBrandInfo struct {
	// 云图品牌id
	BrandNameId *int64 `json:"brand_name_id,omitempty"`
	// cdp品牌id
	CdpBrandId *int64 `json:"cdp_brand_id,omitempty"`
	// 云图品牌名称
	CdpBrandName *string `json:"cdp_brand_name,omitempty"`
	// 电商品牌id
	EcomBrandId *int64 `json:"ecom_brand_id,omitempty"`
	//
	SubBrandNameIds []string `json:"sub_brand_name_ids,omitempty"`
	// 子品牌名称
	SubBrandNames []string `json:"sub_brand_names,omitempty"`
	// 品牌分类id
	YuntuCategoryId *int64 `json:"yuntu_category_id,omitempty"`
}