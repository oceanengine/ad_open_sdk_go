/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAgentGetV2V2FilteringAccountSource
type ReportAgentGetV2V2FilteringAccountSource string

// List of report_agent_get_v2_v2_filtering_account_source
const (
	LUBAN_ACCOUNT_ReportAgentGetV2V2FilteringAccountSource     ReportAgentGetV2V2FilteringAccountSource = "LUBAN_ACCOUNT"
	NORMAL_ADVERTISER_ReportAgentGetV2V2FilteringAccountSource ReportAgentGetV2V2FilteringAccountSource = "NORMAL_ADVERTISER"
)

// Ptr returns reference to report_agent_get_v2_v2_filtering_account_source value
func (v ReportAgentGetV2V2FilteringAccountSource) Ptr() *ReportAgentGetV2V2FilteringAccountSource {
	return &v
}
