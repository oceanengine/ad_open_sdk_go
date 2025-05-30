/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvBiddingListQueryV2FilteringAccountStatus
type AgentAdvBiddingListQueryV2FilteringAccountStatus string

// List of agent_adv_bidding_list_query_v2_filtering_account_status
const (
	STATUS_DISABLE_AgentAdvBiddingListQueryV2FilteringAccountStatus                    AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_DISABLE"
	STATUS_PENDING_CONFIRM_AgentAdvBiddingListQueryV2FilteringAccountStatus            AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_PENDING_CONFIRM"
	STATUS_PENDING_VERIFIED_AgentAdvBiddingListQueryV2FilteringAccountStatus           AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_PENDING_VERIFIED"
	STATUS_CONFIRM_FAIL_AgentAdvBiddingListQueryV2FilteringAccountStatus               AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_CONFIRM_FAIL"
	STATUS_ENABLE_AgentAdvBiddingListQueryV2FilteringAccountStatus                     AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_ENABLE"
	STATUS_CONFIRM_FAIL_END_AgentAdvBiddingListQueryV2FilteringAccountStatus           AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_CONFIRM_FAIL_END"
	STATUS_PENDING_CONFIRM_MODIFY_AgentAdvBiddingListQueryV2FilteringAccountStatus     AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_PENDING_CONFIRM_MODIFY"
	STATUS_CONFIRM_MODIFY_FAIL_AgentAdvBiddingListQueryV2FilteringAccountStatus        AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_CONFIRM_MODIFY_FAIL"
	STATUS_PUNISH_AgentAdvBiddingListQueryV2FilteringAccountStatus                     AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_PUNISH"
	STATUS_WAIT_FOR_BPM_AUDIT_AgentAdvBiddingListQueryV2FilteringAccountStatus         AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_WAIT_FOR_BPM_AUDIT"
	STATUS_SELF_SERVICE_UNAUDITED_AgentAdvBiddingListQueryV2FilteringAccountStatus     AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_SELF_SERVICE_UNAUDITED"
	STATUS_ENABLE_AND_AVATAR_AUDITING_AgentAdvBiddingListQueryV2FilteringAccountStatus AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_ENABLE_AND_AVATAR_AUDITING"
	STATUS_WAIT_FOR_BPM_FILE_CONTACT_AgentAdvBiddingListQueryV2FilteringAccountStatus  AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_WAIT_FOR_BPM_FILE_CONTACT"
	STATUS_WAIT_FOR_ACCOUNT_FEE_AgentAdvBiddingListQueryV2FilteringAccountStatus       AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_WAIT_FOR_ACCOUNT_FEE"
	STATUS_WAIT_FOR_PUBLIC_AUTH_AgentAdvBiddingListQueryV2FilteringAccountStatus       AgentAdvBiddingListQueryV2FilteringAccountStatus = "STATUS_WAIT_FOR_PUBLIC_AUTH"
)

// Ptr returns reference to agent_adv_bidding_list_query_v2_filtering_account_status value
func (v AgentAdvBiddingListQueryV2FilteringAccountStatus) Ptr() *AgentAdvBiddingListQueryV2FilteringAccountStatus {
	return &v
}
