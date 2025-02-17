/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportRtaGetV2ReportType
type ReportRtaGetV2ReportType string

// List of report_rta_get_v2_report_type
const (
	CREATIVE_RTA_PLATFORM_ReportRtaGetV2ReportType   ReportRtaGetV2ReportType = "CREATIVE_RTA_PLATFORM"
	CAMPAIGN_RTA_ReportRtaGetV2ReportType            ReportRtaGetV2ReportType = "CAMPAIGN_RTA"
	ADVERTISER_RTA_ReportRtaGetV2ReportType          ReportRtaGetV2ReportType = "ADVERTISER_RTA"
	CAMPAIGN_RTA_PLATFORM_ReportRtaGetV2ReportType   ReportRtaGetV2ReportType = "CAMPAIGN_RTA_PLATFORM"
	ADVERTISER_RTA_PLATFORM_ReportRtaGetV2ReportType ReportRtaGetV2ReportType = "ADVERTISER_RTA_PLATFORM"
	AD_RTA_PLATFORM_ReportRtaGetV2ReportType         ReportRtaGetV2ReportType = "AD_RTA_PLATFORM"
	CREATIVE_RTA_ReportRtaGetV2ReportType            ReportRtaGetV2ReportType = "CREATIVE_RTA"
	AD_RTA_ReportRtaGetV2ReportType                  ReportRtaGetV2ReportType = "AD_RTA"
)

// Ptr returns reference to report_rta_get_v2_report_type value
func (v ReportRtaGetV2ReportType) Ptr() *ReportRtaGetV2ReportType {
	return &v
}
