/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupListV30DataGroupsStatus
type AdlabGroupListV30DataGroupsStatus string

// List of adlab_group_list_v3.0_data_groups_status
const (
	ADV_BUDGET_EXCEED_AdlabGroupListV30DataGroupsStatus      AdlabGroupListV30DataGroupsStatus = "ADV_BUDGET_EXCEED"
	AUDIT_FAILED_AdlabGroupListV30DataGroupsStatus           AdlabGroupListV30DataGroupsStatus = "AUDIT_FAILED"
	AUTO_STOP_AdlabGroupListV30DataGroupsStatus              AdlabGroupListV30DataGroupsStatus = "AUTO_STOP"
	BALANCE_EXCEED_AdlabGroupListV30DataGroupsStatus         AdlabGroupListV30DataGroupsStatus = "BALANCE_EXCEED"
	CAMPAIGN_DELETED_AdlabGroupListV30DataGroupsStatus       AdlabGroupListV30DataGroupsStatus = "CAMPAIGN_DELETED"
	CAMPAIGN_DISABLE_AdlabGroupListV30DataGroupsStatus       AdlabGroupListV30DataGroupsStatus = "CAMPAIGN_DISABLE"
	CREATING_AdlabGroupListV30DataGroupsStatus               AdlabGroupListV30DataGroupsStatus = "CREATING"
	DELETED_AdlabGroupListV30DataGroupsStatus                AdlabGroupListV30DataGroupsStatus = "DELETED"
	DELIVERY_OK_AdlabGroupListV30DataGroupsStatus            AdlabGroupListV30DataGroupsStatus = "DELIVERY_OK"
	DISABLE_AdlabGroupListV30DataGroupsStatus                AdlabGroupListV30DataGroupsStatus = "DISABLE"
	GROUP_GOODS_CLOSE_AdlabGroupListV30DataGroupsStatus      AdlabGroupListV30DataGroupsStatus = "GROUP_GOODS_CLOSE"
	INITIALIZING_AdlabGroupListV30DataGroupsStatus           AdlabGroupListV30DataGroupsStatus = "INITIALIZING"
	INIT_FAILED_AdlabGroupListV30DataGroupsStatus            AdlabGroupListV30DataGroupsStatus = "INIT_FAILED"
	INVALID_STATUS_AdlabGroupListV30DataGroupsStatus         AdlabGroupListV30DataGroupsStatus = "INVALID_STATUS"
	LACK_OF_POSTERIOR_DATA_AdlabGroupListV30DataGroupsStatus AdlabGroupListV30DataGroupsStatus = "LACK_OF_POSTERIOR_DATA"
	LIVE_CAN_NOT_LAUNCH_AdlabGroupListV30DataGroupsStatus    AdlabGroupListV30DataGroupsStatus = "LIVE_CAN_NOT_LAUNCH"
	NO_SCHEDULE_AdlabGroupListV30DataGroupsStatus            AdlabGroupListV30DataGroupsStatus = "NO_SCHEDULE"
	OUT_OF_BUDGET_AdlabGroupListV30DataGroupsStatus          AdlabGroupListV30DataGroupsStatus = "OUT_OF_BUDGET"
	OUT_OF_CREATIVE_AdlabGroupListV30DataGroupsStatus        AdlabGroupListV30DataGroupsStatus = "OUT_OF_CREATIVE"
	OUT_OF_DAILY_BUDGET_AdlabGroupListV30DataGroupsStatus    AdlabGroupListV30DataGroupsStatus = "OUT_OF_DAILY_BUDGET"
	OUT_OF_DAILY_TIME_AdlabGroupListV30DataGroupsStatus      AdlabGroupListV30DataGroupsStatus = "OUT_OF_DAILY_TIME"
	OUT_OF_QUOTA_AdlabGroupListV30DataGroupsStatus           AdlabGroupListV30DataGroupsStatus = "OUT_OF_QUOTA"
	OUT_OF_TIME_AdlabGroupListV30DataGroupsStatus            AdlabGroupListV30DataGroupsStatus = "OUT_OF_TIME"
	READY_AdlabGroupListV30DataGroupsStatus                  AdlabGroupListV30DataGroupsStatus = "READY"
	RECOVERABLE_ERROR_AdlabGroupListV30DataGroupsStatus      AdlabGroupListV30DataGroupsStatus = "RECOVERABLE_ERROR"
	SYSTEM_ERROR_AdlabGroupListV30DataGroupsStatus           AdlabGroupListV30DataGroupsStatus = "SYSTEM_ERROR"
)

// All allowed values of AdlabGroupListV30DataGroupsStatus enum
var AllowedAdlabGroupListV30DataGroupsStatusEnumValues = []AdlabGroupListV30DataGroupsStatus{
	"ADV_BUDGET_EXCEED",
	"AUDIT_FAILED",
	"AUTO_STOP",
	"BALANCE_EXCEED",
	"CAMPAIGN_DELETED",
	"CAMPAIGN_DISABLE",
	"CREATING",
	"DELETED",
	"DELIVERY_OK",
	"DISABLE",
	"GROUP_GOODS_CLOSE",
	"INITIALIZING",
	"INIT_FAILED",
	"INVALID_STATUS",
	"LACK_OF_POSTERIOR_DATA",
	"LIVE_CAN_NOT_LAUNCH",
	"NO_SCHEDULE",
	"OUT_OF_BUDGET",
	"OUT_OF_CREATIVE",
	"OUT_OF_DAILY_BUDGET",
	"OUT_OF_DAILY_TIME",
	"OUT_OF_QUOTA",
	"OUT_OF_TIME",
	"READY",
	"RECOVERABLE_ERROR",
	"SYSTEM_ERROR",
}

// NewAdlabGroupListV30DataGroupsStatusFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsStatusFromValue(v string) (*AdlabGroupListV30DataGroupsStatus, error) {
	ev := AdlabGroupListV30DataGroupsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsStatus: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsStatus) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_status value
func (v AdlabGroupListV30DataGroupsStatus) Ptr() *AdlabGroupListV30DataGroupsStatus {
	return &v
}
