/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCreativeGetV2FilteringPricing
type ReportCreativeGetV2FilteringPricing string

// List of report_creative_get_v2_filtering_pricing
const (
	PRICING_CPM_ReportCreativeGetV2FilteringPricing      ReportCreativeGetV2FilteringPricing = "PRICING_CPM"
	PRICING_CPC_OCPM_ReportCreativeGetV2FilteringPricing ReportCreativeGetV2FilteringPricing = "PRICING_CPC_OCPM"
	PRICING_OCPM_ReportCreativeGetV2FilteringPricing     ReportCreativeGetV2FilteringPricing = "PRICING_OCPM"
	PRICING_CPA_ReportCreativeGetV2FilteringPricing      ReportCreativeGetV2FilteringPricing = "PRICING_CPA"
	PRICING_OCPC_ReportCreativeGetV2FilteringPricing     ReportCreativeGetV2FilteringPricing = "PRICING_OCPC"
	PRICING_CPV_ReportCreativeGetV2FilteringPricing      ReportCreativeGetV2FilteringPricing = "PRICING_CPV"
	PRICING_CPC_ReportCreativeGetV2FilteringPricing      ReportCreativeGetV2FilteringPricing = "PRICING_CPC"
)

// Ptr returns reference to report_creative_get_v2_filtering_pricing value
func (v ReportCreativeGetV2FilteringPricing) Ptr() *ReportCreativeGetV2FilteringPricing {
	return &v
}
