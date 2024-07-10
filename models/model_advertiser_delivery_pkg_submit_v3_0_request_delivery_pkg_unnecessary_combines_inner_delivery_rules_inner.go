/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInner struct for AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInner
type AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInner struct {
	// 具体的资质信息，数组长度最小为1
	Deliveries []*AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInnerDeliveriesInner `json:"deliveries"`
	// 资质的规则id，来自【推广产品资质规则配置查询接口】
	RuleId int64 `json:"rule_id"`
}
