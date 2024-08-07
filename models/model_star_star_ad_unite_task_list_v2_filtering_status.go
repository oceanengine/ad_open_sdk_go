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

// StarStarAdUniteTaskListV2FilteringStatus
type StarStarAdUniteTaskListV2FilteringStatus string

// List of star_star_ad_unite_task_list_v2_filtering_status
const (
	ALL_StarStarAdUniteTaskListV2FilteringStatus           StarStarAdUniteTaskListV2FilteringStatus = "ALL"
	CANCELED_StarStarAdUniteTaskListV2FilteringStatus      StarStarAdUniteTaskListV2FilteringStatus = "CANCELED"
	FINISHED_StarStarAdUniteTaskListV2FilteringStatus      StarStarAdUniteTaskListV2FilteringStatus = "FINISHED"
	ONGOING_StarStarAdUniteTaskListV2FilteringStatus       StarStarAdUniteTaskListV2FilteringStatus = "ONGOING"
	RECEIVEING_StarStarAdUniteTaskListV2FilteringStatus    StarStarAdUniteTaskListV2FilteringStatus = "RECEIVEING"
	WAIT_EVALUATE_StarStarAdUniteTaskListV2FilteringStatus StarStarAdUniteTaskListV2FilteringStatus = "WAIT_EVALUATE"
	WAIT_PAYMENT_StarStarAdUniteTaskListV2FilteringStatus  StarStarAdUniteTaskListV2FilteringStatus = "WAIT_PAYMENT"
)

// All allowed values of StarStarAdUniteTaskListV2FilteringStatus enum
var AllowedStarStarAdUniteTaskListV2FilteringStatusEnumValues = []StarStarAdUniteTaskListV2FilteringStatus{
	"ALL",
	"CANCELED",
	"FINISHED",
	"ONGOING",
	"RECEIVEING",
	"WAIT_EVALUATE",
	"WAIT_PAYMENT",
}

// NewStarStarAdUniteTaskListV2FilteringStatusFromValue returns a pointer to a valid StarStarAdUniteTaskListV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarStarAdUniteTaskListV2FilteringStatusFromValue(v string) (*StarStarAdUniteTaskListV2FilteringStatus, error) {
	ev := StarStarAdUniteTaskListV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarStarAdUniteTaskListV2FilteringStatus: valid values are %v", v, AllowedStarStarAdUniteTaskListV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarStarAdUniteTaskListV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedStarStarAdUniteTaskListV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_star_ad_unite_task_list_v2_filtering_status value
func (v StarStarAdUniteTaskListV2FilteringStatus) Ptr() *StarStarAdUniteTaskListV2FilteringStatus {
	return &v
}
