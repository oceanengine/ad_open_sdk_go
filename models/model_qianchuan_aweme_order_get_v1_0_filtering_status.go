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

// QianchuanAwemeOrderGetV10FilteringStatus
type QianchuanAwemeOrderGetV10FilteringStatus string

// List of qianchuan_aweme_order_get_v1.0_filtering_status
const (
	AUDIT_QianchuanAwemeOrderGetV10FilteringStatus         QianchuanAwemeOrderGetV10FilteringStatus = "AUDIT"
	BOOK_QianchuanAwemeOrderGetV10FilteringStatus          QianchuanAwemeOrderGetV10FilteringStatus = "BOOK"
	CREATING_QianchuanAwemeOrderGetV10FilteringStatus      QianchuanAwemeOrderGetV10FilteringStatus = "CREATING"
	DELIVERY_OK_QianchuanAwemeOrderGetV10FilteringStatus   QianchuanAwemeOrderGetV10FilteringStatus = "DELIVERY_OK"
	FAILED_QianchuanAwemeOrderGetV10FilteringStatus        QianchuanAwemeOrderGetV10FilteringStatus = "FAILED"
	FINISHED_QianchuanAwemeOrderGetV10FilteringStatus      QianchuanAwemeOrderGetV10FilteringStatus = "FINISHED"
	FROZEN_QianchuanAwemeOrderGetV10FilteringStatus        QianchuanAwemeOrderGetV10FilteringStatus = "FROZEN"
	OFFLINE_AUDIT_QianchuanAwemeOrderGetV10FilteringStatus QianchuanAwemeOrderGetV10FilteringStatus = "OFFLINE_AUDIT"
	OVER_QianchuanAwemeOrderGetV10FilteringStatus          QianchuanAwemeOrderGetV10FilteringStatus = "OVER"
	TIMEOUT_QianchuanAwemeOrderGetV10FilteringStatus       QianchuanAwemeOrderGetV10FilteringStatus = "TIMEOUT"
	UNPAID_QianchuanAwemeOrderGetV10FilteringStatus        QianchuanAwemeOrderGetV10FilteringStatus = "UNPAID"
	UNPAIDCANCEL_QianchuanAwemeOrderGetV10FilteringStatus  QianchuanAwemeOrderGetV10FilteringStatus = "UNPAIDCANCEL"
)

// All allowed values of QianchuanAwemeOrderGetV10FilteringStatus enum
var AllowedQianchuanAwemeOrderGetV10FilteringStatusEnumValues = []QianchuanAwemeOrderGetV10FilteringStatus{
	"AUDIT",
	"BOOK",
	"CREATING",
	"DELIVERY_OK",
	"FAILED",
	"FINISHED",
	"FROZEN",
	"OFFLINE_AUDIT",
	"OVER",
	"TIMEOUT",
	"UNPAID",
	"UNPAIDCANCEL",
}

// NewQianchuanAwemeOrderGetV10FilteringStatusFromValue returns a pointer to a valid QianchuanAwemeOrderGetV10FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderGetV10FilteringStatusFromValue(v string) (*QianchuanAwemeOrderGetV10FilteringStatus, error) {
	ev := QianchuanAwemeOrderGetV10FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderGetV10FilteringStatus: valid values are %v", v, AllowedQianchuanAwemeOrderGetV10FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderGetV10FilteringStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderGetV10FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_get_v1.0_filtering_status value
func (v QianchuanAwemeOrderGetV10FilteringStatus) Ptr() *QianchuanAwemeOrderGetV10FilteringStatus {
	return &v
}
