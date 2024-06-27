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

// AgentAdvCostReportListQueryV2DataListSystemOrigin
type AgentAdvCostReportListQueryV2DataListSystemOrigin int64

// List of agent_adv_cost_report_list_query_v2_data_list_system_origin
const (
	Enum_0_AgentAdvCostReportListQueryV2DataListSystemOrigin AgentAdvCostReportListQueryV2DataListSystemOrigin = 0
	Enum_1_AgentAdvCostReportListQueryV2DataListSystemOrigin AgentAdvCostReportListQueryV2DataListSystemOrigin = 1
)

// All allowed values of AgentAdvCostReportListQueryV2DataListSystemOrigin enum
var AllowedAgentAdvCostReportListQueryV2DataListSystemOriginEnumValues = []AgentAdvCostReportListQueryV2DataListSystemOrigin{
	0,
	1,
}

// NewAgentAdvCostReportListQueryV2DataListSystemOriginFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2DataListSystemOrigin
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2DataListSystemOriginFromValue(v int64) (*AgentAdvCostReportListQueryV2DataListSystemOrigin, error) {
	ev := AgentAdvCostReportListQueryV2DataListSystemOrigin(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2DataListSystemOrigin: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2DataListSystemOriginEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2DataListSystemOrigin) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2DataListSystemOriginEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_data_list_system_origin value
func (v AgentAdvCostReportListQueryV2DataListSystemOrigin) Ptr() *AgentAdvCostReportListQueryV2DataListSystemOrigin {
	return &v
}
