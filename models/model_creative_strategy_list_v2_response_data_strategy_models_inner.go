/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeStrategyListV2ResponseDataStrategyModelsInner struct for CreativeStrategyListV2ResponseDataStrategyModelsInner
type CreativeStrategyListV2ResponseDataStrategyModelsInner struct {
	Info *CreativeStrategyListV2ResponseDataStrategyModelsInnerInfo `json:"info,omitempty"`
	// 策略配置项列表
	StateList []*CreativeStrategyListV2ResponseDataStrategyModelsInnerStateListInner `json:"state_list,omitempty"`
	// 策略id
	StrategyId int64 `json:"strategy_id"`
	// 策略名称
	StrategyName *string                                      `json:"strategy_name,omitempty"`
	Type         CreativeStrategyListV2DataStrategyModelsType `json:"type"`
}
