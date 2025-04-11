/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2DataListPricingCategory
type ReportAdGetV2DataListPricingCategory string

// List of report_ad_get_v2_data_list_pricing_category
const (
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportAdGetV2DataListPricingCategory ReportAdGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND_AND_PRICING"
	PRICING_CATEGORY_BRAND_ReportAdGetV2DataListPricingCategory             ReportAdGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND"
	PRICING_CATEGORY_NOC_ReportAdGetV2DataListPricingCategory               ReportAdGetV2DataListPricingCategory = "PRICING_CATEGORY_NOC"
	PRICING_CATEGORY_FREE_ReportAdGetV2DataListPricingCategory              ReportAdGetV2DataListPricingCategory = "PRICING_CATEGORY_FREE"
	PRICING_CATEGORY_BID_ReportAdGetV2DataListPricingCategory               ReportAdGetV2DataListPricingCategory = "PRICING_CATEGORY_BID"
)

// Ptr returns reference to report_ad_get_v2_data_list_pricing_category value
func (v ReportAdGetV2DataListPricingCategory) Ptr() *ReportAdGetV2DataListPricingCategory {
	return &v
}
