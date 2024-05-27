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

// AgentQueryRiskPromotionListV2FilteringFinalOperatorTag
type AgentQueryRiskPromotionListV2FilteringFinalOperatorTag string

// List of agent_query_risk_promotion_list_v2_filtering_final_operator_tag
const (
	DECREASE_QUANTITY_AgentQueryRiskPromotionListV2FilteringFinalOperatorTag AgentQueryRiskPromotionListV2FilteringFinalOperatorTag = "DECREASE_QUANTITY"
	EMPTY_AgentQueryRiskPromotionListV2FilteringFinalOperatorTag             AgentQueryRiskPromotionListV2FilteringFinalOperatorTag = "EMPTY"
	INCREASE_QUANTITY_AgentQueryRiskPromotionListV2FilteringFinalOperatorTag AgentQueryRiskPromotionListV2FilteringFinalOperatorTag = "INCREASE_QUANTITY"
	SELF_OPERATION_AgentQueryRiskPromotionListV2FilteringFinalOperatorTag    AgentQueryRiskPromotionListV2FilteringFinalOperatorTag = "SELF_OPERATION"
)

// All allowed values of AgentQueryRiskPromotionListV2FilteringFinalOperatorTag enum
var AllowedAgentQueryRiskPromotionListV2FilteringFinalOperatorTagEnumValues = []AgentQueryRiskPromotionListV2FilteringFinalOperatorTag{
	"DECREASE_QUANTITY",
	"EMPTY",
	"INCREASE_QUANTITY",
	"SELF_OPERATION",
}

// NewAgentQueryRiskPromotionListV2FilteringFinalOperatorTagFromValue returns a pointer to a valid AgentQueryRiskPromotionListV2FilteringFinalOperatorTag
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentQueryRiskPromotionListV2FilteringFinalOperatorTagFromValue(v string) (*AgentQueryRiskPromotionListV2FilteringFinalOperatorTag, error) {
	ev := AgentQueryRiskPromotionListV2FilteringFinalOperatorTag(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentQueryRiskPromotionListV2FilteringFinalOperatorTag: valid values are %v", v, AllowedAgentQueryRiskPromotionListV2FilteringFinalOperatorTagEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentQueryRiskPromotionListV2FilteringFinalOperatorTag) IsValid() bool {
	for _, existing := range AllowedAgentQueryRiskPromotionListV2FilteringFinalOperatorTagEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_query_risk_promotion_list_v2_filtering_final_operator_tag value
func (v AgentQueryRiskPromotionListV2FilteringFinalOperatorTag) Ptr() *AgentQueryRiskPromotionListV2FilteringFinalOperatorTag {
	return &v
}
