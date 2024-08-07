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

// AgentInfoV2DataRole
type AgentInfoV2DataRole string

// List of agent_info_v2_data_role
const (
	ROLE_HOTSOON_PROMOTION_ADVERTISER_AgentInfoV2DataRole AgentInfoV2DataRole = "ROLE_HOTSOON_PROMOTION_ADVERTISER"
	ROLE_ECP_ADVERTISER_AgentInfoV2DataRole               AgentInfoV2DataRole = "ROLE_ECP_ADVERTISER"
	ROLE_CHILD_AGENT_AgentInfoV2DataRole                  AgentInfoV2DataRole = "ROLE_CHILD_AGENT"
	ROLE_ADMIN_AgentInfoV2DataRole                        AgentInfoV2DataRole = "ROLE_ADMIN"
	ROLE_DOUDIAN_ADVERTISER_AgentInfoV2DataRole           AgentInfoV2DataRole = "ROLE_DOUDIAN_ADVERTISER"
	ROLE_ECP_INTERNAL_ADVERTISER_AgentInfoV2DataRole      AgentInfoV2DataRole = "ROLE_ECP_INTERNAL_ADVERTISER"
	ROLE_DSP_ADVERTISER_AgentInfoV2DataRole               AgentInfoV2DataRole = "ROLE_DSP_ADVERTISER"
	ROLE_ECP_CHILD_ADVERTISER_AgentInfoV2DataRole         AgentInfoV2DataRole = "ROLE_ECP_CHILD_ADVERTISER"
	ROLE_CHILD_ADVERTISER_AgentInfoV2DataRole             AgentInfoV2DataRole = "ROLE_CHILD_ADVERTISER"
	ROLE_AGENT_ABSTRACT_AgentInfoV2DataRole               AgentInfoV2DataRole = "ROLE_AGENT_ABSTRACT"
	ROLE_AWEME_PROMOTION_ADVERTISER_AgentInfoV2DataRole   AgentInfoV2DataRole = "ROLE_AWEME_PROMOTION_ADVERTISER"
	ROLE_LITE_ADVERTISER_AgentInfoV2DataRole              AgentInfoV2DataRole = "ROLE_LITE_ADVERTISER"
	ROLE_HOTSOON_ADVERTISER_AgentInfoV2DataRole           AgentInfoV2DataRole = "ROLE_HOTSOON_ADVERTISER"
	ROLE_MAJORDOMO_AgentInfoV2DataRole                    AgentInfoV2DataRole = "ROLE_MAJORDOMO"
	ROLE_AGENT_AgentInfoV2DataRole                        AgentInfoV2DataRole = "ROLE_AGENT"
	ROLE_AGENT_SYSTEM_ACCOUNT_AgentInfoV2DataRole         AgentInfoV2DataRole = "ROLE_AGENT_SYSTEM_ACCOUNT"
	ROLE_ADVERTISER_SYSTEM_ACCOUNT_AgentInfoV2DataRole    AgentInfoV2DataRole = "ROLE_ADVERTISER_SYSTEM_ACCOUNT"
	ROLE_INTERNAL_ADV_AgentInfoV2DataRole                 AgentInfoV2DataRole = "ROLE_INTERNAL_ADV"
	ROLE_PGC_ADVERTISER_AgentInfoV2DataRole               AgentInfoV2DataRole = "ROLE_PGC_ADVERTISER"
	ROLE_ADVERTISER_AgentInfoV2DataRole                   AgentInfoV2DataRole = "ROLE_ADVERTISER"
	ROLE_ADVERTISER_ABSTRACT_AgentInfoV2DataRole          AgentInfoV2DataRole = "ROLE_ADVERTISER_ABSTRACT"
)

// All allowed values of AgentInfoV2DataRole enum
var AllowedAgentInfoV2DataRoleEnumValues = []AgentInfoV2DataRole{
	"ROLE_HOTSOON_PROMOTION_ADVERTISER",
	"ROLE_ECP_ADVERTISER",
	"ROLE_CHILD_AGENT",
	"ROLE_ADMIN",
	"ROLE_DOUDIAN_ADVERTISER",
	"ROLE_ECP_INTERNAL_ADVERTISER",
	"ROLE_DSP_ADVERTISER",
	"ROLE_ECP_CHILD_ADVERTISER",
	"ROLE_CHILD_ADVERTISER",
	"ROLE_AGENT_ABSTRACT",
	"ROLE_AWEME_PROMOTION_ADVERTISER",
	"ROLE_LITE_ADVERTISER",
	"ROLE_HOTSOON_ADVERTISER",
	"ROLE_MAJORDOMO",
	"ROLE_AGENT",
	"ROLE_AGENT_SYSTEM_ACCOUNT",
	"ROLE_ADVERTISER_SYSTEM_ACCOUNT",
	"ROLE_INTERNAL_ADV",
	"ROLE_PGC_ADVERTISER",
	"ROLE_ADVERTISER",
	"ROLE_ADVERTISER_ABSTRACT",
}

// NewAgentInfoV2DataRoleFromValue returns a pointer to a valid AgentInfoV2DataRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentInfoV2DataRoleFromValue(v string) (*AgentInfoV2DataRole, error) {
	ev := AgentInfoV2DataRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentInfoV2DataRole: valid values are %v", v, AllowedAgentInfoV2DataRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentInfoV2DataRole) IsValid() bool {
	for _, existing := range AllowedAgentInfoV2DataRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_info_v2_data_role value
func (v AgentInfoV2DataRole) Ptr() *AgentInfoV2DataRole {
	return &v
}
