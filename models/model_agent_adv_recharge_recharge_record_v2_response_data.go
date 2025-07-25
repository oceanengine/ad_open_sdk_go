/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvRechargeRechargeRecordV2ResponseData
type AgentAdvRechargeRechargeRecordV2ResponseData struct {
	// 响应码
	Code *int64 `json:"code,omitempty"`
	// 充值记录
	Data []*AgentAdvRechargeRechargeRecordV2ResponseDataDataInner `json:"data,omitempty"`
	// 响应信息
	Message  *string                                               `json:"message,omitempty"`
	PageInfo *AgentAdvRechargeRechargeRecordV2ResponseDataPageInfo `json:"page_info,omitempty"`
	// 充值金额汇总（元）
	TotalAmount *float64 `json:"total_amount,omitempty"`
}
