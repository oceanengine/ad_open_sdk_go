/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2FilteringPricing
type ReportAdGetV2FilteringPricing string

// List of report_ad_get_v2_filtering_pricing
const (
	PRICING_OCPM_ReportAdGetV2FilteringPricing ReportAdGetV2FilteringPricing = "PRICING_OCPM"
	PRICING_OCPC_ReportAdGetV2FilteringPricing ReportAdGetV2FilteringPricing = "PRICING_OCPC"
	PRICING_ECPC_ReportAdGetV2FilteringPricing ReportAdGetV2FilteringPricing = "PRICING_ECPC"
	PRICING_CPV_ReportAdGetV2FilteringPricing  ReportAdGetV2FilteringPricing = "PRICING_CPV"
	PRICING_CPM_ReportAdGetV2FilteringPricing  ReportAdGetV2FilteringPricing = "PRICING_CPM"
	PRICING_CPC_ReportAdGetV2FilteringPricing  ReportAdGetV2FilteringPricing = "PRICING_CPC"
	PRICING_CPA_ReportAdGetV2FilteringPricing  ReportAdGetV2FilteringPricing = "PRICING_CPA"
)

// Ptr returns reference to report_ad_get_v2_filtering_pricing value
func (v ReportAdGetV2FilteringPricing) Ptr() *ReportAdGetV2FilteringPricing {
	return &v
}
