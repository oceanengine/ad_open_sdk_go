/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvertiserAssignV2Type
type AgentAdvertiserAssignV2Type string

// List of agent_advertiser_assign_v2_type
const (
	BIDDING_AgentAdvertiserAssignV2Type AgentAdvertiserAssignV2Type = "BIDDING"
	BRAND_AgentAdvertiserAssignV2Type   AgentAdvertiserAssignV2Type = "BRAND"
)

// Ptr returns reference to agent_advertiser_assign_v2_type value
func (v AgentAdvertiserAssignV2Type) Ptr() *AgentAdvertiserAssignV2Type {
	return &v
}
