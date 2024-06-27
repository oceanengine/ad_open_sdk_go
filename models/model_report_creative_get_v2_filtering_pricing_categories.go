/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCreativeGetV2FilteringPricingCategories
type ReportCreativeGetV2FilteringPricingCategories string

// List of report_creative_get_v2_filtering_pricing_categories
const (
	PRICING_CATEGORY_BRAND_ReportCreativeGetV2FilteringPricingCategories             ReportCreativeGetV2FilteringPricingCategories = "PRICING_CATEGORY_BRAND"
	PRICING_CATEGORY_BRAND_AND_PRICING_ReportCreativeGetV2FilteringPricingCategories ReportCreativeGetV2FilteringPricingCategories = "PRICING_CATEGORY_BRAND_AND_PRICING"
	PRICING_CATEGORY_BID_ReportCreativeGetV2FilteringPricingCategories               ReportCreativeGetV2FilteringPricingCategories = "PRICING_CATEGORY_BID"
	PRICING_CATEGORY_NOC_ReportCreativeGetV2FilteringPricingCategories               ReportCreativeGetV2FilteringPricingCategories = "PRICING_CATEGORY_NOC"
	PRICING_CATEGORY_FREE_ReportCreativeGetV2FilteringPricingCategories              ReportCreativeGetV2FilteringPricingCategories = "PRICING_CATEGORY_FREE"
)

// All allowed values of ReportCreativeGetV2FilteringPricingCategories enum
var AllowedReportCreativeGetV2FilteringPricingCategoriesEnumValues = []ReportCreativeGetV2FilteringPricingCategories{
	"PRICING_CATEGORY_BRAND",
	"PRICING_CATEGORY_BRAND_AND_PRICING",
	"PRICING_CATEGORY_BID",
	"PRICING_CATEGORY_NOC",
	"PRICING_CATEGORY_FREE",
}

// NewReportCreativeGetV2FilteringPricingCategoriesFromValue returns a pointer to a valid ReportCreativeGetV2FilteringPricingCategories
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2FilteringPricingCategoriesFromValue(v string) (*ReportCreativeGetV2FilteringPricingCategories, error) {
	ev := ReportCreativeGetV2FilteringPricingCategories(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2FilteringPricingCategories: valid values are %v", v, AllowedReportCreativeGetV2FilteringPricingCategoriesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2FilteringPricingCategories) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2FilteringPricingCategoriesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_filtering_pricing_categories value
func (v ReportCreativeGetV2FilteringPricingCategories) Ptr() *ReportCreativeGetV2FilteringPricingCategories {
	return &v
}
