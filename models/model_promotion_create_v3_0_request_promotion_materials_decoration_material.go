/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30RequestPromotionMaterialsDecorationMaterial 家装联盟卡券
type PromotionCreateV30RequestPromotionMaterialsDecorationMaterial struct {
	// 活动ID
	ActivityId *string                                                          `json:"activity_id,omitempty"`
	ImageMode  *PromotionCreateV30PromotionMaterialsDecorationMaterialImageMode `json:"image_mode,omitempty"`
}
