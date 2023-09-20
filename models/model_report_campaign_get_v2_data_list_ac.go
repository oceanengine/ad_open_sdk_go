/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2DataListAc
type ReportCampaignGetV2DataListAc string

// List of report_campaign_get_v2_data_list_ac
const (
	WIFI_ReportCampaignGetV2DataListAc     ReportCampaignGetV2DataListAc = "WIFI"
	Enum_4_G_ReportCampaignGetV2DataListAc ReportCampaignGetV2DataListAc = "4G"
	UNKNOWN_ReportCampaignGetV2DataListAc  ReportCampaignGetV2DataListAc = "unknown"
	Enum_5_G_ReportCampaignGetV2DataListAc ReportCampaignGetV2DataListAc = "5G"
	Enum_2_G_ReportCampaignGetV2DataListAc ReportCampaignGetV2DataListAc = "2G"
	Enum_3_G_ReportCampaignGetV2DataListAc ReportCampaignGetV2DataListAc = "3G"
)

// All allowed values of ReportCampaignGetV2DataListAc enum
var AllowedReportCampaignGetV2DataListAcEnumValues = []ReportCampaignGetV2DataListAc{
	"WIFI",
	"4G",
	"unknown",
	"5G",
	"2G",
	"3G",
}

// NewReportCampaignGetV2DataListAcFromValue returns a pointer to a valid ReportCampaignGetV2DataListAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2DataListAcFromValue(v string) (*ReportCampaignGetV2DataListAc, error) {
	ev := ReportCampaignGetV2DataListAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2DataListAc: valid values are %v", v, AllowedReportCampaignGetV2DataListAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2DataListAc) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2DataListAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_data_list_ac value
func (v ReportCampaignGetV2DataListAc) Ptr() *ReportCampaignGetV2DataListAc {
	return &v
}
