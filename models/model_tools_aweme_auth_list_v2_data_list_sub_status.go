/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAwemeAuthListV2DataListSubStatus
type ToolsAwemeAuthListV2DataListSubStatus string

// List of tools_aweme_auth_list_v2_data_list_sub_status
const (
	INVALID_CANCEL_ToolsAwemeAuthListV2DataListSubStatus   ToolsAwemeAuthListV2DataListSubStatus = "INVALID_CANCEL"
	INVALID_EXPIRED_ToolsAwemeAuthListV2DataListSubStatus  ToolsAwemeAuthListV2DataListSubStatus = "INVALID_EXPIRED"
	INVALID_REJECT_ToolsAwemeAuthListV2DataListSubStatus   ToolsAwemeAuthListV2DataListSubStatus = "INVALID_REJECT"
	INVALID_TIME_OUT_ToolsAwemeAuthListV2DataListSubStatus ToolsAwemeAuthListV2DataListSubStatus = "INVALID_TIME_OUT"
	RENEWING_ToolsAwemeAuthListV2DataListSubStatus         ToolsAwemeAuthListV2DataListSubStatus = "RENEWING"
	RENEW_FAIL_ToolsAwemeAuthListV2DataListSubStatus       ToolsAwemeAuthListV2DataListSubStatus = "RENEW_FAIL"
	RENEW_SUCCESS_ToolsAwemeAuthListV2DataListSubStatus    ToolsAwemeAuthListV2DataListSubStatus = "RENEW_SUCCESS"
)

// All allowed values of ToolsAwemeAuthListV2DataListSubStatus enum
var AllowedToolsAwemeAuthListV2DataListSubStatusEnumValues = []ToolsAwemeAuthListV2DataListSubStatus{
	"INVALID_CANCEL",
	"INVALID_EXPIRED",
	"INVALID_REJECT",
	"INVALID_TIME_OUT",
	"RENEWING",
	"RENEW_FAIL",
	"RENEW_SUCCESS",
}

// NewToolsAwemeAuthListV2DataListSubStatusFromValue returns a pointer to a valid ToolsAwemeAuthListV2DataListSubStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeAuthListV2DataListSubStatusFromValue(v string) (*ToolsAwemeAuthListV2DataListSubStatus, error) {
	ev := ToolsAwemeAuthListV2DataListSubStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeAuthListV2DataListSubStatus: valid values are %v", v, AllowedToolsAwemeAuthListV2DataListSubStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeAuthListV2DataListSubStatus) IsValid() bool {
	for _, existing := range AllowedToolsAwemeAuthListV2DataListSubStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_auth_list_v2_data_list_sub_status value
func (v ToolsAwemeAuthListV2DataListSubStatus) Ptr() *ToolsAwemeAuthListV2DataListSubStatus {
	return &v
}
