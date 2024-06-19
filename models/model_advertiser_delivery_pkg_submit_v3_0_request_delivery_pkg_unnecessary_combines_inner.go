/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkgUnnecessaryCombinesInner struct for AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkgUnnecessaryCombinesInner
type AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkgUnnecessaryCombinesInner struct {
	// 来自【推广产品资质规则配置查询接口】，行业的资质规则中的combine_id
	CombineId int64 `json:"combine_id"`
	// 资质规则的组成类型，数组长度最小为1
	DeliveryRules []*AdvertiserDeliveryPkgSubmitV30RequestDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInner `json:"delivery_rules"`
}
