/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCampaignGetV2FilteringPricingCategories
type ReportCampaignGetV2FilteringPricingCategories string

// List of report_campaign_get_v2_filtering_pricing_categories
const (
	PRICING_CATEGORY_BRAND_ReportCampaignGetV2FilteringPricingCategories             ReportCampaignGetV2FilteringPricingCategories = "PRICING_CATEGORY_BRAND"
	PRICING_CATEGORY_FREE_ReportCampaignGetV2FilteringPricingCategories              ReportCampaignGetV2FilteringPricingCategories = "PRICING_CATEGORY_FREE"
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportCampaignGetV2FilteringPricingCategories ReportCampaignGetV2FilteringPricingCategories = "PRICING_CATEGORY_BRAND_AND_PRICING"
	PRICING_CATEGORY_BID_ReportCampaignGetV2FilteringPricingCategories               ReportCampaignGetV2FilteringPricingCategories = "PRICING_CATEGORY_BID"
	PRICING_CATEGORY_NOC_ReportCampaignGetV2FilteringPricingCategories               ReportCampaignGetV2FilteringPricingCategories = "PRICING_CATEGORY_NOC"
)

// Ptr returns reference to report_campaign_get_v2_filtering_pricing_categories value
func (v ReportCampaignGetV2FilteringPricingCategories) Ptr() *ReportCampaignGetV2FilteringPricingCategories {
	return &v
}
