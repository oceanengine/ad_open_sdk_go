/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdGetV2FilteringCampaignTypes
type ReportAdGetV2FilteringCampaignTypes string

// List of report_ad_get_v2_filtering_campaign_types
const (
	SEARCH_ReportAdGetV2FilteringCampaignTypes  ReportAdGetV2FilteringCampaignTypes = "SEARCH"
	FEED_ReportAdGetV2FilteringCampaignTypes    ReportAdGetV2FilteringCampaignTypes = "FEED"
	CONTENT_ReportAdGetV2FilteringCampaignTypes ReportAdGetV2FilteringCampaignTypes = "CONTENT"
)

// All allowed values of ReportAdGetV2FilteringCampaignTypes enum
var AllowedReportAdGetV2FilteringCampaignTypesEnumValues = []ReportAdGetV2FilteringCampaignTypes{
	"SEARCH",
	"FEED",
	"CONTENT",
}

// NewReportAdGetV2FilteringCampaignTypesFromValue returns a pointer to a valid ReportAdGetV2FilteringCampaignTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2FilteringCampaignTypesFromValue(v string) (*ReportAdGetV2FilteringCampaignTypes, error) {
	ev := ReportAdGetV2FilteringCampaignTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2FilteringCampaignTypes: valid values are %v", v, AllowedReportAdGetV2FilteringCampaignTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2FilteringCampaignTypes) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2FilteringCampaignTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_filtering_campaign_types value
func (v ReportAdGetV2FilteringCampaignTypes) Ptr() *ReportAdGetV2FilteringCampaignTypes {
	return &v
}
