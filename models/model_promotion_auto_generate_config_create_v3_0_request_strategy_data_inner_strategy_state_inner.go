/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionAutoGenerateConfigCreateV30RequestStrategyDataInnerStrategyStateInner struct for PromotionAutoGenerateConfigCreateV30RequestStrategyDataInnerStrategyStateInner
type PromotionAutoGenerateConfigCreateV30RequestStrategyDataInnerStrategyStateInner struct {
	// 配置项key
	StateKey  string                                                                 `json:"state_key"`
	StateType PromotionAutoGenerateConfigCreateV30StrategyDataStrategyStateStateType `json:"state_type"`
	// 配置项值
	StateValue *string `json:"state_value,omitempty"`
}
