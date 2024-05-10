/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StardeliveryTaskDetailV30DataStarTaskSubStatus
type StardeliveryTaskDetailV30DataStarTaskSubStatus string

// List of stardelivery_task_detail_v3.0_data_star_task_sub_status
const (
	AWAITING_ISV_ACCEPT_StardeliveryTaskDetailV30DataStarTaskSubStatus    StardeliveryTaskDetailV30DataStarTaskSubStatus = "AWAITING_ISV_ACCEPT"
	CALCULATING_COSTS_StardeliveryTaskDetailV30DataStarTaskSubStatus      StardeliveryTaskDetailV30DataStarTaskSubStatus = "CALCULATING_COSTS"
	NO_ISV_ACCEPT_StardeliveryTaskDetailV30DataStarTaskSubStatus          StardeliveryTaskDetailV30DataStarTaskSubStatus = "NO_ISV_ACCEPT"
	REJECTED_StardeliveryTaskDetailV30DataStarTaskSubStatus               StardeliveryTaskDetailV30DataStarTaskSubStatus = "REJECTED"
	SUBMISSION_IN_PROGRESS_StardeliveryTaskDetailV30DataStarTaskSubStatus StardeliveryTaskDetailV30DataStarTaskSubStatus = "SUBMISSION_IN_PROGRESS"
	UNDER_REVIEW_StardeliveryTaskDetailV30DataStarTaskSubStatus           StardeliveryTaskDetailV30DataStarTaskSubStatus = "UNDER_REVIEW"
)

// All allowed values of StardeliveryTaskDetailV30DataStarTaskSubStatus enum
var AllowedStardeliveryTaskDetailV30DataStarTaskSubStatusEnumValues = []StardeliveryTaskDetailV30DataStarTaskSubStatus{
	"AWAITING_ISV_ACCEPT",
	"CALCULATING_COSTS",
	"NO_ISV_ACCEPT",
	"REJECTED",
	"SUBMISSION_IN_PROGRESS",
	"UNDER_REVIEW",
}

// NewStardeliveryTaskDetailV30DataStarTaskSubStatusFromValue returns a pointer to a valid StardeliveryTaskDetailV30DataStarTaskSubStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskDetailV30DataStarTaskSubStatusFromValue(v string) (*StardeliveryTaskDetailV30DataStarTaskSubStatus, error) {
	ev := StardeliveryTaskDetailV30DataStarTaskSubStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskDetailV30DataStarTaskSubStatus: valid values are %v", v, AllowedStardeliveryTaskDetailV30DataStarTaskSubStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskDetailV30DataStarTaskSubStatus) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskDetailV30DataStarTaskSubStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_detail_v3.0_data_star_task_sub_status value
func (v StardeliveryTaskDetailV30DataStarTaskSubStatus) Ptr() *StardeliveryTaskDetailV30DataStarTaskSubStatus {
	return &v
}
