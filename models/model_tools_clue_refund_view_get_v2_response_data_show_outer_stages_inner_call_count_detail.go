/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInnerCallCountDetail 通话详情列表
type ToolsClueRefundViewGetV2ResponseDataShowOuterStagesInnerCallCountDetail struct {
	// 已接通话单数
	Answered *int64 `json:"answered,omitempty"`
	// 不满足要求话单数
	Invalid *int64 `json:"invalid,omitempty"`
	// 话单总数
	Total *int64 `json:"total,omitempty"`
	// 满足要求话单数
	Valid *int64 `json:"valid,omitempty"`
}
