/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAgentGetV2V2DataListAccountStatus
type ReportAgentGetV2V2DataListAccountStatus string

// List of report_agent_get_v2_v2_data_list_account_status
const (
	STATUS_CONFIRM_FAIL_END_ReportAgentGetV2V2DataListAccountStatus           ReportAgentGetV2V2DataListAccountStatus = "STATUS_CONFIRM_FAIL_END"
	STATUS_PUNISH_ReportAgentGetV2V2DataListAccountStatus                     ReportAgentGetV2V2DataListAccountStatus = "STATUS_PUNISH"
	STATUS_CONFIRM_MODIFY_FAIL_ReportAgentGetV2V2DataListAccountStatus        ReportAgentGetV2V2DataListAccountStatus = "STATUS_CONFIRM_MODIFY_FAIL"
	STATUS_SELF_SERVICE_UNAUDITED_ReportAgentGetV2V2DataListAccountStatus     ReportAgentGetV2V2DataListAccountStatus = "STATUS_SELF_SERVICE_UNAUDITED"
	STATUS_WAIT_FOR_PUBLIC_AUTH_ReportAgentGetV2V2DataListAccountStatus       ReportAgentGetV2V2DataListAccountStatus = "STATUS_WAIT_FOR_PUBLIC_AUTH"
	STATUS_WAIT_FOR_BPM_AUDIT_ReportAgentGetV2V2DataListAccountStatus         ReportAgentGetV2V2DataListAccountStatus = "STATUS_WAIT_FOR_BPM_AUDIT"
	STATUS_ENABLE_AND_AVATAR_AUDITING_ReportAgentGetV2V2DataListAccountStatus ReportAgentGetV2V2DataListAccountStatus = "STATUS_ENABLE_AND_AVATAR_AUDITING"
	STATUS_WAIT_FOR_BPM_FILE_CONTACT_ReportAgentGetV2V2DataListAccountStatus  ReportAgentGetV2V2DataListAccountStatus = "STATUS_WAIT_FOR_BPM_FILE_CONTACT"
	STATUS_PENDING_CONFIRM_ReportAgentGetV2V2DataListAccountStatus            ReportAgentGetV2V2DataListAccountStatus = "STATUS_PENDING_CONFIRM"
	STATUS_DISABLE_ReportAgentGetV2V2DataListAccountStatus                    ReportAgentGetV2V2DataListAccountStatus = "STATUS_DISABLE"
	STATUS_PENDING_VERIFIED_ReportAgentGetV2V2DataListAccountStatus           ReportAgentGetV2V2DataListAccountStatus = "STATUS_PENDING_VERIFIED"
	STATUS_PENDING_CONFIRM_MODIFY_ReportAgentGetV2V2DataListAccountStatus     ReportAgentGetV2V2DataListAccountStatus = "STATUS_PENDING_CONFIRM_MODIFY"
	STATUS_CONFIRM_FAIL_ReportAgentGetV2V2DataListAccountStatus               ReportAgentGetV2V2DataListAccountStatus = "STATUS_CONFIRM_FAIL"
	STATUS_WAIT_FOR_ACCOUNT_FEE_ReportAgentGetV2V2DataListAccountStatus       ReportAgentGetV2V2DataListAccountStatus = "STATUS_WAIT_FOR_ACCOUNT_FEE"
	STATUS_ENABLE_ReportAgentGetV2V2DataListAccountStatus                     ReportAgentGetV2V2DataListAccountStatus = "STATUS_ENABLE"
)

// Ptr returns reference to report_agent_get_v2_v2_data_list_account_status value
func (v ReportAgentGetV2V2DataListAccountStatus) Ptr() *ReportAgentGetV2V2DataListAccountStatus {
	return &v
}
