/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

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

// Ptr returns reference to qianchuan_report_creative_get_v1.0_filtering_status value
func (v QianchuanReportCreativeGetV10FilteringStatus) Ptr() *QianchuanReportCreativeGetV10FilteringStatus {
	return &v
}
