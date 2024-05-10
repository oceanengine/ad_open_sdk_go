/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvertiserSelectV2Filtering
type AgentAdvertiserSelectV2Filtering struct {
	//
	AdvertiserIds []int64                                     `json:"advertiser_ids,omitempty"`
	CostPeriod    *AgentAdvertiserSelectV2FilteringCostPeriod `json:"cost_period,omitempty"`
	//
	StatCostFenGt *int64 `json:"stat_cost_fen_gt,omitempty"`
}
