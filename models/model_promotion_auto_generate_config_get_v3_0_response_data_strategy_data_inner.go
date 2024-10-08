/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionAutoGenerateConfigGetV30ResponseDataStrategyDataInner struct for PromotionAutoGenerateConfigGetV30ResponseDataStrategyDataInner
type PromotionAutoGenerateConfigGetV30ResponseDataStrategyDataInner struct {
	Info *PromotionAutoGenerateConfigGetV30ResponseDataStrategyDataInnerInfo `json:"info,omitempty"`
	// 策略id
	StrategyId int64 `json:"strategy_id"`
	// 策略名称
	StrategyName *string `json:"strategy_name,omitempty"`
	// 策略配置项列表
	StrategyState []*PromotionAutoGenerateConfigGetV30ResponseDataStrategyDataInnerStrategyStateInner `json:"strategy_state,omitempty"`
	Type          PromotionAutoGenerateConfigGetV30DataStrategyDataType                               `json:"type"`
}
