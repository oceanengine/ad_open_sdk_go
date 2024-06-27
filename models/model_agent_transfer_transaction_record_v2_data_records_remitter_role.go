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

// AgentTransferTransactionRecordV2DataRecordsRemitterRole
type AgentTransferTransactionRecordV2DataRecordsRemitterRole string

// List of agent_transfer_transaction_record_v2_data_records_remitter_role
const (
	ROLE_ADVERTISER_AgentTransferTransactionRecordV2DataRecordsRemitterRole                    AgentTransferTransactionRecordV2DataRecordsRemitterRole = "ROLE_ADVERTISER"
	ROLE_ADVERTISER_SYSTEM_ACCOUNT_AgentTransferTransactionRecordV2DataRecordsRemitterRole     AgentTransferTransactionRecordV2DataRecordsRemitterRole = "ROLE_ADVERTISER_SYSTEM_ACCOUNT"
	ROLE_AGENT_AgentTransferTransactionRecordV2DataRecordsRemitterRole                         AgentTransferTransactionRecordV2DataRecordsRemitterRole = "ROLE_AGENT"
	ROLE_AGENT_SYSTEM_ACCOUNT_AgentTransferTransactionRecordV2DataRecordsRemitterRole          AgentTransferTransactionRecordV2DataRecordsRemitterRole = "ROLE_AGENT_SYSTEM_ACCOUNT"
	ROLE_CHILD_AGENT_AgentTransferTransactionRecordV2DataRecordsRemitterRole                   AgentTransferTransactionRecordV2DataRecordsRemitterRole = "ROLE_CHILD_AGENT"
	ROLE_ECP_VIRTUAL_ADVERTISER_AgentTransferTransactionRecordV2DataRecordsRemitterRole        AgentTransferTransactionRecordV2DataRecordsRemitterRole = "ROLE_ECP_VIRTUAL_ADVERTISER"
	ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER_AgentTransferTransactionRecordV2DataRecordsRemitterRole AgentTransferTransactionRecordV2DataRecordsRemitterRole = "ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER"
	ROLE_VIRTAUL_ADVERTISER_AgentTransferTransactionRecordV2DataRecordsRemitterRole            AgentTransferTransactionRecordV2DataRecordsRemitterRole = "ROLE_VIRTAUL_ADVERTISER"
)

// All allowed values of AgentTransferTransactionRecordV2DataRecordsRemitterRole enum
var AllowedAgentTransferTransactionRecordV2DataRecordsRemitterRoleEnumValues = []AgentTransferTransactionRecordV2DataRecordsRemitterRole{
	"ROLE_ADVERTISER",
	"ROLE_ADVERTISER_SYSTEM_ACCOUNT",
	"ROLE_AGENT",
	"ROLE_AGENT_SYSTEM_ACCOUNT",
	"ROLE_CHILD_AGENT",
	"ROLE_ECP_VIRTUAL_ADVERTISER",
	"ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER",
	"ROLE_VIRTAUL_ADVERTISER",
}

// NewAgentTransferTransactionRecordV2DataRecordsRemitterRoleFromValue returns a pointer to a valid AgentTransferTransactionRecordV2DataRecordsRemitterRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentTransferTransactionRecordV2DataRecordsRemitterRoleFromValue(v string) (*AgentTransferTransactionRecordV2DataRecordsRemitterRole, error) {
	ev := AgentTransferTransactionRecordV2DataRecordsRemitterRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentTransferTransactionRecordV2DataRecordsRemitterRole: valid values are %v", v, AllowedAgentTransferTransactionRecordV2DataRecordsRemitterRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentTransferTransactionRecordV2DataRecordsRemitterRole) IsValid() bool {
	for _, existing := range AllowedAgentTransferTransactionRecordV2DataRecordsRemitterRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_transfer_transaction_record_v2_data_records_remitter_role value
func (v AgentTransferTransactionRecordV2DataRecordsRemitterRole) Ptr() *AgentTransferTransactionRecordV2DataRecordsRemitterRole {
	return &v
}
