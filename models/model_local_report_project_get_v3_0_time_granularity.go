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

// LocalReportProjectGetV30TimeGranularity
type LocalReportProjectGetV30TimeGranularity string

// List of local_report_project_get_v3.0_time_granularity
const (
	TIME_GRANULARITY_DAILY_LocalReportProjectGetV30TimeGranularity  LocalReportProjectGetV30TimeGranularity = "TIME_GRANULARITY_DAILY"
	TIME_GRANULARITY_HOURLY_LocalReportProjectGetV30TimeGranularity LocalReportProjectGetV30TimeGranularity = "TIME_GRANULARITY_HOURLY"
	TIME_GRANULARITY_TOTAL_LocalReportProjectGetV30TimeGranularity  LocalReportProjectGetV30TimeGranularity = "TIME_GRANULARITY_TOTAL"
)

// All allowed values of LocalReportProjectGetV30TimeGranularity enum
var AllowedLocalReportProjectGetV30TimeGranularityEnumValues = []LocalReportProjectGetV30TimeGranularity{
	"TIME_GRANULARITY_DAILY",
	"TIME_GRANULARITY_HOURLY",
	"TIME_GRANULARITY_TOTAL",
}

// NewLocalReportProjectGetV30TimeGranularityFromValue returns a pointer to a valid LocalReportProjectGetV30TimeGranularity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocalReportProjectGetV30TimeGranularityFromValue(v string) (*LocalReportProjectGetV30TimeGranularity, error) {
	ev := LocalReportProjectGetV30TimeGranularity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocalReportProjectGetV30TimeGranularity: valid values are %v", v, AllowedLocalReportProjectGetV30TimeGranularityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocalReportProjectGetV30TimeGranularity) IsValid() bool {
	for _, existing := range AllowedLocalReportProjectGetV30TimeGranularityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to local_report_project_get_v3.0_time_granularity value
func (v LocalReportProjectGetV30TimeGranularity) Ptr() *LocalReportProjectGetV30TimeGranularity {
	return &v
}
