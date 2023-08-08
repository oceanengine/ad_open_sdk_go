/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerBrandInfo
type PromotionListV30ResponseDataListInnerBrandInfo struct {
	// 云图品牌id
	BrandNameId *int64 `json:"brand_name_id,omitempty"`
	// cdp品牌id
	CdpBrandId *int64 `json:"cdp_brand_id,omitempty"`
	//
	CdpBrandName *string `json:"cdp_brand_name,omitempty"`
	// 电商品牌ID
	EcomBrandId *int64 `json:"ecom_brand_id,omitempty"`
	// 子品牌信息id
	SubBrandIds []string `json:"sub_brand_ids,omitempty"`
	// 子品牌信息名称
	SubBrandNames []string `json:"sub_brand_names,omitempty"`
	// 云图分类ID
	YuntuCategoryId *int64 `json:"yuntu_category_id,omitempty"`
}
