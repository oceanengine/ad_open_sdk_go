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

// ReportStardeliveryTaskDataGetV30DataListStarTaskSource
type ReportStardeliveryTaskDataGetV30DataListStarTaskSource string

// List of report_stardelivery_task_data_get_v3.0_data_list_star_task_source
const (
	MY_CREATIONS_ReportStardeliveryTaskDataGetV30DataListStarTaskSource  ReportStardeliveryTaskDataGetV30DataListStarTaskSource = "MY_CREATIONS"
	SHARING_ReportStardeliveryTaskDataGetV30DataListStarTaskSource       ReportStardeliveryTaskDataGetV30DataListStarTaskSource = "SHARING"
	SHATE_EXPIRED_ReportStardeliveryTaskDataGetV30DataListStarTaskSource ReportStardeliveryTaskDataGetV30DataListStarTaskSource = "SHATE_EXPIRED"
)

// All allowed values of ReportStardeliveryTaskDataGetV30DataListStarTaskSource enum
var AllowedReportStardeliveryTaskDataGetV30DataListStarTaskSourceEnumValues = []ReportStardeliveryTaskDataGetV30DataListStarTaskSource{
	"MY_CREATIONS",
	"SHARING",
	"SHATE_EXPIRED",
}

// NewReportStardeliveryTaskDataGetV30DataListStarTaskSourceFromValue returns a pointer to a valid ReportStardeliveryTaskDataGetV30DataListStarTaskSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportStardeliveryTaskDataGetV30DataListStarTaskSourceFromValue(v string) (*ReportStardeliveryTaskDataGetV30DataListStarTaskSource, error) {
	ev := ReportStardeliveryTaskDataGetV30DataListStarTaskSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportStardeliveryTaskDataGetV30DataListStarTaskSource: valid values are %v", v, AllowedReportStardeliveryTaskDataGetV30DataListStarTaskSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportStardeliveryTaskDataGetV30DataListStarTaskSource) IsValid() bool {
	for _, existing := range AllowedReportStardeliveryTaskDataGetV30DataListStarTaskSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_stardelivery_task_data_get_v3.0_data_list_star_task_source value
func (v ReportStardeliveryTaskDataGetV30DataListStarTaskSource) Ptr() *ReportStardeliveryTaskDataGetV30DataListStarTaskSource {
	return &v
}
