/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestPromotionMaterialsDecorationMaterial 家装卡券素材
type PromotionUpdateV30RequestPromotionMaterialsDecorationMaterial struct {
	// 卡券活动ID
	ActivityId *string                                                          `json:"activity_id,omitempty"`
	ImageMode  *PromotionUpdateV30PromotionMaterialsDecorationMaterialImageMode `json:"image_mode,omitempty"`
}
