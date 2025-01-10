/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCreativeGetV2DataListPricingCategory
type ReportCreativeGetV2DataListPricingCategory string

// List of report_creative_get_v2_data_list_pricing_category
const (
	PRICING_CATEGORY_NOC_ReportCreativeGetV2DataListPricingCategory               ReportCreativeGetV2DataListPricingCategory = "PRICING_CATEGORY_NOC"
	PRICING_CATEGORY_BRAND_ReportCreativeGetV2DataListPricingCategory             ReportCreativeGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND"
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportCreativeGetV2DataListPricingCategory ReportCreativeGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND_AND_PRICING"
	PRICING_CATEGORY_BID_ReportCreativeGetV2DataListPricingCategory               ReportCreativeGetV2DataListPricingCategory = "PRICING_CATEGORY_BID"
	PRICING_CATEGORY_FREE_ReportCreativeGetV2DataListPricingCategory              ReportCreativeGetV2DataListPricingCategory = "PRICING_CATEGORY_FREE"
)

// Ptr returns reference to report_creative_get_v2_data_list_pricing_category value
func (v ReportCreativeGetV2DataListPricingCategory) Ptr() *ReportCreativeGetV2DataListPricingCategory {
	return &v
}
