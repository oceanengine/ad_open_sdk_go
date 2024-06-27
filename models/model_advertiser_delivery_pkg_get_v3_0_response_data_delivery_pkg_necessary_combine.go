/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgNecessaryCombine 必填资质模块
type AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgNecessaryCombine struct {
	// 推广类型id，来自【推广产品资质规则配置查询接口】，行业的资质规则中的promotion_type_id
	CombineId int64 `json:"combine_id"`
	// 资质规则
	DeliveryRules []*AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgNecessaryCombineDeliveryRulesInner `json:"delivery_rules"`
	// 推广类型描述
	Description *string `json:"description,omitempty"`
}
