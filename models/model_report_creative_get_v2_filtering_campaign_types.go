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

// ReportCreativeGetV2FilteringCampaignTypes
type ReportCreativeGetV2FilteringCampaignTypes string

// List of report_creative_get_v2_filtering_campaign_types
const (
	FEED_ReportCreativeGetV2FilteringCampaignTypes    ReportCreativeGetV2FilteringCampaignTypes = "FEED"
	CONTENT_ReportCreativeGetV2FilteringCampaignTypes ReportCreativeGetV2FilteringCampaignTypes = "CONTENT"
	SEARCH_ReportCreativeGetV2FilteringCampaignTypes  ReportCreativeGetV2FilteringCampaignTypes = "SEARCH"
)

// All allowed values of ReportCreativeGetV2FilteringCampaignTypes enum
var AllowedReportCreativeGetV2FilteringCampaignTypesEnumValues = []ReportCreativeGetV2FilteringCampaignTypes{
	"FEED",
	"CONTENT",
	"SEARCH",
}

// NewReportCreativeGetV2FilteringCampaignTypesFromValue returns a pointer to a valid ReportCreativeGetV2FilteringCampaignTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2FilteringCampaignTypesFromValue(v string) (*ReportCreativeGetV2FilteringCampaignTypes, error) {
	ev := ReportCreativeGetV2FilteringCampaignTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2FilteringCampaignTypes: valid values are %v", v, AllowedReportCreativeGetV2FilteringCampaignTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2FilteringCampaignTypes) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2FilteringCampaignTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_filtering_campaign_types value
func (v ReportCreativeGetV2FilteringCampaignTypes) Ptr() *ReportCreativeGetV2FilteringCampaignTypes {
	return &v
}
