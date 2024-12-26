/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataPricing
type AdGetV2DataPricing string

// List of ad_get_v2_data_pricing
const (
	PRICING_OCPM_AdGetV2DataPricing     AdGetV2DataPricing = "PRICING_OCPM"
	PRICING_CPC_OCPM_AdGetV2DataPricing AdGetV2DataPricing = "PRICING_CPC_OCPM"
	PRICING_CPM_AdGetV2DataPricing      AdGetV2DataPricing = "PRICING_CPM"
	PRICING_CPC_AdGetV2DataPricing      AdGetV2DataPricing = "PRICING_CPC"
	PRICING_CPV_AdGetV2DataPricing      AdGetV2DataPricing = "PRICING_CPV"
	PRICING_CPA_AdGetV2DataPricing      AdGetV2DataPricing = "PRICING_CPA"
	PRICING_OCPC_AdGetV2DataPricing     AdGetV2DataPricing = "PRICING_OCPC"
)

// Ptr returns reference to ad_get_v2_data_pricing value
func (v AdGetV2DataPricing) Ptr() *AdGetV2DataPricing {
	return &v
}
