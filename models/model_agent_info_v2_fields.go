/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentInfoV2Fields
type AgentInfoV2Fields string

// List of agent_info_v2_fields
const (
	COMPANY_ID_AgentInfoV2Fields     AgentInfoV2Fields = "company_id"
	ACCOUNT_STATUS_AgentInfoV2Fields AgentInfoV2Fields = "account_status"
	CUSTOMER_NAME_AgentInfoV2Fields  AgentInfoV2Fields = "customer_name"
	COMPANY_NAME_AgentInfoV2Fields   AgentInfoV2Fields = "company_name"
	AGENT_ID_AgentInfoV2Fields       AgentInfoV2Fields = "agent_id"
	ROLE_AgentInfoV2Fields           AgentInfoV2Fields = "role"
	AGENT_NAME_AgentInfoV2Fields     AgentInfoV2Fields = "agent_name"
	CREATE_TIME_AgentInfoV2Fields    AgentInfoV2Fields = "create_time"
	CUSTOMER_ID_AgentInfoV2Fields    AgentInfoV2Fields = "customer_id"
)

// Ptr returns reference to agent_info_v2_fields value
func (v AgentInfoV2Fields) Ptr() *AgentInfoV2Fields {
	return &v
}
