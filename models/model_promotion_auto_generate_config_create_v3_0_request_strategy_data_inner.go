/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionAutoGenerateConfigCreateV30RequestStrategyDataInner struct for PromotionAutoGenerateConfigCreateV30RequestStrategyDataInner
type PromotionAutoGenerateConfigCreateV30RequestStrategyDataInner struct {
	// 策略id
	StrategyId int64 `json:"strategy_id"`
	// 策略配置项列表
	StrategyState []*PromotionAutoGenerateConfigCreateV30RequestStrategyDataInnerStrategyStateInner `json:"strategy_state,omitempty"`
}
