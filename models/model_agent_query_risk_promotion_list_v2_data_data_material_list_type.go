/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentQueryRiskPromotionListV2DataDataMaterialListType
type AgentQueryRiskPromotionListV2DataDataMaterialListType string

// List of agent_query_risk_promotion_list_v2_data_data_material_list_type
const (
	IMAGE_AgentQueryRiskPromotionListV2DataDataMaterialListType AgentQueryRiskPromotionListV2DataDataMaterialListType = "IMAGE"
	VIDEO_AgentQueryRiskPromotionListV2DataDataMaterialListType AgentQueryRiskPromotionListV2DataDataMaterialListType = "VIDEO"
	SITE_AgentQueryRiskPromotionListV2DataDataMaterialListType  AgentQueryRiskPromotionListV2DataDataMaterialListType = "SITE"
)

// All allowed values of AgentQueryRiskPromotionListV2DataDataMaterialListType enum
var AllowedAgentQueryRiskPromotionListV2DataDataMaterialListTypeEnumValues = []AgentQueryRiskPromotionListV2DataDataMaterialListType{
	"IMAGE",
	"VIDEO",
	"SITE",
}

// NewAgentQueryRiskPromotionListV2DataDataMaterialListTypeFromValue returns a pointer to a valid AgentQueryRiskPromotionListV2DataDataMaterialListType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentQueryRiskPromotionListV2DataDataMaterialListTypeFromValue(v string) (*AgentQueryRiskPromotionListV2DataDataMaterialListType, error) {
	ev := AgentQueryRiskPromotionListV2DataDataMaterialListType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentQueryRiskPromotionListV2DataDataMaterialListType: valid values are %v", v, AllowedAgentQueryRiskPromotionListV2DataDataMaterialListTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentQueryRiskPromotionListV2DataDataMaterialListType) IsValid() bool {
	for _, existing := range AllowedAgentQueryRiskPromotionListV2DataDataMaterialListTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_query_risk_promotion_list_v2_data_data_material_list_type value
func (v AgentQueryRiskPromotionListV2DataDataMaterialListType) Ptr() *AgentQueryRiskPromotionListV2DataDataMaterialListType {
	return &v
}
