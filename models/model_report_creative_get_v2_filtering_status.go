/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCreativeGetV2FilteringStatus
type ReportCreativeGetV2FilteringStatus string

// List of report_creative_get_v2_filtering_status
const (
	CREATIVE_STATUS_REAUDIT_ReportCreativeGetV2FilteringStatus          ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_REAUDIT"
	CREATIVE_STATUS_AD_AUDIT_DENY_ReportCreativeGetV2FilteringStatus    ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_AD_AUDIT_DENY"
	CREATIVE_STATUS_BUDGET_EXCEED_ReportCreativeGetV2FilteringStatus    ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_BUDGET_EXCEED"
	CREATIVE_STATUS_CAMPAIGN_DISABLE_ReportCreativeGetV2FilteringStatus ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_CAMPAIGN_DISABLE"
	CREATIVE_STATUS_ALL_ReportCreativeGetV2FilteringStatus              ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_ALL"
	CREATIVE_STATUS_DELIVERY_OK_ReportCreativeGetV2FilteringStatus      ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_DELIVERY_OK"
	CREATIVE_STATUS_DONE_ReportCreativeGetV2FilteringStatus             ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_DONE"
	CREATIVE_STATUS_AUDIT_DENY_ReportCreativeGetV2FilteringStatus       ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_AUDIT_DENY"
	CREATIVE_STATUS_NO_SCHEDULE_ReportCreativeGetV2FilteringStatus      ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_NO_SCHEDULE"
	CREATIVE_STATUS_DISABLE_ReportCreativeGetV2FilteringStatus          ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_DISABLE"
	CREATIVE_STATUS_AD_REAUDIT_ReportCreativeGetV2FilteringStatus       ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_AD_REAUDIT"
	CREATIVE_STATUS_BALANCE_EXCEED_ReportCreativeGetV2FilteringStatus   ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_BALANCE_EXCEED"
	CREATIVE_STATUS_NOT_DELETE_ReportCreativeGetV2FilteringStatus       ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_NOT_DELETE"
	CREATIVE_STATUS_AUDIT_ReportCreativeGetV2FilteringStatus            ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_AUDIT"
	CREATIVE_STATUS_NOT_START_ReportCreativeGetV2FilteringStatus        ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_NOT_START"
	CREATIVE_STATUS_DELETE_ReportCreativeGetV2FilteringStatus           ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_DELETE"
	CREATIVE_STATUS_CAMPAIGN_EXCEED_ReportCreativeGetV2FilteringStatus  ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_CAMPAIGN_EXCEED"
	CREATIVE_STATUS_AD_DISABLE_ReportCreativeGetV2FilteringStatus       ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_AD_DISABLE"
	CREATIVE_STATUS_AD_AUDIT_ReportCreativeGetV2FilteringStatus         ReportCreativeGetV2FilteringStatus = "CREATIVE_STATUS_AD_AUDIT"
)

// Ptr returns reference to report_creative_get_v2_filtering_status value
func (v ReportCreativeGetV2FilteringStatus) Ptr() *ReportCreativeGetV2FilteringStatus {
	return &v
}
