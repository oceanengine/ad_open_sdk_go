/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCampaignGetV2FilteringStatus
type ReportCampaignGetV2FilteringStatus string

// List of report_campaign_get_v2_filtering_status
const (
	CAMPAIGN_STATUS_ADVERTISER_BUDGET_EXCEED_ReportCampaignGetV2FilteringStatus ReportCampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_ADVERTISER_BUDGET_EXCEED"
	CAMPAIGN_STATUS_ALL_ReportCampaignGetV2FilteringStatus                      ReportCampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_ALL"
	CAMPAIGN_STATUS_ENABLE_ReportCampaignGetV2FilteringStatus                   ReportCampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_ENABLE"
	CAMPAIGN_STATUS_DELETE_ReportCampaignGetV2FilteringStatus                   ReportCampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_DELETE"
	CAMPAIGN_STATUS_NOT_DELETE_ReportCampaignGetV2FilteringStatus               ReportCampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_NOT_DELETE"
	CAMPAIGN_STATUS_DISABLE_ReportCampaignGetV2FilteringStatus                  ReportCampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_DISABLE"
)

// Ptr returns reference to report_campaign_get_v2_filtering_status value
func (v ReportCampaignGetV2FilteringStatus) Ptr() *ReportCampaignGetV2FilteringStatus {
	return &v
}
