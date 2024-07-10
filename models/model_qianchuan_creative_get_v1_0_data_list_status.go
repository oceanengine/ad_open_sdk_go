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

// QianchuanCreativeGetV10DataListStatus
type QianchuanCreativeGetV10DataListStatus string

// List of qianchuan_creative_get_v1.0_data_list_status
const (
	ADVERTISER_OFFLINE_BUDGET_QianchuanCreativeGetV10DataListStatus     QianchuanCreativeGetV10DataListStatus = "ADVERTISER_OFFLINE_BUDGET"
	ADVERTISER_PRE_OFFLINE_BUDGET_QianchuanCreativeGetV10DataListStatus QianchuanCreativeGetV10DataListStatus = "ADVERTISER_PRE_OFFLINE_BUDGET"
	AD_AUDIT_QianchuanCreativeGetV10DataListStatus                      QianchuanCreativeGetV10DataListStatus = "AD_AUDIT"
	AD_DISABLE_QianchuanCreativeGetV10DataListStatus                    QianchuanCreativeGetV10DataListStatus = "AD_DISABLE"
	AD_EXTERNAL_URL_DISABLE_QianchuanCreativeGetV10DataListStatus       QianchuanCreativeGetV10DataListStatus = "AD_EXTERNAL_URL_DISABLE"
	AD_OFFLINE_AUDIT_QianchuanCreativeGetV10DataListStatus              QianchuanCreativeGetV10DataListStatus = "AD_OFFLINE_AUDIT"
	AD_OFFLINE_BUDGET_QianchuanCreativeGetV10DataListStatus             QianchuanCreativeGetV10DataListStatus = "AD_OFFLINE_BUDGET"
	AD_PRE_OFFLINE_BUDGET_QianchuanCreativeGetV10DataListStatus         QianchuanCreativeGetV10DataListStatus = "AD_PRE_OFFLINE_BUDGET"
	AD_REAUDIT_QianchuanCreativeGetV10DataListStatus                    QianchuanCreativeGetV10DataListStatus = "AD_REAUDIT"
	AUDIT_QianchuanCreativeGetV10DataListStatus                         QianchuanCreativeGetV10DataListStatus = "AUDIT"
	AWEME_ITEM_DISABLED_QianchuanCreativeGetV10DataListStatus           QianchuanCreativeGetV10DataListStatus = "AWEME_ITEM_DISABLED"
	CAMPAIGN_DISABLE_QianchuanCreativeGetV10DataListStatus              QianchuanCreativeGetV10DataListStatus = "CAMPAIGN_DISABLE"
	CAMPAIGN_OFFLINE_BUDGET_QianchuanCreativeGetV10DataListStatus       QianchuanCreativeGetV10DataListStatus = "CAMPAIGN_OFFLINE_BUDGET"
	CAMPAIGN_PRE_OFFLINE_BUDGET_QianchuanCreativeGetV10DataListStatus   QianchuanCreativeGetV10DataListStatus = "CAMPAIGN_PRE_OFFLINE_BUDGET"
	CREATE_QianchuanCreativeGetV10DataListStatus                        QianchuanCreativeGetV10DataListStatus = "CREATE"
	DELETE_QianchuanCreativeGetV10DataListStatus                        QianchuanCreativeGetV10DataListStatus = "DELETE"
	DELIVERY_OK_QianchuanCreativeGetV10DataListStatus                   QianchuanCreativeGetV10DataListStatus = "DELIVERY_OK"
	DISABLE_QianchuanCreativeGetV10DataListStatus                       QianchuanCreativeGetV10DataListStatus = "DISABLE"
	ERROR_QianchuanCreativeGetV10DataListStatus                         QianchuanCreativeGetV10DataListStatus = "ERROR"
	FROZEN_QianchuanCreativeGetV10DataListStatus                        QianchuanCreativeGetV10DataListStatus = "FROZEN"
	LIVE_CAN_NOT_LAUNCH_QianchuanCreativeGetV10DataListStatus           QianchuanCreativeGetV10DataListStatus = "LIVE_CAN_NOT_LAUNCH"
	NO_SCHEDULE_QianchuanCreativeGetV10DataListStatus                   QianchuanCreativeGetV10DataListStatus = "NO_SCHEDULE"
	OFFLINE_AUDIT_QianchuanCreativeGetV10DataListStatus                 QianchuanCreativeGetV10DataListStatus = "OFFLINE_AUDIT"
	OFFLINE_BALANCE_QianchuanCreativeGetV10DataListStatus               QianchuanCreativeGetV10DataListStatus = "OFFLINE_BALANCE"
	PRE_ONLINE_QianchuanCreativeGetV10DataListStatus                    QianchuanCreativeGetV10DataListStatus = "PRE_ONLINE"
	REAUDIT_QianchuanCreativeGetV10DataListStatus                       QianchuanCreativeGetV10DataListStatus = "REAUDIT"
	TIME_DONE_QianchuanCreativeGetV10DataListStatus                     QianchuanCreativeGetV10DataListStatus = "TIME_DONE"
	TIME_NO_REACH_QianchuanCreativeGetV10DataListStatus                 QianchuanCreativeGetV10DataListStatus = "TIME_NO_REACH"
)

// All allowed values of QianchuanCreativeGetV10DataListStatus enum
var AllowedQianchuanCreativeGetV10DataListStatusEnumValues = []QianchuanCreativeGetV10DataListStatus{
	"ADVERTISER_OFFLINE_BUDGET",
	"ADVERTISER_PRE_OFFLINE_BUDGET",
	"AD_AUDIT",
	"AD_DISABLE",
	"AD_EXTERNAL_URL_DISABLE",
	"AD_OFFLINE_AUDIT",
	"AD_OFFLINE_BUDGET",
	"AD_PRE_OFFLINE_BUDGET",
	"AD_REAUDIT",
	"AUDIT",
	"AWEME_ITEM_DISABLED",
	"CAMPAIGN_DISABLE",
	"CAMPAIGN_OFFLINE_BUDGET",
	"CAMPAIGN_PRE_OFFLINE_BUDGET",
	"CREATE",
	"DELETE",
	"DELIVERY_OK",
	"DISABLE",
	"ERROR",
	"FROZEN",
	"LIVE_CAN_NOT_LAUNCH",
	"NO_SCHEDULE",
	"OFFLINE_AUDIT",
	"OFFLINE_BALANCE",
	"PRE_ONLINE",
	"REAUDIT",
	"TIME_DONE",
	"TIME_NO_REACH",
}

// NewQianchuanCreativeGetV10DataListStatusFromValue returns a pointer to a valid QianchuanCreativeGetV10DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10DataListStatusFromValue(v string) (*QianchuanCreativeGetV10DataListStatus, error) {
	ev := QianchuanCreativeGetV10DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10DataListStatus: valid values are %v", v, AllowedQianchuanCreativeGetV10DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10DataListStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_data_list_status value
func (v QianchuanCreativeGetV10DataListStatus) Ptr() *QianchuanCreativeGetV10DataListStatus {
	return &v
}
