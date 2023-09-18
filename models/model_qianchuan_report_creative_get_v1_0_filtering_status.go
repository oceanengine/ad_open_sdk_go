/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportCreativeGetV10FilteringStatus
type QianchuanReportCreativeGetV10FilteringStatus string

// List of qianchuan_report_creative_get_v1.0_filtering_status
const (
	ALL_INCLUDE_DELETED_QianchuanReportCreativeGetV10FilteringStatus  QianchuanReportCreativeGetV10FilteringStatus = "ALL_INCLUDE_DELETED"
	AUDIT_QianchuanReportCreativeGetV10FilteringStatus                QianchuanReportCreativeGetV10FilteringStatus = "AUDIT"
	DELETED_QianchuanReportCreativeGetV10FilteringStatus              QianchuanReportCreativeGetV10FilteringStatus = "DELETED"
	DELIVERY_OK_QianchuanReportCreativeGetV10FilteringStatus          QianchuanReportCreativeGetV10FilteringStatus = "DELIVERY_OK"
	DISABLE_QianchuanReportCreativeGetV10FilteringStatus              QianchuanReportCreativeGetV10FilteringStatus = "DISABLE"
	EXTERNAL_URL_DISABLE_QianchuanReportCreativeGetV10FilteringStatus QianchuanReportCreativeGetV10FilteringStatus = "EXTERNAL_URL_DISABLE"
	FROZEN_QianchuanReportCreativeGetV10FilteringStatus               QianchuanReportCreativeGetV10FilteringStatus = "FROZEN"
	LIVE_ROOM_OFF_QianchuanReportCreativeGetV10FilteringStatus        QianchuanReportCreativeGetV10FilteringStatus = "LIVE_ROOM_OFF"
	NO_SCHEDULE_QianchuanReportCreativeGetV10FilteringStatus          QianchuanReportCreativeGetV10FilteringStatus = "NO_SCHEDULE"
	OFFLINE_AUDIT_QianchuanReportCreativeGetV10FilteringStatus        QianchuanReportCreativeGetV10FilteringStatus = "OFFLINE_AUDIT"
	OFFLINE_BALANCE_QianchuanReportCreativeGetV10FilteringStatus      QianchuanReportCreativeGetV10FilteringStatus = "OFFLINE_BALANCE"
	OFFLINE_BUDGET_QianchuanReportCreativeGetV10FilteringStatus       QianchuanReportCreativeGetV10FilteringStatus = "OFFLINE_BUDGET"
	QUOTA_DISABLE_QianchuanReportCreativeGetV10FilteringStatus        QianchuanReportCreativeGetV10FilteringStatus = "QUOTA_DISABLE"
	REAUDIT_QianchuanReportCreativeGetV10FilteringStatus              QianchuanReportCreativeGetV10FilteringStatus = "REAUDIT"
	SYSTEM_DISABLE_QianchuanReportCreativeGetV10FilteringStatus       QianchuanReportCreativeGetV10FilteringStatus = "SYSTEM_DISABLE"
	TIME_DONE_QianchuanReportCreativeGetV10FilteringStatus            QianchuanReportCreativeGetV10FilteringStatus = "TIME_DONE"
	TIME_NO_REACH_QianchuanReportCreativeGetV10FilteringStatus        QianchuanReportCreativeGetV10FilteringStatus = "TIME_NO_REACH"
)

// All allowed values of QianchuanReportCreativeGetV10FilteringStatus enum
var AllowedQianchuanReportCreativeGetV10FilteringStatusEnumValues = []QianchuanReportCreativeGetV10FilteringStatus{
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
	"SYSTEM_DISABLE",
	"TIME_DONE",
	"TIME_NO_REACH",
}

// NewQianchuanReportCreativeGetV10FilteringStatusFromValue returns a pointer to a valid QianchuanReportCreativeGetV10FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportCreativeGetV10FilteringStatusFromValue(v string) (*QianchuanReportCreativeGetV10FilteringStatus, error) {
	ev := QianchuanReportCreativeGetV10FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportCreativeGetV10FilteringStatus: valid values are %v", v, AllowedQianchuanReportCreativeGetV10FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportCreativeGetV10FilteringStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanReportCreativeGetV10FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_creative_get_v1.0_filtering_status value
func (v QianchuanReportCreativeGetV10FilteringStatus) Ptr() *QianchuanReportCreativeGetV10FilteringStatus {
	return &v
}