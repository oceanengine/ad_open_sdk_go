/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceProvinceV2IdType
type ReportAudienceProvinceV2IdType string

// List of report_audience_province_v2_id_type
const (
	AUDIENCE_STAT_ID_TYPE_ADVERTISER_ReportAudienceProvinceV2IdType ReportAudienceProvinceV2IdType = "AUDIENCE_STAT_ID_TYPE_ADVERTISER"
	AUDIENCE_STAT_ID_TYPE_AD_ReportAudienceProvinceV2IdType         ReportAudienceProvinceV2IdType = "AUDIENCE_STAT_ID_TYPE_AD"
	AUDIENCE_STAT_ID_TYPE_CAMPAIGN_ReportAudienceProvinceV2IdType   ReportAudienceProvinceV2IdType = "AUDIENCE_STAT_ID_TYPE_CAMPAIGN"
)

// Ptr returns reference to report_audience_province_v2_id_type value
func (v ReportAudienceProvinceV2IdType) Ptr() *ReportAudienceProvinceV2IdType {
	return &v
}
