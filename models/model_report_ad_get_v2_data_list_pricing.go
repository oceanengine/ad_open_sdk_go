/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdGetV2DataListPricing
type ReportAdGetV2DataListPricing string

// List of report_ad_get_v2_data_list_pricing
const (
	PRICING_CPA_ReportAdGetV2DataListPricing  ReportAdGetV2DataListPricing = "PRICING_CPA"
	PRICING_CPM_ReportAdGetV2DataListPricing  ReportAdGetV2DataListPricing = "PRICING_CPM"
	PRICING_ECPC_ReportAdGetV2DataListPricing ReportAdGetV2DataListPricing = "PRICING_ECPC"
	PRICING_CPC_ReportAdGetV2DataListPricing  ReportAdGetV2DataListPricing = "PRICING_CPC"
	PRICING_OCPC_ReportAdGetV2DataListPricing ReportAdGetV2DataListPricing = "PRICING_OCPC"
	PRICING_CPV_ReportAdGetV2DataListPricing  ReportAdGetV2DataListPricing = "PRICING_CPV"
	PRICING_OCPM_ReportAdGetV2DataListPricing ReportAdGetV2DataListPricing = "PRICING_OCPM"
)

// All allowed values of ReportAdGetV2DataListPricing enum
var AllowedReportAdGetV2DataListPricingEnumValues = []ReportAdGetV2DataListPricing{
	"PRICING_CPA",
	"PRICING_CPM",
	"PRICING_ECPC",
	"PRICING_CPC",
	"PRICING_OCPC",
	"PRICING_CPV",
	"PRICING_OCPM",
}

// NewReportAdGetV2DataListPricingFromValue returns a pointer to a valid ReportAdGetV2DataListPricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2DataListPricingFromValue(v string) (*ReportAdGetV2DataListPricing, error) {
	ev := ReportAdGetV2DataListPricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2DataListPricing: valid values are %v", v, AllowedReportAdGetV2DataListPricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2DataListPricing) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2DataListPricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_data_list_pricing value
func (v ReportAdGetV2DataListPricing) Ptr() *ReportAdGetV2DataListPricing {
	return &v
}
