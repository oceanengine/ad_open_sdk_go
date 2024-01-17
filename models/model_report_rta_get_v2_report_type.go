/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportRtaGetV2ReportType
type ReportRtaGetV2ReportType string

// List of report_rta_get_v2_report_type
const (
	CREATIVE_RTA_ReportRtaGetV2ReportType            ReportRtaGetV2ReportType = "CREATIVE_RTA"
	ADVERTISER_RTA_PLATFORM_ReportRtaGetV2ReportType ReportRtaGetV2ReportType = "ADVERTISER_RTA_PLATFORM"
	CAMPAIGN_RTA_PLATFORM_ReportRtaGetV2ReportType   ReportRtaGetV2ReportType = "CAMPAIGN_RTA_PLATFORM"
	AD_RTA_ReportRtaGetV2ReportType                  ReportRtaGetV2ReportType = "AD_RTA"
	AD_RTA_PLATFORM_ReportRtaGetV2ReportType         ReportRtaGetV2ReportType = "AD_RTA_PLATFORM"
	CAMPAIGN_RTA_ReportRtaGetV2ReportType            ReportRtaGetV2ReportType = "CAMPAIGN_RTA"
	CREATIVE_RTA_PLATFORM_ReportRtaGetV2ReportType   ReportRtaGetV2ReportType = "CREATIVE_RTA_PLATFORM"
	ADVERTISER_RTA_ReportRtaGetV2ReportType          ReportRtaGetV2ReportType = "ADVERTISER_RTA"
)

// All allowed values of ReportRtaGetV2ReportType enum
var AllowedReportRtaGetV2ReportTypeEnumValues = []ReportRtaGetV2ReportType{
	"CREATIVE_RTA",
	"ADVERTISER_RTA_PLATFORM",
	"CAMPAIGN_RTA_PLATFORM",
	"AD_RTA",
	"AD_RTA_PLATFORM",
	"CAMPAIGN_RTA",
	"CREATIVE_RTA_PLATFORM",
	"ADVERTISER_RTA",
}

// NewReportRtaGetV2ReportTypeFromValue returns a pointer to a valid ReportRtaGetV2ReportType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportRtaGetV2ReportTypeFromValue(v string) (*ReportRtaGetV2ReportType, error) {
	ev := ReportRtaGetV2ReportType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportRtaGetV2ReportType: valid values are %v", v, AllowedReportRtaGetV2ReportTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportRtaGetV2ReportType) IsValid() bool {
	for _, existing := range AllowedReportRtaGetV2ReportTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_rta_get_v2_report_type value
func (v ReportRtaGetV2ReportType) Ptr() *ReportRtaGetV2ReportType {
	return &v
}
