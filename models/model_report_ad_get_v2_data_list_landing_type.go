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
	SHOP_ReportAdGetV2DataListLandingType      ReportAdGetV2DataListLandingType = "SHOP"
	LIVE_ReportAdGetV2DataListLandingType      ReportAdGetV2DataListLandingType = "LIVE"
	QUICK_APP_ReportAdGetV2DataListLandingType ReportAdGetV2DataListLandingType = "QUICK_APP"
	DPA_ReportAdGetV2DataListLandingType       ReportAdGetV2DataListLandingType = "DPA"
	APP_ReportAdGetV2DataListLandingType       ReportAdGetV2DataListLandingType = "APP"
	ARTICLE_ReportAdGetV2DataListLandingType   ReportAdGetV2DataListLandingType = "ARTICLE"
	AWEME_ReportAdGetV2DataListLandingType     ReportAdGetV2DataListLandingType = "AWEME"
	LINK_ReportAdGetV2DataListLandingType      ReportAdGetV2DataListLandingType = "LINK"
	STORE_ReportAdGetV2DataListLandingType     ReportAdGetV2DataListLandingType = "STORE"
	GOODS_ReportAdGetV2DataListLandingType     ReportAdGetV2DataListLandingType = "GOODS"
)

// Ptr returns reference to report_ad_get_v2_data_list_landing_type value
func (v ReportAdGetV2DataListLandingType) Ptr() *ReportAdGetV2DataListLandingType {
	return &v
}
