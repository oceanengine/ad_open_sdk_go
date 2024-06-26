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

// ReportAdGetV2DataListLandingType
type ReportAdGetV2DataListLandingType string

// List of report_ad_get_v2_data_list_landing_type
const (
	APP_ReportAdGetV2DataListLandingType       ReportAdGetV2DataListLandingType = "APP"
	LINK_ReportAdGetV2DataListLandingType      ReportAdGetV2DataListLandingType = "LINK"
	DPA_ReportAdGetV2DataListLandingType       ReportAdGetV2DataListLandingType = "DPA"
	AWEME_ReportAdGetV2DataListLandingType     ReportAdGetV2DataListLandingType = "AWEME"
	QUICK_APP_ReportAdGetV2DataListLandingType ReportAdGetV2DataListLandingType = "QUICK_APP"
	SHOP_ReportAdGetV2DataListLandingType      ReportAdGetV2DataListLandingType = "SHOP"
	ARTICLE_ReportAdGetV2DataListLandingType   ReportAdGetV2DataListLandingType = "ARTICLE"
	STORE_ReportAdGetV2DataListLandingType     ReportAdGetV2DataListLandingType = "STORE"
	LIVE_ReportAdGetV2DataListLandingType      ReportAdGetV2DataListLandingType = "LIVE"
	GOODS_ReportAdGetV2DataListLandingType     ReportAdGetV2DataListLandingType = "GOODS"
)

// All allowed values of ReportAdGetV2DataListLandingType enum
var AllowedReportAdGetV2DataListLandingTypeEnumValues = []ReportAdGetV2DataListLandingType{
	"APP",
	"LINK",
	"DPA",
	"AWEME",
	"QUICK_APP",
	"SHOP",
	"ARTICLE",
	"STORE",
	"LIVE",
	"GOODS",
}

// NewReportAdGetV2DataListLandingTypeFromValue returns a pointer to a valid ReportAdGetV2DataListLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2DataListLandingTypeFromValue(v string) (*ReportAdGetV2DataListLandingType, error) {
	ev := ReportAdGetV2DataListLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2DataListLandingType: valid values are %v", v, AllowedReportAdGetV2DataListLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2DataListLandingType) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2DataListLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_data_list_landing_type value
func (v ReportAdGetV2DataListLandingType) Ptr() *ReportAdGetV2DataListLandingType {
	return &v
}
