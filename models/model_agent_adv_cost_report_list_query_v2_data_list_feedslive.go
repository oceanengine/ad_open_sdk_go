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

// AgentAdvCostReportListQueryV2DataListFeedslive
type AgentAdvCostReportListQueryV2DataListFeedslive int64

// List of agent_adv_cost_report_list_query_v2_data_list_feedslive
const (
	Enum_0_AgentAdvCostReportListQueryV2DataListFeedslive   AgentAdvCostReportListQueryV2DataListFeedslive = 0
	Enum_1_AgentAdvCostReportListQueryV2DataListFeedslive   AgentAdvCostReportListQueryV2DataListFeedslive = 1
	Enum_999_AgentAdvCostReportListQueryV2DataListFeedslive AgentAdvCostReportListQueryV2DataListFeedslive = 999
)

// All allowed values of AgentAdvCostReportListQueryV2DataListFeedslive enum
var AllowedAgentAdvCostReportListQueryV2DataListFeedsliveEnumValues = []AgentAdvCostReportListQueryV2DataListFeedslive{
	0,
	1,
	999,
}

// NewAgentAdvCostReportListQueryV2DataListFeedsliveFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2DataListFeedslive
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2DataListFeedsliveFromValue(v int64) (*AgentAdvCostReportListQueryV2DataListFeedslive, error) {
	ev := AgentAdvCostReportListQueryV2DataListFeedslive(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2DataListFeedslive: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2DataListFeedsliveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2DataListFeedslive) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2DataListFeedsliveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_data_list_feedslive value
func (v AgentAdvCostReportListQueryV2DataListFeedslive) Ptr() *AgentAdvCostReportListQueryV2DataListFeedslive {
	return &v
}
