/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentFundTransferSeqCreateV2ResponseData
type AgentFundTransferSeqCreateV2ResponseData struct {
	// 转账序列号，第二步提交操作需要
	TransferSeq *int64 `json:"transfer_seq,omitempty"`
}
