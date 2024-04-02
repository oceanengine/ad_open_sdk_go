/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportBrandCampaignGetV30PricingType
type ReportBrandCampaignGetV30PricingType int64

// List of report_brand_campaign_get_v3.0_pricing_type
const (
	Enum_1_ReportBrandCampaignGetV30PricingType  ReportBrandCampaignGetV30PricingType = 1
	Enum_10_ReportBrandCampaignGetV30PricingType ReportBrandCampaignGetV30PricingType = 10
	Enum_2_ReportBrandCampaignGetV30PricingType  ReportBrandCampaignGetV30PricingType = 2
	Enum_3_ReportBrandCampaignGetV30PricingType  ReportBrandCampaignGetV30PricingType = 3
	Enum_4_ReportBrandCampaignGetV30PricingType  ReportBrandCampaignGetV30PricingType = 4
	Enum_5_ReportBrandCampaignGetV30PricingType  ReportBrandCampaignGetV30PricingType = 5
	Enum_6_ReportBrandCampaignGetV30PricingType  ReportBrandCampaignGetV30PricingType = 6
	Enum_7_ReportBrandCampaignGetV30PricingType  ReportBrandCampaignGetV30PricingType = 7
	Enum_8_ReportBrandCampaignGetV30PricingType  ReportBrandCampaignGetV30PricingType = 8
	Enum_9_ReportBrandCampaignGetV30PricingType  ReportBrandCampaignGetV30PricingType = 9
)

// All allowed values of ReportBrandCampaignGetV30PricingType enum
var AllowedReportBrandCampaignGetV30PricingTypeEnumValues = []ReportBrandCampaignGetV30PricingType{
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

// NewReportBrandCampaignGetV30PricingTypeFromValue returns a pointer to a valid ReportBrandCampaignGetV30PricingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportBrandCampaignGetV30PricingTypeFromValue(v int64) (*ReportBrandCampaignGetV30PricingType, error) {
	ev := ReportBrandCampaignGetV30PricingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportBrandCampaignGetV30PricingType: valid values are %v", v, AllowedReportBrandCampaignGetV30PricingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportBrandCampaignGetV30PricingType) IsValid() bool {
	for _, existing := range AllowedReportBrandCampaignGetV30PricingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_brand_campaign_get_v3.0_pricing_type value
func (v ReportBrandCampaignGetV30PricingType) Ptr() *ReportBrandCampaignGetV30PricingType {
	return &v
}