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

// ReportAdGetV2FilteringLandingTypes
type ReportAdGetV2FilteringLandingTypes string

// List of report_ad_get_v2_filtering_landing_types
const (
	APP_ReportAdGetV2FilteringLandingTypes       ReportAdGetV2FilteringLandingTypes = "APP"
	DPA_ReportAdGetV2FilteringLandingTypes       ReportAdGetV2FilteringLandingTypes = "DPA"
	AWEME_ReportAdGetV2FilteringLandingTypes     ReportAdGetV2FilteringLandingTypes = "AWEME"
	GOODS_ReportAdGetV2FilteringLandingTypes     ReportAdGetV2FilteringLandingTypes = "GOODS"
	LINK_ReportAdGetV2FilteringLandingTypes      ReportAdGetV2FilteringLandingTypes = "LINK"
	LIVE_ReportAdGetV2FilteringLandingTypes      ReportAdGetV2FilteringLandingTypes = "LIVE"
	STORE_ReportAdGetV2FilteringLandingTypes     ReportAdGetV2FilteringLandingTypes = "STORE"
	SHOP_ReportAdGetV2FilteringLandingTypes      ReportAdGetV2FilteringLandingTypes = "SHOP"
	QUICK_APP_ReportAdGetV2FilteringLandingTypes ReportAdGetV2FilteringLandingTypes = "QUICK_APP"
	ARTICLE_ReportAdGetV2FilteringLandingTypes   ReportAdGetV2FilteringLandingTypes = "ARTICLE"
)

// All allowed values of ReportAdGetV2FilteringLandingTypes enum
var AllowedReportAdGetV2FilteringLandingTypesEnumValues = []ReportAdGetV2FilteringLandingTypes{
	"APP",
	"DPA",
	"AWEME",
	"GOODS",
	"LINK",
	"LIVE",
	"STORE",
	"SHOP",
	"QUICK_APP",
	"ARTICLE",
}

// NewReportAdGetV2FilteringLandingTypesFromValue returns a pointer to a valid ReportAdGetV2FilteringLandingTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2FilteringLandingTypesFromValue(v string) (*ReportAdGetV2FilteringLandingTypes, error) {
	ev := ReportAdGetV2FilteringLandingTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2FilteringLandingTypes: valid values are %v", v, AllowedReportAdGetV2FilteringLandingTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2FilteringLandingTypes) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2FilteringLandingTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_filtering_landing_types value
func (v ReportAdGetV2FilteringLandingTypes) Ptr() *ReportAdGetV2FilteringLandingTypes {
	return &v
}
