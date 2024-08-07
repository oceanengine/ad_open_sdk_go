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

// ReportAudienceCityV2IdType
type ReportAudienceCityV2IdType string

// List of report_audience_city_v2_id_type
const (
	AUDIENCE_STAT_ID_TYPE_AD_ReportAudienceCityV2IdType         ReportAudienceCityV2IdType = "AUDIENCE_STAT_ID_TYPE_AD"
	AUDIENCE_STAT_ID_TYPE_ADVERTISER_ReportAudienceCityV2IdType ReportAudienceCityV2IdType = "AUDIENCE_STAT_ID_TYPE_ADVERTISER"
	AUDIENCE_STAT_ID_TYPE_CAMPAIGN_ReportAudienceCityV2IdType   ReportAudienceCityV2IdType = "AUDIENCE_STAT_ID_TYPE_CAMPAIGN"
)

// All allowed values of ReportAudienceCityV2IdType enum
var AllowedReportAudienceCityV2IdTypeEnumValues = []ReportAudienceCityV2IdType{
	"AUDIENCE_STAT_ID_TYPE_AD",
	"AUDIENCE_STAT_ID_TYPE_ADVERTISER",
	"AUDIENCE_STAT_ID_TYPE_CAMPAIGN",
}

// NewReportAudienceCityV2IdTypeFromValue returns a pointer to a valid ReportAudienceCityV2IdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAudienceCityV2IdTypeFromValue(v string) (*ReportAudienceCityV2IdType, error) {
	ev := ReportAudienceCityV2IdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAudienceCityV2IdType: valid values are %v", v, AllowedReportAudienceCityV2IdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAudienceCityV2IdType) IsValid() bool {
	for _, existing := range AllowedReportAudienceCityV2IdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_audience_city_v2_id_type value
func (v ReportAudienceCityV2IdType) Ptr() *ReportAudienceCityV2IdType {
	return &v
}
