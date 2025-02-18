/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdvertiserGetV2DataListPricing
type ReportAdvertiserGetV2DataListPricing string

// List of report_advertiser_get_v2_data_list_pricing
const (
	PRICING_CPM_ReportAdvertiserGetV2DataListPricing  ReportAdvertiserGetV2DataListPricing = "PRICING_CPM"
	PRICING_CPC_ReportAdvertiserGetV2DataListPricing  ReportAdvertiserGetV2DataListPricing = "PRICING_CPC"
	PRICING_CPV_ReportAdvertiserGetV2DataListPricing  ReportAdvertiserGetV2DataListPricing = "PRICING_CPV"
	PRICING_OCPC_ReportAdvertiserGetV2DataListPricing ReportAdvertiserGetV2DataListPricing = "PRICING_OCPC"
	PRICING_ECPC_ReportAdvertiserGetV2DataListPricing ReportAdvertiserGetV2DataListPricing = "PRICING_ECPC"
	PRICING_OCPM_ReportAdvertiserGetV2DataListPricing ReportAdvertiserGetV2DataListPricing = "PRICING_OCPM"
	PRICING_CPA_ReportAdvertiserGetV2DataListPricing  ReportAdvertiserGetV2DataListPricing = "PRICING_CPA"
)

// Ptr returns reference to report_advertiser_get_v2_data_list_pricing value
func (v ReportAdvertiserGetV2DataListPricing) Ptr() *ReportAdvertiserGetV2DataListPricing {
	return &v
}
