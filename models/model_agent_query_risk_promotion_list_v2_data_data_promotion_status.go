/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AgentQueryRiskPromotionListV2DataDataPromotionStatus
type AgentQueryRiskPromotionListV2DataDataPromotionStatus string

// List of agent_query_risk_promotion_list_v2_data_data_promotion_status
const (
	ADV_OFFLINE_BUDGET_AgentQueryRiskPromotionListV2DataDataPromotionStatus           AgentQueryRiskPromotionListV2DataDataPromotionStatus = "ADV_OFFLINE_BUDGET"
	ADV_PRE_OFFLINE_BUDGET_AgentQueryRiskPromotionListV2DataDataPromotionStatus       AgentQueryRiskPromotionListV2DataDataPromotionStatus = "ADV_PRE_OFFLINE_BUDGET"
	AUDIT_AgentQueryRiskPromotionListV2DataDataPromotionStatus                        AgentQueryRiskPromotionListV2DataDataPromotionStatus = "AUDIT"
	AWEME_ACCOUNT_DISABLED_AgentQueryRiskPromotionListV2DataDataPromotionStatus       AgentQueryRiskPromotionListV2DataDataPromotionStatus = "AWEME_ACCOUNT_DISABLED"
	AWEME_ANCHOR_DISABLED_AgentQueryRiskPromotionListV2DataDataPromotionStatus        AgentQueryRiskPromotionListV2DataDataPromotionStatus = "AWEME_ANCHOR_DISABLED"
	CREATE_AgentQueryRiskPromotionListV2DataDataPromotionStatus                       AgentQueryRiskPromotionListV2DataDataPromotionStatus = "CREATE"
	DELETE_AgentQueryRiskPromotionListV2DataDataPromotionStatus                       AgentQueryRiskPromotionListV2DataDataPromotionStatus = "DELETE"
	DELIVERY_OK_AgentQueryRiskPromotionListV2DataDataPromotionStatus                  AgentQueryRiskPromotionListV2DataDataPromotionStatus = "DELIVERY_OK"
	ERROR_DEFAULT_AgentQueryRiskPromotionListV2DataDataPromotionStatus                AgentQueryRiskPromotionListV2DataDataPromotionStatus = "ERROR_DEFAULT"
	FROZEN_AgentQueryRiskPromotionListV2DataDataPromotionStatus                       AgentQueryRiskPromotionListV2DataDataPromotionStatus = "FROZEN"
	LIVE_ROOM_OFF_AgentQueryRiskPromotionListV2DataDataPromotionStatus                AgentQueryRiskPromotionListV2DataDataPromotionStatus = "LIVE_ROOM_OFF"
	NO_SCHEDULE_AgentQueryRiskPromotionListV2DataDataPromotionStatus                  AgentQueryRiskPromotionListV2DataDataPromotionStatus = "NO_SCHEDULE"
	OFFLINE_AUDIT_AgentQueryRiskPromotionListV2DataDataPromotionStatus                AgentQueryRiskPromotionListV2DataDataPromotionStatus = "OFFLINE_AUDIT"
	PRE_ONLINE_AgentQueryRiskPromotionListV2DataDataPromotionStatus                   AgentQueryRiskPromotionListV2DataDataPromotionStatus = "PRE_ONLINE"
	PRODUCT_OFFLINE_AgentQueryRiskPromotionListV2DataDataPromotionStatus              AgentQueryRiskPromotionListV2DataDataPromotionStatus = "PRODUCT_OFFLINE"
	PROJECT_DISABLE_AgentQueryRiskPromotionListV2DataDataPromotionStatus              AgentQueryRiskPromotionListV2DataDataPromotionStatus = "PROJECT_DISABLE"
	PROJECT_OFFLINE_BUDGET_AgentQueryRiskPromotionListV2DataDataPromotionStatus       AgentQueryRiskPromotionListV2DataDataPromotionStatus = "PROJECT_OFFLINE_BUDGET"
	PROJECT_PRE_OFFLINE_BUDGET_AgentQueryRiskPromotionListV2DataDataPromotionStatus   AgentQueryRiskPromotionListV2DataDataPromotionStatus = "PROJECT_PRE_OFFLINE_BUDGET"
	PROMOTION_DISABLE_AgentQueryRiskPromotionListV2DataDataPromotionStatus            AgentQueryRiskPromotionListV2DataDataPromotionStatus = "PROMOTION_DISABLE"
	PROMOTION_OFFLINE_BALANCE_AgentQueryRiskPromotionListV2DataDataPromotionStatus    AgentQueryRiskPromotionListV2DataDataPromotionStatus = "PROMOTION_OFFLINE_BALANCE"
	PROMOTION_OFFLINE_BUDGET_AgentQueryRiskPromotionListV2DataDataPromotionStatus     AgentQueryRiskPromotionListV2DataDataPromotionStatus = "PROMOTION_OFFLINE_BUDGET"
	PROMOTION_PRE_OFFLINE_BUDGET_AgentQueryRiskPromotionListV2DataDataPromotionStatus AgentQueryRiskPromotionListV2DataDataPromotionStatus = "PROMOTION_PRE_OFFLINE_BUDGET"
	PROMOTION_QUOTA_LIMIT_AgentQueryRiskPromotionListV2DataDataPromotionStatus        AgentQueryRiskPromotionListV2DataDataPromotionStatus = "PROMOTION_QUOTA_LIMIT"
	RE_AUDIT_AgentQueryRiskPromotionListV2DataDataPromotionStatus                     AgentQueryRiskPromotionListV2DataDataPromotionStatus = "RE_AUDIT"
	TIME_DONE_AgentQueryRiskPromotionListV2DataDataPromotionStatus                    AgentQueryRiskPromotionListV2DataDataPromotionStatus = "TIME_DONE"
	TIME_NO_REACH_AgentQueryRiskPromotionListV2DataDataPromotionStatus                AgentQueryRiskPromotionListV2DataDataPromotionStatus = "TIME_NO_REACH"
)

