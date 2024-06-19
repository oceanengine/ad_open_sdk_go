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

// ReportAgentGetV2V2DataListCustomerType
type ReportAgentGetV2V2DataListCustomerType string

// List of report_agent_get_v2_v2_data_list_customer_type
const (
	OVERSEA_VIRTUAL_ReportAgentGetV2V2DataListCustomerType ReportAgentGetV2V2DataListCustomerType = "OVERSEA_VIRTUAL"
	BRANCH_ReportAgentGetV2V2DataListCustomerType          ReportAgentGetV2V2DataListCustomerType = "BRANCH"
	INTERNAL_ReportAgentGetV2V2DataListCustomerType        ReportAgentGetV2V2DataListCustomerType = "INTERNAL"
	LA_ReportAgentGetV2V2DataListCustomerType              ReportAgentGetV2V2DataListCustomerType = "LA"
	EXCHANGE_ReportAgentGetV2V2DataListCustomerType        ReportAgentGetV2V2DataListCustomerType = "EXCHANGE"
	I18_N_ADV_ReportAgentGetV2V2DataListCustomerType       ReportAgentGetV2V2DataListCustomerType = "I18N_ADV"
	VIRTUAL_LA_ReportAgentGetV2V2DataListCustomerType      ReportAgentGetV2V2DataListCustomerType = "VIRTUAL_LA"
	I18_N_AGENT_ReportAgentGetV2V2DataListCustomerType     ReportAgentGetV2V2DataListCustomerType = "I18N_AGENT"
	LA_AGENT_ReportAgentGetV2V2DataListCustomerType        ReportAgentGetV2V2DataListCustomerType = "LA_AGENT"
	KA_AGENT_ReportAgentGetV2V2DataListCustomerType        ReportAgentGetV2V2DataListCustomerType = "KA_AGENT"
	REGION_ReportAgentGetV2V2DataListCustomerType          ReportAgentGetV2V2DataListCustomerType = "REGION"
	INTERNET_ReportAgentGetV2V2DataListCustomerType        ReportAgentGetV2V2DataListCustomerType = "INTERNET"
	KA_ReportAgentGetV2V2DataListCustomerType              ReportAgentGetV2V2DataListCustomerType = "KA"
	VIRTUAL_KA_ReportAgentGetV2V2DataListCustomerType      ReportAgentGetV2V2DataListCustomerType = "VIRTUAL_KA"
	SELF_SERVICE_ReportAgentGetV2V2DataListCustomerType    ReportAgentGetV2V2DataListCustomerType = "SELF_SERVICE"
	GAME_ReportAgentGetV2V2DataListCustomerType            ReportAgentGetV2V2DataListCustomerType = "GAME"
	DSP_ReportAgentGetV2V2DataListCustomerType             ReportAgentGetV2V2DataListCustomerType = "DSP"
	VERTICAL_ReportAgentGetV2V2DataListCustomerType        ReportAgentGetV2V2DataListCustomerType = "VERTICAL"
	SMB_ReportAgentGetV2V2DataListCustomerType             ReportAgentGetV2V2DataListCustomerType = "SMB"
	VIRTUAL_SMB_ReportAgentGetV2V2DataListCustomerType     ReportAgentGetV2V2DataListCustomerType = "VIRTUAL_SMB"
	SMB_AGENT_ReportAgentGetV2V2DataListCustomerType       ReportAgentGetV2V2DataListCustomerType = "SMB_AGENT"
	LOCAL_ReportAgentGetV2V2DataListCustomerType           ReportAgentGetV2V2DataListCustomerType = "LOCAL"
	VIRTUAL_CUS_ReportAgentGetV2V2DataListCustomerType     ReportAgentGetV2V2DataListCustomerType = "VIRTUAL_CUS"
)

// All allowed values of ReportAgentGetV2V2DataListCustomerType enum
var AllowedReportAgentGetV2V2DataListCustomerTypeEnumValues = []ReportAgentGetV2V2DataListCustomerType{
	"OVERSEA_VIRTUAL",
	"BRANCH",
	"INTERNAL",
	"LA",
	"EXCHANGE",
	"I18N_ADV",
	"VIRTUAL_LA",
	"I18N_AGENT",
	"LA_AGENT",
	"KA_AGENT",
	"REGION",
	"INTERNET",
	"KA",
	"VIRTUAL_KA",
	"SELF_SERVICE",
	"GAME",
	"DSP",
	"VERTICAL",
	"SMB",
	"VIRTUAL_SMB",
	"SMB_AGENT",
	"LOCAL",
	"VIRTUAL_CUS",
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
