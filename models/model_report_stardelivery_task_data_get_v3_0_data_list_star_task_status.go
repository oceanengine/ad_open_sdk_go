/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportStardeliveryTaskDataGetV30DataListStarTaskStatus
type ReportStardeliveryTaskDataGetV30DataListStarTaskStatus string

// List of report_stardelivery_task_data_get_v3.0_data_list_star_task_status
const (
	BILLING_IN_PROGRESS_ReportStardeliveryTaskDataGetV30DataListStarTaskStatus ReportStardeliveryTaskDataGetV30DataListStarTaskStatus = "BILLING_IN_PROGRESS"
	CANCELLED_ReportStardeliveryTaskDataGetV30DataListStarTaskStatus           ReportStardeliveryTaskDataGetV30DataListStarTaskStatus = "CANCELLED"
	COMPLETED_ReportStardeliveryTaskDataGetV30DataListStarTaskStatus           ReportStardeliveryTaskDataGetV30DataListStarTaskStatus = "COMPLETED"
	PENDING_LAUNCH_ReportStardeliveryTaskDataGetV30DataListStarTaskStatus      ReportStardeliveryTaskDataGetV30DataListStarTaskStatus = "PENDING_LAUNCH"
	PREPARATION_ReportStardeliveryTaskDataGetV30DataListStarTaskStatus         ReportStardeliveryTaskDataGetV30DataListStarTaskStatus = "PREPARATION"
)

// All allowed values of ReportStardeliveryTaskDataGetV30DataListStarTaskStatus enum
var AllowedReportStardeliveryTaskDataGetV30DataListStarTaskStatusEnumValues = []ReportStardeliveryTaskDataGetV30DataListStarTaskStatus{
	"BILLING_IN_PROGRESS",
	"CANCELLED",
	"COMPLETED",
	"PENDING_LAUNCH",
	"PREPARATION",
}

// NewReportStardeliveryTaskDataGetV30DataListStarTaskStatusFromValue returns a pointer to a valid ReportStardeliveryTaskDataGetV30DataListStarTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportStardeliveryTaskDataGetV30DataListStarTaskStatusFromValue(v string) (*ReportStardeliveryTaskDataGetV30DataListStarTaskStatus, error) {
	ev := ReportStardeliveryTaskDataGetV30DataListStarTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportStardeliveryTaskDataGetV30DataListStarTaskStatus: valid values are %v", v, AllowedReportStardeliveryTaskDataGetV30DataListStarTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportStardeliveryTaskDataGetV30DataListStarTaskStatus) IsValid() bool {
	for _, existing := range AllowedReportStardeliveryTaskDataGetV30DataListStarTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_stardelivery_task_data_get_v3.0_data_list_star_task_status value
func (v ReportStardeliveryTaskDataGetV30DataListStarTaskStatus) Ptr() *ReportStardeliveryTaskDataGetV30DataListStarTaskStatus {
	return &v
}
