/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserInfoV2DataStatus
type AdvertiserInfoV2DataStatus string

// List of advertiser_info_v2_data_status
const (
	STATUS_DISABLE_AdvertiserInfoV2DataStatus                AdvertiserInfoV2DataStatus = "STATUS_DISABLE"
	STATUS_LIMIT_AdvertiserInfoV2DataStatus                  AdvertiserInfoV2DataStatus = "STATUS_LIMIT"
	STATUS_PENDING_CONFIRM_AdvertiserInfoV2DataStatus        AdvertiserInfoV2DataStatus = "STATUS_PENDING_CONFIRM"
	STATUS_WAIT_FOR_BPM_AUDIT_AdvertiserInfoV2DataStatus     AdvertiserInfoV2DataStatus = "STATUS_WAIT_FOR_BPM_AUDIT"
	STATUS_CONFIRM_FAIL_END_AdvertiserInfoV2DataStatus       AdvertiserInfoV2DataStatus = "STATUS_CONFIRM_FAIL_END"
	STATUS_ENABLE_AdvertiserInfoV2DataStatus                 AdvertiserInfoV2DataStatus = "STATUS_ENABLE"
	STATUS_SELF_SERVICE_UNAUDITED_AdvertiserInfoV2DataStatus AdvertiserInfoV2DataStatus = "STATUS_SELF_SERVICE_UNAUDITED"
	STATUS_WAIT_FOR_PUBLIC_AUTH_AdvertiserInfoV2DataStatus   AdvertiserInfoV2DataStatus = "STATUS_WAIT_FOR_PUBLIC_AUTH"
	STATUS_CONFIRM_MODIFY_FAIL_AdvertiserInfoV2DataStatus    AdvertiserInfoV2DataStatus = "STATUS_CONFIRM_MODIFY_FAIL"
	STATUS_CONFIRM_FAIL_AdvertiserInfoV2DataStatus           AdvertiserInfoV2DataStatus = "STATUS_CONFIRM_FAIL"
	STATUS_PENDING_VERIFIED_AdvertiserInfoV2DataStatus       AdvertiserInfoV2DataStatus = "STATUS_PENDING_VERIFIED"
	STATUS_PENDING_CONFIRM_MODIFY_AdvertiserInfoV2DataStatus AdvertiserInfoV2DataStatus = "STATUS_PENDING_CONFIRM_MODIFY"
)

// All allowed values of AdvertiserInfoV2DataStatus enum
var AllowedAdvertiserInfoV2DataStatusEnumValues = []AdvertiserInfoV2DataStatus{
	"STATUS_DISABLE",
	"STATUS_LIMIT",
	"STATUS_PENDING_CONFIRM",
	"STATUS_WAIT_FOR_BPM_AUDIT",
	"STATUS_CONFIRM_FAIL_END",
	"STATUS_ENABLE",
	"STATUS_SELF_SERVICE_UNAUDITED",
	"STATUS_WAIT_FOR_PUBLIC_AUTH",
	"STATUS_CONFIRM_MODIFY_FAIL",
	"STATUS_CONFIRM_FAIL",
	"STATUS_PENDING_VERIFIED",
	"STATUS_PENDING_CONFIRM_MODIFY",
}

// NewAdvertiserInfoV2DataStatusFromValue returns a pointer to a valid AdvertiserInfoV2DataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserInfoV2DataStatusFromValue(v string) (*AdvertiserInfoV2DataStatus, error) {
	ev := AdvertiserInfoV2DataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserInfoV2DataStatus: valid values are %v", v, AllowedAdvertiserInfoV2DataStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserInfoV2DataStatus) IsValid() bool {
	for _, existing := range AllowedAdvertiserInfoV2DataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_info_v2_data_status value
func (v AdvertiserInfoV2DataStatus) Ptr() *AdvertiserInfoV2DataStatus {
	return &v
}