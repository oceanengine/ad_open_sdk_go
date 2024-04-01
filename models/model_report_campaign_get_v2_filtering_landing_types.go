/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2FilteringLandingTypes
type ReportCampaignGetV2FilteringLandingTypes string

// List of report_campaign_get_v2_filtering_landing_types
const (
	LIVE_ReportCampaignGetV2FilteringLandingTypes      ReportCampaignGetV2FilteringLandingTypes = "LIVE"
	STORE_ReportCampaignGetV2FilteringLandingTypes     ReportCampaignGetV2FilteringLandingTypes = "STORE"
	SHOP_ReportCampaignGetV2FilteringLandingTypes      ReportCampaignGetV2FilteringLandingTypes = "SHOP"
	GOODS_ReportCampaignGetV2FilteringLandingTypes     ReportCampaignGetV2FilteringLandingTypes = "GOODS"
	ARTICLE_ReportCampaignGetV2FilteringLandingTypes   ReportCampaignGetV2FilteringLandingTypes = "ARTICLE"
	LINK_ReportCampaignGetV2FilteringLandingTypes      ReportCampaignGetV2FilteringLandingTypes = "LINK"
	DPA_ReportCampaignGetV2FilteringLandingTypes       ReportCampaignGetV2FilteringLandingTypes = "DPA"
	AWEME_ReportCampaignGetV2FilteringLandingTypes     ReportCampaignGetV2FilteringLandingTypes = "AWEME"
	APP_ReportCampaignGetV2FilteringLandingTypes       ReportCampaignGetV2FilteringLandingTypes = "APP"
	QUICK_APP_ReportCampaignGetV2FilteringLandingTypes ReportCampaignGetV2FilteringLandingTypes = "QUICK_APP"
)

// All allowed values of ReportCampaignGetV2FilteringLandingTypes enum
var AllowedReportCampaignGetV2FilteringLandingTypesEnumValues = []ReportCampaignGetV2FilteringLandingTypes{
	"LIVE",
	"STORE",
	"SHOP",
	"GOODS",
	"ARTICLE",
	"LINK",
	"DPA",
	"AWEME",
	"APP",
	"QUICK_APP",
}

// NewReportCampaignGetV2FilteringLandingTypesFromValue returns a pointer to a valid ReportCampaignGetV2FilteringLandingTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2FilteringLandingTypesFromValue(v string) (*ReportCampaignGetV2FilteringLandingTypes, error) {
	ev := ReportCampaignGetV2FilteringLandingTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2FilteringLandingTypes: valid values are %v", v, AllowedReportCampaignGetV2FilteringLandingTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2FilteringLandingTypes) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2FilteringLandingTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_filtering_landing_types value
func (v ReportCampaignGetV2FilteringLandingTypes) Ptr() *ReportCampaignGetV2FilteringLandingTypes {
	return &v
}
