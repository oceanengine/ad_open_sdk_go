/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportBrandCampaignGetV30LandingType
type ReportBrandCampaignGetV30LandingType int64

// List of report_brand_campaign_get_v3.0_landing_type
const (
	Enum_1_ReportBrandCampaignGetV30LandingType  ReportBrandCampaignGetV30LandingType = 1
	Enum_10_ReportBrandCampaignGetV30LandingType ReportBrandCampaignGetV30LandingType = 10
	Enum_11_ReportBrandCampaignGetV30LandingType ReportBrandCampaignGetV30LandingType = 11
	Enum_12_ReportBrandCampaignGetV30LandingType ReportBrandCampaignGetV30LandingType = 12
	Enum_13_ReportBrandCampaignGetV30LandingType ReportBrandCampaignGetV30LandingType = 13
	Enum_2_ReportBrandCampaignGetV30LandingType  ReportBrandCampaignGetV30LandingType = 2
	Enum_3_ReportBrandCampaignGetV30LandingType  ReportBrandCampaignGetV30LandingType = 3
	Enum_4_ReportBrandCampaignGetV30LandingType  ReportBrandCampaignGetV30LandingType = 4
	Enum_5_ReportBrandCampaignGetV30LandingType  ReportBrandCampaignGetV30LandingType = 5
	Enum_6_ReportBrandCampaignGetV30LandingType  ReportBrandCampaignGetV30LandingType = 6
	Enum_7_ReportBrandCampaignGetV30LandingType  ReportBrandCampaignGetV30LandingType = 7
	Enum_8_ReportBrandCampaignGetV30LandingType  ReportBrandCampaignGetV30LandingType = 8
	Enum_9_ReportBrandCampaignGetV30LandingType  ReportBrandCampaignGetV30LandingType = 9
)

// All allowed values of ReportBrandCampaignGetV30LandingType enum
var AllowedReportBrandCampaignGetV30LandingTypeEnumValues = []ReportBrandCampaignGetV30LandingType{
	1,
	10,
	11,
	12,
	13,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
}

// NewReportBrandCampaignGetV30LandingTypeFromValue returns a pointer to a valid ReportBrandCampaignGetV30LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportBrandCampaignGetV30LandingTypeFromValue(v int64) (*ReportBrandCampaignGetV30LandingType, error) {
	ev := ReportBrandCampaignGetV30LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportBrandCampaignGetV30LandingType: valid values are %v", v, AllowedReportBrandCampaignGetV30LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportBrandCampaignGetV30LandingType) IsValid() bool {
	for _, existing := range AllowedReportBrandCampaignGetV30LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_brand_campaign_get_v3.0_landing_type value
func (v ReportBrandCampaignGetV30LandingType) Ptr() *ReportBrandCampaignGetV30LandingType {
	return &v
}
