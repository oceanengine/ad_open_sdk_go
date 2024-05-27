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

// AgentQueryRiskPromotionListV2DataDataFinalOperatorTag
type AgentQueryRiskPromotionListV2DataDataFinalOperatorTag string

// List of agent_query_risk_promotion_list_v2_data_data_final_operator_tag
const (
	DECREASE_QUANTITY_AgentQueryRiskPromotionListV2DataDataFinalOperatorTag AgentQueryRiskPromotionListV2DataDataFinalOperatorTag = "DECREASE_QUANTITY"
	EMPTY_AgentQueryRiskPromotionListV2DataDataFinalOperatorTag             AgentQueryRiskPromotionListV2DataDataFinalOperatorTag = "EMPTY"
	INCREASE_QUANTITY_AgentQueryRiskPromotionListV2DataDataFinalOperatorTag AgentQueryRiskPromotionListV2DataDataFinalOperatorTag = "INCREASE_QUANTITY"
	SELF_OPERATION_AgentQueryRiskPromotionListV2DataDataFinalOperatorTag    AgentQueryRiskPromotionListV2DataDataFinalOperatorTag = "SELF_OPERATION"
)

// All allowed values of AgentQueryRiskPromotionListV2DataDataFinalOperatorTag enum
var AllowedAgentQueryRiskPromotionListV2DataDataFinalOperatorTagEnumValues = []AgentQueryRiskPromotionListV2DataDataFinalOperatorTag{
	"DECREASE_QUANTITY",
	"EMPTY",
	"INCREASE_QUANTITY",
	"SELF_OPERATION",
}

// NewAgentQueryRiskPromotionListV2DataDataFinalOperatorTagFromValue returns a pointer to a valid AgentQueryRiskPromotionListV2DataDataFinalOperatorTag
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentQueryRiskPromotionListV2DataDataFinalOperatorTagFromValue(v string) (*AgentQueryRiskPromotionListV2DataDataFinalOperatorTag, error) {
	ev := AgentQueryRiskPromotionListV2DataDataFinalOperatorTag(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentQueryRiskPromotionListV2DataDataFinalOperatorTag: valid values are %v", v, AllowedAgentQueryRiskPromotionListV2DataDataFinalOperatorTagEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentQueryRiskPromotionListV2DataDataFinalOperatorTag) IsValid() bool {
	for _, existing := range AllowedAgentQueryRiskPromotionListV2DataDataFinalOperatorTagEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_query_risk_promotion_list_v2_data_data_final_operator_tag value
func (v AgentQueryRiskPromotionListV2DataDataFinalOperatorTag) Ptr() *AgentQueryRiskPromotionListV2DataDataFinalOperatorTag {
	return &v
}
