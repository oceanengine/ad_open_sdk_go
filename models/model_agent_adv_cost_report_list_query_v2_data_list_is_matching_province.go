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

// AgentAdvCostReportListQueryV2DataListIsMatchingProvince
type AgentAdvCostReportListQueryV2DataListIsMatchingProvince int64

// List of agent_adv_cost_report_list_query_v2_data_list_is_matching_province
const (
	Enum_0_AgentAdvCostReportListQueryV2DataListIsMatchingProvince AgentAdvCostReportListQueryV2DataListIsMatchingProvince = 0
	Enum_1_AgentAdvCostReportListQueryV2DataListIsMatchingProvince AgentAdvCostReportListQueryV2DataListIsMatchingProvince = 1
	Enum_2_AgentAdvCostReportListQueryV2DataListIsMatchingProvince AgentAdvCostReportListQueryV2DataListIsMatchingProvince = 2
)

// All allowed values of AgentAdvCostReportListQueryV2DataListIsMatchingProvince enum
var AllowedAgentAdvCostReportListQueryV2DataListIsMatchingProvinceEnumValues = []AgentAdvCostReportListQueryV2DataListIsMatchingProvince{
	0,
	1,
	2,
}

// NewAgentAdvCostReportListQueryV2DataListIsMatchingProvinceFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2DataListIsMatchingProvince
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2DataListIsMatchingProvinceFromValue(v int64) (*AgentAdvCostReportListQueryV2DataListIsMatchingProvince, error) {
	ev := AgentAdvCostReportListQueryV2DataListIsMatchingProvince(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2DataListIsMatchingProvince: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2DataListIsMatchingProvinceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2DataListIsMatchingProvince) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2DataListIsMatchingProvinceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_data_list_is_matching_province value
func (v AgentAdvCostReportListQueryV2DataListIsMatchingProvince) Ptr() *AgentAdvCostReportListQueryV2DataListIsMatchingProvince {
	return &v
}
