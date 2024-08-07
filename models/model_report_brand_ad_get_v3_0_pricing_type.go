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

// ReportBrandAdGetV30PricingType
type ReportBrandAdGetV30PricingType int64

// List of report_brand_ad_get_v3.0_pricing_type
const (
	Enum_1_ReportBrandAdGetV30PricingType  ReportBrandAdGetV30PricingType = 1
	Enum_10_ReportBrandAdGetV30PricingType ReportBrandAdGetV30PricingType = 10
	Enum_2_ReportBrandAdGetV30PricingType  ReportBrandAdGetV30PricingType = 2
	Enum_3_ReportBrandAdGetV30PricingType  ReportBrandAdGetV30PricingType = 3
	Enum_4_ReportBrandAdGetV30PricingType  ReportBrandAdGetV30PricingType = 4
	Enum_5_ReportBrandAdGetV30PricingType  ReportBrandAdGetV30PricingType = 5
	Enum_6_ReportBrandAdGetV30PricingType  ReportBrandAdGetV30PricingType = 6
	Enum_7_ReportBrandAdGetV30PricingType  ReportBrandAdGetV30PricingType = 7
	Enum_8_ReportBrandAdGetV30PricingType  ReportBrandAdGetV30PricingType = 8
	Enum_9_ReportBrandAdGetV30PricingType  ReportBrandAdGetV30PricingType = 9
)

// All allowed values of ReportBrandAdGetV30PricingType enum
var AllowedReportBrandAdGetV30PricingTypeEnumValues = []ReportBrandAdGetV30PricingType{
	1,
	10,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
}

// NewReportBrandAdGetV30PricingTypeFromValue returns a pointer to a valid ReportBrandAdGetV30PricingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportBrandAdGetV30PricingTypeFromValue(v int64) (*ReportBrandAdGetV30PricingType, error) {
	ev := ReportBrandAdGetV30PricingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportBrandAdGetV30PricingType: valid values are %v", v, AllowedReportBrandAdGetV30PricingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportBrandAdGetV30PricingType) IsValid() bool {
	for _, existing := range AllowedReportBrandAdGetV30PricingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_brand_ad_get_v3.0_pricing_type value
func (v ReportBrandAdGetV30PricingType) Ptr() *ReportBrandAdGetV30PricingType {
	return &v
}
