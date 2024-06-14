/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus
type StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus string

// List of star_vas_get_boost_group_list_v2_data_boost_group_infos_task_infos_audit_status
const (
	AUDITING_StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus     StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus = "AUDITING"
	AUDIT_FAILED_StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus = "AUDIT_FAILED"
	AUDIT_PASS_StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus   StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus = "AUDIT_PASS"
)

// All allowed values of StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus enum
var AllowedStarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatusEnumValues = []StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus{
	"AUDITING",
	"AUDIT_FAILED",
	"AUDIT_PASS",
}

// NewStarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatusFromValue returns a pointer to a valid StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatusFromValue(v string) (*StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus, error) {
	ev := StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus: valid values are %v", v, AllowedStarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus) IsValid() bool {
	for _, existing := range AllowedStarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_vas_get_boost_group_list_v2_data_boost_group_infos_task_infos_audit_status value
func (v StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus) Ptr() *StarVasGetBoostGroupListV2DataBoostGroupInfosTaskInfosAuditStatus {
	return &v
}