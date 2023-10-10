/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentAdvertiserCopyV2DataCopyStatus
type AgentAdvertiserCopyV2DataCopyStatus int64

// List of agent_advertiser_copy_v2_data_copy_status
const (
	Enum_1_AgentAdvertiserCopyV2DataCopyStatus AgentAdvertiserCopyV2DataCopyStatus = 1
	Enum_2_AgentAdvertiserCopyV2DataCopyStatus AgentAdvertiserCopyV2DataCopyStatus = 2
	Enum_3_AgentAdvertiserCopyV2DataCopyStatus AgentAdvertiserCopyV2DataCopyStatus = 3
)

// All allowed values of AgentAdvertiserCopyV2DataCopyStatus enum
var AllowedAgentAdvertiserCopyV2DataCopyStatusEnumValues = []AgentAdvertiserCopyV2DataCopyStatus{
	1,
	2,
	3,
}

// NewAgentAdvertiserCopyV2DataCopyStatusFromValue returns a pointer to a valid AgentAdvertiserCopyV2DataCopyStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvertiserCopyV2DataCopyStatusFromValue(v int64) (*AgentAdvertiserCopyV2DataCopyStatus, error) {
	ev := AgentAdvertiserCopyV2DataCopyStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvertiserCopyV2DataCopyStatus: valid values are %v", v, AllowedAgentAdvertiserCopyV2DataCopyStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvertiserCopyV2DataCopyStatus) IsValid() bool {
	for _, existing := range AllowedAgentAdvertiserCopyV2DataCopyStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_advertiser_copy_v2_data_copy_status value
func (v AgentAdvertiserCopyV2DataCopyStatus) Ptr() *AgentAdvertiserCopyV2DataCopyStatus {
	return &v
}
