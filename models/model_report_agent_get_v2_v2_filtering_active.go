/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAgentGetV2V2FilteringActive
type ReportAgentGetV2V2FilteringActive string

// List of report_agent_get_v2_v2_filtering_active
const (
	SILENT_ReportAgentGetV2V2FilteringActive ReportAgentGetV2V2FilteringActive = "SILENT"
	ACTIVE_ReportAgentGetV2V2FilteringActive ReportAgentGetV2V2FilteringActive = "ACTIVE"
	ALL_ReportAgentGetV2V2FilteringActive    ReportAgentGetV2V2FilteringActive = "ALL"
)

// Ptr returns reference to report_agent_get_v2_v2_filtering_active value
func (v ReportAgentGetV2V2FilteringActive) Ptr() *ReportAgentGetV2V2FilteringActive {
	return &v
}
