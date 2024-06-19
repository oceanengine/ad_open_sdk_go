/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentTransferTransactionRecordV2FilteringPayeeType
type AgentTransferTransactionRecordV2FilteringPayeeType string

// List of agent_transfer_transaction_record_v2_filtering_payee_type
const (
	ROLE_ADVERTISER_AgentTransferTransactionRecordV2FilteringPayeeType                    AgentTransferTransactionRecordV2FilteringPayeeType = "ROLE_ADVERTISER"
	ROLE_AGENT_AgentTransferTransactionRecordV2FilteringPayeeType                         AgentTransferTransactionRecordV2FilteringPayeeType = "ROLE_AGENT"
	ROLE_CHILD_AGENT_AgentTransferTransactionRecordV2FilteringPayeeType                   AgentTransferTransactionRecordV2FilteringPayeeType = "ROLE_CHILD_AGENT"
	ROLE_ECP_VIRTUAL_ADVERTISER_AgentTransferTransactionRecordV2FilteringPayeeType        AgentTransferTransactionRecordV2FilteringPayeeType = "ROLE_ECP_VIRTUAL_ADVERTISER"
	ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER_AgentTransferTransactionRecordV2FilteringPayeeType AgentTransferTransactionRecordV2FilteringPayeeType = "ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER"
	ROLE_VIRTAUL_ADVERTISER_AgentTransferTransactionRecordV2FilteringPayeeType            AgentTransferTransactionRecordV2FilteringPayeeType = "ROLE_VIRTAUL_ADVERTISER"
)

// All allowed values of AgentTransferTransactionRecordV2FilteringPayeeType enum
var AllowedAgentTransferTransactionRecordV2FilteringPayeeTypeEnumValues = []AgentTransferTransactionRecordV2FilteringPayeeType{
	"ROLE_ADVERTISER",
	"ROLE_AGENT",
	"ROLE_CHILD_AGENT",
	"ROLE_ECP_VIRTUAL_ADVERTISER",
	"ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER",
	"ROLE_VIRTAUL_ADVERTISER",
}

// NewAgentTransferTransactionRecordV2FilteringPayeeTypeFromValue returns a pointer to a valid AgentTransferTransactionRecordV2FilteringPayeeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentTransferTransactionRecordV2FilteringPayeeTypeFromValue(v string) (*AgentTransferTransactionRecordV2FilteringPayeeType, error) {
	ev := AgentTransferTransactionRecordV2FilteringPayeeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentTransferTransactionRecordV2FilteringPayeeType: valid values are %v", v, AllowedAgentTransferTransactionRecordV2FilteringPayeeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentTransferTransactionRecordV2FilteringPayeeType) IsValid() bool {
	for _, existing := range AllowedAgentTransferTransactionRecordV2FilteringPayeeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_transfer_transaction_record_v2_filtering_payee_type value
func (v AgentTransferTransactionRecordV2FilteringPayeeType) Ptr() *AgentTransferTransactionRecordV2FilteringPayeeType {
	return &v
}
