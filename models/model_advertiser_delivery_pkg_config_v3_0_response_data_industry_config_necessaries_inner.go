/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigNecessariesInner struct for AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigNecessariesInner
type AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigNecessariesInner struct {
	// 推广类型id
	PromotionTypeId int64 `json:"promotion_type_id"`
	// 推广类名称
	PromotionTypeName *string `json:"promotion_type_name,omitempty"`
	// 具体的资质规则
	Rules []*AdvertiserDeliveryPkgConfigV30ResponseDataIndustryConfigNecessariesInnerRulesInner `json:"rules,omitempty"`
}
