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

// ReportBrandAdvertiserGetV30PricingType
type ReportBrandAdvertiserGetV30PricingType int64

// List of report_brand_advertiser_get_v3.0_pricing_type
const (
	Enum_1_ReportBrandAdvertiserGetV30PricingType  ReportBrandAdvertiserGetV30PricingType = 1
	Enum_10_ReportBrandAdvertiserGetV30PricingType ReportBrandAdvertiserGetV30PricingType = 10
	Enum_2_ReportBrandAdvertiserGetV30PricingType  ReportBrandAdvertiserGetV30PricingType = 2
	Enum_3_ReportBrandAdvertiserGetV30PricingType  ReportBrandAdvertiserGetV30PricingType = 3
	Enum_4_ReportBrandAdvertiserGetV30PricingType  ReportBrandAdvertiserGetV30PricingType = 4
	Enum_5_ReportBrandAdvertiserGetV30PricingType  ReportBrandAdvertiserGetV30PricingType = 5
	Enum_6_ReportBrandAdvertiserGetV30PricingType  ReportBrandAdvertiserGetV30PricingType = 6
	Enum_7_ReportBrandAdvertiserGetV30PricingType  ReportBrandAdvertiserGetV30PricingType = 7
	Enum_8_ReportBrandAdvertiserGetV30PricingType  ReportBrandAdvertiserGetV30PricingType = 8
	Enum_9_ReportBrandAdvertiserGetV30PricingType  ReportBrandAdvertiserGetV30PricingType = 9
)

// All allowed values of ReportBrandAdvertiserGetV30PricingType enum
var AllowedReportBrandAdvertiserGetV30PricingTypeEnumValues = []ReportBrandAdvertiserGetV30PricingType{
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

// NewReportBrandAdvertiserGetV30PricingTypeFromValue returns a pointer to a valid ReportBrandAdvertiserGetV30PricingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportBrandAdvertiserGetV30PricingTypeFromValue(v int64) (*ReportBrandAdvertiserGetV30PricingType, error) {
	ev := ReportBrandAdvertiserGetV30PricingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportBrandAdvertiserGetV30PricingType: valid values are %v", v, AllowedReportBrandAdvertiserGetV30PricingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportBrandAdvertiserGetV30PricingType) IsValid() bool {
	for _, existing := range AllowedReportBrandAdvertiserGetV30PricingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_brand_advertiser_get_v3.0_pricing_type value
func (v ReportBrandAdvertiserGetV30PricingType) Ptr() *ReportBrandAdvertiserGetV30PricingType {
	return &v
}
