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

// ReportCreativeGetV2FilteringLandingType
type ReportCreativeGetV2FilteringLandingType string

// List of report_creative_get_v2_filtering_landing_type
const (
	APP_ReportCreativeGetV2FilteringLandingType       ReportCreativeGetV2FilteringLandingType = "APP"
	LINK_ReportCreativeGetV2FilteringLandingType      ReportCreativeGetV2FilteringLandingType = "LINK"
	DPA_ReportCreativeGetV2FilteringLandingType       ReportCreativeGetV2FilteringLandingType = "DPA"
	AWEME_ReportCreativeGetV2FilteringLandingType     ReportCreativeGetV2FilteringLandingType = "AWEME"
	QUICK_APP_ReportCreativeGetV2FilteringLandingType ReportCreativeGetV2FilteringLandingType = "QUICK_APP"
	SHOP_ReportCreativeGetV2FilteringLandingType      ReportCreativeGetV2FilteringLandingType = "SHOP"
	ARTICLE_ReportCreativeGetV2FilteringLandingType   ReportCreativeGetV2FilteringLandingType = "ARTICLE"
	STORE_ReportCreativeGetV2FilteringLandingType     ReportCreativeGetV2FilteringLandingType = "STORE"
	LIVE_ReportCreativeGetV2FilteringLandingType      ReportCreativeGetV2FilteringLandingType = "LIVE"
	GOODS_ReportCreativeGetV2FilteringLandingType     ReportCreativeGetV2FilteringLandingType = "GOODS"
)

// All allowed values of ReportCreativeGetV2FilteringLandingType enum
var AllowedReportCreativeGetV2FilteringLandingTypeEnumValues = []ReportCreativeGetV2FilteringLandingType{
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

// NewReportCreativeGetV2FilteringLandingTypeFromValue returns a pointer to a valid ReportCreativeGetV2FilteringLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2FilteringLandingTypeFromValue(v string) (*ReportCreativeGetV2FilteringLandingType, error) {
	ev := ReportCreativeGetV2FilteringLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2FilteringLandingType: valid values are %v", v, AllowedReportCreativeGetV2FilteringLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2FilteringLandingType) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2FilteringLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_filtering_landing_type value
func (v ReportCreativeGetV2FilteringLandingType) Ptr() *ReportCreativeGetV2FilteringLandingType {
	return &v
}
