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

// ReportAgentGetV2V2DataListCustomerType
type ReportAgentGetV2V2DataListCustomerType string

// List of report_agent_get_v2_v2_data_list_customer_type
const (
	SELF_SERVICE_ReportAgentGetV2V2DataListCustomerType    ReportAgentGetV2V2DataListCustomerType = "SELF_SERVICE"
	DSP_ReportAgentGetV2V2DataListCustomerType             ReportAgentGetV2V2DataListCustomerType = "DSP"
	KA_AGENT_ReportAgentGetV2V2DataListCustomerType        ReportAgentGetV2V2DataListCustomerType = "KA_AGENT"
	INTERNAL_ReportAgentGetV2V2DataListCustomerType        ReportAgentGetV2V2DataListCustomerType = "INTERNAL"
	INTERNET_ReportAgentGetV2V2DataListCustomerType        ReportAgentGetV2V2DataListCustomerType = "INTERNET"
	EXCHANGE_ReportAgentGetV2V2DataListCustomerType        ReportAgentGetV2V2DataListCustomerType = "EXCHANGE"
	VIRTUAL_LA_ReportAgentGetV2V2DataListCustomerType      ReportAgentGetV2V2DataListCustomerType = "VIRTUAL_LA"
	GAME_ReportAgentGetV2V2DataListCustomerType            ReportAgentGetV2V2DataListCustomerType = "GAME"
	VIRTUAL_KA_ReportAgentGetV2V2DataListCustomerType      ReportAgentGetV2V2DataListCustomerType = "VIRTUAL_KA"
	BRANCH_ReportAgentGetV2V2DataListCustomerType          ReportAgentGetV2V2DataListCustomerType = "BRANCH"
	SMB_ReportAgentGetV2V2DataListCustomerType             ReportAgentGetV2V2DataListCustomerType = "SMB"
	VERTICAL_ReportAgentGetV2V2DataListCustomerType        ReportAgentGetV2V2DataListCustomerType = "VERTICAL"
	I18_N_AGENT_ReportAgentGetV2V2DataListCustomerType     ReportAgentGetV2V2DataListCustomerType = "I18N_AGENT"
	OVERSEA_VIRTUAL_ReportAgentGetV2V2DataListCustomerType ReportAgentGetV2V2DataListCustomerType = "OVERSEA_VIRTUAL"
	I18_N_ADV_ReportAgentGetV2V2DataListCustomerType       ReportAgentGetV2V2DataListCustomerType = "I18N_ADV"
	LOCAL_ReportAgentGetV2V2DataListCustomerType           ReportAgentGetV2V2DataListCustomerType = "LOCAL"
	REGION_ReportAgentGetV2V2DataListCustomerType          ReportAgentGetV2V2DataListCustomerType = "REGION"
	VIRTUAL_CUS_ReportAgentGetV2V2DataListCustomerType     ReportAgentGetV2V2DataListCustomerType = "VIRTUAL_CUS"
	LA_ReportAgentGetV2V2DataListCustomerType              ReportAgentGetV2V2DataListCustomerType = "LA"
	KA_ReportAgentGetV2V2DataListCustomerType              ReportAgentGetV2V2DataListCustomerType = "KA"
	VIRTUAL_SMB_ReportAgentGetV2V2DataListCustomerType     ReportAgentGetV2V2DataListCustomerType = "VIRTUAL_SMB"
	SMB_AGENT_ReportAgentGetV2V2DataListCustomerType       ReportAgentGetV2V2DataListCustomerType = "SMB_AGENT"
	LA_AGENT_ReportAgentGetV2V2DataListCustomerType        ReportAgentGetV2V2DataListCustomerType = "LA_AGENT"
)

// All allowed values of ReportAgentGetV2V2DataListCustomerType enum
var AllowedReportAgentGetV2V2DataListCustomerTypeEnumValues = []ReportAgentGetV2V2DataListCustomerType{
	"SELF_SERVICE",
	"DSP",
	"KA_AGENT",
	"INTERNAL",
	"INTERNET",
	"EXCHANGE",
	"VIRTUAL_LA",
	"GAME",
	"VIRTUAL_KA",
	"BRANCH",
	"SMB",
	"VERTICAL",
	"I18N_AGENT",
	"OVERSEA_VIRTUAL",
	"I18N_ADV",
	"LOCAL",
	"REGION",
	"VIRTUAL_CUS",
	"LA",
	"KA",
	"VIRTUAL_SMB",
	"SMB_AGENT",
	"LA_AGENT",
}

// NewReportAgentGetV2V2DataListCustomerTypeFromValue returns a pointer to a valid ReportAgentGetV2V2DataListCustomerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAgentGetV2V2DataListCustomerTypeFromValue(v string) (*ReportAgentGetV2V2DataListCustomerType, error) {
	ev := ReportAgentGetV2V2DataListCustomerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAgentGetV2V2DataListCustomerType: valid values are %v", v, AllowedReportAgentGetV2V2DataListCustomerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAgentGetV2V2DataListCustomerType) IsValid() bool {
	for _, existing := range AllowedReportAgentGetV2V2DataListCustomerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_agent_get_v2_v2_data_list_customer_type value
func (v ReportAgentGetV2V2DataListCustomerType) Ptr() *ReportAgentGetV2V2DataListCustomerType {
	return &v
}
