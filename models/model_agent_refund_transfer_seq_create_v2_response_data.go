/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentRefundTransferSeqCreateV2ResponseData
type AgentRefundTransferSeqCreateV2ResponseData struct {
	// 退款序列号，第二步提交操作需要
	RefundSeq *int64 `json:"refund_seq,omitempty"`
}
