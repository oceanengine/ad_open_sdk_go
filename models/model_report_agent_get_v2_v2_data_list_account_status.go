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

// ReportAgentGetV2V2DataListAccountStatus
type ReportAgentGetV2V2DataListAccountStatus string

// List of report_agent_get_v2_v2_data_list_account_status
const (
	STATUS_SELF_SERVICE_UNAUDITED_ReportAgentGetV2V2DataListAccountStatus     ReportAgentGetV2V2DataListAccountStatus = "STATUS_SELF_SERVICE_UNAUDITED"
	STATUS_ENABLE_ReportAgentGetV2V2DataListAccountStatus                     ReportAgentGetV2V2DataListAccountStatus = "STATUS_ENABLE"
	STATUS_WAIT_FOR_BPM_AUDIT_ReportAgentGetV2V2DataListAccountStatus         ReportAgentGetV2V2DataListAccountStatus = "STATUS_WAIT_FOR_BPM_AUDIT"
	STATUS_WAIT_FOR_PUBLIC_AUTH_ReportAgentGetV2V2DataListAccountStatus       ReportAgentGetV2V2DataListAccountStatus = "STATUS_WAIT_FOR_PUBLIC_AUTH"
	STATUS_PENDING_VERIFIED_ReportAgentGetV2V2DataListAccountStatus           ReportAgentGetV2V2DataListAccountStatus = "STATUS_PENDING_VERIFIED"
	STATUS_WAIT_FOR_ACCOUNT_FEE_ReportAgentGetV2V2DataListAccountStatus       ReportAgentGetV2V2DataListAccountStatus = "STATUS_WAIT_FOR_ACCOUNT_FEE"
	STATUS_ENABLE_AND_AVATAR_AUDITING_ReportAgentGetV2V2DataListAccountStatus ReportAgentGetV2V2DataListAccountStatus = "STATUS_ENABLE_AND_AVATAR_AUDITING"
	STATUS_PENDING_CONFIRM_ReportAgentGetV2V2DataListAccountStatus            ReportAgentGetV2V2DataListAccountStatus = "STATUS_PENDING_CONFIRM"
	STATUS_PENDING_CONFIRM_MODIFY_ReportAgentGetV2V2DataListAccountStatus     ReportAgentGetV2V2DataListAccountStatus = "STATUS_PENDING_CONFIRM_MODIFY"
	STATUS_WAIT_FOR_BPM_FILE_CONTACT_ReportAgentGetV2V2DataListAccountStatus  ReportAgentGetV2V2DataListAccountStatus = "STATUS_WAIT_FOR_BPM_FILE_CONTACT"
	STATUS_CONFIRM_FAIL_ReportAgentGetV2V2DataListAccountStatus               ReportAgentGetV2V2DataListAccountStatus = "STATUS_CONFIRM_FAIL"
	STATUS_CONFIRM_MODIFY_FAIL_ReportAgentGetV2V2DataListAccountStatus        ReportAgentGetV2V2DataListAccountStatus = "STATUS_CONFIRM_MODIFY_FAIL"
	STATUS_PUNISH_ReportAgentGetV2V2DataListAccountStatus                     ReportAgentGetV2V2DataListAccountStatus = "STATUS_PUNISH"
	STATUS_DISABLE_ReportAgentGetV2V2DataListAccountStatus                    ReportAgentGetV2V2DataListAccountStatus = "STATUS_DISABLE"
	STATUS_CONFIRM_FAIL_END_ReportAgentGetV2V2DataListAccountStatus           ReportAgentGetV2V2DataListAccountStatus = "STATUS_CONFIRM_FAIL_END"
)

// All allowed values of ReportAgentGetV2V2DataListAccountStatus enum
var AllowedReportAgentGetV2V2DataListAccountStatusEnumValues = []ReportAgentGetV2V2DataListAccountStatus{
	"STATUS_SELF_SERVICE_UNAUDITED",
	"STATUS_ENABLE",
	"STATUS_WAIT_FOR_BPM_AUDIT",
	"STATUS_WAIT_FOR_PUBLIC_AUTH",
	"STATUS_PENDING_VERIFIED",
	"STATUS_WAIT_FOR_ACCOUNT_FEE",
	"STATUS_ENABLE_AND_AVATAR_AUDITING",
	"STATUS_PENDING_CONFIRM",
	"STATUS_PENDING_CONFIRM_MODIFY",
	"STATUS_WAIT_FOR_BPM_FILE_CONTACT",
	"STATUS_CONFIRM_FAIL",
	"STATUS_CONFIRM_MODIFY_FAIL",
	"STATUS_PUNISH",
	"STATUS_DISABLE",
	"STATUS_CONFIRM_FAIL_END",
}

// NewReportAgentGetV2V2DataListAccountStatusFromValue returns a pointer to a valid ReportAgentGetV2V2DataListAccountStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAgentGetV2V2DataListAccountStatusFromValue(v string) (*ReportAgentGetV2V2DataListAccountStatus, error) {
	ev := ReportAgentGetV2V2DataListAccountStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAgentGetV2V2DataListAccountStatus: valid values are %v", v, AllowedReportAgentGetV2V2DataListAccountStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAgentGetV2V2DataListAccountStatus) IsValid() bool {
	for _, existing := range AllowedReportAgentGetV2V2DataListAccountStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_agent_get_v2_v2_data_list_account_status value
func (v ReportAgentGetV2V2DataListAccountStatus) Ptr() *ReportAgentGetV2V2DataListAccountStatus {
	return &v
}
