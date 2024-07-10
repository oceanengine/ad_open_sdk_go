/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgUnnecessaryCombinesInner struct for AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgUnnecessaryCombinesInner
type AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgUnnecessaryCombinesInner struct {
	// 来自【推广产品资质规则配置查询接口】，行业的资质规则中的combine_id
	CombineId int64 `json:"combine_id"`
	// 资质规则
	DeliveryRules []*AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInner `json:"delivery_rules"`
	// 选填资质描述
	Description *string `json:"description,omitempty"`
}
