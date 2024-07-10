/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeAutoGenerateConfigGetV2ResponseDataStrategyDataInner struct for CreativeAutoGenerateConfigGetV2ResponseDataStrategyDataInner
type CreativeAutoGenerateConfigGetV2ResponseDataStrategyDataInner struct {
	Info *CreativeAutoGenerateConfigGetV2ResponseDataStrategyDataInnerInfo `json:"info,omitempty"`
	// 策略id
	StrategyId int64 `json:"strategy_id"`
	// 策略名称
	StrategyName *string `json:"strategy_name,omitempty"`
	// 策略配置项列表
	StrategyState []*CreativeAutoGenerateConfigGetV2ResponseDataStrategyDataInnerStrategyStateInner `json:"strategy_state,omitempty"`
	Type          CreativeAutoGenerateConfigGetV2DataStrategyDataType                               `json:"type"`
}
