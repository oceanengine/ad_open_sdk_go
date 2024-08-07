/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2DataListCampaignType
type ReportCampaignGetV2DataListCampaignType string

// List of report_campaign_get_v2_data_list_campaign_type
const (
	SEARCH_ReportCampaignGetV2DataListCampaignType ReportCampaignGetV2DataListCampaignType = "SEARCH"
	FEED_ReportCampaignGetV2DataListCampaignType   ReportCampaignGetV2DataListCampaignType = "FEED"
)

// All allowed values of ReportCampaignGetV2DataListCampaignType enum
var AllowedReportCampaignGetV2DataListCampaignTypeEnumValues = []ReportCampaignGetV2DataListCampaignType{
	"SEARCH",
	"FEED",
}

// NewReportCampaignGetV2DataListCampaignTypeFromValue returns a pointer to a valid ReportCampaignGetV2DataListCampaignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2DataListCampaignTypeFromValue(v string) (*ReportCampaignGetV2DataListCampaignType, error) {
	ev := ReportCampaignGetV2DataListCampaignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2DataListCampaignType: valid values are %v", v, AllowedReportCampaignGetV2DataListCampaignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2DataListCampaignType) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2DataListCampaignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_data_list_campaign_type value
func (v ReportCampaignGetV2DataListCampaignType) Ptr() *ReportCampaignGetV2DataListCampaignType {
	return &v
}
