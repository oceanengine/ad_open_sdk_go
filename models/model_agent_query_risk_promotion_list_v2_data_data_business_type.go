/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentQueryRiskPromotionListV2DataDataBusinessType
type AgentQueryRiskPromotionListV2DataDataBusinessType string

// List of agent_query_risk_promotion_list_v2_data_data_business_type
const (
	AD_AgentQueryRiskPromotionListV2DataDataBusinessType AgentQueryRiskPromotionListV2DataDataBusinessType = "AD"
)

// All allowed values of AgentQueryRiskPromotionListV2DataDataBusinessType enum
var AllowedAgentQueryRiskPromotionListV2DataDataBusinessTypeEnumValues = []AgentQueryRiskPromotionListV2DataDataBusinessType{
	"AD",
}

// NewAgentQueryRiskPromotionListV2DataDataBusinessTypeFromValue returns a pointer to a valid AgentQueryRiskPromotionListV2DataDataBusinessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentQueryRiskPromotionListV2DataDataBusinessTypeFromValue(v string) (*AgentQueryRiskPromotionListV2DataDataBusinessType, error) {
	ev := AgentQueryRiskPromotionListV2DataDataBusinessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentQueryRiskPromotionListV2DataDataBusinessType: valid values are %v", v, AllowedAgentQueryRiskPromotionListV2DataDataBusinessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentQueryRiskPromotionListV2DataDataBusinessType) IsValid() bool {
	for _, existing := range AllowedAgentQueryRiskPromotionListV2DataDataBusinessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_query_risk_promotion_list_v2_data_data_business_type value
func (v AgentQueryRiskPromotionListV2DataDataBusinessType) Ptr() *AgentQueryRiskPromotionListV2DataDataBusinessType {
	return &v
}
