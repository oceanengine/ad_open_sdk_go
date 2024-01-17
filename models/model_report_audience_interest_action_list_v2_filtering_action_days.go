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

// ReportAudienceInterestActionListV2FilteringActionDays
type ReportAudienceInterestActionListV2FilteringActionDays int64

// List of report_audience_interest_action_list_v2_filtering_action_days
const (
	Enum_7_ReportAudienceInterestActionListV2FilteringActionDays   ReportAudienceInterestActionListV2FilteringActionDays = 7
	Enum_15_ReportAudienceInterestActionListV2FilteringActionDays  ReportAudienceInterestActionListV2FilteringActionDays = 15
	Enum_30_ReportAudienceInterestActionListV2FilteringActionDays  ReportAudienceInterestActionListV2FilteringActionDays = 30
	Enum_60_ReportAudienceInterestActionListV2FilteringActionDays  ReportAudienceInterestActionListV2FilteringActionDays = 60
	Enum_90_ReportAudienceInterestActionListV2FilteringActionDays  ReportAudienceInterestActionListV2FilteringActionDays = 90
	Enum_180_ReportAudienceInterestActionListV2FilteringActionDays ReportAudienceInterestActionListV2FilteringActionDays = 180
	Enum_365_ReportAudienceInterestActionListV2FilteringActionDays ReportAudienceInterestActionListV2FilteringActionDays = 365
)

// All allowed values of ReportAudienceInterestActionListV2FilteringActionDays enum
var AllowedReportAudienceInterestActionListV2FilteringActionDaysEnumValues = []ReportAudienceInterestActionListV2FilteringActionDays{
	7,
	15,
	30,
	60,
	90,
	180,
	365,
}

// NewReportAudienceInterestActionListV2FilteringActionDaysFromValue returns a pointer to a valid ReportAudienceInterestActionListV2FilteringActionDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAudienceInterestActionListV2FilteringActionDaysFromValue(v int64) (*ReportAudienceInterestActionListV2FilteringActionDays, error) {
	ev := ReportAudienceInterestActionListV2FilteringActionDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAudienceInterestActionListV2FilteringActionDays: valid values are %v", v, AllowedReportAudienceInterestActionListV2FilteringActionDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAudienceInterestActionListV2FilteringActionDays) IsValid() bool {
	for _, existing := range AllowedReportAudienceInterestActionListV2FilteringActionDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_audience_interest_action_list_v2_filtering_action_days value
func (v ReportAudienceInterestActionListV2FilteringActionDays) Ptr() *ReportAudienceInterestActionListV2FilteringActionDays {
	return &v
}
