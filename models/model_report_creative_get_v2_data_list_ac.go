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

// ReportCreativeGetV2DataListAc
type ReportCreativeGetV2DataListAc string

// List of report_creative_get_v2_data_list_ac
const (
	Enum_3_G_ReportCreativeGetV2DataListAc ReportCreativeGetV2DataListAc = "3G"
	WIFI_ReportCreativeGetV2DataListAc     ReportCreativeGetV2DataListAc = "WIFI"
	Enum_5_G_ReportCreativeGetV2DataListAc ReportCreativeGetV2DataListAc = "5G"
	Enum_2_G_ReportCreativeGetV2DataListAc ReportCreativeGetV2DataListAc = "2G"
	Enum_4_G_ReportCreativeGetV2DataListAc ReportCreativeGetV2DataListAc = "4G"
	UNKNOWN_ReportCreativeGetV2DataListAc  ReportCreativeGetV2DataListAc = "unknown"
)

// All allowed values of ReportCreativeGetV2DataListAc enum
var AllowedReportCreativeGetV2DataListAcEnumValues = []ReportCreativeGetV2DataListAc{
	"3G",
	"WIFI",
	"5G",
	"2G",
	"4G",
	"unknown",
}

// NewReportCreativeGetV2DataListAcFromValue returns a pointer to a valid ReportCreativeGetV2DataListAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2DataListAcFromValue(v string) (*ReportCreativeGetV2DataListAc, error) {
	ev := ReportCreativeGetV2DataListAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2DataListAc: valid values are %v", v, AllowedReportCreativeGetV2DataListAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2DataListAc) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2DataListAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_data_list_ac value
func (v ReportCreativeGetV2DataListAc) Ptr() *ReportCreativeGetV2DataListAc {
	return &v
}
