/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeAutoGenerateConfigGetV2ResponseDataStrategyDataInnerStrategyStateInner struct for CreativeAutoGenerateConfigGetV2ResponseDataStrategyDataInnerStrategyStateInner
type CreativeAutoGenerateConfigGetV2ResponseDataStrategyDataInnerStrategyStateInner struct {
	// 配置项key
	StateKey  string                                                                `json:"state_key"`
	StateType CreativeAutoGenerateConfigGetV2DataStrategyDataStrategyStateStateType `json:"state_type"`
	// 配置项值
	StateValue *string `json:"state_value,omitempty"`
}
