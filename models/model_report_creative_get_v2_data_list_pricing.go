/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCreativeGetV2DataListPricing
type ReportCreativeGetV2DataListPricing string

// List of report_creative_get_v2_data_list_pricing
const (
	PRICING_CPM_ReportCreativeGetV2DataListPricing  ReportCreativeGetV2DataListPricing = "PRICING_CPM"
	PRICING_CPA_ReportCreativeGetV2DataListPricing  ReportCreativeGetV2DataListPricing = "PRICING_CPA"
	PRICING_CPC_ReportCreativeGetV2DataListPricing  ReportCreativeGetV2DataListPricing = "PRICING_CPC"
	PRICING_OCPM_ReportCreativeGetV2DataListPricing ReportCreativeGetV2DataListPricing = "PRICING_OCPM"
	PRICING_ECPC_ReportCreativeGetV2DataListPricing ReportCreativeGetV2DataListPricing = "PRICING_ECPC"
	PRICING_CPV_ReportCreativeGetV2DataListPricing  ReportCreativeGetV2DataListPricing = "PRICING_CPV"
	PRICING_OCPC_ReportCreativeGetV2DataListPricing ReportCreativeGetV2DataListPricing = "PRICING_OCPC"
)

// All allowed values of ReportCreativeGetV2DataListPricing enum
var AllowedReportCreativeGetV2DataListPricingEnumValues = []ReportCreativeGetV2DataListPricing{
	"PRICING_CPM",
	"PRICING_CPA",
	"PRICING_CPC",
	"PRICING_OCPM",
	"PRICING_ECPC",
	"PRICING_CPV",
	"PRICING_OCPC",
}

// NewReportCreativeGetV2DataListPricingFromValue returns a pointer to a valid ReportCreativeGetV2DataListPricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2DataListPricingFromValue(v string) (*ReportCreativeGetV2DataListPricing, error) {
	ev := ReportCreativeGetV2DataListPricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2DataListPricing: valid values are %v", v, AllowedReportCreativeGetV2DataListPricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2DataListPricing) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2DataListPricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_data_list_pricing value
func (v ReportCreativeGetV2DataListPricing) Ptr() *ReportCreativeGetV2DataListPricing {
	return &v
}
