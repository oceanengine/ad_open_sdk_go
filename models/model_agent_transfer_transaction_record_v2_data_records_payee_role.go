/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentTransferTransactionRecordV2DataRecordsPayeeRole
type AgentTransferTransactionRecordV2DataRecordsPayeeRole string

// List of agent_transfer_transaction_record_v2_data_records_payee_role
const (
	ROLE_ADVERTISER_AgentTransferTransactionRecordV2DataRecordsPayeeRole                    AgentTransferTransactionRecordV2DataRecordsPayeeRole = "ROLE_ADVERTISER"
	ROLE_ADVERTISER_SYSTEM_ACCOUNT_AgentTransferTransactionRecordV2DataRecordsPayeeRole     AgentTransferTransactionRecordV2DataRecordsPayeeRole = "ROLE_ADVERTISER_SYSTEM_ACCOUNT"
	ROLE_AGENT_AgentTransferTransactionRecordV2DataRecordsPayeeRole                         AgentTransferTransactionRecordV2DataRecordsPayeeRole = "ROLE_AGENT"
	ROLE_AGENT_SYSTEM_ACCOUNT_AgentTransferTransactionRecordV2DataRecordsPayeeRole          AgentTransferTransactionRecordV2DataRecordsPayeeRole = "ROLE_AGENT_SYSTEM_ACCOUNT"
	ROLE_CHILD_AGENT_AgentTransferTransactionRecordV2DataRecordsPayeeRole                   AgentTransferTransactionRecordV2DataRecordsPayeeRole = "ROLE_CHILD_AGENT"
	ROLE_ECP_VIRTUAL_ADVERTISER_AgentTransferTransactionRecordV2DataRecordsPayeeRole        AgentTransferTransactionRecordV2DataRecordsPayeeRole = "ROLE_ECP_VIRTUAL_ADVERTISER"
	ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER_AgentTransferTransactionRecordV2DataRecordsPayeeRole AgentTransferTransactionRecordV2DataRecordsPayeeRole = "ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER"
	ROLE_VIRTAUL_ADVERTISER_AgentTransferTransactionRecordV2DataRecordsPayeeRole            AgentTransferTransactionRecordV2DataRecordsPayeeRole = "ROLE_VIRTAUL_ADVERTISER"
)

// All allowed values of AgentTransferTransactionRecordV2DataRecordsPayeeRole enum
var AllowedAgentTransferTransactionRecordV2DataRecordsPayeeRoleEnumValues = []AgentTransferTransactionRecordV2DataRecordsPayeeRole{
	"ROLE_ADVERTISER",
	"ROLE_ADVERTISER_SYSTEM_ACCOUNT",
	"ROLE_AGENT",
	"ROLE_AGENT_SYSTEM_ACCOUNT",
	"ROLE_CHILD_AGENT",
	"ROLE_ECP_VIRTUAL_ADVERTISER",
	"ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER",
	"ROLE_VIRTAUL_ADVERTISER",
}

// NewAgentTransferTransactionRecordV2DataRecordsPayeeRoleFromValue returns a pointer to a valid AgentTransferTransactionRecordV2DataRecordsPayeeRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentTransferTransactionRecordV2DataRecordsPayeeRoleFromValue(v string) (*AgentTransferTransactionRecordV2DataRecordsPayeeRole, error) {
	ev := AgentTransferTransactionRecordV2DataRecordsPayeeRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentTransferTransactionRecordV2DataRecordsPayeeRole: valid values are %v", v, AllowedAgentTransferTransactionRecordV2DataRecordsPayeeRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentTransferTransactionRecordV2DataRecordsPayeeRole) IsValid() bool {
	for _, existing := range AllowedAgentTransferTransactionRecordV2DataRecordsPayeeRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_transfer_transaction_record_v2_data_records_payee_role value
func (v AgentTransferTransactionRecordV2DataRecordsPayeeRole) Ptr() *AgentTransferTransactionRecordV2DataRecordsPayeeRole {
	return &v
}
