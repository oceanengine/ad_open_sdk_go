/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

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

// All allowed values of AgentFundTransferSeqCreateV2TransferType enum
var AllowedAgentFundTransferSeqCreateV2TransferTypeEnumValues = []AgentFundTransferSeqCreateV2TransferType{
	"CASH_DEFAULT",
	"CREDIT_BIDDING",
	"CREDIT_BRAND",
	"CREDIT_GENERAL",
	"GRANT_GENERAL",
	"PREPAY_BIDDING",
	"PREPAY_BRAND",
	"PREPAY_GENERAL",
}

// NewAgentFundTransferSeqCreateV2TransferTypeFromValue returns a pointer to a valid AgentFundTransferSeqCreateV2TransferType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentFundTransferSeqCreateV2TransferTypeFromValue(v string) (*AgentFundTransferSeqCreateV2TransferType, error) {
	ev := AgentFundTransferSeqCreateV2TransferType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentFundTransferSeqCreateV2TransferType: valid values are %v", v, AllowedAgentFundTransferSeqCreateV2TransferTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentFundTransferSeqCreateV2TransferType) IsValid() bool {
	for _, existing := range AllowedAgentFundTransferSeqCreateV2TransferTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_fund_transfer_seq_create_v2_transfer_type value
func (v AgentFundTransferSeqCreateV2TransferType) Ptr() *AgentFundTransferSeqCreateV2TransferType {
	return &v
}
