/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentFundTransferSeqCreateV2TransferType
type AgentFundTransferSeqCreateV2TransferType string

// List of agent_fund_transfer_seq_create_v2_transfer_type
const (
	CASH_DEFAULT_AgentFundTransferSeqCreateV2TransferType   AgentFundTransferSeqCreateV2TransferType = "CASH_DEFAULT"
	CREDIT_BIDDING_AgentFundTransferSeqCreateV2TransferType AgentFundTransferSeqCreateV2TransferType = "CREDIT_BIDDING"
	CREDIT_BRAND_AgentFundTransferSeqCreateV2TransferType   AgentFundTransferSeqCreateV2TransferType = "CREDIT_BRAND"
	CREDIT_GENERAL_AgentFundTransferSeqCreateV2TransferType AgentFundTransferSeqCreateV2TransferType = "CREDIT_GENERAL"
	GRANT_GENERAL_AgentFundTransferSeqCreateV2TransferType  AgentFundTransferSeqCreateV2TransferType = "GRANT_GENERAL"
	PREPAY_BIDDING_AgentFundTransferSeqCreateV2TransferType AgentFundTransferSeqCreateV2TransferType = "PREPAY_BIDDING"
	PREPAY_BRAND_AgentFundTransferSeqCreateV2TransferType   AgentFundTransferSeqCreateV2TransferType = "PREPAY_BRAND"
	PREPAY_GENERAL_AgentFundTransferSeqCreateV2TransferType AgentFundTransferSeqCreateV2TransferType = "PREPAY_GENERAL"
)

// Ptr returns reference to agent_fund_transfer_seq_create_v2_transfer_type value
func (v AgentFundTransferSeqCreateV2TransferType) Ptr() *AgentFundTransferSeqCreateV2TransferType {
	return &v
}
