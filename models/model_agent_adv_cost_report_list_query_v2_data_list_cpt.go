/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentAdvCostReportListQueryV2DataListCpt
type AgentAdvCostReportListQueryV2DataListCpt int64

// List of agent_adv_cost_report_list_query_v2_data_list_cpt
const (
	Enum_0_AgentAdvCostReportListQueryV2DataListCpt AgentAdvCostReportListQueryV2DataListCpt = 0
	Enum_1_AgentAdvCostReportListQueryV2DataListCpt AgentAdvCostReportListQueryV2DataListCpt = 1
)

// All allowed values of AgentAdvCostReportListQueryV2DataListCpt enum
var AllowedAgentAdvCostReportListQueryV2DataListCptEnumValues = []AgentAdvCostReportListQueryV2DataListCpt{
	0,
	1,
}

// NewAgentAdvCostReportListQueryV2DataListCptFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2DataListCpt
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2DataListCptFromValue(v int64) (*AgentAdvCostReportListQueryV2DataListCpt, error) {
	ev := AgentAdvCostReportListQueryV2DataListCpt(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2DataListCpt: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2DataListCptEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2DataListCpt) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2DataListCptEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_data_list_cpt value
func (v AgentAdvCostReportListQueryV2DataListCpt) Ptr() *AgentAdvCostReportListQueryV2DataListCpt {
	return &v
}
