/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2FilteringPricingCategories
type ReportAdGetV2FilteringPricingCategories string

// List of report_ad_get_v2_filtering_pricing_categories
const (
	PRICING_CATEGORY_NOC_ReportAdGetV2FilteringPricingCategories               ReportAdGetV2FilteringPricingCategories = "PRICING_CATEGORY_NOC"
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportAdGetV2FilteringPricingCategories ReportAdGetV2FilteringPricingCategories = "PRICING_CATEGORY_BRAND_AND_PRICING"
	PRICING_CATEGORY_FREE_ReportAdGetV2FilteringPricingCategories              ReportAdGetV2FilteringPricingCategories = "PRICING_CATEGORY_FREE"
	PRICING_CATEGORY_BID_ReportAdGetV2FilteringPricingCategories               ReportAdGetV2FilteringPricingCategories = "PRICING_CATEGORY_BID"
	PRICING_CATEGORY_BRAND_ReportAdGetV2FilteringPricingCategories             ReportAdGetV2FilteringPricingCategories = "PRICING_CATEGORY_BRAND"
)

// Ptr returns reference to report_ad_get_v2_filtering_pricing_categories value
func (v ReportAdGetV2FilteringPricingCategories) Ptr() *ReportAdGetV2FilteringPricingCategories {
	return &v
}
