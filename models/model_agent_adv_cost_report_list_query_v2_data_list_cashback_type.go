/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentAdvCostReportListQueryV2DataListCashbackType
type AgentAdvCostReportListQueryV2DataListCashbackType int64

// List of agent_adv_cost_report_list_query_v2_data_list_cashback_type
const (
	Enum_MINUS_1_AgentAdvCostReportListQueryV2DataListCashbackType AgentAdvCostReportListQueryV2DataListCashbackType = -1
	Enum_0_AgentAdvCostReportListQueryV2DataListCashbackType       AgentAdvCostReportListQueryV2DataListCashbackType = 0
	Enum_1_AgentAdvCostReportListQueryV2DataListCashbackType       AgentAdvCostReportListQueryV2DataListCashbackType = 1
	Enum_2_AgentAdvCostReportListQueryV2DataListCashbackType       AgentAdvCostReportListQueryV2DataListCashbackType = 2
	Enum_3_AgentAdvCostReportListQueryV2DataListCashbackType       AgentAdvCostReportListQueryV2DataListCashbackType = 3
	Enum_4_AgentAdvCostReportListQueryV2DataListCashbackType       AgentAdvCostReportListQueryV2DataListCashbackType = 4
	Enum_5_AgentAdvCostReportListQueryV2DataListCashbackType       AgentAdvCostReportListQueryV2DataListCashbackType = 5
	Enum_6_AgentAdvCostReportListQueryV2DataListCashbackType       AgentAdvCostReportListQueryV2DataListCashbackType = 6
	Enum_7_AgentAdvCostReportListQueryV2DataListCashbackType       AgentAdvCostReportListQueryV2DataListCashbackType = 7
	Enum_8_AgentAdvCostReportListQueryV2DataListCashbackType       AgentAdvCostReportListQueryV2DataListCashbackType = 8
	Enum_9_AgentAdvCostReportListQueryV2DataListCashbackType       AgentAdvCostReportListQueryV2DataListCashbackType = 9
)

// All allowed values of AgentAdvCostReportListQueryV2DataListCashbackType enum
var AllowedAgentAdvCostReportListQueryV2DataListCashbackTypeEnumValues = []AgentAdvCostReportListQueryV2DataListCashbackType{
	-1,
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
}

// NewAgentAdvCostReportListQueryV2DataListCashbackTypeFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2DataListCashbackType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2DataListCashbackTypeFromValue(v int64) (*AgentAdvCostReportListQueryV2DataListCashbackType, error) {
	ev := AgentAdvCostReportListQueryV2DataListCashbackType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2DataListCashbackType: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2DataListCashbackTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2DataListCashbackType) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2DataListCashbackTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_data_list_cashback_type value
func (v AgentAdvCostReportListQueryV2DataListCashbackType) Ptr() *AgentAdvCostReportListQueryV2DataListCashbackType {
	return &v
}
