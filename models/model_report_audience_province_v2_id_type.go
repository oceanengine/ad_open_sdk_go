/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAudienceProvinceV2IdType
type ReportAudienceProvinceV2IdType string

// List of report_audience_province_v2_id_type
const (
	AUDIENCE_STAT_ID_TYPE_CAMPAIGN_ReportAudienceProvinceV2IdType   ReportAudienceProvinceV2IdType = "AUDIENCE_STAT_ID_TYPE_CAMPAIGN"
	AUDIENCE_STAT_ID_TYPE_ADVERTISER_ReportAudienceProvinceV2IdType ReportAudienceProvinceV2IdType = "AUDIENCE_STAT_ID_TYPE_ADVERTISER"
	AUDIENCE_STAT_ID_TYPE_AD_ReportAudienceProvinceV2IdType         ReportAudienceProvinceV2IdType = "AUDIENCE_STAT_ID_TYPE_AD"
)

// All allowed values of ReportAudienceProvinceV2IdType enum
var AllowedReportAudienceProvinceV2IdTypeEnumValues = []ReportAudienceProvinceV2IdType{
	"AUDIENCE_STAT_ID_TYPE_CAMPAIGN",
	"AUDIENCE_STAT_ID_TYPE_ADVERTISER",
	"AUDIENCE_STAT_ID_TYPE_AD",
}

// NewReportAudienceProvinceV2IdTypeFromValue returns a pointer to a valid ReportAudienceProvinceV2IdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAudienceProvinceV2IdTypeFromValue(v string) (*ReportAudienceProvinceV2IdType, error) {
	ev := ReportAudienceProvinceV2IdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAudienceProvinceV2IdType: valid values are %v", v, AllowedReportAudienceProvinceV2IdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAudienceProvinceV2IdType) IsValid() bool {
	for _, existing := range AllowedReportAudienceProvinceV2IdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_audience_province_v2_id_type value
func (v ReportAudienceProvinceV2IdType) Ptr() *ReportAudienceProvinceV2IdType {
	return &v
}
