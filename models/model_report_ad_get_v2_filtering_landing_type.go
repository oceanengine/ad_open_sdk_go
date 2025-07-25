/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2FilteringLandingType
type ReportAdGetV2FilteringLandingType string

// List of report_ad_get_v2_filtering_landing_type
const (
	ARTICLE_ReportAdGetV2FilteringLandingType   ReportAdGetV2FilteringLandingType = "ARTICLE"
	DPA_ReportAdGetV2FilteringLandingType       ReportAdGetV2FilteringLandingType = "DPA"
	QUICK_APP_ReportAdGetV2FilteringLandingType ReportAdGetV2FilteringLandingType = "QUICK_APP"
	GOODS_ReportAdGetV2FilteringLandingType     ReportAdGetV2FilteringLandingType = "GOODS"
	STORE_ReportAdGetV2FilteringLandingType     ReportAdGetV2FilteringLandingType = "STORE"
	APP_ReportAdGetV2FilteringLandingType       ReportAdGetV2FilteringLandingType = "APP"
	AWEME_ReportAdGetV2FilteringLandingType     ReportAdGetV2FilteringLandingType = "AWEME"
	LINK_ReportAdGetV2FilteringLandingType      ReportAdGetV2FilteringLandingType = "LINK"
	SHOP_ReportAdGetV2FilteringLandingType      ReportAdGetV2FilteringLandingType = "SHOP"
	LIVE_ReportAdGetV2FilteringLandingType      ReportAdGetV2FilteringLandingType = "LIVE"
)

// Ptr returns reference to report_ad_get_v2_filtering_landing_type value
func (v ReportAdGetV2FilteringLandingType) Ptr() *ReportAdGetV2FilteringLandingType {
	return &v
}
