/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCreativeGetV2DataListPricingCategory
type ReportCreativeGetV2DataListPricingCategory string

// List of report_creative_get_v2_data_list_pricing_category
const (
	PRICING_CATEGORY_BRAND_ReportCreativeGetV2DataListPricingCategory             ReportCreativeGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND"
	PRICING_CATEGORY_NOC_ReportCreativeGetV2DataListPricingCategory               ReportCreativeGetV2DataListPricingCategory = "PRICING_CATEGORY_NOC"
	PRICING_CATEGORY_FREE_ReportCreativeGetV2DataListPricingCategory              ReportCreativeGetV2DataListPricingCategory = "PRICING_CATEGORY_FREE"
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportCreativeGetV2DataListPricingCategory ReportCreativeGetV2DataListPricingCategory = "PRICING_CATEGORY_BRAND_AND_PRICING"
	PRICING_CATEGORY_BID_ReportCreativeGetV2DataListPricingCategory               ReportCreativeGetV2DataListPricingCategory = "PRICING_CATEGORY_BID"
)

// All allowed values of ReportCreativeGetV2DataListPricingCategory enum
var AllowedReportCreativeGetV2DataListPricingCategoryEnumValues = []ReportCreativeGetV2DataListPricingCategory{
	"PRICING_CATEGORY_BRAND",
	"PRICING_CATEGORY_NOC",
	"PRICING_CATEGORY_FREE",
	"PRICING_CATEGORY_BRAND_AND_PRICING",
	"PRICING_CATEGORY_BID",
}

// NewReportCreativeGetV2DataListPricingCategoryFromValue returns a pointer to a valid ReportCreativeGetV2DataListPricingCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2DataListPricingCategoryFromValue(v string) (*ReportCreativeGetV2DataListPricingCategory, error) {
	ev := ReportCreativeGetV2DataListPricingCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2DataListPricingCategory: valid values are %v", v, AllowedReportCreativeGetV2DataListPricingCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2DataListPricingCategory) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2DataListPricingCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_data_list_pricing_category value
func (v ReportCreativeGetV2DataListPricingCategory) Ptr() *ReportCreativeGetV2DataListPricingCategory {
	return &v
}
