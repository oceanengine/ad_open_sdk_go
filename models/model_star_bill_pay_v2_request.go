/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarBillPayV2Request struct for StarBillPayV2Request
type StarBillPayV2Request struct {
	// 需求ID。仅供下单后首次付款使用
	CampaignId *int64 `json:"campaign_id,omitempty"`
	// 任务ID。仅供预付单付尾款使用
	OrderId *int64 `json:"order_id,omitempty"`
	// 星图客户ID
	StarId int64 `json:"star_id"`
}