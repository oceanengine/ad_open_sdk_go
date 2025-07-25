/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2FilteringStatus
type ReportAdGetV2FilteringStatus string

// List of report_ad_get_v2_filtering_status
const (
	AD_STATUS_CREATE_ReportAdGetV2FilteringStatus           ReportAdGetV2FilteringStatus = "AD_STATUS_CREATE"
	AD_STATUS_AUDIT_DENY_ReportAdGetV2FilteringStatus       ReportAdGetV2FilteringStatus = "AD_STATUS_AUDIT_DENY"
	AD_STATUS_NO_SCHEDULE_ReportAdGetV2FilteringStatus      ReportAdGetV2FilteringStatus = "AD_STATUS_NO_SCHEDULE"
	AD_STATUS_DONE_ReportAdGetV2FilteringStatus             ReportAdGetV2FilteringStatus = "AD_STATUS_DONE"
	AD_STATUS_CAMPAIGN_DISABLE_ReportAdGetV2FilteringStatus ReportAdGetV2FilteringStatus = "AD_STATUS_CAMPAIGN_DISABLE"
	AD_STATUS_AUDIT_ReportAdGetV2FilteringStatus            ReportAdGetV2FilteringStatus = "AD_STATUS_AUDIT"
	AD_STATUS_NOT_START_ReportAdGetV2FilteringStatus        ReportAdGetV2FilteringStatus = "AD_STATUS_NOT_START"
	AD_STATUS_ALL_ReportAdGetV2FilteringStatus              ReportAdGetV2FilteringStatus = "AD_STATUS_ALL"
	AD_STATUS_NOT_DELETE_ReportAdGetV2FilteringStatus       ReportAdGetV2FilteringStatus = "AD_STATUS_NOT_DELETE"
	AD_STATUS_CAMPAIGN_EXCEED_ReportAdGetV2FilteringStatus  ReportAdGetV2FilteringStatus = "AD_STATUS_CAMPAIGN_EXCEED"
	AD_STATUS_DELETE_ReportAdGetV2FilteringStatus           ReportAdGetV2FilteringStatus = "AD_STATUS_DELETE"
	AD_STATUS_BUDGET_EXCEED_ReportAdGetV2FilteringStatus    ReportAdGetV2FilteringStatus = "AD_STATUS_BUDGET_EXCEED"
	AD_STATUS_REAUDIT_ReportAdGetV2FilteringStatus          ReportAdGetV2FilteringStatus = "AD_STATUS_REAUDIT"
	AD_STATUS_DISABLE_ReportAdGetV2FilteringStatus          ReportAdGetV2FilteringStatus = "AD_STATUS_DISABLE"
	AD_STATUS_BALANCE_EXCEED_ReportAdGetV2FilteringStatus   ReportAdGetV2FilteringStatus = "AD_STATUS_BALANCE_EXCEED"
	AD_STATUS_DELIVERY_OK_ReportAdGetV2FilteringStatus      ReportAdGetV2FilteringStatus = "AD_STATUS_DELIVERY_OK"
)

// Ptr returns reference to report_ad_get_v2_filtering_status value
func (v ReportAdGetV2FilteringStatus) Ptr() *ReportAdGetV2FilteringStatus {
	return &v
}
