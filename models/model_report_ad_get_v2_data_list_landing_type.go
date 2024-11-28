/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2DataListLandingType
type ReportAdGetV2DataListLandingType string

// List of report_ad_get_v2_data_list_landing_type
const (
	LIVE_ReportAdGetV2DataListLandingType      ReportAdGetV2DataListLandingType = "LIVE"
	STORE_ReportAdGetV2DataListLandingType     ReportAdGetV2DataListLandingType = "STORE"
	SHOP_ReportAdGetV2DataListLandingType      ReportAdGetV2DataListLandingType = "SHOP"
	ARTICLE_ReportAdGetV2DataListLandingType   ReportAdGetV2DataListLandingType = "ARTICLE"
	GOODS_ReportAdGetV2DataListLandingType     ReportAdGetV2DataListLandingType = "GOODS"
	LINK_ReportAdGetV2DataListLandingType      ReportAdGetV2DataListLandingType = "LINK"
	AWEME_ReportAdGetV2DataListLandingType     ReportAdGetV2DataListLandingType = "AWEME"
	QUICK_APP_ReportAdGetV2DataListLandingType ReportAdGetV2DataListLandingType = "QUICK_APP"
	APP_ReportAdGetV2DataListLandingType       ReportAdGetV2DataListLandingType = "APP"
	DPA_ReportAdGetV2DataListLandingType       ReportAdGetV2DataListLandingType = "DPA"
)

// Ptr returns reference to report_ad_get_v2_data_list_landing_type value
func (v ReportAdGetV2DataListLandingType) Ptr() *ReportAdGetV2DataListLandingType {
	return &v
}
