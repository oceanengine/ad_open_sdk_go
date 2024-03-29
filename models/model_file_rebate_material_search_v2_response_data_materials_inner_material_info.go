/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileRebateMaterialSearchV2ResponseDataMaterialsInnerMaterialInfo
type FileRebateMaterialSearchV2ResponseDataMaterialsInnerMaterialInfo struct {
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	AutoAdIds []int64 `json:"auto_ad_ids,omitempty"`
	//
	AutoPromotionIds []int64 `json:"auto_promotion_ids,omitempty"`
	//
	Cost *float64 `json:"cost,omitempty"`
	//
	IsValidVideoMaterial *int32 `json:"is_valid_video_material,omitempty"`
	//
	MaterialCreateTime *string `json:"material_create_time,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	MaterialIsEffective *int32 `json:"material_is_effective,omitempty"`
	//
	MaterialTags []string `json:"material_tags,omitempty"`
	//
	PolicyCost *float64 `json:"policy_cost,omitempty"`
	//
	PromotionIds []int64 `json:"promotion_ids,omitempty"`
	//
	RowId *int64 `json:"row_id,omitempty"`
}
