/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataStatus
type QianchuanAdDetailGetV10DataStatus string

// List of qianchuan_ad_detail_get_v1.0_data_status
const (
	ADVERTISER_OFFLINE_BUDGET_QianchuanAdDetailGetV10DataStatus     QianchuanAdDetailGetV10DataStatus = "ADVERTISER_OFFLINE_BUDGET"
	ADVERTISER_PRE_OFFLINE_BUDGET_QianchuanAdDetailGetV10DataStatus QianchuanAdDetailGetV10DataStatus = "ADVERTISER_PRE_OFFLINE_BUDGET"
	AUDIT_QianchuanAdDetailGetV10DataStatus                         QianchuanAdDetailGetV10DataStatus = "AUDIT"
	AUDIT_STATUS_ERROR_QianchuanAdDetailGetV10DataStatus            QianchuanAdDetailGetV10DataStatus = "AUDIT_STATUS_ERROR"
	AWEME_ACCOUNT_DISABLED_QianchuanAdDetailGetV10DataStatus        QianchuanAdDetailGetV10DataStatus = "AWEME_ACCOUNT_DISABLED"
	CAMPAIGN_DISABLE_QianchuanAdDetailGetV10DataStatus              QianchuanAdDetailGetV10DataStatus = "CAMPAIGN_DISABLE"
	CAMPAIGN_OFFLINE_BUDGET_QianchuanAdDetailGetV10DataStatus       QianchuanAdDetailGetV10DataStatus = "CAMPAIGN_OFFLINE_BUDGET"
	CAMPAIGN_PRE_OFFLINE_BUDGET_QianchuanAdDetailGetV10DataStatus   QianchuanAdDetailGetV10DataStatus = "CAMPAIGN_PRE_OFFLINE_BUDGET"
	CREATE_QianchuanAdDetailGetV10DataStatus                        QianchuanAdDetailGetV10DataStatus = "CREATE"
	DELETE_QianchuanAdDetailGetV10DataStatus                        QianchuanAdDetailGetV10DataStatus = "DELETE"
	DELIVERY_OK_QianchuanAdDetailGetV10DataStatus                   QianchuanAdDetailGetV10DataStatus = "DELIVERY_OK"
	DISABLE_QianchuanAdDetailGetV10DataStatus                       QianchuanAdDetailGetV10DataStatus = "DISABLE"
	DRAFT_QianchuanAdDetailGetV10DataStatus                         QianchuanAdDetailGetV10DataStatus = "DRAFT"
	ERROR_QianchuanAdDetailGetV10DataStatus                         QianchuanAdDetailGetV10DataStatus = "ERROR"
	EXTERNAL_URL_DISABLE_QianchuanAdDetailGetV10DataStatus          QianchuanAdDetailGetV10DataStatus = "EXTERNAL_URL_DISABLE"
	FROZEN_QianchuanAdDetailGetV10DataStatus                        QianchuanAdDetailGetV10DataStatus = "FROZEN"
	LIVE_ROOM_OFF_QianchuanAdDetailGetV10DataStatus                 QianchuanAdDetailGetV10DataStatus = "LIVE_ROOM_OFF"
	NO_SCHEDULE_QianchuanAdDetailGetV10DataStatus                   QianchuanAdDetailGetV10DataStatus = "NO_SCHEDULE"
	OFFLINE_AUDIT_QianchuanAdDetailGetV10DataStatus                 QianchuanAdDetailGetV10DataStatus = "OFFLINE_AUDIT"
	OFFLINE_BALANCE_QianchuanAdDetailGetV10DataStatus               QianchuanAdDetailGetV10DataStatus = "OFFLINE_BALANCE"
	OFFLINE_BUDGET_QianchuanAdDetailGetV10DataStatus                QianchuanAdDetailGetV10DataStatus = "OFFLINE_BUDGET"
	PRE_OFFLINE_BUDGET_QianchuanAdDetailGetV10DataStatus            QianchuanAdDetailGetV10DataStatus = "PRE_OFFLINE_BUDGET"
	PRE_ONLINE_QianchuanAdDetailGetV10DataStatus                    QianchuanAdDetailGetV10DataStatus = "PRE_ONLINE"
	QUOTA_DISABLE_QianchuanAdDetailGetV10DataStatus                 QianchuanAdDetailGetV10DataStatus = "QUOTA_DISABLE"
	REAUDIT_QianchuanAdDetailGetV10DataStatus                       QianchuanAdDetailGetV10DataStatus = "REAUDIT"
	ROI2_DISABLE_QianchuanAdDetailGetV10DataStatus                  QianchuanAdDetailGetV10DataStatus = "ROI2_DISABLE"
	SYSTEM_DISABLE_QianchuanAdDetailGetV10DataStatus                QianchuanAdDetailGetV10DataStatus = "SYSTEM_DISABLE"
	TIME_DONE_QianchuanAdDetailGetV10DataStatus                     QianchuanAdDetailGetV10DataStatus = "TIME_DONE"
	TIME_NO_REACH_QianchuanAdDetailGetV10DataStatus                 QianchuanAdDetailGetV10DataStatus = "TIME_NO_REACH"
)

// All allowed values of QianchuanAdDetailGetV10DataStatus enum
var AllowedQianchuanAdDetailGetV10DataStatusEnumValues = []QianchuanAdDetailGetV10DataStatus{
	"ADVERTISER_OFFLINE_BUDGET",
	"ADVERTISER_PRE_OFFLINE_BUDGET",
	"AUDIT",
	"AUDIT_STATUS_ERROR",
	"AWEME_ACCOUNT_DISABLED",
	"CAMPAIGN_DISABLE",
	"CAMPAIGN_OFFLINE_BUDGET",
	"CAMPAIGN_PRE_OFFLINE_BUDGET",
	"CREATE",
	"DELETE",
	"DELIVERY_OK",
	"DISABLE",
	"DRAFT",
	"ERROR",
	"EXTERNAL_URL_DISABLE",
	"FROZEN",
	"LIVE_ROOM_OFF",
	"NO_SCHEDULE",
	"OFFLINE_AUDIT",
	"OFFLINE_BALANCE",
	"OFFLINE_BUDGET",
	"PRE_OFFLINE_BUDGET",
	"PRE_ONLINE",
	"QUOTA_DISABLE",
	"REAUDIT",
	"ROI2_DISABLE",
	"SYSTEM_DISABLE",
	"TIME_DONE",
	"TIME_NO_REACH",
}

// NewQianchuanAdDetailGetV10DataStatusFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataStatusFromValue(v string) (*QianchuanAdDetailGetV10DataStatus, error) {
	ev := QianchuanAdDetailGetV10DataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataStatus: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_status value
func (v QianchuanAdDetailGetV10DataStatus) Ptr() *QianchuanAdDetailGetV10DataStatus {
	return &v
}
