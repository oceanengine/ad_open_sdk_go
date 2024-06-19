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

// StardeliveryTaskDetailV30DataStarTaskStatus
type StardeliveryTaskDetailV30DataStarTaskStatus string

// List of stardelivery_task_detail_v3.0_data_star_task_status
const (
	BILLING_IN_PROGRESS_StardeliveryTaskDetailV30DataStarTaskStatus StardeliveryTaskDetailV30DataStarTaskStatus = "BILLING_IN_PROGRESS"
	CANCELLED_StardeliveryTaskDetailV30DataStarTaskStatus           StardeliveryTaskDetailV30DataStarTaskStatus = "CANCELLED"
	COMPLETED_StardeliveryTaskDetailV30DataStarTaskStatus           StardeliveryTaskDetailV30DataStarTaskStatus = "COMPLETED"
	PENDING_LAUNCH_StardeliveryTaskDetailV30DataStarTaskStatus      StardeliveryTaskDetailV30DataStarTaskStatus = "PENDING_LAUNCH"
	PREPARATION_StardeliveryTaskDetailV30DataStarTaskStatus         StardeliveryTaskDetailV30DataStarTaskStatus = "PREPARATION"
)

// All allowed values of StardeliveryTaskDetailV30DataStarTaskStatus enum
var AllowedStardeliveryTaskDetailV30DataStarTaskStatusEnumValues = []StardeliveryTaskDetailV30DataStarTaskStatus{
	"BILLING_IN_PROGRESS",
	"CANCELLED",
	"COMPLETED",
	"PENDING_LAUNCH",
	"PREPARATION",
}

// NewStardeliveryTaskDetailV30DataStarTaskStatusFromValue returns a pointer to a valid StardeliveryTaskDetailV30DataStarTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskDetailV30DataStarTaskStatusFromValue(v string) (*StardeliveryTaskDetailV30DataStarTaskStatus, error) {
	ev := StardeliveryTaskDetailV30DataStarTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskDetailV30DataStarTaskStatus: valid values are %v", v, AllowedStardeliveryTaskDetailV30DataStarTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskDetailV30DataStarTaskStatus) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskDetailV30DataStarTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_detail_v3.0_data_star_task_status value
func (v StardeliveryTaskDetailV30DataStarTaskStatus) Ptr() *StardeliveryTaskDetailV30DataStarTaskStatus {
	return &v
}
