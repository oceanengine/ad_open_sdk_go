/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus
type AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus string

// List of agent_transfer_transaction_record_v2_data_records_transfer_target_pay_status
const (
	NO_TRANSFER_AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus      AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus = "NO_TRANSFER"
	TRANSFER_FAILED_AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus  AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus = "TRANSFER_FAILED"
	TRANSFER_ING_AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus     AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus = "TRANSFER_ING"
	TRANSFER_PART_AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus    AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus = "TRANSFER_PART"
	TRANSFER_SUCCESS_AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus = "TRANSFER_SUCCESS"
)

// All allowed values of AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus enum
var AllowedAgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatusEnumValues = []AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus{
	"NO_TRANSFER",
	"TRANSFER_FAILED",
	"TRANSFER_ING",
	"TRANSFER_PART",
	"TRANSFER_SUCCESS",
}

// NewAgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatusFromValue returns a pointer to a valid AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatusFromValue(v string) (*AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus, error) {
	ev := AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus: valid values are %v", v, AllowedAgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus) IsValid() bool {
	for _, existing := range AllowedAgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_transfer_transaction_record_v2_data_records_transfer_target_pay_status value
func (v AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus) Ptr() *AgentTransferTransactionRecordV2DataRecordsTransferTargetPayStatus {
	return &v
}
