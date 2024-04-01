/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdGetV10FilteringStatus
type QianchuanAdGetV10FilteringStatus string

// List of qianchuan_ad_get_v1.0_filtering_status
const (
	ALL_QianchuanAdGetV10FilteringStatus                  QianchuanAdGetV10FilteringStatus = "ALL"
	ALL_INCLUDE_DELETED_QianchuanAdGetV10FilteringStatus  QianchuanAdGetV10FilteringStatus = "ALL_INCLUDE_DELETED"
	AUDIT_QianchuanAdGetV10FilteringStatus                QianchuanAdGetV10FilteringStatus = "AUDIT"
	DELETED_QianchuanAdGetV10FilteringStatus              QianchuanAdGetV10FilteringStatus = "DELETED"
	DELIVERY_OK_QianchuanAdGetV10FilteringStatus          QianchuanAdGetV10FilteringStatus = "DELIVERY_OK"
	DISABLE_QianchuanAdGetV10FilteringStatus              QianchuanAdGetV10FilteringStatus = "DISABLE"
	EXTERNAL_URL_DISABLE_QianchuanAdGetV10FilteringStatus QianchuanAdGetV10FilteringStatus = "EXTERNAL_URL_DISABLE"
	FROZEN_QianchuanAdGetV10FilteringStatus               QianchuanAdGetV10FilteringStatus = "FROZEN"
	LIVE_ROOM_OFF_QianchuanAdGetV10FilteringStatus        QianchuanAdGetV10FilteringStatus = "LIVE_ROOM_OFF"
	NO_SCHEDULE_QianchuanAdGetV10FilteringStatus          QianchuanAdGetV10FilteringStatus = "NO_SCHEDULE"
	OFFLINE_AUDIT_QianchuanAdGetV10FilteringStatus        QianchuanAdGetV10FilteringStatus = "OFFLINE_AUDIT"
	OFFLINE_BALANCE_QianchuanAdGetV10FilteringStatus      QianchuanAdGetV10FilteringStatus = "OFFLINE_BALANCE"
	OFFLINE_BUDGET_QianchuanAdGetV10FilteringStatus       QianchuanAdGetV10FilteringStatus = "OFFLINE_BUDGET"
	QUOTA_DISABLE_QianchuanAdGetV10FilteringStatus        QianchuanAdGetV10FilteringStatus = "QUOTA_DISABLE"
	REAUDIT_QianchuanAdGetV10FilteringStatus              QianchuanAdGetV10FilteringStatus = "REAUDIT"
	ROI2_DISABLE_QianchuanAdGetV10FilteringStatus         QianchuanAdGetV10FilteringStatus = "ROI2_DISABLE"
	SYSTEM_DISABLE_QianchuanAdGetV10FilteringStatus       QianchuanAdGetV10FilteringStatus = "SYSTEM_DISABLE"
	TIME_DONE_QianchuanAdGetV10FilteringStatus            QianchuanAdGetV10FilteringStatus = "TIME_DONE"
	TIME_NO_REACH_QianchuanAdGetV10FilteringStatus        QianchuanAdGetV10FilteringStatus = "TIME_NO_REACH"
)

// All allowed values of QianchuanAdGetV10FilteringStatus enum
var AllowedQianchuanAdGetV10FilteringStatusEnumValues = []QianchuanAdGetV10FilteringStatus{
	"ALL",
	"ALL_INCLUDE_DELETED",
	"AUDIT",
	"DELETED",
	"DELIVERY_OK",
	"DISABLE",
	"EXTERNAL_URL_DISABLE",
	"FROZEN",
	"LIVE_ROOM_OFF",
	"NO_SCHEDULE",
	"OFFLINE_AUDIT",
	"OFFLINE_BALANCE",
	"OFFLINE_BUDGET",
	"QUOTA_DISABLE",
	"REAUDIT",
	"ROI2_DISABLE",
	"SYSTEM_DISABLE",
	"TIME_DONE",
	"TIME_NO_REACH",
}

// NewQianchuanAdGetV10FilteringStatusFromValue returns a pointer to a valid QianchuanAdGetV10FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10FilteringStatusFromValue(v string) (*QianchuanAdGetV10FilteringStatus, error) {
	ev := QianchuanAdGetV10FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10FilteringStatus: valid values are %v", v, AllowedQianchuanAdGetV10FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10FilteringStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_filtering_status value
func (v QianchuanAdGetV10FilteringStatus) Ptr() *QianchuanAdGetV10FilteringStatus {
	return &v
}
