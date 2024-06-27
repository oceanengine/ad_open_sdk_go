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

// AgentQueryRiskPromotionListV2BusinessType
type AgentQueryRiskPromotionListV2BusinessType string

// List of agent_query_risk_promotion_list_v2_business_type
const (
	AD_AgentQueryRiskPromotionListV2BusinessType AgentQueryRiskPromotionListV2BusinessType = "AD"
)

// All allowed values of AgentQueryRiskPromotionListV2BusinessType enum
var AllowedAgentQueryRiskPromotionListV2BusinessTypeEnumValues = []AgentQueryRiskPromotionListV2BusinessType{
	"AD",
}

// NewAgentQueryRiskPromotionListV2BusinessTypeFromValue returns a pointer to a valid AgentQueryRiskPromotionListV2BusinessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentQueryRiskPromotionListV2BusinessTypeFromValue(v string) (*AgentQueryRiskPromotionListV2BusinessType, error) {
	ev := AgentQueryRiskPromotionListV2BusinessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentQueryRiskPromotionListV2BusinessType: valid values are %v", v, AllowedAgentQueryRiskPromotionListV2BusinessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentQueryRiskPromotionListV2BusinessType) IsValid() bool {
	for _, existing := range AllowedAgentQueryRiskPromotionListV2BusinessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_query_risk_promotion_list_v2_business_type value
func (v AgentQueryRiskPromotionListV2BusinessType) Ptr() *AgentQueryRiskPromotionListV2BusinessType {
	return &v
}
