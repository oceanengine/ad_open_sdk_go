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

// AgentAdvertiserCopyV2ReportType
type AgentAdvertiserCopyV2ReportType string

// List of agent_advertiser_copy_v2_report_type
const (
	EMPTY_AgentAdvertiserCopyV2ReportType                      AgentAdvertiserCopyV2ReportType = "EMPTY"
	INCREASE_QUANTITY_AgentAdvertiserCopyV2ReportType          AgentAdvertiserCopyV2ReportType = "INCREASE_QUANTITY"
	SELF_OPERATION_AgentAdvertiserCopyV2ReportType             AgentAdvertiserCopyV2ReportType = "SELF_OPERATION"
	INCREASE_QUANTITY_QUANTITY_AgentAdvertiserCopyV2ReportType AgentAdvertiserCopyV2ReportType = "INCREASE_QUANTITY_QUANTITY"
)

// All allowed values of AgentAdvertiserCopyV2ReportType enum
var AllowedAgentAdvertiserCopyV2ReportTypeEnumValues = []AgentAdvertiserCopyV2ReportType{
	"EMPTY",
	"INCREASE_QUANTITY",
	"SELF_OPERATION",
	"INCREASE_QUANTITY_QUANTITY",
}

// NewAgentAdvertiserCopyV2ReportTypeFromValue returns a pointer to a valid AgentAdvertiserCopyV2ReportType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvertiserCopyV2ReportTypeFromValue(v string) (*AgentAdvertiserCopyV2ReportType, error) {
	ev := AgentAdvertiserCopyV2ReportType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvertiserCopyV2ReportType: valid values are %v", v, AllowedAgentAdvertiserCopyV2ReportTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvertiserCopyV2ReportType) IsValid() bool {
	for _, existing := range AllowedAgentAdvertiserCopyV2ReportTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_advertiser_copy_v2_report_type value
func (v AgentAdvertiserCopyV2ReportType) Ptr() *AgentAdvertiserCopyV2ReportType {
	return &v
}
