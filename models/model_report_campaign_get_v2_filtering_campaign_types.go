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

// ReportCampaignGetV2FilteringCampaignTypes
type ReportCampaignGetV2FilteringCampaignTypes string

// List of report_campaign_get_v2_filtering_campaign_types
const (
	FEED_ReportCampaignGetV2FilteringCampaignTypes    ReportCampaignGetV2FilteringCampaignTypes = "FEED"
	SEARCH_ReportCampaignGetV2FilteringCampaignTypes  ReportCampaignGetV2FilteringCampaignTypes = "SEARCH"
	CONTENT_ReportCampaignGetV2FilteringCampaignTypes ReportCampaignGetV2FilteringCampaignTypes = "CONTENT"
)

// All allowed values of ReportCampaignGetV2FilteringCampaignTypes enum
var AllowedReportCampaignGetV2FilteringCampaignTypesEnumValues = []ReportCampaignGetV2FilteringCampaignTypes{
	"FEED",
	"SEARCH",
	"CONTENT",
}

// NewReportCampaignGetV2FilteringCampaignTypesFromValue returns a pointer to a valid ReportCampaignGetV2FilteringCampaignTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2FilteringCampaignTypesFromValue(v string) (*ReportCampaignGetV2FilteringCampaignTypes, error) {
	ev := ReportCampaignGetV2FilteringCampaignTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2FilteringCampaignTypes: valid values are %v", v, AllowedReportCampaignGetV2FilteringCampaignTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2FilteringCampaignTypes) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2FilteringCampaignTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_filtering_campaign_types value
func (v ReportCampaignGetV2FilteringCampaignTypes) Ptr() *ReportCampaignGetV2FilteringCampaignTypes {
	return &v
}