// All allowed values of AgentQueryRiskPromotionListV2DataDataPromotionStatus enum
var AllowedAgentQueryRiskPromotionListV2DataDataPromotionStatusEnumValues = []AgentQueryRiskPromotionListV2DataDataPromotionStatus{
	"ADV_OFFLINE_BUDGET",
	"ADV_PRE_OFFLINE_BUDGET",
	"AUDIT",
	"AWEME_ACCOUNT_DISABLED",
	"AWEME_ANCHOR_DISABLED",
	"CREATE",
	"DELETE",
	"DELIVERY_OK",
	"ERROR_DEFAULT",
	"FROZEN",
	"LIVE_ROOM_OFF",
	"NO_SCHEDULE",
	"OFFLINE_AUDIT",
	"PRE_ONLINE",
	"PRODUCT_OFFLINE",
	"PROJECT_DISABLE",
	"PROJECT_OFFLINE_BUDGET",
	"PROJECT_PRE_OFFLINE_BUDGET",
	"PROMOTION_DISABLE",
	"PROMOTION_OFFLINE_BALANCE",
	"PROMOTION_OFFLINE_BUDGET",
	"PROMOTION_PRE_OFFLINE_BUDGET",
	"PROMOTION_QUOTA_LIMIT",
	"RE_AUDIT",
	"TIME_DONE",
	"TIME_NO_REACH",
}

// NewAgentQueryRiskPromotionListV2DataDataPromotionStatusFromValue returns a pointer to a valid AgentQueryRiskPromotionListV2DataDataPromotionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentQueryRiskPromotionListV2DataDataPromotionStatusFromValue(v string) (*AgentQueryRiskPromotionListV2DataDataPromotionStatus, error) {
	ev := AgentQueryRiskPromotionListV2DataDataPromotionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentQueryRiskPromotionListV2DataDataPromotionStatus: valid values are %v", v, AllowedAgentQueryRiskPromotionListV2DataDataPromotionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentQueryRiskPromotionListV2DataDataPromotionStatus) IsValid() bool {
	for _, existing := range AllowedAgentQueryRiskPromotionListV2DataDataPromotionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to agent_query_risk_promotion_list_v2_data_data_promotion_status value
func (v AgentQueryRiskPromotionListV2DataDataPromotionStatus) Ptr() *AgentQueryRiskPromotionListV2DataDataPromotionStatus {
	return &v
}
