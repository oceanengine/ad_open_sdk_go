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

// ReportAdGetV2FilteringLandingType
type ReportAdGetV2FilteringLandingType string

// List of report_ad_get_v2_filtering_landing_type
const (
	APP_ReportAdGetV2FilteringLandingType       ReportAdGetV2FilteringLandingType = "APP"
	LINK_ReportAdGetV2FilteringLandingType      ReportAdGetV2FilteringLandingType = "LINK"
	DPA_ReportAdGetV2FilteringLandingType       ReportAdGetV2FilteringLandingType = "DPA"
	AWEME_ReportAdGetV2FilteringLandingType     ReportAdGetV2FilteringLandingType = "AWEME"
	QUICK_APP_ReportAdGetV2FilteringLandingType ReportAdGetV2FilteringLandingType = "QUICK_APP"
	SHOP_ReportAdGetV2FilteringLandingType      ReportAdGetV2FilteringLandingType = "SHOP"
	ARTICLE_ReportAdGetV2FilteringLandingType   ReportAdGetV2FilteringLandingType = "ARTICLE"
	STORE_ReportAdGetV2FilteringLandingType     ReportAdGetV2FilteringLandingType = "STORE"
	LIVE_ReportAdGetV2FilteringLandingType      ReportAdGetV2FilteringLandingType = "LIVE"
	GOODS_ReportAdGetV2FilteringLandingType     ReportAdGetV2FilteringLandingType = "GOODS"
)

// All allowed values of ReportAdGetV2FilteringLandingType enum
var AllowedReportAdGetV2FilteringLandingTypeEnumValues = []ReportAdGetV2FilteringLandingType{
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

// NewReportAdGetV2FilteringLandingTypeFromValue returns a pointer to a valid ReportAdGetV2FilteringLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2FilteringLandingTypeFromValue(v string) (*ReportAdGetV2FilteringLandingType, error) {
	ev := ReportAdGetV2FilteringLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2FilteringLandingType: valid values are %v", v, AllowedReportAdGetV2FilteringLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2FilteringLandingType) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2FilteringLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_filtering_landing_type value
func (v ReportAdGetV2FilteringLandingType) Ptr() *ReportAdGetV2FilteringLandingType {
	return &v
}
