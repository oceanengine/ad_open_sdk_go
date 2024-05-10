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

// StarStarAdUniteTaskDetailV2DataAuditStatus
type StarStarAdUniteTaskDetailV2DataAuditStatus string

// List of star_star_ad_unite_task_detail_v2_data_audit_status
const (
	CONFIRM_StarStarAdUniteTaskDetailV2DataAuditStatus         StarStarAdUniteTaskDetailV2DataAuditStatus = "CONFIRM"
	CONFIRM_FAIL_StarStarAdUniteTaskDetailV2DataAuditStatus    StarStarAdUniteTaskDetailV2DataAuditStatus = "CONFIRM_FAIL"
	PENDING_CONFIRM_StarStarAdUniteTaskDetailV2DataAuditStatus StarStarAdUniteTaskDetailV2DataAuditStatus = "PENDING_CONFIRM"
)

// All allowed values of StarStarAdUniteTaskDetailV2DataAuditStatus enum
var AllowedStarStarAdUniteTaskDetailV2DataAuditStatusEnumValues = []StarStarAdUniteTaskDetailV2DataAuditStatus{
	"CONFIRM",
	"CONFIRM_FAIL",
	"PENDING_CONFIRM",
}

// NewStarStarAdUniteTaskDetailV2DataAuditStatusFromValue returns a pointer to a valid StarStarAdUniteTaskDetailV2DataAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarStarAdUniteTaskDetailV2DataAuditStatusFromValue(v string) (*StarStarAdUniteTaskDetailV2DataAuditStatus, error) {
	ev := StarStarAdUniteTaskDetailV2DataAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarStarAdUniteTaskDetailV2DataAuditStatus: valid values are %v", v, AllowedStarStarAdUniteTaskDetailV2DataAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarStarAdUniteTaskDetailV2DataAuditStatus) IsValid() bool {
	for _, existing := range AllowedStarStarAdUniteTaskDetailV2DataAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_star_ad_unite_task_detail_v2_data_audit_status value
func (v StarStarAdUniteTaskDetailV2DataAuditStatus) Ptr() *StarStarAdUniteTaskDetailV2DataAuditStatus {
	return &v
}
