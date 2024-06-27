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

// AgentAdvCostReportListQueryV2FilteringEcommerceType
type AgentAdvCostReportListQueryV2FilteringEcommerceType int64

// List of agent_adv_cost_report_list_query_v2_filtering_ecommerce_type
const (
	Enum_1_AgentAdvCostReportListQueryV2FilteringEcommerceType AgentAdvCostReportListQueryV2FilteringEcommerceType = 1
	Enum_2_AgentAdvCostReportListQueryV2FilteringEcommerceType AgentAdvCostReportListQueryV2FilteringEcommerceType = 2
	Enum_3_AgentAdvCostReportListQueryV2FilteringEcommerceType AgentAdvCostReportListQueryV2FilteringEcommerceType = 3
	Enum_4_AgentAdvCostReportListQueryV2FilteringEcommerceType AgentAdvCostReportListQueryV2FilteringEcommerceType = 4
	Enum_5_AgentAdvCostReportListQueryV2FilteringEcommerceType AgentAdvCostReportListQueryV2FilteringEcommerceType = 5
)

// All allowed values of AgentAdvCostReportListQueryV2FilteringEcommerceType enum
var AllowedAgentAdvCostReportListQueryV2FilteringEcommerceTypeEnumValues = []AgentAdvCostReportListQueryV2FilteringEcommerceType{
	1,
	2,
	3,
	4,
	5,
}

// NewAgentAdvCostReportListQueryV2FilteringEcommerceTypeFromValue returns a pointer to a valid AgentAdvCostReportListQueryV2FilteringEcommerceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentAdvCostReportListQueryV2FilteringEcommerceTypeFromValue(v int64) (*AgentAdvCostReportListQueryV2FilteringEcommerceType, error) {
	ev := AgentAdvCostReportListQueryV2FilteringEcommerceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentAdvCostReportListQueryV2FilteringEcommerceType: valid values are %v", v, AllowedAgentAdvCostReportListQueryV2FilteringEcommerceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentAdvCostReportListQueryV2FilteringEcommerceType) IsValid() bool {
	for _, existing := range AllowedAgentAdvCostReportListQueryV2FilteringEcommerceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_adv_cost_report_list_query_v2_filtering_ecommerce_type value
func (v AgentAdvCostReportListQueryV2FilteringEcommerceType) Ptr() *AgentAdvCostReportListQueryV2FilteringEcommerceType {
	return &v
}
