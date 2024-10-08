/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentFundTransferSeqCreateV2Request struct for AgentFundTransferSeqCreateV2Request
type AgentFundTransferSeqCreateV2Request struct {
	// 广告主账户ID
	AccountId int64 `json:"account_id"`
	// 代理商账户ID
	AgentId int64 `json:"agent_id"`
	// 转账金额，单位元，支持两位小时
	Amount       float64                                  `json:"amount"`
	TransferType AgentFundTransferSeqCreateV2TransferType `json:"transfer_type"`
}
