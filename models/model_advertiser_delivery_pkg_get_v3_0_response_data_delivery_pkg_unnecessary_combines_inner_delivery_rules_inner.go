/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInner struct for AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInner
type AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInner struct {
	// 资质的具体信息
	Deliveries []*AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInnerDeliveriesInner `json:"deliveries"`
	// 原子规则id
	RuleId int64 `json:"rule_id"`
}
