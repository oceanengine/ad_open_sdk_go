/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateAssignV2RequestDemandInfoComponentInfo 组件信息
type StarDemandCreateAssignV2RequestDemandInfoComponentInfo struct {
	// 行业组件ID
	IndustryComponentId *int64 `json:"industry_component_id,omitempty"`
	// Link组件ID（落地页组件） 目前只支持1个
	LinkComponentIds []int64 `json:"link_component_ids,omitempty"`
	// 直播引流组件ID
	LiveAttractComponentId *int64 `json:"live_attract_component_id,omitempty"`
}
