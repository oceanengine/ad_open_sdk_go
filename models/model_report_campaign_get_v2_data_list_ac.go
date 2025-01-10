/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCampaignGetV2DataListAc
type ReportCampaignGetV2DataListAc string

// List of report_campaign_get_v2_data_list_ac
const (
	Enum_5_G_ReportCampaignGetV2DataListAc ReportCampaignGetV2DataListAc = "5G"
	Enum_3_G_ReportCampaignGetV2DataListAc ReportCampaignGetV2DataListAc = "3G"
	UNKNOWN_ReportCampaignGetV2DataListAc  ReportCampaignGetV2DataListAc = "unknown"
	Enum_4_G_ReportCampaignGetV2DataListAc ReportCampaignGetV2DataListAc = "4G"
	WIFI_ReportCampaignGetV2DataListAc     ReportCampaignGetV2DataListAc = "WIFI"
	Enum_2_G_ReportCampaignGetV2DataListAc ReportCampaignGetV2DataListAc = "2G"
)

// Ptr returns reference to report_campaign_get_v2_data_list_ac value
func (v ReportCampaignGetV2DataListAc) Ptr() *ReportCampaignGetV2DataListAc {
	return &v
}
