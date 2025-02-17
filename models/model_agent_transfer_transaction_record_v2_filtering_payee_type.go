/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentTransferTransactionRecordV2FilteringPayeeType
type AgentTransferTransactionRecordV2FilteringPayeeType string

// List of agent_transfer_transaction_record_v2_filtering_payee_type
const (
	ROLE_ADVERTISER_AgentTransferTransactionRecordV2FilteringPayeeType                    AgentTransferTransactionRecordV2FilteringPayeeType = "ROLE_ADVERTISER"
	ROLE_AGENT_AgentTransferTransactionRecordV2FilteringPayeeType                         AgentTransferTransactionRecordV2FilteringPayeeType = "ROLE_AGENT"
	ROLE_CHILD_AGENT_AgentTransferTransactionRecordV2FilteringPayeeType                   AgentTransferTransactionRecordV2FilteringPayeeType = "ROLE_CHILD_AGENT"
	ROLE_ECP_VIRTUAL_ADVERTISER_AgentTransferTransactionRecordV2FilteringPayeeType        AgentTransferTransactionRecordV2FilteringPayeeType = "ROLE_ECP_VIRTUAL_ADVERTISER"
	ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER_AgentTransferTransactionRecordV2FilteringPayeeType AgentTransferTransactionRecordV2FilteringPayeeType = "ROLE_LOCAL_LIFE_VIRTUAL_ADVERTISER"
	ROLE_SHARE_WALLET_AgentTransferTransactionRecordV2FilteringPayeeType                  AgentTransferTransactionRecordV2FilteringPayeeType = "ROLE_SHARE_WALLET"
	ROLE_VIRTAUL_ADVERTISER_AgentTransferTransactionRecordV2FilteringPayeeType            AgentTransferTransactionRecordV2FilteringPayeeType = "ROLE_VIRTAUL_ADVERTISER"
)

// Ptr returns reference to agent_transfer_transaction_record_v2_filtering_payee_type value
func (v AgentTransferTransactionRecordV2FilteringPayeeType) Ptr() *AgentTransferTransactionRecordV2FilteringPayeeType {
	return &v
}
