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

// LocalReportProjectGetV30FilteringCampaignType
type LocalReportProjectGetV30FilteringCampaignType string

// List of local_report_project_get_v3.0_filtering_campaign_type
const (
	GENERAL_LocalReportProjectGetV30FilteringCampaignType   LocalReportProjectGetV30FilteringCampaignType = "GENERAL"
	SEARCHING_LocalReportProjectGetV30FilteringCampaignType LocalReportProjectGetV30FilteringCampaignType = "SEARCHING"
)

// All allowed values of LocalReportProjectGetV30FilteringCampaignType enum
var AllowedLocalReportProjectGetV30FilteringCampaignTypeEnumValues = []LocalReportProjectGetV30FilteringCampaignType{
	"GENERAL",
	"SEARCHING",
}

// NewLocalReportProjectGetV30FilteringCampaignTypeFromValue returns a pointer to a valid LocalReportProjectGetV30FilteringCampaignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocalReportProjectGetV30FilteringCampaignTypeFromValue(v string) (*LocalReportProjectGetV30FilteringCampaignType, error) {
	ev := LocalReportProjectGetV30FilteringCampaignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocalReportProjectGetV30FilteringCampaignType: valid values are %v", v, AllowedLocalReportProjectGetV30FilteringCampaignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocalReportProjectGetV30FilteringCampaignType) IsValid() bool {
	for _, existing := range AllowedLocalReportProjectGetV30FilteringCampaignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to local_report_project_get_v3.0_filtering_campaign_type value
func (v LocalReportProjectGetV30FilteringCampaignType) Ptr() *LocalReportProjectGetV30FilteringCampaignType {
	return &v
}
